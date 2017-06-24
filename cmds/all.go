package cmds

import (
	"fmt"

	"github.com/mgm702/gitinfo/lib"
	"github.com/urfave/cli"
)

func LoadCmds(a *cli.App) {
	UrlCmd(a)
	IssuesCmd(a)
	StarsCmd(a)
	ForksCmd(a)
	SizeCmd(a)
	DescriptionCmd(a)
}

func AllInfo(c *cli.Context) {
	repo := lib.GatherInfo(c)
	printAll(repo)
}

func printAll(repo lib.RepoInfo) {
	fmt.Println("--------------")
	fmt.Println("|ID          |", repo.Id)
	fmt.Println("|NAME        |", repo.Name)
	fmt.Println("|HOMEPAGE    |", repo.Homepage)
	fmt.Println("|DESCRIPTION |", repo.Description)
	fmt.Println("--------------")
	fmt.Println("|URL         |", repo.Url)
	fmt.Println("|CLONE URL   |", repo.CloneUrl)
	fmt.Println("|SSH URL     |", repo.SshUrl)
	fmt.Println("--------------")
	// Convert the Size to something readable
	fmt.Println("|SIZE (KB)   |", readableSize(repo.Size))
	fmt.Println("|OPEN ISSUES |", repo.Issues)
	fmt.Println("|STARGAZERS  |", repo.Stars)
	fmt.Println("|FORKS       |", repo.Forks)
	fmt.Println("--------------")
}
