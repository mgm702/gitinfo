package cmds

import (
	"fmt"
	_ "os"

	_ "github.com/mgm702/gitinfo/lib"
	"github.com/urfave/cli"
)

func AllCmds(a *cli.App) {
	// Here we will call to each of the commands in
	// there files and build the command list to
	// return to the main application
	DescriptionCmd(a)
}

func AllInfo() {
	// if this function is called all info is loaded
	// and displayed to user
	fmt.Println("Hello, gathering all the info!")
}

// Action:       helpCommand.Action,

//var helpCommand = Command{
//Name:      "help",
//Aliases:   []string{"h"},
//Usage:     "Shows a list of commands or help for one command",
//ArgsUsage: "[command]",
//Action: func(c *Context) error {
//args := c.Args()
//if args.Present() {
//return ShowCommandHelp(c, args.First())
//}

//ShowAppHelp(c)
//return nil
//},
//}
