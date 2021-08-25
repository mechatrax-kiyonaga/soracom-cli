// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgListIpAddressMapEntriesCmdVpgId holds value of 'vpg_id' option
var VpgListIpAddressMapEntriesCmdVpgId string

func init() {
	VpgListIpAddressMapEntriesCmd.Flags().StringVar(&VpgListIpAddressMapEntriesCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))
	VpgCmd.AddCommand(VpgListIpAddressMapEntriesCmd)
}

// VpgListIpAddressMapEntriesCmd defines 'list-ip-address-map-entries' subcommand
var VpgListIpAddressMapEntriesCmd = &cobra.Command{
	Use:   "list-ip-address-map-entries",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/ip_address_map:get:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/ip_address_map:get:description`),
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

		param, err := collectVpgListIpAddressMapEntriesCmdParams(ac)
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

func collectVpgListIpAddressMapEntriesCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgListIpAddressMapEntriesCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForVpgListIpAddressMapEntriesCmd("/virtual_private_gateways/{vpg_id}/ip_address_map"),
		query:  buildQueryForVpgListIpAddressMapEntriesCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgListIpAddressMapEntriesCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgListIpAddressMapEntriesCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgListIpAddressMapEntriesCmd() url.Values {
	result := url.Values{}

	return result
}
