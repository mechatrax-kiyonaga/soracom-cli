// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// EmailsIssueAddEmailTokenCmdEmail holds value of 'email' option
var EmailsIssueAddEmailTokenCmdEmail string

// EmailsIssueAddEmailTokenCmdPassword holds value of 'password' option
var EmailsIssueAddEmailTokenCmdPassword string

// EmailsIssueAddEmailTokenCmdBody holds contents of request body to be sent
var EmailsIssueAddEmailTokenCmdBody string

func init() {
	EmailsIssueAddEmailTokenCmd.Flags().StringVar(&EmailsIssueAddEmailTokenCmdEmail, "email", "", TRAPI("Email address to be added"))

	EmailsIssueAddEmailTokenCmd.Flags().StringVar(&EmailsIssueAddEmailTokenCmdPassword, "password", "", TRAPI("Password of the operator"))

	EmailsIssueAddEmailTokenCmd.Flags().StringVar(&EmailsIssueAddEmailTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	EmailsCmd.AddCommand(EmailsIssueAddEmailTokenCmd)
}

// EmailsIssueAddEmailTokenCmd defines 'issue-add-email-token' subcommand
var EmailsIssueAddEmailTokenCmd = &cobra.Command{
	Use:   "issue-add-email-token",
	Short: TRAPI("/operators/add_email_token/issue:post:summary"),
	Long:  TRAPI(`/operators/add_email_token/issue:post:description`) + "\n\n" + createLinkToAPIReference("Email", "issueAddEmailToken"),
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

		param, err := collectEmailsIssueAddEmailTokenCmdParams(ac)
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

func collectEmailsIssueAddEmailTokenCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForEmailsIssueAddEmailTokenCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredStringParameterIsSupplied("email", "email", "body", parsedBody, EmailsIssueAddEmailTokenCmdEmail)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("password", "password", "body", parsedBody, EmailsIssueAddEmailTokenCmdPassword)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForEmailsIssueAddEmailTokenCmd("/operators/add_email_token/issue"),
		query:       buildQueryForEmailsIssueAddEmailTokenCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForEmailsIssueAddEmailTokenCmd(path string) string {

	return path
}

func buildQueryForEmailsIssueAddEmailTokenCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForEmailsIssueAddEmailTokenCmd() (string, error) {
	var result map[string]interface{}

	if EmailsIssueAddEmailTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(EmailsIssueAddEmailTokenCmdBody, "@") {
			fname := strings.TrimPrefix(EmailsIssueAddEmailTokenCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if EmailsIssueAddEmailTokenCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(EmailsIssueAddEmailTokenCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if EmailsIssueAddEmailTokenCmdEmail != "" {
		result["email"] = EmailsIssueAddEmailTokenCmdEmail
	}

	if EmailsIssueAddEmailTokenCmdPassword != "" {
		result["password"] = EmailsIssueAddEmailTokenCmdPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
