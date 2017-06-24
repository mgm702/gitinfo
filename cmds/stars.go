package cmds

import (
	"fmt"

	"github.com/mgm702/gitinfo/lib"
	"github.com/urfave/cli"
)

func StarsCmd(a *cli.App) {
	stars := cli.Command{
		Name:    "stars",
		Aliases: []string{"s"},
		Usage:   "Stars of a project",
		Action:  starsInfo,
	}
	a.Commands = append(a.Commands, stars)
}

func starsInfo(c *cli.Context) {
	repo := lib.GatherInfo(c)
	printStars(repo)
}

func printStars(repo lib.RepoInfo) {
	fmt.Println("--------")
	fmt.Println("|STARS | ", repo.Stars)
	fmt.Println("--------")
}
