package main

import (
	"os"

	"github.com/gschauer/cli-cmp/chat"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("chat", "A command-line chat application.")
	app.Version(chat.Version)

	app.Flag("config", "Location of client config file.").Short('c').Default("~/.chat.conf").Envar("CHAT_CONFIG").String()
	app.Flag("timeout", "Idle timeout in minutes.").Default("30").Int()
	app.Flag("log-level", "Set the log level (debug|info|warn|error).").Short('l').Default("info").Enum("debug", "info", "warn", "error")

	connect := app.Command("connect", "Connect to chat server.")
	server := connect.Flag("server", "Server address.").Short('s').Default("127.0.0.1").IP()
	channels := connect.Arg("channels", "Channels to open.").Required().Strings()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case connect.FullCommand():
		_ = chat.Connect(server.String(), *channels)
	}
}
