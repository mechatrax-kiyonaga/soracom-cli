// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersPermissionsDeleteCmdOperatorId holds value of 'operator_id' option
var UsersPermissionsDeleteCmdOperatorId string

// UsersPermissionsDeleteCmdUserName holds value of 'user_name' option
var UsersPermissionsDeleteCmdUserName string

func init() {
	UsersPermissionsDeleteCmd.Flags().StringVar(&UsersPermissionsDeleteCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersPermissionsDeleteCmd.Flags().StringVar(&UsersPermissionsDeleteCmdUserName, "user-name", "", TRAPI("SAM user name"))
	UsersPermissionsCmd.AddCommand(UsersPermissionsDeleteCmd)
}

// UsersPermissionsDeleteCmd defines 'delete' subcommand
var UsersPermissionsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/permission:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/permission:delete:description`),
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

		param, err := collectUsersPermissionsDeleteCmdParams(ac)
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

func collectUsersPermissionsDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if UsersPermissionsDeleteCmdOperatorId == "" {
		UsersPermissionsDeleteCmdOperatorId = ac.OperatorID
	}

	err = checkIfRequiredStringParameterIsSupplied("user_name", "user-name", "path", parsedBody, UsersPermissionsDeleteCmdUserName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForUsersPermissionsDeleteCmd("/operators/{operator_id}/users/{user_name}/permission"),
		query:  buildQueryForUsersPermissionsDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersPermissionsDeleteCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersPermissionsDeleteCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersPermissionsDeleteCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersPermissionsDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
