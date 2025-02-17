// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsListSubscriptionContainersCmdIccid holds value of 'iccid' option
var SimsListSubscriptionContainersCmdIccid string

// SimsListSubscriptionContainersCmdSimId holds value of 'sim_id' option
var SimsListSubscriptionContainersCmdSimId string

func init() {
	SimsListSubscriptionContainersCmd.Flags().StringVar(&SimsListSubscriptionContainersCmdIccid, "iccid", "", TRAPI("Iccid of the target profile"))

	SimsListSubscriptionContainersCmd.Flags().StringVar(&SimsListSubscriptionContainersCmdSimId, "sim-id", "", TRAPI("Sim Id of the target SIM."))
	SimsCmd.AddCommand(SimsListSubscriptionContainersCmd)
}

// SimsListSubscriptionContainersCmd defines 'list-subscription-containers' subcommand
var SimsListSubscriptionContainersCmd = &cobra.Command{
	Use:   "list-subscription-containers",
	Short: TRAPI("/sims/{sim_id}/profiles/{iccid}/subscription_containers:get:summary"),
	Long:  TRAPI(`/sims/{sim_id}/profiles/{iccid}/subscription_containers:get:description`) + "\n\n" + createLinkToAPIReference("Sim", "listSubscriptionContainers"),
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

		param, err := collectSimsListSubscriptionContainersCmdParams(ac)
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

func collectSimsListSubscriptionContainersCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("iccid", "iccid", "path", parsedBody, SimsListSubscriptionContainersCmdIccid)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsListSubscriptionContainersCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSimsListSubscriptionContainersCmd("/sims/{sim_id}/profiles/{iccid}/subscription_containers"),
		query:  buildQueryForSimsListSubscriptionContainersCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsListSubscriptionContainersCmd(path string) string {

	escapedIccid := url.PathEscape(SimsListSubscriptionContainersCmdIccid)

	path = strReplace(path, "{"+"iccid"+"}", escapedIccid, -1)

	escapedSimId := url.PathEscape(SimsListSubscriptionContainersCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsListSubscriptionContainersCmd() url.Values {
	result := url.Values{}

	return result
}
