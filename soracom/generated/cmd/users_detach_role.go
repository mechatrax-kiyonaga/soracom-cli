// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersDetachRoleCmdOperatorId holds value of 'operator_id' option
var UsersDetachRoleCmdOperatorId string

// UsersDetachRoleCmdRoleId holds value of 'role_id' option
var UsersDetachRoleCmdRoleId string

// UsersDetachRoleCmdUserName holds value of 'user_name' option
var UsersDetachRoleCmdUserName string

func init() {
	UsersDetachRoleCmd.Flags().StringVar(&UsersDetachRoleCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersDetachRoleCmd.Flags().StringVar(&UsersDetachRoleCmdRoleId, "role-id", "", TRAPI("role_id"))

	UsersDetachRoleCmd.MarkFlagRequired("role-id")

	UsersDetachRoleCmd.Flags().StringVar(&UsersDetachRoleCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersDetachRoleCmd.MarkFlagRequired("user-name")

	UsersCmd.AddCommand(UsersDetachRoleCmd)
}

// UsersDetachRoleCmd defines 'detach-role' subcommand
var UsersDetachRoleCmd = &cobra.Command{
	Use:   "detach-role",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/roles/{role_id}:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/roles/{role_id}:delete:description`),
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

		param, err := collectUsersDetachRoleCmdParams(ac)
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

func collectUsersDetachRoleCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersDetachRoleCmdOperatorId == "" {
		UsersDetachRoleCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForUsersDetachRoleCmd("/operators/{operator_id}/users/{user_name}/roles/{role_id}"),
		query:  buildQueryForUsersDetachRoleCmd(),
	}, nil
}

func buildPathForUsersDetachRoleCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersDetachRoleCmdOperatorId)

	path = strings.Replace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedRoleId := url.PathEscape(UsersDetachRoleCmdRoleId)

	path = strings.Replace(path, "{"+"role_id"+"}", escapedRoleId, -1)

	escapedUserName := url.PathEscape(UsersDetachRoleCmdUserName)

	path = strings.Replace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersDetachRoleCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
