// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OrdersConfirmCmdOrderId holds value of 'order_id' option
var OrdersConfirmCmdOrderId string

func init() {
	OrdersConfirmCmd.Flags().StringVar(&OrdersConfirmCmdOrderId, "order-id", "", TRAPI("order_id"))
	OrdersCmd.AddCommand(OrdersConfirmCmd)
}

// OrdersConfirmCmd defines 'confirm' subcommand
var OrdersConfirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: TRAPI("/orders/{order_id}/confirm:put:summary"),
	Long:  TRAPI(`/orders/{order_id}/confirm:put:description`),
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

		param, err := collectOrdersConfirmCmdParams(ac)
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

func collectOrdersConfirmCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("order_id", "order-id", "path", parsedBody, OrdersConfirmCmdOrderId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "PUT",
		path:   buildPathForOrdersConfirmCmd("/orders/{order_id}/confirm"),
		query:  buildQueryForOrdersConfirmCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOrdersConfirmCmd(path string) string {

	escapedOrderId := url.PathEscape(OrdersConfirmCmdOrderId)

	path = strReplace(path, "{"+"order_id"+"}", escapedOrderId, -1)

	return path
}

func buildQueryForOrdersConfirmCmd() url.Values {
	result := url.Values{}

	return result
}
