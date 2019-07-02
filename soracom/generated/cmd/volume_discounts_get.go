// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VolumeDiscountsGetCmdContractId holds value of 'contract_id' option
var VolumeDiscountsGetCmdContractId string

func init() {
	VolumeDiscountsGetCmd.Flags().StringVar(&VolumeDiscountsGetCmdContractId, "contract-id", "", TRAPI("contract_id"))

	VolumeDiscountsGetCmd.MarkFlagRequired("contract-id")

	VolumeDiscountsCmd.AddCommand(VolumeDiscountsGetCmd)
}

// VolumeDiscountsGetCmd defines 'get' subcommand
var VolumeDiscountsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/volume_discounts/{contract_id}:get:summary"),
	Long:  TRAPI(`/volume_discounts/{contract_id}:get:description`),
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

		param, err := collectVolumeDiscountsGetCmdParams(ac)
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

func collectVolumeDiscountsGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForVolumeDiscountsGetCmd("/volume_discounts/{contract_id}"),
		query:  buildQueryForVolumeDiscountsGetCmd(),
	}, nil
}

func buildPathForVolumeDiscountsGetCmd(path string) string {

	escapedContractId := url.PathEscape(VolumeDiscountsGetCmdContractId)

	path = strings.Replace(path, "{"+"contract_id"+"}", escapedContractId, -1)

	return path
}

func buildQueryForVolumeDiscountsGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
