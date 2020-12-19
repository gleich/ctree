package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Matt-Gleich/ctree/pkg/lights"
	"github.com/Matt-Gleich/ctree/pkg/tree"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/inancgumus/screen"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ctree",
		Usage: "🎄 Christmas tree right from your terminal!",
		Action: func(c *cli.Context) error {
			t := tree.MsgBase
			for {
				screen.Clear()
				screen.MoveTopLeft()
				base := tree.ApplyColors(t)
				full := lights.ApplyColorsToLights(base)
				fmt.Println(full)
				time.Sleep(time.Second * 3)
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		statuser.Error("Failed to parse flags and run main command", err, 1)
	}
}
