package cmds

import (
	"fmt"

	"github.com/mgm702/gitinfo/lib"
	"github.com/urfave/cli"
)

func UrlCmd(a *cli.App) {
	url := cli.Command{
		Name:    "url",
		Aliases: []string{"u"},
		Usage:   "Prints the Urls of a project",
		Action:  urlInfo,
	}
	a.Commands = append(a.Commands, url)
}

func urlInfo(c *cli.Context) {
	repo := lib.GatherInfo(c)
	printUrls(repo)
}

func printUrls(repo lib.RepoInfo) {
	fmt.Println("------------")
	fmt.Println("|URL       |", repo.Url)
	fmt.Println("|CLONE URL |", repo.CloneUrl)
	fmt.Println("|SSH URL   |", repo.SshUrl)
	fmt.Println("------------")
}
