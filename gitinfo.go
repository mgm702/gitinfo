package main

import (
	_ "fmt"
	"os"

	_ "github.com/caarlos0/spin"
	"github.com/mgm702/gitinfo/cmds"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"

	app.EnableBashCompletion = true
	cmds.AllCmds(app)

	app.Action = func(c *cli.Context) error {
		cmds.AllInfo()
		return nil
	}
	app.Run(os.Args)
}
