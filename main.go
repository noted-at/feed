package main

import (
	"os"

	"github.com/agentio/slink/pkg/slink"
	"github.com/noted-at/feed/internal/server"
	"github.com/spf13/cobra"
)

func main() {
	if err := cmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func cmd() *cobra.Command {
	var loglevel string
	cmd := &cobra.Command{
		Use:   "feed",
		Short: "Run the feed server",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := slink.SetLogLevel(loglevel); err != nil {
				return err
			}
			return server.Run()
		},
	}
	cmd.Flags().StringVarP(&loglevel, "log-level", "l", "warn", "log level (debug, info, warn, error, fatal)")
	return cmd
}
