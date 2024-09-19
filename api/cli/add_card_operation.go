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

// makeOperationAdminsAddCardCmd returns a command to handle operation addCard
func makeOperationAdminsAddCardCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "addCard",
		Short: `Adds an card to the system`,
		RunE:  runOperationAdminsAddCard,
	}

	if err := registerOperationAdminsAddCardParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAdminsAddCard uses cmd flags to call endpoint api
func runOperationAdminsAddCard(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := admins.NewAddCardParams()
	if err, _ = retrieveOperationAdminsAddCardCardFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAdminsAddCardResult(appCli.Admins.AddCard(params))
	if err != nil {
		return err
	}

	if !debug {
		fmt.Println(msgStr)
	}

	return nil
}

// registerOperationAdminsAddCardParamFlags registers all flags needed to fill params
func registerOperationAdminsAddCardParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAdminsAddCardCardParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAdminsAddCardCardParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var flagCardName string
	if cmdPrefix == "" {
		flagCardName = "card"
	} else {
		flagCardName = fmt.Sprintf("%v.card", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(flagCardName, "", `Optional json string for [card]. card to add`)

	// add flags for body
	if err := registerModelCardFlags(0, "card", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationAdminsAddCardCardFlag(m *admins.AddCardParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("card") {
		// Read card string from cmd and unmarshal
		flagCardValueStr, err := cmd.Flags().GetString("card")
		if err != nil {
			return err, false
		}

		flagCardValue := models.Card{}
		if err := json.Unmarshal([]byte(flagCardValueStr), &flagCardValue); err != nil {
			return fmt.Errorf("cannot unmarshal card string in models.Card: %v", err), false
		}
		m.Card = &flagCardValue
	}
	flagCardModel := m.Card
	if swag.IsZero(flagCardModel) {
		flagCardModel = &models.Card{}
	}
	err, added := retrieveModelCardFlags(0, flagCardModel, "card", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Card = flagCardModel
	}

	if dryRun && debug {
		flagCardValueDebugBytes, err := json.Marshal(m.Card)
		if err != nil {
			return err, false
		}
		logDebugf("Card dry-run payload: %v", string(flagCardValueDebugBytes))
	}

	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationAdminsAddCardResult parses request result and return the string content
func parseOperationAdminsAddCardResult(resp0 *admins.AddCardCreated, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning addCardCreated is not supported

		// Non schema case: warning addCardBadRequest is not supported

		// Non schema case: warning addCardConflict is not supported

		return "", respErr
	}

	// warning: non schema response addCardCreated is not supported by go-swagger cli yet.

	return "", nil
}