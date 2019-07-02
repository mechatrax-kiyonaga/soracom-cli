// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// CredentialsDeleteCmdCredentialsId holds value of 'credentials_id' option
var CredentialsDeleteCmdCredentialsId string

func init() {
	CredentialsDeleteCmd.Flags().StringVar(&CredentialsDeleteCmdCredentialsId, "credentials-id", "", TRAPI("Credentials ID"))

	CredentialsDeleteCmd.MarkFlagRequired("credentials-id")

	CredentialsCmd.AddCommand(CredentialsDeleteCmd)
}

// CredentialsDeleteCmd defines 'delete' subcommand
var CredentialsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/credentials/{credentials_id}:delete:summary"),
	Long:  TRAPI(`/credentials/{credentials_id}:delete:description`),
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

		param, err := collectCredentialsDeleteCmdParams(ac)
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

func collectCredentialsDeleteCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForCredentialsDeleteCmd("/credentials/{credentials_id}"),
		query:  buildQueryForCredentialsDeleteCmd(),
	}, nil
}

func buildPathForCredentialsDeleteCmd(path string) string {

	escapedCredentialsId := url.PathEscape(CredentialsDeleteCmdCredentialsId)

	path = strings.Replace(path, "{"+"credentials_id"+"}", escapedCredentialsId, -1)

	return path
}

func buildQueryForCredentialsDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
