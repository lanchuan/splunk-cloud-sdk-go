// Package action -- generated by scloudgen
// !! DO NOT EDIT !!
//
package action

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/auth"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/flags"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/jsonx"
	model "github.com/splunk/splunk-cloud-sdk-go/services/action"
)

// CreateAction Creates an action template.
func CreateAction(cmd *cobra.Command, args []string) error {

	return nil

}

// CreateActionEmailAction Creates an action template.
func CreateActionEmailAction(cmd *cobra.Command, args []string) error {


	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Parse all flags
	var addresses []string
	err = flags.ParseFlag(cmd.Flags(), "addresses", &addresses)
	if err != nil {
		return fmt.Errorf(`error parsing "addresses": ` + err.Error())
	}
	var bodyDefault string
	body := &bodyDefault
	err = flags.ParseFlag(cmd.Flags(), "body", &body)
	if err != nil {
		return fmt.Errorf(`error parsing "body": ` + err.Error())
	}
	var bodyPlainTextDefault string
	bodyPlainText := &bodyPlainTextDefault
	err = flags.ParseFlag(cmd.Flags(), "body-plain-text", &bodyPlainText)
	if err != nil {
		return fmt.Errorf(`error parsing "body-plain-text": ` + err.Error())
	}
	var fromNameDefault string
	fromName := &fromNameDefault
	err = flags.ParseFlag(cmd.Flags(), "from-name", &fromName)
	if err != nil {
		return fmt.Errorf(`error parsing "from-name": ` + err.Error())
	}
	var kind model.ActionKind
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var members []string
	err = flags.ParseFlag(cmd.Flags(), "members", &members)
	if err != nil {
		return fmt.Errorf(`error parsing "members": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var subjectDefault string
	subject := &subjectDefault
	err = flags.ParseFlag(cmd.Flags(), "subject", &subject)
	if err != nil {
		return fmt.Errorf(`error parsing "subject": ` + err.Error())
	}
	var titleDefault string
	title := &titleDefault
	err = flags.ParseFlag(cmd.Flags(), "title", &title)
	if err != nil {
		return fmt.Errorf(`error parsing "title": ` + err.Error())
	}

	// Form the request body
	generated_request_body := model.EmailAction{
		Addresses: addresses,
		Body: body,
		BodyPlainText: bodyPlainText,
		FromName: fromName,
		Kind: kind,
		Members: members,
		Name: name,
		Subject: subject,
		Title: title,
	}
	resp, err := client.ActionService.CreateAction(model.MakeActionFromEmailAction(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil

}

// CreateActionWebhookAction Creates an action template.
func CreateActionWebhookAction(cmd *cobra.Command, args []string) error {


	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Parse all flags
	var kind model.ActionKind
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var titleDefault string
	title := &titleDefault
	err = flags.ParseFlag(cmd.Flags(), "title", &title)
	if err != nil {
		return fmt.Errorf(`error parsing "title": ` + err.Error())
	}
	var webhookHeaders map[string][]string
	err = flags.ParseFlag(cmd.Flags(), "webhook-headers", &webhookHeaders)
	if err != nil {
		return fmt.Errorf(`error parsing "webhook-headers": ` + err.Error())
	}
	var webhookPayload string
	err = flags.ParseFlag(cmd.Flags(), "webhook-payload", &webhookPayload)
	if err != nil {
		return fmt.Errorf(`error parsing "webhook-payload": ` + err.Error())
	}
	var webhookUrl string
	err = flags.ParseFlag(cmd.Flags(), "webhook-url", &webhookUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "webhook-url": ` + err.Error())
	}

	// Form the request body
	generated_request_body := model.WebhookAction{
		Kind: kind,
		Name: name,
		Title: title,
		WebhookHeaders: webhookHeaders,
		WebhookPayload: webhookPayload,
		WebhookUrl: webhookUrl,
	}
	resp, err := client.ActionService.CreateAction(model.MakeActionFromWebhookAction(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil

}

// DeleteAction Removes an action template.
func DeleteAction(cmd *cobra.Command, args []string) error {


	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Parse all flags
	var action_name string
	err = flags.ParseFlag(cmd.Flags(), "action-name", &action_name)
	if err != nil {
		return fmt.Errorf(`error parsing "action-name": ` + err.Error())
	}

	err = client.ActionService.DeleteAction(action_name)
	if err != nil {
		return err
	}

	return nil

}

// GetAction Returns a specific action template.
func GetAction(cmd *cobra.Command, args []string) error {


	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Parse all flags
	var action_name string
	err = flags.ParseFlag(cmd.Flags(), "action-name", &action_name)
	if err != nil {
		return fmt.Errorf(`error parsing "action-name": ` + err.Error())
	}

	resp, err := client.ActionService.GetAction(action_name)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil

}

// GetActionStatus Returns the status of an action that was invoked. The status is available for 4 days after the last status change.
func GetActionStatus(cmd *cobra.Command, args []string) error {


	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Parse all flags
	var action_name string
	err = flags.ParseFlag(cmd.Flags(), "action-name", &action_name)
	if err != nil {
		return fmt.Errorf(`error parsing "action-name": ` + err.Error())
	}
	var status_id string
	err = flags.ParseFlag(cmd.Flags(), "status-id", &status_id)
	if err != nil {
		return fmt.Errorf(`error parsing "status-id": ` + err.Error())
	}

	resp, err := client.ActionService.GetActionStatus(action_name, status_id)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil

}

// GetActionStatusDetails Returns the status details of the invoked email action. The status is available for 4 days after the last status change.
func GetActionStatusDetails(cmd *cobra.Command, args []string) error {


	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Parse all flags
	var action_name string
	err = flags.ParseFlag(cmd.Flags(), "action-name", &action_name)
	if err != nil {
		return fmt.Errorf(`error parsing "action-name": ` + err.Error())
	}
	var status_id string
	err = flags.ParseFlag(cmd.Flags(), "status-id", &status_id)
	if err != nil {
		return fmt.Errorf(`error parsing "status-id": ` + err.Error())
	}

	resp, err := client.ActionService.GetActionStatusDetails(action_name, status_id)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil

}

// GetPublicWebhookKeys Returns an array of one or two webhook keys. The first key is active. The second key, if present, is expired.
func GetPublicWebhookKeys(cmd *cobra.Command, args []string) error {


	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}


	resp, err := client.ActionService.GetPublicWebhookKeys()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil

}

// ListActions Returns the list of action templates.
func ListActions(cmd *cobra.Command, args []string) error {


	client, err := auth.GetClient()
	if err != nil {
		return err
	}


	resp, err := client.ActionService.ListActions()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil

}

// TriggerAction Invokes an action.
func TriggerAction(cmd *cobra.Command, args []string) error {


	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Parse all flags
	var action_name string
	err = flags.ParseFlag(cmd.Flags(), "action-name", &action_name)
	if err != nil {
		return fmt.Errorf(`error parsing "action-name": ` + err.Error())
	}
	var addresses []string
	err = flags.ParseFlag(cmd.Flags(), "addresses", &addresses)
	if err != nil {
		return fmt.Errorf(`error parsing "addresses": ` + err.Error())
	}
	var kindDefault model.TriggerEventKind
	kind := &kindDefault
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var members []string
	err = flags.ParseFlag(cmd.Flags(), "members", &members)
	if err != nil {
		return fmt.Errorf(`error parsing "members": ` + err.Error())
	}
	var payloadDefault model.RawJSONPayload
	payload := &payloadDefault
	err = flags.ParseFlag(cmd.Flags(), "payload", &payload)
	if err != nil {
		return fmt.Errorf(`error parsing "payload": ` + err.Error())
	}
	var triggerConditionDefault string
	triggerCondition := &triggerConditionDefault
	err = flags.ParseFlag(cmd.Flags(), "trigger-condition", &triggerCondition)
	if err != nil {
		return fmt.Errorf(`error parsing "trigger-condition": ` + err.Error())
	}
	var triggerNameDefault string
	triggerName := &triggerNameDefault
	err = flags.ParseFlag(cmd.Flags(), "trigger-name", &triggerName)
	if err != nil {
		return fmt.Errorf(`error parsing "trigger-name": ` + err.Error())
	}
	var ttlSecondsDefault int32
	ttlSeconds := &ttlSecondsDefault
	err = flags.ParseFlag(cmd.Flags(), "ttl-seconds", &ttlSeconds)
	if err != nil {
		return fmt.Errorf(`error parsing "ttl-seconds": ` + err.Error())
	}

	// Form the request body
	generated_request_body := model.TriggerEvent{
		ActionMetadata: &model.TriggerEventActionMetadata{
			Addresses: addresses,
			Members: members,
		},
		Kind: kind,
		Payload: payload,
		TriggerCondition: triggerCondition,
		TriggerName: triggerName,
		TtlSeconds: ttlSeconds,
	}
	err = client.ActionService.TriggerAction(action_name, generated_request_body)
	if err != nil {
		return err
	}

	return nil

}

// UpdateAction Modifies an action template.
func UpdateAction(cmd *cobra.Command, args []string) error {

	return nil

}

// UpdateActionEmailActionMutable Modifies an action template.
func UpdateActionEmailActionMutable(cmd *cobra.Command, args []string) error {


	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Parse all flags
	var action_name string
	err = flags.ParseFlag(cmd.Flags(), "action-name", &action_name)
	if err != nil {
		return fmt.Errorf(`error parsing "action-name": ` + err.Error())
	}
	var addresses []string
	err = flags.ParseFlag(cmd.Flags(), "addresses", &addresses)
	if err != nil {
		return fmt.Errorf(`error parsing "addresses": ` + err.Error())
	}
	var bodyDefault string
	body := &bodyDefault
	err = flags.ParseFlag(cmd.Flags(), "body", &body)
	if err != nil {
		return fmt.Errorf(`error parsing "body": ` + err.Error())
	}
	var bodyPlainTextDefault string
	bodyPlainText := &bodyPlainTextDefault
	err = flags.ParseFlag(cmd.Flags(), "body-plain-text", &bodyPlainText)
	if err != nil {
		return fmt.Errorf(`error parsing "body-plain-text": ` + err.Error())
	}
	var fromNameDefault string
	fromName := &fromNameDefault
	err = flags.ParseFlag(cmd.Flags(), "from-name", &fromName)
	if err != nil {
		return fmt.Errorf(`error parsing "from-name": ` + err.Error())
	}
	var members []string
	err = flags.ParseFlag(cmd.Flags(), "members", &members)
	if err != nil {
		return fmt.Errorf(`error parsing "members": ` + err.Error())
	}
	var subjectDefault string
	subject := &subjectDefault
	err = flags.ParseFlag(cmd.Flags(), "subject", &subject)
	if err != nil {
		return fmt.Errorf(`error parsing "subject": ` + err.Error())
	}
	var titleDefault string
	title := &titleDefault
	err = flags.ParseFlag(cmd.Flags(), "title", &title)
	if err != nil {
		return fmt.Errorf(`error parsing "title": ` + err.Error())
	}

	// Form the request body
	generated_request_body := model.EmailActionMutable{
		Addresses: addresses,
		Body: body,
		BodyPlainText: bodyPlainText,
		FromName: fromName,
		Members: members,
		Subject: subject,
		Title: title,
	}
	resp, err := client.ActionService.UpdateAction(action_name,model.MakeActionMutableFromEmailActionMutable(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil

}

// UpdateActionWebhookActionMutable Modifies an action template.
func UpdateActionWebhookActionMutable(cmd *cobra.Command, args []string) error {


	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Parse all flags
	var action_name string
	err = flags.ParseFlag(cmd.Flags(), "action-name", &action_name)
	if err != nil {
		return fmt.Errorf(`error parsing "action-name": ` + err.Error())
	}
	var titleDefault string
	title := &titleDefault
	err = flags.ParseFlag(cmd.Flags(), "title", &title)
	if err != nil {
		return fmt.Errorf(`error parsing "title": ` + err.Error())
	}
	var webhookHeaders map[string][]string
	err = flags.ParseFlag(cmd.Flags(), "webhook-headers", &webhookHeaders)
	if err != nil {
		return fmt.Errorf(`error parsing "webhook-headers": ` + err.Error())
	}
	var webhookPayloadDefault string
	webhookPayload := &webhookPayloadDefault
	err = flags.ParseFlag(cmd.Flags(), "webhook-payload", &webhookPayload)
	if err != nil {
		return fmt.Errorf(`error parsing "webhook-payload": ` + err.Error())
	}
	var webhookUrlDefault string
	webhookUrl := &webhookUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "webhook-url", &webhookUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "webhook-url": ` + err.Error())
	}

	// Form the request body
	generated_request_body := model.WebhookActionMutable{
		Title: title,
		WebhookHeaders: webhookHeaders,
		WebhookPayload: webhookPayload,
		WebhookUrl: webhookUrl,
	}
	resp, err := client.ActionService.UpdateAction(action_name,model.MakeActionMutableFromWebhookActionMutable(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil

}

