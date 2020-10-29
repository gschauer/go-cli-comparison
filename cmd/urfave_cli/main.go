package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gschauer/cli-cmp/chat"
	"github.com/urfave/cli/v2"
)

func main() {
	connectCmd := &cli.Command{
		Name:      "connect",
		Usage:     "Connect to chat server.",
		ArgsUsage: "CHANNEL...",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "server",
				Aliases:  []string{"s"},
				Usage:    "Server address.",
				Value:    "127.0.0.1",
			},
		},
		Before: func(c *cli.Context) error {
			if c.Args().First() == "" {
				return errors.New("no channels")
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			return chat.Connect(c.Value("server").(string), c.Args().Slice())
		},
	}

	app := &cli.App{
		Name:    "chat",
		Usage:   "A command-line chat application.",
		Version: chat.Version,
		Commands: []*cli.Command{
			connectCmd,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "config",
				Aliases:   []string{"c"},
				Usage:     "Location of client config file.",
				EnvVars:   []string{"CHAT_CONFIG_FILE"},
				TakesFile: true,
				Value:     "~/.chat.conf",
			},
			&cli.IntFlag{
				Name:  "timeout",
				Usage: "Idle timeout in minutes.",
				Value: 30,
			},
			&cli.GenericFlag{
				Name:    "log-level",
				Aliases: []string{"l"},
				Usage:   "Set the log level (debug|info|warn|error).",
				Value: &EnumValue{
					Enum:    []string{"debug", "info", "warn", "error"},
					Default: "info",
				},
			},
		},
	}

	_ = app.Run(os.Args)
}

type EnumValue struct {
	Enum     []string
	Default  string
	selected string
}

func (e *EnumValue) Set(value string) error {
	for _, enum := range e.Enum {
		if enum == value {
			e.selected = value
			return nil
		}
	}

	return fmt.Errorf("allowed values are %s", strings.Join(e.Enum, ", "))
}

func (e EnumValue) String() string {
	if e.selected == "" {
		return e.Default
	}
	return e.selected
}
