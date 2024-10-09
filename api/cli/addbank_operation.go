// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"bank-microservice/api/client/admins"
	"bank-microservice/api/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationAdminsAddbankCmd returns a command to handle operation addbank
func makeOperationAdminsAddbankCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "addbank",
		Short: `Adds an bank to the system`,
		RunE:  runOperationAdminsAddbank,
	}

	if err := registerOperationAdminsAddbankParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAdminsAddbank uses cmd flags to call endpoint api
func runOperationAdminsAddbank(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := admins.NewAddbankParams()
	if err, _ = retrieveOperationAdminsAddbankBankFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAdminsAddbankResult(appCli.Admins.Addbank(params))
	if err != nil {
		return err
	}

	if !debug {
		fmt.Println(msgStr)
	}

	return nil
}

// registerOperationAdminsAddbankParamFlags registers all flags needed to fill params
func registerOperationAdminsAddbankParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAdminsAddbankBankParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAdminsAddbankBankParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var flagBankName string
	if cmdPrefix == "" {
		flagBankName = "bank"
	} else {
		flagBankName = fmt.Sprintf("%v.bank", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(flagBankName, "", `Optional json string for [bank]. bank to add`)

	// add flags for body
	if err := registerModelBankFlags(0, "bank", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationAdminsAddbankBankFlag(m *admins.AddbankParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("bank") {
		// Read bank string from cmd and unmarshal
		flagBankValueStr, err := cmd.Flags().GetString("bank")
		if err != nil {
			return err, false
		}

		flagBankValue := models.Bank{}
		if err := json.Unmarshal([]byte(flagBankValueStr), &flagBankValue); err != nil {
			return fmt.Errorf("cannot unmarshal bank string in models.Bank: %v", err), false
		}
		m.Bank = &flagBankValue
	}
	flagBankModel := m.Bank
	if swag.IsZero(flagBankModel) {
		flagBankModel = &models.Bank{}
	}
	err, added := retrieveModelBankFlags(0, flagBankModel, "bank", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Bank = flagBankModel
	}

	if dryRun && debug {
		flagBankValueDebugBytes, err := json.Marshal(m.Bank)
		if err != nil {
			return err, false
		}
		logDebugf("Bank dry-run payload: %v", string(flagBankValueDebugBytes))
	}

	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationAdminsAddbankResult parses request result and return the string content
func parseOperationAdminsAddbankResult(resp0 *admins.AddbankCreated, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning addbankCreated is not supported

		// Non schema case: warning addbankBadRequest is not supported

		// Non schema case: warning addbankConflict is not supported

		return "", respErr
	}

	// warning: non schema response addbankCreated is not supported by go-swagger cli yet.

	return "", nil
}
