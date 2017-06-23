package cmds

import (
	_ "fmt"

	"github.com/mgm702/gitinfo/lib"
	"github.com/urfave/cli"
)

func DescriptionCmd(a *cli.App) {
	desc := cli.Command{
		Name:    "description",
		Aliases: []string{"d"},
		Usage:   "returns the description of a project",
		Action:  descriptionInfo,
	}
	a.Commands = append(a.Commands, desc)
}

func descriptionInfo(c *cli.Context) {
	lib.GatherInfo(c)
}
