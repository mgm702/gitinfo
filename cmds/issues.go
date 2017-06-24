package cmds

import (
	"fmt"

	"github.com/mgm702/gitinfo/lib"
	"github.com/urfave/cli"
)

func IssuesCmd(a *cli.App) {
	issues := cli.Command{
		Name:    "issues",
		Aliases: []string{"i"},
		Usage:   "Issues of a project",
		Action:  issuesInfo,
	}
	a.Commands = append(a.Commands, issues)
}

func issuesInfo(c *cli.Context) {
	repo := lib.GatherInfo(c)
	printIssues(repo)
}

func printIssues(repo lib.RepoInfo) {
	fmt.Println("--------------")
	fmt.Println("|OPEN ISSUES | ", repo.Issues)
	fmt.Println("--------------")
}
