package cmds

import (
	"fmt"

	"github.com/mgm702/gitinfo/lib"
	"github.com/urfave/cli"
)

func DescriptionCmd(a *cli.App) {
	desc := cli.Command{
		Name:    "description",
		Aliases: []string{"d"},
		Usage:   "Description of a project",
		Action:  descriptionInfo,
	}
	a.Commands = append(a.Commands, desc)
}

func descriptionInfo(c *cli.Context) {
	repo := lib.GatherInfo(c)
	printDesc(repo)
}

func printDesc(repo lib.RepoInfo) {
	fmt.Println("--------------")
	fmt.Println("|DESCRIPTION | ", repo.Description)
	fmt.Println("--------------")
}
