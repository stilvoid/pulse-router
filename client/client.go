package client

import (
	"fmt"
	"strings"

	"github.com/lawl/pulseaudio"
)

var client *pulseaudio.Client
var Sources = make(SourceMap)
var Sinks = make(SinkMap)
var Modules = make(ModuleMap)

func init() {
	var err error
	client, err = pulseaudio.NewClient()
	if err != nil {
		panic(err)
	}

	Refresh()
}

func Refresh() {
	// Sources
	sources, err := client.Sources()
	if err != nil {
		panic(err)
	}

	Sources = make(SourceMap)
	for _, source := range sources {
		Sources[int(source.Index)] = Source(source)
	}

	// Sinks
	sinks, err := client.Sinks()
	if err != nil {
		panic(err)
	}

	Sinks = make(SinkMap)
	for _, sink := range sinks {
		Sinks[int(sink.Index)] = Sink(sink)
	}

	// Modules
	modules, err := client.ModuleList()
	if err != nil {
		panic(err)
	}

	Modules = make(ModuleMap)
	for _, module := range modules {
		Modules[int(module.Index)] = Module(module)
	}
}

func String() string {
	out := strings.Builder{}

	out.WriteString("Sources:\n")
	out.WriteString(indentString(fmt.Sprint(Sources), "  "))

	out.WriteString("Sinks:\n")
	out.WriteString(indentString(fmt.Sprint(Sinks), "  "))

	out.WriteString("Modules:\n")
	out.WriteString(indentString(fmt.Sprint(Modules), "  "))

	return out.String()
}
