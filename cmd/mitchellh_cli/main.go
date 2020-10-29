package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gschauer/cli-cmp/chat"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("chat", chat.Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"connect": func() (cli.Command, error) {
			return &Connect{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}

type Connect struct {
}

func (*Connect) Help() string {
	return "Connect to chat server."
}

func (c *Connect) Run(args []string) int {
	set := flag.NewFlagSet("connect", flag.ContinueOnError)
	server := set.String("server", "127.0.0.1", "Server address.")

	if err := set.Parse(args); err == nil {
		if err = chat.Connect(*server, set.Args()); err == nil {
			return 0
		}
	}

	fmt.Println(c.Help())
	return 1
}

func (c *Connect) Synopsis() string {
	return c.Help()
}
