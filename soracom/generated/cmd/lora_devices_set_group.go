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

// LoraDevicesSetGroupCmdDeviceId holds value of 'device_id' option
var LoraDevicesSetGroupCmdDeviceId string

// LoraDevicesSetGroupCmdGroupId holds value of 'groupId' option
var LoraDevicesSetGroupCmdGroupId string

// LoraDevicesSetGroupCmdOperatorId holds value of 'operatorId' option
var LoraDevicesSetGroupCmdOperatorId string

// LoraDevicesSetGroupCmdCreatedTime holds value of 'createdTime' option
var LoraDevicesSetGroupCmdCreatedTime int64

// LoraDevicesSetGroupCmdLastModifiedTime holds value of 'lastModifiedTime' option
var LoraDevicesSetGroupCmdLastModifiedTime int64

// LoraDevicesSetGroupCmdBody holds contents of request body to be sent
var LoraDevicesSetGroupCmdBody string

func init() {
	LoraDevicesSetGroupCmd.Flags().StringVar(&LoraDevicesSetGroupCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRa device."))

	LoraDevicesSetGroupCmd.Flags().StringVar(&LoraDevicesSetGroupCmdGroupId, "group-id", "", TRAPI(""))

	LoraDevicesSetGroupCmd.Flags().StringVar(&LoraDevicesSetGroupCmdOperatorId, "operator-id", "", TRAPI(""))

	LoraDevicesSetGroupCmd.Flags().Int64Var(&LoraDevicesSetGroupCmdCreatedTime, "created-time", 0, TRAPI(""))

	LoraDevicesSetGroupCmd.Flags().Int64Var(&LoraDevicesSetGroupCmdLastModifiedTime, "last-modified-time", 0, TRAPI(""))

	LoraDevicesSetGroupCmd.Flags().StringVar(&LoraDevicesSetGroupCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LoraDevicesCmd.AddCommand(LoraDevicesSetGroupCmd)
}

// LoraDevicesSetGroupCmd defines 'set-group' subcommand
var LoraDevicesSetGroupCmd = &cobra.Command{
	Use:   "set-group",
	Short: TRAPI("/lora_devices/{device_id}/set_group:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/set_group:post:description`) + "\n\n" + createLinkToAPIReference("LoraDevice", "setLoraDeviceGroup"),
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

		param, err := collectLoraDevicesSetGroupCmdParams(ac)
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

func collectLoraDevicesSetGroupCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLoraDevicesSetGroupCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, LoraDevicesSetGroupCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraDevicesSetGroupCmd("/lora_devices/{device_id}/set_group"),
		query:       buildQueryForLoraDevicesSetGroupCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraDevicesSetGroupCmd(path string) string {

	escapedDeviceId := url.PathEscape(LoraDevicesSetGroupCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesSetGroupCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLoraDevicesSetGroupCmd() (string, error) {
	var result map[string]interface{}

	if LoraDevicesSetGroupCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LoraDevicesSetGroupCmdBody, "@") {
			fname := strings.TrimPrefix(LoraDevicesSetGroupCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LoraDevicesSetGroupCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraDevicesSetGroupCmdBody)
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

	if LoraDevicesSetGroupCmdGroupId != "" {
		result["groupId"] = LoraDevicesSetGroupCmdGroupId
	}

	if LoraDevicesSetGroupCmdOperatorId != "" {
		result["operatorId"] = LoraDevicesSetGroupCmdOperatorId
	}

	if LoraDevicesSetGroupCmdCreatedTime != 0 {
		result["createdTime"] = LoraDevicesSetGroupCmdCreatedTime
	}

	if LoraDevicesSetGroupCmdLastModifiedTime != 0 {
		result["lastModifiedTime"] = LoraDevicesSetGroupCmdLastModifiedTime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
