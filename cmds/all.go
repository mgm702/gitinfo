package cmds

import (
	"fmt"

	_ "github.com/mgm702/gitinfo/lib"
	"github.com/urfave/cli"
)

func AllCmds(a *cli.App) {
	// Here we will call to each of the commands in
	// there files and build the command list to
	// return to the main application

	//StarsCmd(a)
	//ForksCmd(a)
	//IssuesCmd(a)
	//UrlCmd(a)
	DescriptionCmd(a)
}

func AllInfo(c *cli.Context) {
	// if this function is called all info is loaded
	// and displayed to user
	fmt.Printf("Hello from AllInfo: %q\n", c.Args().Get(0))
}
