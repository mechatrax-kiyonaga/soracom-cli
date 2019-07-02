// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersPasswordConfiguredCmdOperatorId holds value of 'operator_id' option
var UsersPasswordConfiguredCmdOperatorId string

// UsersPasswordConfiguredCmdUserName holds value of 'user_name' option
var UsersPasswordConfiguredCmdUserName string

func init() {
	UsersPasswordConfiguredCmd.Flags().StringVar(&UsersPasswordConfiguredCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersPasswordConfiguredCmd.Flags().StringVar(&UsersPasswordConfiguredCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersPasswordConfiguredCmd.MarkFlagRequired("user-name")

	UsersPasswordCmd.AddCommand(UsersPasswordConfiguredCmd)
}

// UsersPasswordConfiguredCmd defines 'configured' subcommand
var UsersPasswordConfiguredCmd = &cobra.Command{
	Use:   "configured",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/password:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/password:get:description`),
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

		param, err := collectUsersPasswordConfiguredCmdParams(ac)
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

func collectUsersPasswordConfiguredCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersPasswordConfiguredCmdOperatorId == "" {
		UsersPasswordConfiguredCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersPasswordConfiguredCmd("/operators/{operator_id}/users/{user_name}/password"),
		query:  buildQueryForUsersPasswordConfiguredCmd(),
	}, nil
}

func buildPathForUsersPasswordConfiguredCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersPasswordConfiguredCmdOperatorId)

	path = strings.Replace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersPasswordConfiguredCmdUserName)

	path = strings.Replace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersPasswordConfiguredCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
