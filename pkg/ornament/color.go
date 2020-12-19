package ornament

import (
	"math/rand"

	"github.com/wayneashleyberry/truecolor/pkg/color"
)

// Apply a random color to a ornament.
func GenerateColor(in string) string {
	var (
		selection = rand.Intn(5)
		c         *color.Message
	)

	switch selection {
	case 0:
		// Pink
		c = color.Color(247, 108, 246)
	case 1:
		// Light Blue
		c = color.Color(85, 202, 255)
	case 2:
		// Purple
		c = color.Color(100, 0, 255)
	case 3:
		// Red
		c = color.Color(236, 21, 0)
	case 4:
		// Orange
		c = color.Color(243, 213, 0)
	case 5:
		// Light Green
		c = color.Color(100, 255, 24)
	}
	return c.Sprint(in)
}
