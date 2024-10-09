// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"bank-microservice/api/client/admins"

	"github.com/spf13/cobra"
)

// makeOperationAdminsDelСardCmd returns a command to handle operation delСard
func makeOperationAdminsDelСardCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "delСard",
		Short: `Remove an card from the system`,
		RunE:  runOperationAdminsDelСard,
	}

	if err := registerOperationAdminsDelСardParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAdminsDelСard uses cmd flags to call endpoint api
func runOperationAdminsDelСard(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := admins.NewDelСardParams()
	if err, _ = retrieveOperationAdminsDelСardInnFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAdminsDelСardResult(appCli.Admins.DelСard(params))
	if err != nil {
		return err
	}

	if !debug {
		fmt.Println(msgStr)
	}

	return nil
}

// registerOperationAdminsDelСardParamFlags registers all flags needed to fill params
func registerOperationAdminsDelСardParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAdminsDelСardInnParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAdminsDelСardInnParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	flagInnDescription := `Required. card to remove`

	var flagInnName string
	if cmdPrefix == "" {
		flagInnName = "inn"
	} else {
		flagInnName = fmt.Sprintf("%v.inn", cmdPrefix)
	}

	var flagInnDefault int64

	_ = cmd.PersistentFlags().Int64(flagInnName, flagInnDefault, flagInnDescription)

	return nil
}

func retrieveOperationAdminsDelСardInnFlag(m *admins.DelСardParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("inn") {

		var flagInnName string
		if cmdPrefix == "" {
			flagInnName = "inn"
		} else {
			flagInnName = fmt.Sprintf("%v.inn", cmdPrefix)
		}

		flagInnValue, err := cmd.Flags().GetInt64(flagInnName)
		if err != nil {
			return err, false
		}
		m.Inn = flagInnValue

	}

	return nil, retAdded
}

// parseOperationAdminsDelСardResult parses request result and return the string content
func parseOperationAdminsDelСardResult(resp0 *admins.DelСardNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning delСardNoContent is not supported

		// Non schema case: warning delСardBadRequest is not supported

		// Non schema case: warning delСardConflict is not supported

		return "", respErr
	}

	// warning: non schema response delСardNoContent is not supported by go-swagger cli yet.

	return "", nil
}
