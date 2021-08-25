// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgUnsetRedirectionCmdVpgId holds value of 'vpg_id' option
var VpgUnsetRedirectionCmdVpgId string

func init() {
	VpgUnsetRedirectionCmd.Flags().StringVar(&VpgUnsetRedirectionCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))
	VpgCmd.AddCommand(VpgUnsetRedirectionCmd)
}

// VpgUnsetRedirectionCmd defines 'unset-redirection' subcommand
var VpgUnsetRedirectionCmd = &cobra.Command{
	Use:   "unset-redirection",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/junction/unset_redirection:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/junction/unset_redirection:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
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

		param, err := collectVpgUnsetRedirectionCmdParams(ac)
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
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectVpgUnsetRedirectionCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgUnsetRedirectionCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgUnsetRedirectionCmd("/virtual_private_gateways/{vpg_id}/junction/unset_redirection"),
		query:  buildQueryForVpgUnsetRedirectionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgUnsetRedirectionCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgUnsetRedirectionCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgUnsetRedirectionCmd() url.Values {
	result := url.Values{}

	return result
}
