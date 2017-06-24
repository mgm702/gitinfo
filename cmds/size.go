package cmds

import (
	"fmt"

	"github.com/mgm702/gitinfo/lib"
	"github.com/urfave/cli"
)

func SizeCmd(a *cli.App) {
	size := cli.Command{
		Name:    "size",
		Aliases: []string{"k"},
		Usage:   "Size of a project",
		Action:  sizeInfo,
	}
	a.Commands = append(a.Commands, size)
}

func sizeInfo(c *cli.Context) {
	repo := lib.GatherInfo(c)
	printSize(repo)
}

func printSize(repo lib.RepoInfo) {
	fmt.Println("------------")
	fmt.Println("|SIZE (KB) | ", readableSize(repo.Size))
	fmt.Println("------------")
}

func readableSize(size int) int {
	return size
}
