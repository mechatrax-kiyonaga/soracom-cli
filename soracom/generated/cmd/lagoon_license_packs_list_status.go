// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LagoonLicensePacksListStatusCmdOutputJSONL indicates to output with jsonl format
var LagoonLicensePacksListStatusCmdOutputJSONL bool

func init() {
	LagoonLicensePacksListStatusCmd.Flags().BoolVar(&LagoonLicensePacksListStatusCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	LagoonLicensePacksCmd.AddCommand(LagoonLicensePacksListStatusCmd)
}

// LagoonLicensePacksListStatusCmd defines 'list-status' subcommand
var LagoonLicensePacksListStatusCmd = &cobra.Command{
	Use:   "list-status",
	Short: TRAPI("/lagoon/license_packs:get:summary"),
	Long:  TRAPI(`/lagoon/license_packs:get:description`) + "\n\n" + createLinkToAPIReference("Lagoon", "listLagoonLicensePackStatus"),
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) > 0 {
			return fmt.Errorf("unexpected arguments passed => %v", args)
		}

		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectLagoonLicensePacksListStatusCmdParams(ac)
		if err != nil {
			return err
		}

		body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			if LagoonLicensePacksListStatusCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectLagoonLicensePacksListStatusCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLagoonLicensePacksListStatusCmd("/lagoon/license_packs"),
		query:  buildQueryForLagoonLicensePacksListStatusCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonLicensePacksListStatusCmd(path string) string {

	return path
}

func buildQueryForLagoonLicensePacksListStatusCmd() url.Values {
	result := url.Values{}

	return result
}
