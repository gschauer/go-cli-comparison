package main

import (
	"fmt"
	"os"

	"github.com/gschauer/cli-cmp/chat"
	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		Config   *string `long:"config" short:"c" description:"Location of client config file." default:"~/.chat.conf" value-name:"FILE"`
		Timeout  uint    `long:"timeout" description:"The idle timeout in minutes." default:"30"`
		LogLevel string  `long:"log-level" short:"l" description:"Set the log level (debug|info|warn|error)." choice:"debug" choice:"info" choice:"warn" choice:"error"`
		Version  bool    `long:"version" short:"v" description:"Show application version."`

		ConnectCmd struct {
			Server     string `long:"server" short:"s" description:"The server name." default:"127.0.0.1" required:"true"`
			Positional struct {
				Channels []string `description:"The channels to open." required:"true"`
			} `positional-args:"yes"`
		} `command:"connect" description:"Connect to the server."`
	}

	p := flags.NewParser(&opts, flags.Default)
	p.SubcommandsOptional = true

	_, err := p.Parse()
	if err != nil {
		p.WriteHelp(os.Stderr)
		os.Exit(1)
	}

	if p.Active == nil {
		if p.FindOptionByShortName('v').IsSet() {
			fmt.Println(chat.Version)
			os.Exit(0)
		}

		p.WriteHelp(os.Stderr)
		return
	}

	switch p.Active.Name {
	case "connect":
		_ = chat.Connect(opts.ConnectCmd.Server, opts.ConnectCmd.Positional.Channels)
	}
}
