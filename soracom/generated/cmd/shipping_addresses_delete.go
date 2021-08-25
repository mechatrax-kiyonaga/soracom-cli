// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// ShippingAddressesDeleteCmdOperatorId holds value of 'operator_id' option
var ShippingAddressesDeleteCmdOperatorId string

// ShippingAddressesDeleteCmdShippingAddressId holds value of 'shipping_address_id' option
var ShippingAddressesDeleteCmdShippingAddressId string

func init() {
	ShippingAddressesDeleteCmd.Flags().StringVar(&ShippingAddressesDeleteCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	ShippingAddressesDeleteCmd.Flags().StringVar(&ShippingAddressesDeleteCmdShippingAddressId, "shipping-address-id", "", TRAPI("shipping_address_id"))
	ShippingAddressesCmd.AddCommand(ShippingAddressesDeleteCmd)
}

// ShippingAddressesDeleteCmd defines 'delete' subcommand
var ShippingAddressesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/operators/{operator_id}/shipping_addresses/{shipping_address_id}:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/shipping_addresses/{shipping_address_id}:delete:description`),
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

		param, err := collectShippingAddressesDeleteCmdParams(ac)
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

func collectShippingAddressesDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if ShippingAddressesDeleteCmdOperatorId == "" {
		ShippingAddressesDeleteCmdOperatorId = ac.OperatorID
	}

	err = checkIfRequiredStringParameterIsSupplied("shipping_address_id", "shipping-address-id", "path", parsedBody, ShippingAddressesDeleteCmdShippingAddressId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForShippingAddressesDeleteCmd("/operators/{operator_id}/shipping_addresses/{shipping_address_id}"),
		query:  buildQueryForShippingAddressesDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForShippingAddressesDeleteCmd(path string) string {

	escapedOperatorId := url.PathEscape(ShippingAddressesDeleteCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedShippingAddressId := url.PathEscape(ShippingAddressesDeleteCmdShippingAddressId)

	path = strReplace(path, "{"+"shipping_address_id"+"}", escapedShippingAddressId, -1)

	return path
}

func buildQueryForShippingAddressesDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
