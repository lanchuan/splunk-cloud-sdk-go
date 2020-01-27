// Package search -- generated by scloudgen
// !! DO NOT EDIT !!
//
package search

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/pkg/search"
)

// createJob -- Creates a search job.
var createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "Creates a search job.",
	RunE:  impl.CreateJob,
}

// getJob -- Return the search job with the specified search ID (SID).
var getJobCmd = &cobra.Command{
	Use:   "get-job",
	Short: "Return the search job with the specified search ID (SID).",
	RunE:  impl.GetJob,
}

// listEventsSummary -- Return events summary, for search ID (SID) search.
var listEventsSummaryCmd = &cobra.Command{
	Use:   "list-events-summary",
	Short: "Return events summary, for search ID (SID) search.",
	RunE:  impl.ListEventsSummary,
}

// listFieldsSummary -- Return fields stats summary of the events to-date, for search ID (SID) search.
var listFieldsSummaryCmd = &cobra.Command{
	Use:   "list-fields-summary",
	Short: "Return fields stats summary of the events to-date, for search ID (SID) search.",
	RunE:  impl.ListFieldsSummary,
}

// listJobs -- Return the matching list of search jobs.
var listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "Return the matching list of search jobs.",
	RunE:  impl.ListJobs,
}

// listPreviewResults -- Return the preview search results for the job with the specified search ID (SID). Can be used when a job is running to return interim results.
var listPreviewResultsCmd = &cobra.Command{
	Use:   "list-preview-results",
	Short: "Return the preview search results for the job with the specified search ID (SID). Can be used when a job is running to return interim results.",
	RunE:  impl.ListPreviewResults,
}

// listResults -- Return the search results for the job with the specified search ID (SID).
var listResultsCmd = &cobra.Command{
	Use:   "list-results",
	Short: "Return the search results for the job with the specified search ID (SID).",
	RunE:  impl.ListResults,
}

// listTimeBuckets -- Return event distribution over time of the untransformed events read to-date, for search ID(SID) search.
var listTimeBucketsCmd = &cobra.Command{
	Use:   "list-time-buckets",
	Short: "Return event distribution over time of the untransformed events read to-date, for search ID(SID) search.",
	RunE:  impl.ListTimeBuckets,
}

// updateJob -- Update the search job with the specified search ID (SID) with an action.
var updateJobCmd = &cobra.Command{
	Use:   "update-job",
	Short: "Update the search job with the specified search ID (SID) with an action.",
	RunE:  impl.UpdateJob,
}

