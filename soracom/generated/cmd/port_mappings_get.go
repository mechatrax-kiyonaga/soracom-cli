// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// PortMappingsGetCmdImsi holds value of 'imsi' option
var PortMappingsGetCmdImsi string

func init() {
	PortMappingsGetCmd.Flags().StringVar(&PortMappingsGetCmdImsi, "imsi", "", TRAPI("Target subscriber IMSI."))
	PortMappingsCmd.AddCommand(PortMappingsGetCmd)
}

// PortMappingsGetCmd defines 'get' subcommand
var PortMappingsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/port_mappings/subscribers/{imsi}:get:summary"),
	Long:  TRAPI(`/port_mappings/subscribers/{imsi}:get:description`),
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

		param, err := collectPortMappingsGetCmdParams(ac)
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

func collectPortMappingsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, PortMappingsGetCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForPortMappingsGetCmd("/port_mappings/subscribers/{imsi}"),
		query:  buildQueryForPortMappingsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForPortMappingsGetCmd(path string) string {

	escapedImsi := url.PathEscape(PortMappingsGetCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForPortMappingsGetCmd() url.Values {
	result := url.Values{}

	return result
}
