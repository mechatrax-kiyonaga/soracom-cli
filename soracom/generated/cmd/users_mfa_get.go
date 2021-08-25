// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersMfaGetCmdOperatorId holds value of 'operator_id' option
var UsersMfaGetCmdOperatorId string

// UsersMfaGetCmdUserName holds value of 'user_name' option
var UsersMfaGetCmdUserName string

func init() {
	UsersMfaGetCmd.Flags().StringVar(&UsersMfaGetCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersMfaGetCmd.Flags().StringVar(&UsersMfaGetCmdUserName, "user-name", "", TRAPI("SAM user name"))
	UsersMfaCmd.AddCommand(UsersMfaGetCmd)
}

// UsersMfaGetCmd defines 'get' subcommand
var UsersMfaGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/mfa:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/mfa:get:description`),
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

		param, err := collectUsersMfaGetCmdParams(ac)
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

func collectUsersMfaGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if UsersMfaGetCmdOperatorId == "" {
		UsersMfaGetCmdOperatorId = ac.OperatorID
	}

	err = checkIfRequiredStringParameterIsSupplied("user_name", "user-name", "path", parsedBody, UsersMfaGetCmdUserName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersMfaGetCmd("/operators/{operator_id}/users/{user_name}/mfa"),
		query:  buildQueryForUsersMfaGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersMfaGetCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersMfaGetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersMfaGetCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersMfaGetCmd() url.Values {
	result := url.Values{}

	return result
}
