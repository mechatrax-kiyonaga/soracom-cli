// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsListGatewaysCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var LoraNetworkSetsListGatewaysCmdLastEvaluatedKey string

// LoraNetworkSetsListGatewaysCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsListGatewaysCmdNsId string

// LoraNetworkSetsListGatewaysCmdLimit holds value of 'limit' option
var LoraNetworkSetsListGatewaysCmdLimit int64

// LoraNetworkSetsListGatewaysCmdPaginate indicates to do pagination or not
var LoraNetworkSetsListGatewaysCmdPaginate bool

func init() {
	LoraNetworkSetsListGatewaysCmd.Flags().StringVar(&LoraNetworkSetsListGatewaysCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The ID of the last gateway retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next device onward."))

	LoraNetworkSetsListGatewaysCmd.Flags().StringVar(&LoraNetworkSetsListGatewaysCmdNsId, "ns-id", "", TRAPI("ID of the target LoRa network set."))

	LoraNetworkSetsListGatewaysCmd.Flags().Int64Var(&LoraNetworkSetsListGatewaysCmdLimit, "limit", 0, TRAPI("Maximum number of LoRa gateways to retrieve."))

	LoraNetworkSetsListGatewaysCmd.Flags().BoolVar(&LoraNetworkSetsListGatewaysCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsListGatewaysCmd)
}

// LoraNetworkSetsListGatewaysCmd defines 'list-gateways' subcommand
var LoraNetworkSetsListGatewaysCmd = &cobra.Command{
	Use:   "list-gateways",
	Short: TRAPI("/lora_network_sets/{ns_id}/gateways:get:summary"),
	Long:  TRAPI(`/lora_network_sets/{ns_id}/gateways:get:description`),
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

		param, err := collectLoraNetworkSetsListGatewaysCmdParams(ac)
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

func collectLoraNetworkSetsListGatewaysCmdParams(ac *apiClient) (*apiParams, error) {

	if LoraNetworkSetsListGatewaysCmdNsId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "ns-id")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraNetworkSetsListGatewaysCmd("/lora_network_sets/{ns_id}/gateways"),
		query:  buildQueryForLoraNetworkSetsListGatewaysCmd(),

		doPagination:                      LoraNetworkSetsListGatewaysCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraNetworkSetsListGatewaysCmd(path string) string {

	escapedNsId := url.PathEscape(LoraNetworkSetsListGatewaysCmdNsId)

	path = strReplace(path, "{"+"ns_id"+"}", escapedNsId, -1)

	return path
}

func buildQueryForLoraNetworkSetsListGatewaysCmd() url.Values {
	result := url.Values{}

	if LoraNetworkSetsListGatewaysCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", LoraNetworkSetsListGatewaysCmdLastEvaluatedKey)
	}

	if LoraNetworkSetsListGatewaysCmdLimit != 0 {
		result.Add("limit", sprintf("%d", LoraNetworkSetsListGatewaysCmdLimit))
	}

	return result
}
