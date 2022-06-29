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

// SimsSendSmsCmdPayload holds value of 'payload' option
var SimsSendSmsCmdPayload string

// SimsSendSmsCmdSimId holds value of 'sim_id' option
var SimsSendSmsCmdSimId string

// SimsSendSmsCmdEncodingType holds value of 'encodingType' option
var SimsSendSmsCmdEncodingType int64

// SimsSendSmsCmdBody holds contents of request body to be sent
var SimsSendSmsCmdBody string

func init() {
	SimsSendSmsCmd.Flags().StringVar(&SimsSendSmsCmdPayload, "payload", "", TRAPI(""))

	SimsSendSmsCmd.Flags().StringVar(&SimsSendSmsCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsSendSmsCmd.Flags().Int64Var(&SimsSendSmsCmdEncodingType, "encoding-type", 2, TRAPI("Encoding type of the message body. `1` indicates the body is `DCS_7BIT` that only supports single byte characters. `2` is `DCS_UCS2` that supports multi-byte text. When omitted, it is treated as `2` (`DCS_UCS2`)."))

	SimsSendSmsCmd.Flags().StringVar(&SimsSendSmsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SimsCmd.AddCommand(SimsSendSmsCmd)
}

// SimsSendSmsCmd defines 'send-sms' subcommand
var SimsSendSmsCmd = &cobra.Command{
	Use:   "send-sms",
	Short: TRAPI("/sims/{sim_id}/send_sms:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/send_sms:post:description`),
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

		param, err := collectSimsSendSmsCmdParams(ac)
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

func collectSimsSendSmsCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSimsSendSmsCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsSendSmsCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSimsSendSmsCmd("/sims/{sim_id}/send_sms"),
		query:       buildQueryForSimsSendSmsCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsSendSmsCmd(path string) string {

	escapedSimId := url.PathEscape(SimsSendSmsCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsSendSmsCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSimsSendSmsCmd() (string, error) {
	var result map[string]interface{}

	if SimsSendSmsCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SimsSendSmsCmdBody, "@") {
			fname := strings.TrimPrefix(SimsSendSmsCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SimsSendSmsCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SimsSendSmsCmdBody)
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

	if SimsSendSmsCmdPayload != "" {
		result["payload"] = SimsSendSmsCmdPayload
	}

	if SimsSendSmsCmdEncodingType != 2 {
		result["encodingType"] = SimsSendSmsCmdEncodingType
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
