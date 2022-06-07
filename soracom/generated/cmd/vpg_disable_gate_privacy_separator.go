// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgDisableGatePrivacySeparatorCmdVpgId holds value of 'vpg_id' option
var VpgDisableGatePrivacySeparatorCmdVpgId string

func init() {
	VpgDisableGatePrivacySeparatorCmd.Flags().StringVar(&VpgDisableGatePrivacySeparatorCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))
	VpgCmd.AddCommand(VpgDisableGatePrivacySeparatorCmd)
}

// VpgDisableGatePrivacySeparatorCmd defines 'disable-gate-privacy-separator' subcommand
var VpgDisableGatePrivacySeparatorCmd = &cobra.Command{
	Use:   "disable-gate-privacy-separator",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/gate/disable_privacy_separator:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/gate/disable_privacy_separator:post:description`),
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

		param, err := collectVpgDisableGatePrivacySeparatorCmdParams(ac)
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

func collectVpgDisableGatePrivacySeparatorCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgDisableGatePrivacySeparatorCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgDisableGatePrivacySeparatorCmd("/virtual_private_gateways/{vpg_id}/gate/disable_privacy_separator"),
		query:  buildQueryForVpgDisableGatePrivacySeparatorCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgDisableGatePrivacySeparatorCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgDisableGatePrivacySeparatorCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgDisableGatePrivacySeparatorCmd() url.Values {
	result := url.Values{}

	return result
}
