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

// SubscribersVerifyTransferTokenCmdToken holds value of 'token' option
var SubscribersVerifyTransferTokenCmdToken string

// SubscribersVerifyTransferTokenCmdBody holds contents of request body to be sent
var SubscribersVerifyTransferTokenCmdBody string

func init() {
	SubscribersVerifyTransferTokenCmd.Flags().StringVar(&SubscribersVerifyTransferTokenCmdToken, "token", "", TRAPI(""))

	SubscribersVerifyTransferTokenCmd.Flags().StringVar(&SubscribersVerifyTransferTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SubscribersCmd.AddCommand(SubscribersVerifyTransferTokenCmd)
}

// SubscribersVerifyTransferTokenCmd defines 'verify-transfer-token' subcommand
var SubscribersVerifyTransferTokenCmd = &cobra.Command{
	Use:   "verify-transfer-token",
	Short: TRAPI("/subscribers/transfer_token/verify:post:summary"),
	Long:  TRAPI(`/subscribers/transfer_token/verify:post:description`),
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

		param, err := collectSubscribersVerifyTransferTokenCmdParams(ac)
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

func collectSubscribersVerifyTransferTokenCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSubscribersVerifyTransferTokenCmd()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &parsedBody)
	if err != nil {
		return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
	}
	contentType := "application/json"

	err = checkIfRequiredStringParameterIsSupplied("token", "token", "body", parsedBody, SubscribersVerifyTransferTokenCmdToken)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersVerifyTransferTokenCmd("/subscribers/transfer_token/verify"),
		query:       buildQueryForSubscribersVerifyTransferTokenCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersVerifyTransferTokenCmd(path string) string {

	return path
}

func buildQueryForSubscribersVerifyTransferTokenCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSubscribersVerifyTransferTokenCmd() (string, error) {
	var result map[string]interface{}

	if SubscribersVerifyTransferTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SubscribersVerifyTransferTokenCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersVerifyTransferTokenCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SubscribersVerifyTransferTokenCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersVerifyTransferTokenCmdBody)
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

	if SubscribersVerifyTransferTokenCmdToken != "" {
		result["token"] = SubscribersVerifyTransferTokenCmdToken
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
