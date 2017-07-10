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
	app.Usage = "CLI which helps retrieve Github Repo info"
	app.Version = "1.0.0"

	app.EnableBashCompletion = true
	cmds.LoadCmds(app)

	// default action which returns all information
	app.Action = func(c *cli.Context) error {
		cmds.AllInfo(c)
		return nil
	}
	app.Run(os.Args)
}
