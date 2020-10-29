package main

import (
	"fmt"
	"net"

	"github.com/alecthomas/kong"
	"github.com/gschauer/cli-cmp/chat"
)

func main() {
	cli := CLI{
		Globals: Globals{},
	}

	ctx := kong.Parse(&cli,
		kong.Name("chat"),
		kong.Description("A command-line chat application."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"DEFAULT_CONFIG": "~/.chat.conf",
			"version":        chat.Version,
		})

	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}

type CLI struct {
	Globals
	Connect ConnectCmd `cmd help:"Connect to chat server."`
}

type Globals struct {
	Config   string      `short:"c" help:"Location of client config file. (default: ${DEFAULT_CONFIG})." default:"${DEFAULT_CONFIG}" env:"~/.chat.conf" type:"path"`
	Timeout  uint        `help:"Idle timeout in minutes." default:"30"`
	LogLevel string      `name:"log-level" short:"l" help:"Set the log level (debug|info|warn|error)." default:"info" enum:"debug,info,warn,error"`
	Version  VersionFlag `help:"Show application version."`
}

type ConnectCmd struct {
	Server   net.IP   `required:"" short:"s" help:"Server address." default:"127.0.0.1"`
	Channels []string `arg:"" required:"" help:"Channels to open."`
}

func (c *ConnectCmd) Run(*Globals) error {
	return chat.Connect(c.Server.String(), c.Channels)
}

type VersionFlag bool

func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}
