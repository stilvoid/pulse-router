package client

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lawl/pulseaudio"
)

type Module pulseaudio.Module

func (s Module) String() string {
	out := strings.Builder{}

	out.WriteString(fmt.Sprintf("Name: %s\n", s.Name))
	out.WriteString(fmt.Sprintf("Argument: %s\n", s.Argument))
	out.WriteString(fmt.Sprintf("Used: %d\n", s.NUsed))
	out.WriteString(fmt.Sprintf("Properties:\n"))

	for key, value := range s.PropList {
		out.WriteString(fmt.Sprintf("  %s: %s\n", key, value))
	}

	return out.String()
}

type ModuleMap map[int]Module

func (s ModuleMap) String() string {
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
