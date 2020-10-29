package main

import (
	"fmt"
	"os"

	"github.com/gschauer/cli-cmp/chat"
	"github.com/spf13/cobra"
)

func main() {
	var cfgFile string

	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "chat",
		Short: "A command-line chat application.",
		Run: func(cmd *cobra.Command, args []string) {
			if ok, err := cmd.LocalNonPersistentFlags().GetBool("version"); err == nil && ok {
				fmt.Println(chat.Version)
				os.Exit(0)
			}
			_ = cmd.Help()
		},
	}

	var connectCmd = &cobra.Command{
		Use:   "connect",
		Short: "Connect to chat server.",
		Run: func(cmd *cobra.Command, args []string) {
			svr, _ := cmd.Flags().GetString("server")
			_ = chat.Connect(svr, args)
		},
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "~/.chat.conf", "Location of client config file.")
	rootCmd.PersistentFlags().Int("timeout", 30, "Idle timeout in minutes.")
	rootCmd.PersistentFlags().StringP("log-level", "l", "info", "Set the log level (debug|info|warn|error).")

	// Cobra also supports local flags, which will only run when this action is called directly.
	rootCmd.Flags().BoolP("version", "v", false, "Show application version.")

	rootCmd.AddCommand(connectCmd)

	connectCmd.Args = cobra.MinimumNArgs(1)
	connectCmd.Flags().StringP("server", "s", "127.0.0.1", "Server address.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
