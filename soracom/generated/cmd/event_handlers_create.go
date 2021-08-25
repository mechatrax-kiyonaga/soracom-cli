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

// EventHandlersCreateCmdDescription holds value of 'description' option
var EventHandlersCreateCmdDescription string

// EventHandlersCreateCmdName holds value of 'name' option
var EventHandlersCreateCmdName string

// EventHandlersCreateCmdStatus holds value of 'status' option
var EventHandlersCreateCmdStatus string

// EventHandlersCreateCmdTargetGroupId holds value of 'targetGroupId' option
var EventHandlersCreateCmdTargetGroupId string

// EventHandlersCreateCmdTargetImsi holds value of 'targetImsi' option
var EventHandlersCreateCmdTargetImsi string

// EventHandlersCreateCmdTargetOperatorId holds value of 'targetOperatorId' option
var EventHandlersCreateCmdTargetOperatorId string

// EventHandlersCreateCmdTargetSimId holds value of 'targetSimId' option
var EventHandlersCreateCmdTargetSimId string

// EventHandlersCreateCmdBody holds contents of request body to be sent
var EventHandlersCreateCmdBody string

func init() {
	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdDescription, "description", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdName, "name", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdStatus, "status", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdTargetGroupId, "target-group-id", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdTargetImsi, "target-imsi", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdTargetOperatorId, "target-operator-id", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdTargetSimId, "target-sim-id", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	EventHandlersCmd.AddCommand(EventHandlersCreateCmd)
}

// EventHandlersCreateCmd defines 'create' subcommand
var EventHandlersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/event_handlers:post:summary"),
	Long:  TRAPI(`/event_handlers:post:description`),
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

		param, err := collectEventHandlersCreateCmdParams(ac)
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

func collectEventHandlersCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForEventHandlersCreateCmd()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &parsedBody)
	if err != nil {
		return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
	}
	contentType := "application/json"

	err = checkIfRequiredStringParameterIsSupplied("status", "status", "body", parsedBody, EventHandlersCreateCmdStatus)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForEventHandlersCreateCmd("/event_handlers"),
		query:       buildQueryForEventHandlersCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForEventHandlersCreateCmd(path string) string {

	return path
}

func buildQueryForEventHandlersCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForEventHandlersCreateCmd() (string, error) {
	var result map[string]interface{}

	if EventHandlersCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(EventHandlersCreateCmdBody, "@") {
			fname := strings.TrimPrefix(EventHandlersCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if EventHandlersCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(EventHandlersCreateCmdBody)
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

	if EventHandlersCreateCmdDescription != "" {
		result["description"] = EventHandlersCreateCmdDescription
	}

	if EventHandlersCreateCmdName != "" {
		result["name"] = EventHandlersCreateCmdName
	}

	if EventHandlersCreateCmdStatus != "" {
		result["status"] = EventHandlersCreateCmdStatus
	}

	if EventHandlersCreateCmdTargetGroupId != "" {
		result["targetGroupId"] = EventHandlersCreateCmdTargetGroupId
	}

	if EventHandlersCreateCmdTargetImsi != "" {
		result["targetImsi"] = EventHandlersCreateCmdTargetImsi
	}

	if EventHandlersCreateCmdTargetOperatorId != "" {
		result["targetOperatorId"] = EventHandlersCreateCmdTargetOperatorId
	}

	if EventHandlersCreateCmdTargetSimId != "" {
		result["targetSimId"] = EventHandlersCreateCmdTargetSimId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
