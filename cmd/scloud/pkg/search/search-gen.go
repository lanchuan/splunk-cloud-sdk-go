// Package search -- generated by scloudgen
// !! DO NOT EDIT !!
//
package search

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/flags"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	model "github.com/splunk/splunk-cloud-sdk-go/services/search"
)

// CreateJob Creates a search job.
func CreateJob(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var allowSideEffectsDefault bool
	allowSideEffects := &allowSideEffectsDefault
	err = flags.ParseFlag(cmd.Flags(), "allow-side-effects", &allowSideEffects)
	if err != nil {
		return fmt.Errorf(`error parsing "allow-side-effects": ` + err.Error())
	}
	var collectEventSummaryDefault bool
	collectEventSummary := &collectEventSummaryDefault
	err = flags.ParseFlag(cmd.Flags(), "collect-event-summary", &collectEventSummary)
	if err != nil {
		return fmt.Errorf(`error parsing "collect-event-summary": ` + err.Error())
	}
	var collectFieldSummaryDefault bool
	collectFieldSummary := &collectFieldSummaryDefault
	err = flags.ParseFlag(cmd.Flags(), "collect-field-summary", &collectFieldSummary)
	if err != nil {
		return fmt.Errorf(`error parsing "collect-field-summary": ` + err.Error())
	}
	var collectTimeBucketsDefault bool
	collectTimeBuckets := &collectTimeBucketsDefault
	err = flags.ParseFlag(cmd.Flags(), "collect-time-buckets", &collectTimeBuckets)
	if err != nil {
		return fmt.Errorf(`error parsing "collect-time-buckets": ` + err.Error())
	}
	var earliestDefault string
	earliest := &earliestDefault
	err = flags.ParseFlag(cmd.Flags(), "earliest", &earliest)
	if err != nil {
		return fmt.Errorf(`error parsing "earliest": ` + err.Error())
	}
	var enablePreviewDefault bool
	enablePreview := &enablePreviewDefault
	err = flags.ParseFlag(cmd.Flags(), "enable-preview", &enablePreview)
	if err != nil {
		return fmt.Errorf(`error parsing "enable-preview": ` + err.Error())
	}
	var extractAllFieldsDefault bool
	extractAllFields := &extractAllFieldsDefault
	err = flags.ParseFlag(cmd.Flags(), "extract-all-fields", &extractAllFields)
	if err != nil {
		return fmt.Errorf(`error parsing "extract-all-fields": ` + err.Error())
	}
	var latestDefault string
	latest := &latestDefault
	err = flags.ParseFlag(cmd.Flags(), "latest", &latest)
	if err != nil {
		return fmt.Errorf(`error parsing "latest": ` + err.Error())
	}
	var maxTimeDefault float32
	maxTime := &maxTimeDefault
	err = flags.ParseFlag(cmd.Flags(), "max-time", &maxTime)
	if err != nil {
		return fmt.Errorf(`error parsing "max-time": ` + err.Error())
	}
	var messages []model.Message
	err = flags.ParseFlag(cmd.Flags(), "messages", &messages)
	if err != nil {
		return fmt.Errorf(`error parsing "messages": ` + err.Error())
	}
	var moduleDefault string
	module := &moduleDefault
	err = flags.ParseFlag(cmd.Flags(), "module", &module)
	if err != nil {
		return fmt.Errorf(`error parsing "module": ` + err.Error())
	}
	var query string
	err = flags.ParseFlag(cmd.Flags(), "query", &query)
	if err != nil {
		return fmt.Errorf(`error parsing "query": ` + err.Error())
	}
	var relativeTimeAnchorDefault string
	relativeTimeAnchor := &relativeTimeAnchorDefault
	err = flags.ParseFlag(cmd.Flags(), "relative-time-anchor", &relativeTimeAnchor)
	if err != nil {
		return fmt.Errorf(`error parsing "relative-time-anchor": ` + err.Error())
	}
	var requiredFreshnessDefault float32
	requiredFreshness := &requiredFreshnessDefault
	err = flags.ParseFlag(cmd.Flags(), "required-freshness", &requiredFreshness)
	if err != nil {
		return fmt.Errorf(`error parsing "required-freshness": ` + err.Error())
	}
	var statusDefault model.SearchStatus
	status := &statusDefault
	err = flags.ParseFlag(cmd.Flags(), "status", &status)
	if err != nil {
		return fmt.Errorf(`error parsing "status": ` + err.Error())
	}
	var timezone interface{}
	err = flags.ParseFlag(cmd.Flags(), "timezone", &timezone)
	if err != nil {
		return fmt.Errorf(`error parsing "timezone": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.SearchJob{

		AllowSideEffects:    allowSideEffects,
		CollectEventSummary: collectEventSummary,
		CollectFieldSummary: collectFieldSummary,
		CollectTimeBuckets:  collectTimeBuckets,
		EnablePreview:       enablePreview,
		ExtractAllFields:    extractAllFields,
		MaxTime:             maxTime,
		Messages:            messages,
		Module:              module,
		Query:               query,
		QueryParameters: &model.QueryParameters{
			Earliest:           earliest,
			Latest:             latest,
			RelativeTimeAnchor: relativeTimeAnchor,
			Timezone:           timezone,
		},
		RequiredFreshness: requiredFreshness,
		Status:            status,
	}

	resp, err := client.SearchService.CreateJob(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetJob Return the search job with the specified search ID (SID).
func GetJob(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}

	resp, err := client.SearchService.GetJob(sid)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListEventsSummary Return events summary, for search ID (SID) search.
func ListEventsSummary(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var countDefault float32
	count := &countDefault
	err = flags.ParseFlag(cmd.Flags(), "count", &count)
	if err != nil {
		return fmt.Errorf(`error parsing "count": ` + err.Error())
	}
	var earliest string
	err = flags.ParseFlag(cmd.Flags(), "earliest", &earliest)
	if err != nil {
		return fmt.Errorf(`error parsing "earliest": ` + err.Error())
	}
	var field string
	err = flags.ParseFlag(cmd.Flags(), "field", &field)
	if err != nil {
		return fmt.Errorf(`error parsing "field": ` + err.Error())
	}
	var latest string
	err = flags.ParseFlag(cmd.Flags(), "latest", &latest)
	if err != nil {
		return fmt.Errorf(`error parsing "latest": ` + err.Error())
	}
	var offsetDefault float32
	offset := &offsetDefault
	err = flags.ParseFlag(cmd.Flags(), "offset", &offset)
	if err != nil {
		return fmt.Errorf(`error parsing "offset": ` + err.Error())
	}
	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListEventsSummaryQueryParams{}
	generated_query.Count = count
	generated_query.Earliest = earliest
	generated_query.Field = field
	generated_query.Latest = latest
	generated_query.Offset = offset

	resp, err := client.SearchService.ListEventsSummary(sid, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListFieldsSummary Return fields stats summary of the events to-date, for search ID (SID) search.
func ListFieldsSummary(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var earliest string
	err = flags.ParseFlag(cmd.Flags(), "earliest", &earliest)
	if err != nil {
		return fmt.Errorf(`error parsing "earliest": ` + err.Error())
	}
	var latest string
	err = flags.ParseFlag(cmd.Flags(), "latest", &latest)
	if err != nil {
		return fmt.Errorf(`error parsing "latest": ` + err.Error())
	}
	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListFieldsSummaryQueryParams{}
	generated_query.Earliest = earliest
	generated_query.Latest = latest

	resp, err := client.SearchService.ListFieldsSummary(sid, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListJobs Return the matching list of search jobs.
func ListJobs(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var countDefault float32
	count := &countDefault
	err = flags.ParseFlag(cmd.Flags(), "count", &count)
	if err != nil {
		return fmt.Errorf(`error parsing "count": ` + err.Error())
	}
	var filter string
	err = flags.ParseFlag(cmd.Flags(), "filter", &filter)
	if err != nil {
		return fmt.Errorf(`error parsing "filter": ` + err.Error())
	}
	var status model.SearchStatus
	err = flags.ParseFlag(cmd.Flags(), "status", &status)
	if err != nil {
		return fmt.Errorf(`error parsing "status": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListJobsQueryParams{}
	generated_query.Count = count
	generated_query.Filter = filter
	generated_query.Status = status

	resp, err := client.SearchService.ListJobs(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListPreviewResults Return the preview search results for the job with the specified search ID (SID). Can be used when a job is running to return interim results.
func ListPreviewResults(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var countDefault float32
	count := &countDefault
	err = flags.ParseFlag(cmd.Flags(), "count", &count)
	if err != nil {
		return fmt.Errorf(`error parsing "count": ` + err.Error())
	}
	var offsetDefault float32
	offset := &offsetDefault
	err = flags.ParseFlag(cmd.Flags(), "offset", &offset)
	if err != nil {
		return fmt.Errorf(`error parsing "offset": ` + err.Error())
	}
	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListPreviewResultsQueryParams{}
	generated_query.Count = count
	generated_query.Offset = offset

	resp, err := client.SearchService.ListPreviewResults(sid, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListResults Return the search results for the job with the specified search ID (SID).
func ListResults(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var countDefault float32
	count := &countDefault
	err = flags.ParseFlag(cmd.Flags(), "count", &count)
	if err != nil {
		return fmt.Errorf(`error parsing "count": ` + err.Error())
	}
	var field string
	err = flags.ParseFlag(cmd.Flags(), "field", &field)
	if err != nil {
		return fmt.Errorf(`error parsing "field": ` + err.Error())
	}
	var offsetDefault float32
	offset := &offsetDefault
	err = flags.ParseFlag(cmd.Flags(), "offset", &offset)
	if err != nil {
		return fmt.Errorf(`error parsing "offset": ` + err.Error())
	}
	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListResultsQueryParams{}
	generated_query.Count = count
	generated_query.Field = field
	generated_query.Offset = offset

	resp, err := client.SearchService.ListResults(sid, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListTimeBuckets Return event distribution over time of the untransformed events read to-date, for search ID(SID) search.
func ListTimeBuckets(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}

	resp, err := client.SearchService.ListTimeBuckets(sid)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// UpdateJob Update the search job with the specified search ID (SID) with an action.
func UpdateJob(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}
	var status model.UpdateJobStatus
	err = flags.ParseFlag(cmd.Flags(), "status", &status)
	if err != nil {
		return fmt.Errorf(`error parsing "status": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.UpdateJob{

		Status: status,
	}

	resp, err := client.SearchService.UpdateJob(sid, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}