func init() {
	searchCmd.AddCommand(createJobCmd)

	var createJobQuery string
	createJobCmd.Flags().StringVar(&createJobQuery, "query", "", "This is a required parameter. The SPL search string.")
	createJobCmd.MarkFlagRequired("query")

	var createJobAllowSideEffects string
	createJobCmd.Flags().StringVar(&createJobAllowSideEffects, "allow-side-effects", "", "Specifies whether a search that contains commands with side effects (with possible security risks) is allowed to run. type: boolean default: false")

	var createJobCollectEventSummary string
	createJobCmd.Flags().StringVar(&createJobCollectEventSummary, "collect-event-summary", "false", "Specified whether a search is allowed to collect events summary during the run time.")

	var createJobCollectFieldSummary string
	createJobCmd.Flags().StringVar(&createJobCollectFieldSummary, "collect-field-summary", "false", "Specified whether a search is allowed to collect Fields summary during the run time.")

	var createJobCollectTimeBuckets string
	createJobCmd.Flags().StringVar(&createJobCollectTimeBuckets, "collect-time-buckets", "false", "Specified whether a search is allowed to collect Timeline Buckets summary during the run time.")

	var createJobEarliest string
	createJobCmd.Flags().StringVar(&createJobEarliest, "earliest", "", "The earliest time, in absolute or relative format, to retrieve events.  When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%!F(MISSING)T%!T(MISSING).%!Q(MISSING)) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.")

	var createJobEnablePreview string
	createJobCmd.Flags().StringVar(&createJobEnablePreview, "enable-preview", "false", "Specified whether a search is allowed to collect preview results during the run time.")

	var createJobExtractAllFields string
	createJobCmd.Flags().StringVar(&createJobExtractAllFields, "extract-all-fields", "false", "Specifies whether the Search service should extract all of the available fields in the data, including fields not mentioned in the SPL for the search job. Set to 'false' for better search peformance.")

	var createJobLatest string
	createJobCmd.Flags().StringVar(&createJobLatest, "latest", "", "The latest time, in absolute or relative format, to retrieve events.  When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%!F(MISSING)T%!T(MISSING).%!Q(MISSING)) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.")

	var createJobMaxTime float32
	createJobCmd.Flags().Float32Var(&createJobMaxTime, "max-time", 0.0, "The number of seconds to run the search before finalizing the search. The maximum value is 21600 seconds (6 hours).")

	var createJobMessages string
	createJobCmd.Flags().StringVar(&createJobMessages, "messages", "", "")

	var createJobModule string
	createJobCmd.Flags().StringVar(&createJobModule, "module", "", "The module to run the search in. The default module is used if a module is not specified.")

	var createJobRelativeTimeAnchor string
	createJobCmd.Flags().StringVar(&createJobRelativeTimeAnchor, "relative-time-anchor", "", "Relative values for the 'earliest' and 'latest' parameters snap to the unit that you specify.  For example, if 'earliest' is set to -d@d, the unit is day. If the 'relativeTimeAnchor' is is set to '1994-11-05T13:15:30Z'  then 'resolvedEarliest' is snapped to '1994-11-05T00:00:00Z', which is the day. Hours, minutes, and seconds are dropped.  If no 'relativeTimeAnchor' is specified, the default value is set to the time the search job was created.")

	var createJobStatus string
	createJobCmd.Flags().StringVar(&createJobStatus, "status", "", "The current status of the search job. The valid status values are 'running', 'done', 'canceled', and 'failed'. can accept values running, done, canceled, failed")

	var createJobTimezone string
	createJobCmd.Flags().StringVar(&createJobTimezone, "timezone", "", "The timezone that relative time specifiers are based off of. Timezone only applies to relative time literals  for 'earliest' and 'latest'. If UNIX time or UTC format is used for 'earliest' and 'latest', this field is ignored. For the list of supported timezone formats, see https://docs.splunk.com/Documentation/Splunk/latest/Data/Applytimezoneoffsetstotimestamps#zoneinfo_.28TZ.29_database type: string default: GMT")

	searchCmd.AddCommand(getJobCmd)

	var getJobSid string
	getJobCmd.Flags().StringVar(&getJobSid, "sid", "", "This is a required parameter. The search ID.")
	getJobCmd.MarkFlagRequired("sid")

	searchCmd.AddCommand(listEventsSummaryCmd)

	var listEventsSummarySid string
	listEventsSummaryCmd.Flags().StringVar(&listEventsSummarySid, "sid", "", "This is a required parameter. The search ID.")
	listEventsSummaryCmd.MarkFlagRequired("sid")

	var listEventsSummaryCount float32
	listEventsSummaryCmd.Flags().Float32Var(&listEventsSummaryCount, "count", 0.0, "The maximum number of entries to return. Set to 0 to return all available entries.")

	var listEventsSummaryEarliest string
	listEventsSummaryCmd.Flags().StringVar(&listEventsSummaryEarliest, "earliest", "", "The earliest time filter, in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%!F(MISSING)T%!T(MISSING).%!Q(MISSING)) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.")

	var listEventsSummaryField string
	listEventsSummaryCmd.Flags().StringVar(&listEventsSummaryField, "field", "", "A field to return for the result set. You can specify multiple fields of comma-separated values if multiple fields are required.")

	var listEventsSummaryLatest string
	listEventsSummaryCmd.Flags().StringVar(&listEventsSummaryLatest, "latest", "", "The latest time filter in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%!F(MISSING)T%!T(MISSING).%!Q(MISSING)) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.")

	var listEventsSummaryOffset float32
	listEventsSummaryCmd.Flags().Float32Var(&listEventsSummaryOffset, "offset", 0.0, "Index of first item to return.")

	searchCmd.AddCommand(listFieldsSummaryCmd)

	var listFieldsSummarySid string
	listFieldsSummaryCmd.Flags().StringVar(&listFieldsSummarySid, "sid", "", "This is a required parameter. The search ID.")
	listFieldsSummaryCmd.MarkFlagRequired("sid")

	var listFieldsSummaryEarliest string
	listFieldsSummaryCmd.Flags().StringVar(&listFieldsSummaryEarliest, "earliest", "", "The earliest time filter, in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%!F(MISSING)T%!T(MISSING).%!Q(MISSING)) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.")

	var listFieldsSummaryLatest string
	listFieldsSummaryCmd.Flags().StringVar(&listFieldsSummaryLatest, "latest", "", "The latest time filter in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%!F(MISSING)T%!T(MISSING).%!Q(MISSING)) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.")

	searchCmd.AddCommand(listJobsCmd)

	var listJobsCount float32
	listJobsCmd.Flags().Float32Var(&listJobsCount, "count", 0.0, "The maximum number of jobs that you want to return the status entries for.")

	var listJobsStatus string
	listJobsCmd.Flags().StringVar(&listJobsStatus, "status", "", "Filter the list of jobs by status. Valid status values are 'running', 'done', 'canceled', or 'failed'.")

	searchCmd.AddCommand(listPreviewResultsCmd)

	var listPreviewResultsSid string
	listPreviewResultsCmd.Flags().StringVar(&listPreviewResultsSid, "sid", "", "This is a required parameter. The search ID.")
	listPreviewResultsCmd.MarkFlagRequired("sid")

	var listPreviewResultsCount float32
	listPreviewResultsCmd.Flags().Float32Var(&listPreviewResultsCount, "count", 0.0, "The maximum number of entries to return. Set to 0 to return all available entries.")

	var listPreviewResultsOffset float32
	listPreviewResultsCmd.Flags().Float32Var(&listPreviewResultsOffset, "offset", 0.0, "Index of first item to return.")

	searchCmd.AddCommand(listResultsCmd)

	var listResultsSid string
	listResultsCmd.Flags().StringVar(&listResultsSid, "sid", "", "This is a required parameter. The search ID.")
	listResultsCmd.MarkFlagRequired("sid")

	var listResultsCount float32
	listResultsCmd.Flags().Float32Var(&listResultsCount, "count", 0.0, "The maximum number of entries to return. Set to 0 to return all available entries.")

	var listResultsField string
	listResultsCmd.Flags().StringVar(&listResultsField, "field", "", "A field to return for the result set. You can specify multiple fields of comma-separated values if multiple fields are required.")

	var listResultsOffset float32
	listResultsCmd.Flags().Float32Var(&listResultsOffset, "offset", 0.0, "Index of first item to return.")

	searchCmd.AddCommand(listTimeBucketsCmd)

	var listTimeBucketsSid string
	listTimeBucketsCmd.Flags().StringVar(&listTimeBucketsSid, "sid", "", "This is a required parameter. The search ID.")
	listTimeBucketsCmd.MarkFlagRequired("sid")

	searchCmd.AddCommand(updateJobCmd)

	var updateJobSid string
	updateJobCmd.Flags().StringVar(&updateJobSid, "sid", "", "This is a required parameter. The search ID.")
	updateJobCmd.MarkFlagRequired("sid")

	var updateJobStatus string
	updateJobCmd.Flags().StringVar(&updateJobStatus, "status", "", "This is a required parameter. The status to PATCH to an existing search job. The only status values you can PATCH are 'canceled' and 'finalized'. You can PATCH the 'canceled' status only to a search job that is running. can accept values canceled, finalized")
	updateJobCmd.MarkFlagRequired("status")

}
