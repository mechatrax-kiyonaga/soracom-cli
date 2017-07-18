package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(VpgCmd)
}

// VpgCmd defines 'vpg' subcommand
var VpgCmd = &cobra.Command{
	Use:   "vpg",
	Short: TRCLI("cli.vpg.summary"),
	Long:  TRCLI(`cli.vpg.description`),
}
