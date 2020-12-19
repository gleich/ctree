package lights

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/wayneashleyberry/truecolor/pkg/color"
)

func ApplyColorsToLights(t string) string {
	applied := t

	for i := 1; i < 20; i++ {
		iStr := fmt.Sprint(i) + "O"
		applied = strings.Replace(applied, iStr, generateColor("O"), 1)
	}

	return applied
}

// Apply a random color for the light.
func generateColor(in string) string {
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
