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
		Usage: "🎄 A Christmas tree right from your terminal!\nRepository: https://github.com/Matt-Gleich/ctree",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "no-refresh",
				Usage: "Just output the tree once normally",
			},
		},
		Action: func(c *cli.Context) error {
			base := tree.ApplyColors(tree.MsgBase)
			if c.Bool("no-refresh") {
				fmt.Println(lights.ApplyColorsToLights(base))
				os.Exit(0)
			}

			for {
				screen.Clear()
				screen.MoveTopLeft()
				fmt.Println(lights.ApplyColorsToLights(base))
				time.Sleep(time.Second * 3)
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		statuser.Error("Failed to parse flags and run main command", err, 1)
	}
}
