// Package collect -- generated by scloudgen
// !! DO NOT EDIT !!
//
package collect

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/flags"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	model "github.com/splunk/splunk-cloud-sdk-go/services/collect"
)

// CreateExecution Creates an execution for a scheduled job based on the job ID.
func CreateExecution(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var jobId string
	err = flags.ParseFlag(cmd.Flags(), "job-id", &jobId)
	if err != nil {
		return fmt.Errorf(`error parsing "job-id": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.CollectService.CreateExecution(jobId)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// CreateJob This API returns `403` if the number of collect workers is over a certain limit.
func CreateJob(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var connectorID string
	err = flags.ParseFlag(cmd.Flags(), "connector-ID", &connectorID)
	if err != nil {
		return fmt.Errorf(`error parsing "connector-ID": ` + err.Error())
	}
	var eventExtraFields []model.EventExtraField
	err = flags.ParseFlag(cmd.Flags(), "event-extra-fields", &eventExtraFields)
	if err != nil {
		return fmt.Errorf(`error parsing "event-extra-fields": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var parameters map[string]interface{}
	err = flags.ParseFlag(cmd.Flags(), "parameters", &parameters)
	if err != nil {
		return fmt.Errorf(`error parsing "parameters": ` + err.Error())
	}
	var scalePolicy model.ScalePolicy
	err = flags.ParseFlag(cmd.Flags(), "scale-policy", &scalePolicy)
	if err != nil {
		return fmt.Errorf(`error parsing "scale-policy": ` + err.Error())
	}
	var schedule string
	err = flags.ParseFlag(cmd.Flags(), "schedule", &schedule)
	if err != nil {
		return fmt.Errorf(`error parsing "schedule": ` + err.Error())
	}
	var scheduledDefault bool
	scheduled := &scheduledDefault
	err = flags.ParseFlag(cmd.Flags(), "scheduled", &scheduled)
	if err != nil {
		return fmt.Errorf(`error parsing "scheduled": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.Job{

		ConnectorID:      connectorID,
		EventExtraFields: eventExtraFields,
		Name:             name,
		Parameters:       parameters,
		ScalePolicy:      scalePolicy,
		Schedule:         schedule,
		Scheduled:        scheduled,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.CollectService.CreateJob(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// DeleteJob Removes a job based on the job ID.
func DeleteJob(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var jobId string
	err = flags.ParseFlag(cmd.Flags(), "job-id", &jobId)
	if err != nil {
		return fmt.Errorf(`error parsing "job-id": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.CollectService.DeleteJob(jobId)
	if err != nil {
		return err
	}

	return nil
}

// DeleteJobs Removes all jobs on a tenant.
func DeleteJobs(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.CollectService.DeleteJobs()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetExecution Returns the execution details based on the execution ID and job ID.
func GetExecution(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var executionUid string
	err = flags.ParseFlag(cmd.Flags(), "execution-uid", &executionUid)
	if err != nil {
		return fmt.Errorf(`error parsing "execution-uid": ` + err.Error())
	}
	var jobId string
	err = flags.ParseFlag(cmd.Flags(), "job-id", &jobId)
	if err != nil {
		return fmt.Errorf(`error parsing "job-id": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.CollectService.GetExecution(jobId, executionUid)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetJob Returns a job based on the job ID.
func GetJob(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var jobId string
	err = flags.ParseFlag(cmd.Flags(), "job-id", &jobId)
	if err != nil {
		return fmt.Errorf(`error parsing "job-id": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.CollectService.GetJob(jobId)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListJobs Returns a list of all jobs that belong to a tenant.
func ListJobs(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var connectorId string
	err = flags.ParseFlag(cmd.Flags(), "connector-id", &connectorId)
	if err != nil {
		return fmt.Errorf(`error parsing "connector-id": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListJobsQueryParams{}
	generated_query.ConnectorId = connectorId

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.CollectService.ListJobs(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// PatchExecution Modifies an execution based on the job ID.
func PatchExecution(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var executionUid string
	err = flags.ParseFlag(cmd.Flags(), "execution-uid", &executionUid)
	if err != nil {
		return fmt.Errorf(`error parsing "execution-uid": ` + err.Error())
	}
	var jobId string
	err = flags.ParseFlag(cmd.Flags(), "job-id", &jobId)
	if err != nil {
		return fmt.Errorf(`error parsing "job-id": ` + err.Error())
	}
	var status model.ExecutionPatchStatus
	err = flags.ParseFlag(cmd.Flags(), "status", &status)
	if err != nil {
		return fmt.Errorf(`error parsing "status": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.ExecutionPatch{

		Status: status,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.CollectService.PatchExecution(jobId, executionUid, generated_request_body)
	if err != nil {
		return err
	}

	return nil
}

// PatchJob This API returns `403` if the number of collect workers is over a certain limit.
func PatchJob(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var connectorIDDefault string
	connectorID := &connectorIDDefault
	err = flags.ParseFlag(cmd.Flags(), "connector-ID", &connectorID)
	if err != nil {
		return fmt.Errorf(`error parsing "connector-ID": ` + err.Error())
	}
	var eventExtraFields []model.EventExtraField
	err = flags.ParseFlag(cmd.Flags(), "event-extra-fields", &eventExtraFields)
	if err != nil {
		return fmt.Errorf(`error parsing "event-extra-fields": ` + err.Error())
	}
	var jobId string
	err = flags.ParseFlag(cmd.Flags(), "job-id", &jobId)
	if err != nil {
		return fmt.Errorf(`error parsing "job-id": ` + err.Error())
	}
	var nameDefault string
	name := &nameDefault
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var parameters map[string]interface{}
	err = flags.ParseFlag(cmd.Flags(), "parameters", &parameters)
	if err != nil {
		return fmt.Errorf(`error parsing "parameters": ` + err.Error())
	}
	var scalePolicyDefault model.ScalePolicy
	scalePolicy := &scalePolicyDefault
	err = flags.ParseFlag(cmd.Flags(), "scale-policy", &scalePolicy)
	if err != nil {
		return fmt.Errorf(`error parsing "scale-policy": ` + err.Error())
	}
	var scheduleDefault string
	schedule := &scheduleDefault
	err = flags.ParseFlag(cmd.Flags(), "schedule", &schedule)
	if err != nil {
		return fmt.Errorf(`error parsing "schedule": ` + err.Error())
	}
	var scheduledDefault bool
	scheduled := &scheduledDefault
	err = flags.ParseFlag(cmd.Flags(), "scheduled", &scheduled)
	if err != nil {
		return fmt.Errorf(`error parsing "scheduled": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.JobPatch{

		ConnectorID:      connectorID,
		EventExtraFields: eventExtraFields,
		Name:             name,
		Parameters:       parameters,
		ScalePolicy:      scalePolicy,
		Schedule:         schedule,
		Scheduled:        scheduled,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.CollectService.PatchJob(jobId, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// PatchJobs This is a non-atomic operation and the results are returned as a list with each job patch result as its element. This API returns `200 OK` regardless of how many jobs were successfully patched. You must read the response body to find out if all jobs are patched. When the API is called, the `jobIDs` or `connectorID` must be specified. Do not specify more than one of them at the same time. This API returns `403` if the number of collect workers is over a certain limit.
func PatchJobs(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var connectorIDDefault string
	connectorID := &connectorIDDefault
	err = flags.ParseFlag(cmd.Flags(), "connector-ID", &connectorID)
	if err != nil {
		return fmt.Errorf(`error parsing "connector-ID": ` + err.Error())
	}
	var connectorId string
	err = flags.ParseFlag(cmd.Flags(), "connector-id", &connectorId)
	if err != nil {
		return fmt.Errorf(`error parsing "connector-id": ` + err.Error())
	}
	var eventExtraFields []model.EventExtraField
	err = flags.ParseFlag(cmd.Flags(), "event-extra-fields", &eventExtraFields)
	if err != nil {
		return fmt.Errorf(`error parsing "event-extra-fields": ` + err.Error())
	}
	var jobIDs []string
	err = flags.ParseFlag(cmd.Flags(), "job-ids", &jobIDs)
	if err != nil {
		return fmt.Errorf(`error parsing "job-ids": ` + err.Error())
	}
	var scalePolicyDefault model.ScalePolicy
	scalePolicy := &scalePolicyDefault
	err = flags.ParseFlag(cmd.Flags(), "scale-policy", &scalePolicy)
	if err != nil {
		return fmt.Errorf(`error parsing "scale-policy": ` + err.Error())
	}
	// Form query params
	generated_query := model.PatchJobsQueryParams{}
	generated_query.ConnectorId = connectorId
	generated_query.JobIDs = jobIDs

	// Form the request body
	generated_request_body := model.JobsPatch{

		ConnectorID:      connectorID,
		EventExtraFields: eventExtraFields,
		ScalePolicy:      scalePolicy,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.CollectService.PatchJobs(generated_request_body, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}
