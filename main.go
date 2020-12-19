package main

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/ctree/pkg/tree"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ctree",
		Usage: "🎄 Christmas tree right from your terminal!",
		Action: func(c *cli.Context) error {
			t := tree.MsgBase
			gc := tree.ApplyColors(t)
			fmt.Println(gc)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		statuser.Error("Failed to parse flags and run main command", err, 1)
	}
}
