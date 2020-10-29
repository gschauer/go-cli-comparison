package main

import (
	"fmt"
	"os"

	"github.com/gschauer/cli-cmp/chat"
	"github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("chat", "A command-line chat application.")

	app.StringOpt("c config", "~/.chat.conf", "Location of client config file.")
	app.IntOpt("timeout", 30, "Idle timeout in minutes.")
	app.StringOpt("l log-level", "info", "Set the log level (debug|info|warn|error).")
	version := app.BoolOpt("version", false, "Show application version.")

	app.Command("connect", "Connect to chat server.", connect)

	// Specify the action to execute when the app is invoked correctly
	app.Action = func() {
		if *version {
			fmt.Println(chat.Version)
			os.Exit(0)
		}

		app.PrintHelp()
	}

	_ = app.Run(os.Args)
}

func connect(c *cli.Cmd) {
	c.Spec = "[ -s=<server> ] CHANNELS..."
	server := c.StringOpt("s server", "127.0.0.1", "Server address.")
	channels := c.StringsArg("CHANNELS", nil, "Channels to open.")

	c.Action = func() {
		_ = chat.Connect(*server, *channels)
	}
}
