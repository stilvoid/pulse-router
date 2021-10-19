package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const DEBUG = false

type Value interface{}

type Map struct {
	Key      string
	Value    string
	Children []Value
}

var keyIsh = regexp.MustCompile(`^([^:]+):(?: (.+))?$`)
var equIsh = regexp.MustCompile(`^([^=]+) = (.+)$`)

func getIndent(line string) (string, int) {
	text := strings.TrimLeft(line, " ")
	indent := len(line) - len(text)

	return text, indent
}

func debug(args ...interface{}) {
	if DEBUG {
		fmt.Fprintln(os.Stderr, args...)
	}
}

func parseValue(lines []string, indent int) (Value, []string) {
	if len(lines) == 0 {
		debug("DONE")
		return nil, lines
	}

	debug("LINE:", lines[0], len(lines))

	line, newIndent := getIndent(lines[0])

	// Skip zero-level lines
	if newIndent == 0 {
		debug("Ignoring:", lines[0])
		return parseValue(lines[1:], indent)
	}

	// Finished, time to ascend
	if newIndent <= indent {
		debug("Ascending")
		return nil, lines
	}

	m := keyIsh.FindStringSubmatch(line)
	if m == nil {
		m = equIsh.FindStringSubmatch(line)
	}
	if m == nil {
		//panic(fmt.Errorf("Can't parse: %s", line))
		return parseValue(lines[1:], indent)
	}

	children, lines := parseChildren(lines[1:], newIndent)

	return Map{
		Key:      m[1],
		Value:    m[2],
		Children: children,
	}, lines
}

func parseChildren(lines []string, indent int) ([]Value, []string) {
	children := make([]Value, 0)

	var child Value
	for {
		child, lines = parseValue(lines, indent)
		if child == nil {
			break
		}

		children = append(children, child)
	}

	return children, lines
}

func ParseValue(lines []string) Value {
	children, _ := parseChildren(lines, 0)

	return Map{
		Children: children,
	}
}

func main() {
	in, _ := ioutil.ReadFile("input")
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(in)), "\t", "        "), "\n")

	val := ParseValue(lines)

	out, _ := json.MarshalIndent(val, "", "    ")
	fmt.Println(string(out))
}
