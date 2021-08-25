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

// LoraDevicesSendDataCmdData holds value of 'data' option
var LoraDevicesSendDataCmdData string

// LoraDevicesSendDataCmdDeviceId holds value of 'device_id' option
var LoraDevicesSendDataCmdDeviceId string

// LoraDevicesSendDataCmdFPort holds value of 'fPort' option
var LoraDevicesSendDataCmdFPort int64

// LoraDevicesSendDataCmdBody holds contents of request body to be sent
var LoraDevicesSendDataCmdBody string

func init() {
	LoraDevicesSendDataCmd.Flags().StringVar(&LoraDevicesSendDataCmdData, "data", "", TRAPI(""))

	LoraDevicesSendDataCmd.Flags().StringVar(&LoraDevicesSendDataCmdDeviceId, "device-id", "", TRAPI("ID of the recipient device."))

	LoraDevicesSendDataCmd.Flags().Int64Var(&LoraDevicesSendDataCmdFPort, "f-port", 2, TRAPI(""))

	LoraDevicesSendDataCmd.Flags().StringVar(&LoraDevicesSendDataCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LoraDevicesCmd.AddCommand(LoraDevicesSendDataCmd)
}

// LoraDevicesSendDataCmd defines 'send-data' subcommand
var LoraDevicesSendDataCmd = &cobra.Command{
	Use:   "send-data",
	Short: TRAPI("/lora_devices/{device_id}/data:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/data:post:description`),
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

		param, err := collectLoraDevicesSendDataCmdParams(ac)
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

func collectLoraDevicesSendDataCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLoraDevicesSendDataCmd()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &parsedBody)
	if err != nil {
		return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
	}
	contentType := "application/json"

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, LoraDevicesSendDataCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraDevicesSendDataCmd("/lora_devices/{device_id}/data"),
		query:       buildQueryForLoraDevicesSendDataCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraDevicesSendDataCmd(path string) string {

	escapedDeviceId := url.PathEscape(LoraDevicesSendDataCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesSendDataCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLoraDevicesSendDataCmd() (string, error) {
	var result map[string]interface{}

	if LoraDevicesSendDataCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LoraDevicesSendDataCmdBody, "@") {
			fname := strings.TrimPrefix(LoraDevicesSendDataCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LoraDevicesSendDataCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraDevicesSendDataCmdBody)
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

	if LoraDevicesSendDataCmdData != "" {
		result["data"] = LoraDevicesSendDataCmdData
	}

	if LoraDevicesSendDataCmdFPort != 2 {
		result["fPort"] = LoraDevicesSendDataCmdFPort
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
