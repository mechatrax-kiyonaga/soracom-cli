package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	LagoonLicensePacksCmd.AddCommand(LagoonLicensePacksListStatusCmd)
}

// LagoonLicensePacksListStatusCmd defines 'list-status' subcommand
var LagoonLicensePacksListStatusCmd = &cobra.Command{
	Use:   "list-status",
	Short: TRAPI("/lagoon/license_packs:get:summary"),
	Long:  TRAPI(`/lagoon/license_packs:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectLagoonLicensePacksListStatusCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectLagoonLicensePacksListStatusCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLagoonLicensePacksListStatusCmd("/lagoon/license_packs"),
		query:  buildQueryForLagoonLicensePacksListStatusCmd(),
	}, nil
}

func buildPathForLagoonLicensePacksListStatusCmd(path string) string {

	return path
}

func buildQueryForLagoonLicensePacksListStatusCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
