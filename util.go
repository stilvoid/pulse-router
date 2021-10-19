package main

import (
	"strings"
)

func indentString(s, indent string) string {
	parts := strings.Split(s, "\n")
	for i, part := range parts {
		if strings.TrimSpace(parts[i]) != "" {
			parts[i] = indent + part
		}
	}

	return strings.Join(parts, "\n")
}
