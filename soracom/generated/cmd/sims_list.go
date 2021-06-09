// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SimsListCmdLastEvaluatedKey string

// SimsListCmdLimit holds value of 'limit' option
var SimsListCmdLimit int64

// SimsListCmdPaginate indicates to do pagination or not
var SimsListCmdPaginate bool

func init() {
	SimsListCmd.Flags().StringVar(&SimsListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The ID of the last SIM retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next SIM onward."))

	SimsListCmd.Flags().Int64Var(&SimsListCmdLimit, "limit", 0, TRAPI("Maximum number of SIMs to retrieve. Setting a limit does not guarantee the number of sims returned in the response (i.e. the response may contain fewer sims than the specified limit)."))

	SimsListCmd.Flags().BoolVar(&SimsListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	SimsCmd.AddCommand(SimsListCmd)
}

// SimsListCmd defines 'list' subcommand
var SimsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/sims:get:summary"),
	Long:  TRAPI(`/sims:get:description`),
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

		param, err := collectSimsListCmdParams(ac)
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

func collectSimsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSimsListCmd("/sims"),
		query:  buildQueryForSimsListCmd(),

		doPagination:                      SimsListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsListCmd(path string) string {

	return path
}

func buildQueryForSimsListCmd() url.Values {
	result := url.Values{}

	if SimsListCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SimsListCmdLastEvaluatedKey)
	}

	if SimsListCmdLimit != 0 {
		result.Add("limit", sprintf("%d", SimsListCmdLimit))
	}

	return result
}
