// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// PortMappingsDeleteCmdIpAddress holds value of 'ip_address' option
var PortMappingsDeleteCmdIpAddress string

// PortMappingsDeleteCmdPort holds value of 'port' option
var PortMappingsDeleteCmdPort string

func init() {
	PortMappingsDeleteCmd.Flags().StringVar(&PortMappingsDeleteCmdIpAddress, "ip-address", "", TRAPI("IP address of the target port mapping entry"))

	PortMappingsDeleteCmd.Flags().StringVar(&PortMappingsDeleteCmdPort, "port", "", TRAPI("Port of the target port mapping entry"))
	PortMappingsCmd.AddCommand(PortMappingsDeleteCmd)
}

// PortMappingsDeleteCmd defines 'delete' subcommand
var PortMappingsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/port_mappings/{ip_address}/{port}:delete:summary"),
	Long:  TRAPI(`/port_mappings/{ip_address}/{port}:delete:description`) + "\n\n" + createLinkToAPIReference("PortMapping", "deletePortMapping"),
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

		param, err := collectPortMappingsDeleteCmdParams(ac)
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

func collectPortMappingsDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("ip_address", "ip-address", "path", parsedBody, PortMappingsDeleteCmdIpAddress)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("port", "port", "path", parsedBody, PortMappingsDeleteCmdPort)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForPortMappingsDeleteCmd("/port_mappings/{ip_address}/{port}"),
		query:  buildQueryForPortMappingsDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForPortMappingsDeleteCmd(path string) string {

	escapedIpAddress := url.PathEscape(PortMappingsDeleteCmdIpAddress)

	path = strReplace(path, "{"+"ip_address"+"}", escapedIpAddress, -1)

	escapedPort := url.PathEscape(PortMappingsDeleteCmdPort)

	path = strReplace(path, "{"+"port"+"}", escapedPort, -1)

	return path
}

func buildQueryForPortMappingsDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
