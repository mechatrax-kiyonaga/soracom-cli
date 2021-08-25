// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// GroupsDeleteTagCmdGroupId holds value of 'group_id' option
var GroupsDeleteTagCmdGroupId string

// GroupsDeleteTagCmdTagName holds value of 'tag_name' option
var GroupsDeleteTagCmdTagName string

func init() {
	GroupsDeleteTagCmd.Flags().StringVar(&GroupsDeleteTagCmdGroupId, "group-id", "", TRAPI("Target group ID."))

	GroupsDeleteTagCmd.Flags().StringVar(&GroupsDeleteTagCmdTagName, "tag-name", "", TRAPI("Tag name to be deleted. (This will be part of a URL path, so it needs to be percent-encoded. In JavaScript, specify the name after it has been encoded using encodeURIComponent().)"))
	GroupsCmd.AddCommand(GroupsDeleteTagCmd)
}

// GroupsDeleteTagCmd defines 'delete-tag' subcommand
var GroupsDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TRAPI("/groups/{group_id}/tags/{tag_name}:delete:summary"),
	Long:  TRAPI(`/groups/{group_id}/tags/{tag_name}:delete:description`),
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

		param, err := collectGroupsDeleteTagCmdParams(ac)
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

func collectGroupsDeleteTagCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("group_id", "group-id", "path", parsedBody, GroupsDeleteTagCmdGroupId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("tag_name", "tag-name", "path", parsedBody, GroupsDeleteTagCmdTagName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForGroupsDeleteTagCmd("/groups/{group_id}/tags/{tag_name}"),
		query:  buildQueryForGroupsDeleteTagCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGroupsDeleteTagCmd(path string) string {

	escapedGroupId := url.PathEscape(GroupsDeleteTagCmdGroupId)

	path = strReplace(path, "{"+"group_id"+"}", escapedGroupId, -1)

	escapedTagName := url.PathEscape(GroupsDeleteTagCmdTagName)

	path = strReplace(path, "{"+"tag_name"+"}", escapedTagName, -1)

	return path
}

func buildQueryForGroupsDeleteTagCmd() url.Values {
	result := url.Values{}

	return result
}
