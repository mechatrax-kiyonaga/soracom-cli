package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(SubscribersCmd)
}

// SubscribersCmd defines 'subscribers' subcommand
var SubscribersCmd = &cobra.Command{
	Use:   "subscribers",
	Short: TRCLI("cli.subscribers.summary"),
	Long:  TRCLI(`cli.subscribers.description`),
}
