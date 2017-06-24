package cmds

import (
	"fmt"

	"github.com/mgm702/gitinfo/lib"
	"github.com/urfave/cli"
)

func ForksCmd(a *cli.App) {
	forks := cli.Command{
		Name:    "forks",
		Aliases: []string{"f"},
		Usage:   "Forks for a project",
		Action:  forksInfo,
	}
	a.Commands = append(a.Commands, forks)
}

func forksInfo(c *cli.Context) {
	repo := lib.GatherInfo(c)
	printForks(repo)
}

func printForks(repo lib.RepoInfo) {
	fmt.Println("--------")
	fmt.Println("|FORKS | ", repo.Forks)
	fmt.Println("--------")
}
