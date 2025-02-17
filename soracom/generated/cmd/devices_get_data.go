// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DevicesGetDataCmdDeviceId holds value of 'device_id' option
var DevicesGetDataCmdDeviceId string

// DevicesGetDataCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var DevicesGetDataCmdLastEvaluatedKey string

// DevicesGetDataCmdSort holds value of 'sort' option
var DevicesGetDataCmdSort string

// DevicesGetDataCmdFrom holds value of 'from' option
var DevicesGetDataCmdFrom int64

// DevicesGetDataCmdLimit holds value of 'limit' option
var DevicesGetDataCmdLimit int64

// DevicesGetDataCmdTo holds value of 'to' option
var DevicesGetDataCmdTo int64

// DevicesGetDataCmdPaginate indicates to do pagination or not
var DevicesGetDataCmdPaginate bool

// DevicesGetDataCmdOutputJSONL indicates to output with jsonl format
var DevicesGetDataCmdOutputJSONL bool

func init() {
	DevicesGetDataCmd.Flags().StringVar(&DevicesGetDataCmdDeviceId, "device-id", "", TRAPI("Device ID of the target subscriber that generated data entries."))

	DevicesGetDataCmd.Flags().StringVar(&DevicesGetDataCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of `time` in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	DevicesGetDataCmd.Flags().StringVar(&DevicesGetDataCmdSort, "sort", "desc", TRAPI("Sort order of the data entries. Either descending (latest data entry first) or ascending (oldest data entry first)."))

	DevicesGetDataCmd.Flags().Int64Var(&DevicesGetDataCmdFrom, "from", 0, TRAPI("Start time for the data entries search range (unixtime in milliseconds)."))

	DevicesGetDataCmd.Flags().Int64Var(&DevicesGetDataCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve."))

	DevicesGetDataCmd.Flags().Int64Var(&DevicesGetDataCmdTo, "to", 0, TRAPI("End time for the data entries search range (unixtime in milliseconds)."))

	DevicesGetDataCmd.Flags().BoolVar(&DevicesGetDataCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	DevicesGetDataCmd.Flags().BoolVar(&DevicesGetDataCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	DevicesCmd.AddCommand(DevicesGetDataCmd)
}

// DevicesGetDataCmd defines 'get-data' subcommand
var DevicesGetDataCmd = &cobra.Command{
	Use:   "get-data",
	Short: TRAPI("/devices/{device_id}/data:get:summary"),
	Long:  TRAPI(`/devices/{device_id}/data:get:description`) + "\n\n" + createLinkToAPIReference("Device", "getDataFromDevice"),
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

		param, err := collectDevicesGetDataCmdParams(ac)
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
			if DevicesGetDataCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectDevicesGetDataCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, DevicesGetDataCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesGetDataCmd("/devices/{device_id}/data"),
		query:  buildQueryForDevicesGetDataCmd(),

		doPagination:                      DevicesGetDataCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesGetDataCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesGetDataCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForDevicesGetDataCmd() url.Values {
	result := url.Values{}

	if DevicesGetDataCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", DevicesGetDataCmdLastEvaluatedKey)
	}

	if DevicesGetDataCmdSort != "desc" {
		result.Add("sort", DevicesGetDataCmdSort)
	}

	if DevicesGetDataCmdFrom != 0 {
		result.Add("from", sprintf("%d", DevicesGetDataCmdFrom))
	}

	if DevicesGetDataCmdLimit != 0 {
		result.Add("limit", sprintf("%d", DevicesGetDataCmdLimit))
	}

	if DevicesGetDataCmdTo != 0 {
		result.Add("to", sprintf("%d", DevicesGetDataCmdTo))
	}

	return result
}
