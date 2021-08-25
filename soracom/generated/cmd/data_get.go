// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DataGetCmdImsi holds value of 'imsi' option
var DataGetCmdImsi string

// DataGetCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var DataGetCmdLastEvaluatedKey string

// DataGetCmdSort holds value of 'sort' option
var DataGetCmdSort string

// DataGetCmdFrom holds value of 'from' option
var DataGetCmdFrom int64

// DataGetCmdLimit holds value of 'limit' option
var DataGetCmdLimit int64

// DataGetCmdTo holds value of 'to' option
var DataGetCmdTo int64

// DataGetCmdPaginate indicates to do pagination or not
var DataGetCmdPaginate bool

func init() {
	DataGetCmd.Flags().StringVar(&DataGetCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber that generated data entries."))

	DataGetCmd.Flags().StringVar(&DataGetCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of `time` in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	DataGetCmd.Flags().StringVar(&DataGetCmdSort, "sort", "desc", TRAPI("Sort order of the data entries. Either descending (latest data entry first) or ascending (oldest data entry first)."))

	DataGetCmd.Flags().Int64Var(&DataGetCmdFrom, "from", 0, TRAPI("Start time for the data entries search range (unixtime in milliseconds)."))

	DataGetCmd.Flags().Int64Var(&DataGetCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve."))

	DataGetCmd.Flags().Int64Var(&DataGetCmdTo, "to", 0, TRAPI("End time for the data entries search range (unixtime in milliseconds)."))

	DataGetCmd.Flags().BoolVar(&DataGetCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	DataCmd.AddCommand(DataGetCmd)
}

// DataGetCmd defines 'get' subcommand
var DataGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/subscribers/{imsi}/data:get:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/data:get:description`),
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

		param, err := collectDataGetCmdParams(ac)
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

func collectDataGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, DataGetCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForDataGetCmd("/subscribers/{imsi}/data"),
		query:  buildQueryForDataGetCmd(),

		doPagination:                      DataGetCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDataGetCmd(path string) string {

	escapedImsi := url.PathEscape(DataGetCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForDataGetCmd() url.Values {
	result := url.Values{}

	if DataGetCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", DataGetCmdLastEvaluatedKey)
	}

	if DataGetCmdSort != "desc" {
		result.Add("sort", DataGetCmdSort)
	}

	if DataGetCmdFrom != 0 {
		result.Add("from", sprintf("%d", DataGetCmdFrom))
	}

	if DataGetCmdLimit != 0 {
		result.Add("limit", sprintf("%d", DataGetCmdLimit))
	}

	if DataGetCmdTo != 0 {
		result.Add("to", sprintf("%d", DataGetCmdTo))
	}

	return result
}
