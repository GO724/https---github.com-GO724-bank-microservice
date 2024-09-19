// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"bank-microservice/api/client/admins"

	"github.com/spf13/cobra"
)

// makeOperationAdminsDelPersonCmd returns a command to handle operation delPerson
func makeOperationAdminsDelPersonCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "delPerson",
		Short: `Remove an person from the system`,
		RunE:  runOperationAdminsDelPerson,
	}

	if err := registerOperationAdminsDelPersonParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAdminsDelPerson uses cmd flags to call endpoint api
func runOperationAdminsDelPerson(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := admins.NewDelPersonParams()
	if err, _ = retrieveOperationAdminsDelPersonInnFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAdminsDelPersonResult(appCli.Admins.DelPerson(params))
	if err != nil {
		return err
	}

	if !debug {
		fmt.Println(msgStr)
	}

	return nil
}

// registerOperationAdminsDelPersonParamFlags registers all flags needed to fill params
func registerOperationAdminsDelPersonParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAdminsDelPersonInnParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAdminsDelPersonInnParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	flagInnDescription := `Required. Person to remove`

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

func retrieveOperationAdminsDelPersonInnFlag(m *admins.DelPersonParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationAdminsDelPersonResult parses request result and return the string content
func parseOperationAdminsDelPersonResult(resp0 *admins.DelPersonNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning delPersonNoContent is not supported

		// Non schema case: warning delPersonBadRequest is not supported

		// Non schema case: warning delPersonConflict is not supported

		return "", respErr
	}

	// warning: non schema response delPersonNoContent is not supported by go-swagger cli yet.

	return "", nil
}