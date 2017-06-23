package main

import (
	_ "fmt"
	"os"

	"github.com/mgm702/gitinfo/cmds"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Gitinfo"
	app.Version = "0.0.1"

	app.EnableBashCompletion = true
	cmds.AllCmds(app)

	// default action which returns all information
	app.Action = func(c *cli.Context) error {
		cmds.AllInfo(c)
		return nil
	}
	app.Run(os.Args)
}
