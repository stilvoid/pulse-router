package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lawl/pulseaudio"
)

type Source pulseaudio.Source

func (s Source) String() string {
	out := strings.Builder{}

	out.WriteString(fmt.Sprintf("Name: %s\n", s.Name))
	out.WriteString(fmt.Sprintf("Description: %s\n", s.Description))
	out.WriteString(fmt.Sprintf("Module: %s\n", Client.Modules[int(s.ModuleIndex)].Name))
	out.WriteString(fmt.Sprintf("Channels: %s\n", fmt.Sprintf("%#v", s.ChannelMap)))

	return out.String()
}

type Sources map[int]Source

func (s Sources) String() string {
	ids := make([]int, 0)
	for id := range s {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	out := strings.Builder{}

	for _, id := range ids {
		out.WriteString(fmt.Sprintf("%d:\n", id))
		out.WriteString(indentString(fmt.Sprint(s[id]), "  "))
	}

	return out.String()
}
