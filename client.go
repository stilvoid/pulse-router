package main

import (
	"fmt"
	"strings"

	"github.com/lawl/pulseaudio"
)

var Client pulseClient

func init() {
	client, err := pulseaudio.NewClient()
	if err != nil {
		panic(err)
	}

	Client = pulseClient{
		client:  client,
		Sources: make(map[int]Source),
		Sinks:   make(map[int]Sink),
		Modules: make(map[int]Module),
	}
	Client.Refresh()
}

type pulseClient struct {
	client  *pulseaudio.Client
	Sources Sources
	Sinks   Sinks
	Modules Modules
}

func (p *pulseClient) Refresh() {
	// Sources
	sources, err := p.client.Sources()
	if err != nil {
		panic(err)
	}

	p.Sources = make(map[int]Source)
	for _, source := range sources {
		p.Sources[int(source.Index)] = Source(source)
	}

	// Sinks
	sinks, err := p.client.Sinks()
	if err != nil {
		panic(err)
	}

	p.Sinks = make(map[int]Sink)
	for _, sink := range sinks {
		p.Sinks[int(sink.Index)] = Sink(sink)
	}

	// Modules
	modules, err := p.client.ModuleList()
	if err != nil {
		panic(err)
	}

	p.Modules = make(map[int]Module)
	for _, module := range modules {
		p.Modules[int(module.Index)] = Module(module)
	}
}

func (p pulseClient) String() string {
	out := strings.Builder{}

	out.WriteString("Sources:\n")
	out.WriteString(indentString(fmt.Sprint(p.Sources), "  "))

	out.WriteString("Sinks:\n")
	out.WriteString(indentString(fmt.Sprint(p.Sinks), "  "))

	out.WriteString("Modules:\n")
	out.WriteString(indentString(fmt.Sprint(p.Modules), "  "))

	return out.String()
}
