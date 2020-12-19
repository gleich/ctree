package tree

import (
	"strings"

	"github.com/wayneashleyberry/truecolor/pkg/color"
)

// Apply the colors to the tree.
func ApplyColors(base string) string {
	var (
		yellow  = color.Color(255, 255, 0)
		green   = color.Color(0, 255, 0)
		brown   = color.Color(210, 105, 30)
		applied = base
	)

	applied = applyColorToChars(applied, yellow, "y", []string{
		"|", "-", "A", "+", "*",
	})

	applied = applyColorToChars(applied, green, "g", []string{
		"/", "\\", "=",
	})

	applied = applyColorToChars(applied, brown, "b", []string{
		"|", "_",
	})

	return applied
}

// Apply a certain colors to all the chars that start with a character representing that char.
func applyColorToChars(t string, c *color.Message, cChar string, chars []string) string {
	applied := t
	for _, char := range chars {
		applied = strings.ReplaceAll(applied, cChar+char, c.Sprint(char))
	}
	return applied
}
