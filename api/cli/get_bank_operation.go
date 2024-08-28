// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"bank-microservice/api/client/admins"

	"github.com/spf13/cobra"
)

// makeOperationAdminsGetBankCmd returns a command to handle operation getBank
func makeOperationAdminsGetBankCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getBank",
		Short: `Gets an bank from the system`,
		RunE:  runOperationAdminsGetBank,
	}

	if err := registerOperationAdminsGetBankParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAdminsGetBank uses cmd flags to call endpoint api
func runOperationAdminsGetBank(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := admins.NewGetBankParams()
	if err, _ = retrieveOperationAdminsGetBankInnFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAdminsGetBankResult(appCli.Admins.GetBank(params))
	if err != nil {
		return err
	}

	if !debug {
		fmt.Println(msgStr)
	}

	return nil
}

// registerOperationAdminsGetBankParamFlags registers all flags needed to fill params
func registerOperationAdminsGetBankParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAdminsGetBankInnParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAdminsGetBankInnParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	flagInnDescription := `Required. get bank`

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

func retrieveOperationAdminsGetBankInnFlag(m *admins.GetBankParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationAdminsGetBankResult parses request result and return the string content
func parseOperationAdminsGetBankResult(resp0 *admins.GetBankCreated, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning getBankCreated is not supported

		// Non schema case: warning getBankBadRequest is not supported

		// Non schema case: warning getBankConflict is not supported

		return "", respErr
	}

	// warning: non schema response getBankCreated is not supported by go-swagger cli yet.

	return "", nil
}
