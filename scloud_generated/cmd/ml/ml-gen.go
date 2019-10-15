// Package ml -- generated by scloudgen
// !! DO NOT EDIT !!
//
package ml

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/pkg/ml"
)

// createWorkflow -- Creates a workflow configuration.
var createWorkflowCmd = &cobra.Command{
	Use:   "create-workflow",
	Short: "Creates a workflow configuration.",
	RunE:  impl.CreateWorkflow,
}

// createWorkflowBuild -- Creates a workflow build.
var createWorkflowBuildCmd = &cobra.Command{
	Use:   "create-workflow-build",
	Short: "Creates a workflow build.",
	RunE:  impl.CreateWorkflowBuild,
}

// createWorkflowDeployment -- Creates a workflow deployment.
var createWorkflowDeploymentCmd = &cobra.Command{
	Use:   "create-workflow-deployment",
	Short: "Creates a workflow deployment.",
	RunE:  impl.CreateWorkflowDeployment,
}

// createWorkflowInference -- Creates a workflow inference request.
var createWorkflowInferenceCmd = &cobra.Command{
	Use:   "create-workflow-inference",
	Short: "Creates a workflow inference request.",
	RunE:  impl.CreateWorkflowInference,
}

// createWorkflowRun -- Creates a workflow run.
var createWorkflowRunCmd = &cobra.Command{
	Use:   "create-workflow-run",
	Short: "Creates a workflow run.",
	RunE:  impl.CreateWorkflowRun,
}

// createWorkflowStreamDeployment -- Creates a workflow streaming deployment.
var createWorkflowStreamDeploymentCmd = &cobra.Command{
	Use:   "create-workflow-stream-deployment",
	Short: "Creates a workflow streaming deployment.",
	RunE:  impl.CreateWorkflowStreamDeployment,
}

// deleteWorkflow -- Removes a workflow configuration.
var deleteWorkflowCmd = &cobra.Command{
	Use:   "delete-workflow",
	Short: "Removes a workflow configuration.",
	RunE:  impl.DeleteWorkflow,
}

// deleteWorkflowBuild -- Removes a workflow build.
var deleteWorkflowBuildCmd = &cobra.Command{
	Use:   "delete-workflow-build",
	Short: "Removes a workflow build.",
	RunE:  impl.DeleteWorkflowBuild,
}

// deleteWorkflowDeployment -- Removes a workflow deployment.
var deleteWorkflowDeploymentCmd = &cobra.Command{
	Use:   "delete-workflow-deployment",
	Short: "Removes a workflow deployment.",
	RunE:  impl.DeleteWorkflowDeployment,
}

// deleteWorkflowRun -- Removes a workflow run.
var deleteWorkflowRunCmd = &cobra.Command{
	Use:   "delete-workflow-run",
	Short: "Removes a workflow run.",
	RunE:  impl.DeleteWorkflowRun,
}

// deleteWorkflowStreamDeployment -- Removes a workflow streaming deployment.
var deleteWorkflowStreamDeploymentCmd = &cobra.Command{
	Use:   "delete-workflow-stream-deployment",
	Short: "Removes a workflow streaming deployment.",
	RunE:  impl.DeleteWorkflowStreamDeployment,
}

// getWorkflow -- Returns a workflow configuration.
var getWorkflowCmd = &cobra.Command{
	Use:   "get-workflow",
	Short: "Returns a workflow configuration.",
	RunE:  impl.GetWorkflow,
}

// getWorkflowBuild -- Returns the status of a workflow build.
var getWorkflowBuildCmd = &cobra.Command{
	Use:   "get-workflow-build",
	Short: "Returns the status of a workflow build.",
	RunE:  impl.GetWorkflowBuild,
}

// getWorkflowBuildError -- Returns a list of workflow errors.
var getWorkflowBuildErrorCmd = &cobra.Command{
	Use:   "get-workflow-build-error",
	Short: "Returns a list of workflow errors.",
	RunE:  impl.GetWorkflowBuildError,
}

// getWorkflowBuildLog -- Returns the logs from a workflow build.
var getWorkflowBuildLogCmd = &cobra.Command{
	Use:   "get-workflow-build-log",
	Short: "Returns the logs from a workflow build.",
	RunE:  impl.GetWorkflowBuildLog,
}

// getWorkflowDeployment -- Returns the status of a workflow deployment.
var getWorkflowDeploymentCmd = &cobra.Command{
	Use:   "get-workflow-deployment",
	Short: "Returns the status of a workflow deployment.",
	RunE:  impl.GetWorkflowDeployment,
}

// getWorkflowDeploymentError -- Returns a list of workflow deployment errors.
var getWorkflowDeploymentErrorCmd = &cobra.Command{
	Use:   "get-workflow-deployment-error",
	Short: "Returns a list of workflow deployment errors.",
	RunE:  impl.GetWorkflowDeploymentError,
}

// getWorkflowDeploymentLog -- Returns the logs from a workflow deployment.
var getWorkflowDeploymentLogCmd = &cobra.Command{
	Use:   "get-workflow-deployment-log",
	Short: "Returns the logs from a workflow deployment.",
	RunE:  impl.GetWorkflowDeploymentLog,
}

// getWorkflowRun -- Returns the status of a workflow run.
var getWorkflowRunCmd = &cobra.Command{
	Use:   "get-workflow-run",
	Short: "Returns the status of a workflow run.",
	RunE:  impl.GetWorkflowRun,
}

// getWorkflowRunError -- Returns the errors for a workflow run.
var getWorkflowRunErrorCmd = &cobra.Command{
	Use:   "get-workflow-run-error",
	Short: "Returns the errors for a workflow run.",
	RunE:  impl.GetWorkflowRunError,
}

// getWorkflowRunLog -- Returns the logs for a workflow run.
var getWorkflowRunLogCmd = &cobra.Command{
	Use:   "get-workflow-run-log",
	Short: "Returns the logs for a workflow run.",
	RunE:  impl.GetWorkflowRunLog,
}

// getWorkflowStreamDeployment -- Returns the status of a workflow streaming deployment.
var getWorkflowStreamDeploymentCmd = &cobra.Command{
	Use:   "get-workflow-stream-deployment",
	Short: "Returns the status of a workflow streaming deployment.",
	RunE:  impl.GetWorkflowStreamDeployment,
}

// listWorkflowBuilds -- Returns a list of workflow builds.
var listWorkflowBuildsCmd = &cobra.Command{
	Use:   "list-workflow-builds",
	Short: "Returns a list of workflow builds.",
	RunE:  impl.ListWorkflowBuilds,
}

// listWorkflowDeployments -- Returns a list of workflow deployments.
var listWorkflowDeploymentsCmd = &cobra.Command{
	Use:   "list-workflow-deployments",
	Short: "Returns a list of workflow deployments.",
	RunE:  impl.ListWorkflowDeployments,
}

// listWorkflowRuns -- Returns a list of workflow runs.
var listWorkflowRunsCmd = &cobra.Command{
	Use:   "list-workflow-runs",
	Short: "Returns a list of workflow runs.",
	RunE:  impl.ListWorkflowRuns,
}

// listWorkflows -- Returns a list of workflow configurations.
var listWorkflowsCmd = &cobra.Command{
	Use:   "list-workflows",
	Short: "Returns a list of workflow configurations.",
	RunE:  impl.ListWorkflows,
}

func init() {
	mlCmd.AddCommand(createWorkflowCmd)

	var createWorkflowTasks string
	createWorkflowCmd.Flags().StringVar(&createWorkflowTasks, "tasks", "", "This is a required parameter.")
	createWorkflowCmd.MarkFlagRequired("tasks")

	var createWorkflowCreationTime string
	createWorkflowCmd.Flags().StringVar(&createWorkflowCreationTime, "creation-time", "", "")
	var createWorkflowId string
	createWorkflowCmd.Flags().StringVar(&createWorkflowId, "id", "", "")
	var createWorkflowName string
	createWorkflowCmd.Flags().StringVar(&createWorkflowName, "name", "", "")

	mlCmd.AddCommand(createWorkflowBuildCmd)

	var createWorkflowBuildId string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildId, "id", "", "This is a required parameter.The workflow ID.")
	createWorkflowBuildCmd.MarkFlagRequired("id")
	var createWorkflowBuildInput string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildInput, "input", "", "This is a required parameter.")
	createWorkflowBuildCmd.MarkFlagRequired("input")

	var createWorkflowBuildCreationTime string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildCreationTime, "creation-time", "", "")
	var createWorkflowBuildEndTime string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildEndTime, "end-time", "", "")
	var createWorkflowBuildName string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildName, "name", "", "")
	var createWorkflowBuildOutput string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildOutput, "output", "", "")
	var createWorkflowBuildPipelineSummary string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildPipelineSummary, "pipeline-summary", "", "")
	var createWorkflowBuildStartTime string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildStartTime, "start-time", "", "")
	var createWorkflowBuildStatus string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildStatus, "status", "", "json string representing ")
	var createWorkflowBuildTimeoutSecs int32
	createWorkflowBuildCmd.Flags().Int32Var(&createWorkflowBuildTimeoutSecs, "timeout-secs", 0, "Number of seconds before a workflow build times out.")
	var createWorkflowBuildTrainingScore string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildTrainingScore, "training-score", "", "")
	var createWorkflowBuildValidationOption string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildValidationOption, "validation-option", "", "Represents which type of validation to use in the workflow along with any parameters if specified. If this is not included, no validation is done (all data is used for training). Default parameter values are used if no `option` is specified.")
	var createWorkflowBuildValidationScore string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildValidationScore, "validation-score", "", "The validation score whose type is specified by the user in `validationOption`.")
	var createWorkflowBuildWorkflow string
	createWorkflowBuildCmd.Flags().StringVar(&createWorkflowBuildWorkflow, "workflow", "", "")

	mlCmd.AddCommand(createWorkflowDeploymentCmd)

	var createWorkflowDeploymentBuildId string
	createWorkflowDeploymentCmd.Flags().StringVar(&createWorkflowDeploymentBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	createWorkflowDeploymentCmd.MarkFlagRequired("build-id")
	var createWorkflowDeploymentId string
	createWorkflowDeploymentCmd.Flags().StringVar(&createWorkflowDeploymentId, "id", "", "This is a required parameter.The workflow ID.")
	createWorkflowDeploymentCmd.MarkFlagRequired("id")
	var createWorkflowDeploymentSpec string
	createWorkflowDeploymentCmd.Flags().StringVar(&createWorkflowDeploymentSpec, "spec", "", "This is a required parameter.")
	createWorkflowDeploymentCmd.MarkFlagRequired("spec")

	var createWorkflowDeploymentCreationTime string
	createWorkflowDeploymentCmd.Flags().StringVar(&createWorkflowDeploymentCreationTime, "creation-time", "", "")
	var createWorkflowDeploymentEndTime string
	createWorkflowDeploymentCmd.Flags().StringVar(&createWorkflowDeploymentEndTime, "end-time", "", "")
	var createWorkflowDeploymentName string
	createWorkflowDeploymentCmd.Flags().StringVar(&createWorkflowDeploymentName, "name", "", "")
	var createWorkflowDeploymentStartTime string
	createWorkflowDeploymentCmd.Flags().StringVar(&createWorkflowDeploymentStartTime, "start-time", "", "")
	var createWorkflowDeploymentStatus string
	createWorkflowDeploymentCmd.Flags().StringVar(&createWorkflowDeploymentStatus, "status", "", "json string representing ")
	var createWorkflowDeploymentWorkflowBuild string
	createWorkflowDeploymentCmd.Flags().StringVar(&createWorkflowDeploymentWorkflowBuild, "workflow-build", "", "")

	mlCmd.AddCommand(createWorkflowInferenceCmd)

	var createWorkflowInferenceBuildId string
	createWorkflowInferenceCmd.Flags().StringVar(&createWorkflowInferenceBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	createWorkflowInferenceCmd.MarkFlagRequired("build-id")
	var createWorkflowInferenceDeploymentId string
	createWorkflowInferenceCmd.Flags().StringVar(&createWorkflowInferenceDeploymentId, "deployment-id", "", "This is a required parameter.The workflow deployment ID.")
	createWorkflowInferenceCmd.MarkFlagRequired("deployment-id")
	var createWorkflowInferenceId string
	createWorkflowInferenceCmd.Flags().StringVar(&createWorkflowInferenceId, "id", "", "This is a required parameter.The workflow ID.")
	createWorkflowInferenceCmd.MarkFlagRequired("id")
	var createWorkflowInferenceInput string
	createWorkflowInferenceCmd.Flags().StringVar(&createWorkflowInferenceInput, "input", "", "This is a required parameter.")
	createWorkflowInferenceCmd.MarkFlagRequired("input")

	var createWorkflowInferenceOutput string
	createWorkflowInferenceCmd.Flags().StringVar(&createWorkflowInferenceOutput, "output", "", "")

	mlCmd.AddCommand(createWorkflowRunCmd)

	var createWorkflowRunBuildId string
	createWorkflowRunCmd.Flags().StringVar(&createWorkflowRunBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	createWorkflowRunCmd.MarkFlagRequired("build-id")
	var createWorkflowRunId string
	createWorkflowRunCmd.Flags().StringVar(&createWorkflowRunId, "id", "", "This is a required parameter.The workflow ID.")
	createWorkflowRunCmd.MarkFlagRequired("id")
	var createWorkflowRunInput string
	createWorkflowRunCmd.Flags().StringVar(&createWorkflowRunInput, "input", "", "This is a required parameter.")
	createWorkflowRunCmd.MarkFlagRequired("input")
	var createWorkflowRunOutput string
	createWorkflowRunCmd.Flags().StringVar(&createWorkflowRunOutput, "output", "", "This is a required parameter.")
	createWorkflowRunCmd.MarkFlagRequired("output")

	var createWorkflowRunCreationTime string
	createWorkflowRunCmd.Flags().StringVar(&createWorkflowRunCreationTime, "creation-time", "", "")
	var createWorkflowRunEndTime string
	createWorkflowRunCmd.Flags().StringVar(&createWorkflowRunEndTime, "end-time", "", "")
	var createWorkflowRunEvaluate bool
	createWorkflowRunCmd.Flags().BoolVar(&createWorkflowRunEvaluate, "evaluate", false, "Determine whether to evaluate the prediction.")
	var createWorkflowRunName string
	createWorkflowRunCmd.Flags().StringVar(&createWorkflowRunName, "name", "", "")
	var createWorkflowRunPredictionScore string
	createWorkflowRunCmd.Flags().StringVar(&createWorkflowRunPredictionScore, "prediction-score", "", "")
	var createWorkflowRunStartTime string
	createWorkflowRunCmd.Flags().StringVar(&createWorkflowRunStartTime, "start-time", "", "")
	var createWorkflowRunStatus string
	createWorkflowRunCmd.Flags().StringVar(&createWorkflowRunStatus, "status", "", "json string representing ")
	var createWorkflowRunTimeoutSecs int32
	createWorkflowRunCmd.Flags().Int32Var(&createWorkflowRunTimeoutSecs, "timeout-secs", 0, "Number of seconds before a workflow run times out.")
	var createWorkflowRunWorkflowBuild string
	createWorkflowRunCmd.Flags().StringVar(&createWorkflowRunWorkflowBuild, "workflow-build", "", "")

	mlCmd.AddCommand(createWorkflowStreamDeploymentCmd)

	var createWorkflowStreamDeploymentBuildId string
	createWorkflowStreamDeploymentCmd.Flags().StringVar(&createWorkflowStreamDeploymentBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	createWorkflowStreamDeploymentCmd.MarkFlagRequired("build-id")
	var createWorkflowStreamDeploymentId string
	createWorkflowStreamDeploymentCmd.Flags().StringVar(&createWorkflowStreamDeploymentId, "id", "", "This is a required parameter.The workflow ID.")
	createWorkflowStreamDeploymentCmd.MarkFlagRequired("id")
	var createWorkflowStreamDeploymentInput string
	createWorkflowStreamDeploymentCmd.Flags().StringVar(&createWorkflowStreamDeploymentInput, "input", "", "This is a required parameter.")
	createWorkflowStreamDeploymentCmd.MarkFlagRequired("input")
	var createWorkflowStreamDeploymentOutput string
	createWorkflowStreamDeploymentCmd.Flags().StringVar(&createWorkflowStreamDeploymentOutput, "output", "", "This is a required parameter.")
	createWorkflowStreamDeploymentCmd.MarkFlagRequired("output")

	var createWorkflowStreamDeploymentCreationTime string
	createWorkflowStreamDeploymentCmd.Flags().StringVar(&createWorkflowStreamDeploymentCreationTime, "creation-time", "", "")
	var createWorkflowStreamDeploymentEndTime string
	createWorkflowStreamDeploymentCmd.Flags().StringVar(&createWorkflowStreamDeploymentEndTime, "end-time", "", "")
	var createWorkflowStreamDeploymentName string
	createWorkflowStreamDeploymentCmd.Flags().StringVar(&createWorkflowStreamDeploymentName, "name", "", "")
	var createWorkflowStreamDeploymentSpec string
	createWorkflowStreamDeploymentCmd.Flags().StringVar(&createWorkflowStreamDeploymentSpec, "spec", "", "")
	var createWorkflowStreamDeploymentStartTime string
	createWorkflowStreamDeploymentCmd.Flags().StringVar(&createWorkflowStreamDeploymentStartTime, "start-time", "", "")
	var createWorkflowStreamDeploymentStatus string
	createWorkflowStreamDeploymentCmd.Flags().StringVar(&createWorkflowStreamDeploymentStatus, "status", "", "json string representing ")
	var createWorkflowStreamDeploymentWorkflowBuild string
	createWorkflowStreamDeploymentCmd.Flags().StringVar(&createWorkflowStreamDeploymentWorkflowBuild, "workflow-build", "", "")

	mlCmd.AddCommand(deleteWorkflowCmd)

	var deleteWorkflowId string
	deleteWorkflowCmd.Flags().StringVar(&deleteWorkflowId, "id", "", "This is a required parameter.The workflow ID.")
	deleteWorkflowCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(deleteWorkflowBuildCmd)

	var deleteWorkflowBuildBuildId string
	deleteWorkflowBuildCmd.Flags().StringVar(&deleteWorkflowBuildBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	deleteWorkflowBuildCmd.MarkFlagRequired("build-id")
	var deleteWorkflowBuildId string
	deleteWorkflowBuildCmd.Flags().StringVar(&deleteWorkflowBuildId, "id", "", "This is a required parameter.The workflow ID.")
	deleteWorkflowBuildCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(deleteWorkflowDeploymentCmd)

	var deleteWorkflowDeploymentBuildId string
	deleteWorkflowDeploymentCmd.Flags().StringVar(&deleteWorkflowDeploymentBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	deleteWorkflowDeploymentCmd.MarkFlagRequired("build-id")
	var deleteWorkflowDeploymentDeploymentId string
	deleteWorkflowDeploymentCmd.Flags().StringVar(&deleteWorkflowDeploymentDeploymentId, "deployment-id", "", "This is a required parameter.The workflow deployment ID.")
	deleteWorkflowDeploymentCmd.MarkFlagRequired("deployment-id")
	var deleteWorkflowDeploymentId string
	deleteWorkflowDeploymentCmd.Flags().StringVar(&deleteWorkflowDeploymentId, "id", "", "This is a required parameter.The workflow ID.")
	deleteWorkflowDeploymentCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(deleteWorkflowRunCmd)

	var deleteWorkflowRunBuildId string
	deleteWorkflowRunCmd.Flags().StringVar(&deleteWorkflowRunBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	deleteWorkflowRunCmd.MarkFlagRequired("build-id")
	var deleteWorkflowRunId string
	deleteWorkflowRunCmd.Flags().StringVar(&deleteWorkflowRunId, "id", "", "This is a required parameter.The workflow ID.")
	deleteWorkflowRunCmd.MarkFlagRequired("id")
	var deleteWorkflowRunRunId string
	deleteWorkflowRunCmd.Flags().StringVar(&deleteWorkflowRunRunId, "run-id", "", "This is a required parameter.The workflow run ID.")
	deleteWorkflowRunCmd.MarkFlagRequired("run-id")

	mlCmd.AddCommand(deleteWorkflowStreamDeploymentCmd)

	var deleteWorkflowStreamDeploymentBuildId string
	deleteWorkflowStreamDeploymentCmd.Flags().StringVar(&deleteWorkflowStreamDeploymentBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	deleteWorkflowStreamDeploymentCmd.MarkFlagRequired("build-id")
	var deleteWorkflowStreamDeploymentId string
	deleteWorkflowStreamDeploymentCmd.Flags().StringVar(&deleteWorkflowStreamDeploymentId, "id", "", "This is a required parameter.The workflow ID.")
	deleteWorkflowStreamDeploymentCmd.MarkFlagRequired("id")
	var deleteWorkflowStreamDeploymentStreamDeploymentId string
	deleteWorkflowStreamDeploymentCmd.Flags().StringVar(&deleteWorkflowStreamDeploymentStreamDeploymentId, "stream-deployment-id", "", "This is a required parameter.The workflow streaming deployment ID.")
	deleteWorkflowStreamDeploymentCmd.MarkFlagRequired("stream-deployment-id")

	mlCmd.AddCommand(getWorkflowCmd)

	var getWorkflowId string
	getWorkflowCmd.Flags().StringVar(&getWorkflowId, "id", "", "This is a required parameter.The workflow ID.")
	getWorkflowCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(getWorkflowBuildCmd)

	var getWorkflowBuildBuildId string
	getWorkflowBuildCmd.Flags().StringVar(&getWorkflowBuildBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	getWorkflowBuildCmd.MarkFlagRequired("build-id")
	var getWorkflowBuildId string
	getWorkflowBuildCmd.Flags().StringVar(&getWorkflowBuildId, "id", "", "This is a required parameter.The workflow ID.")
	getWorkflowBuildCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(getWorkflowBuildErrorCmd)

	var getWorkflowBuildErrorBuildId string
	getWorkflowBuildErrorCmd.Flags().StringVar(&getWorkflowBuildErrorBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	getWorkflowBuildErrorCmd.MarkFlagRequired("build-id")
	var getWorkflowBuildErrorId string
	getWorkflowBuildErrorCmd.Flags().StringVar(&getWorkflowBuildErrorId, "id", "", "This is a required parameter.The workflow ID.")
	getWorkflowBuildErrorCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(getWorkflowBuildLogCmd)

	var getWorkflowBuildLogBuildId string
	getWorkflowBuildLogCmd.Flags().StringVar(&getWorkflowBuildLogBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	getWorkflowBuildLogCmd.MarkFlagRequired("build-id")
	var getWorkflowBuildLogId string
	getWorkflowBuildLogCmd.Flags().StringVar(&getWorkflowBuildLogId, "id", "", "This is a required parameter.The workflow ID.")
	getWorkflowBuildLogCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(getWorkflowDeploymentCmd)

	var getWorkflowDeploymentBuildId string
	getWorkflowDeploymentCmd.Flags().StringVar(&getWorkflowDeploymentBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	getWorkflowDeploymentCmd.MarkFlagRequired("build-id")
	var getWorkflowDeploymentDeploymentId string
	getWorkflowDeploymentCmd.Flags().StringVar(&getWorkflowDeploymentDeploymentId, "deployment-id", "", "This is a required parameter.The workflow deployment ID.")
	getWorkflowDeploymentCmd.MarkFlagRequired("deployment-id")
	var getWorkflowDeploymentId string
	getWorkflowDeploymentCmd.Flags().StringVar(&getWorkflowDeploymentId, "id", "", "This is a required parameter.The workflow ID.")
	getWorkflowDeploymentCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(getWorkflowDeploymentErrorCmd)

	var getWorkflowDeploymentErrorBuildId string
	getWorkflowDeploymentErrorCmd.Flags().StringVar(&getWorkflowDeploymentErrorBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	getWorkflowDeploymentErrorCmd.MarkFlagRequired("build-id")
	var getWorkflowDeploymentErrorDeploymentId string
	getWorkflowDeploymentErrorCmd.Flags().StringVar(&getWorkflowDeploymentErrorDeploymentId, "deployment-id", "", "This is a required parameter.The workflow deployment ID.")
	getWorkflowDeploymentErrorCmd.MarkFlagRequired("deployment-id")
	var getWorkflowDeploymentErrorId string
	getWorkflowDeploymentErrorCmd.Flags().StringVar(&getWorkflowDeploymentErrorId, "id", "", "This is a required parameter.The workflow ID.")
	getWorkflowDeploymentErrorCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(getWorkflowDeploymentLogCmd)

	var getWorkflowDeploymentLogBuildId string
	getWorkflowDeploymentLogCmd.Flags().StringVar(&getWorkflowDeploymentLogBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	getWorkflowDeploymentLogCmd.MarkFlagRequired("build-id")
	var getWorkflowDeploymentLogDeploymentId string
	getWorkflowDeploymentLogCmd.Flags().StringVar(&getWorkflowDeploymentLogDeploymentId, "deployment-id", "", "This is a required parameter.The workflow deployment ID.")
	getWorkflowDeploymentLogCmd.MarkFlagRequired("deployment-id")
	var getWorkflowDeploymentLogId string
	getWorkflowDeploymentLogCmd.Flags().StringVar(&getWorkflowDeploymentLogId, "id", "", "This is a required parameter.The workflow ID.")
	getWorkflowDeploymentLogCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(getWorkflowRunCmd)

	var getWorkflowRunBuildId string
	getWorkflowRunCmd.Flags().StringVar(&getWorkflowRunBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	getWorkflowRunCmd.MarkFlagRequired("build-id")
	var getWorkflowRunId string
	getWorkflowRunCmd.Flags().StringVar(&getWorkflowRunId, "id", "", "This is a required parameter.The workflow ID.")
	getWorkflowRunCmd.MarkFlagRequired("id")
	var getWorkflowRunRunId string
	getWorkflowRunCmd.Flags().StringVar(&getWorkflowRunRunId, "run-id", "", "This is a required parameter.The workflow run ID.")
	getWorkflowRunCmd.MarkFlagRequired("run-id")

	mlCmd.AddCommand(getWorkflowRunErrorCmd)

	var getWorkflowRunErrorBuildId string
	getWorkflowRunErrorCmd.Flags().StringVar(&getWorkflowRunErrorBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	getWorkflowRunErrorCmd.MarkFlagRequired("build-id")
	var getWorkflowRunErrorId string
	getWorkflowRunErrorCmd.Flags().StringVar(&getWorkflowRunErrorId, "id", "", "This is a required parameter.The workflow ID.")
	getWorkflowRunErrorCmd.MarkFlagRequired("id")
	var getWorkflowRunErrorRunId string
	getWorkflowRunErrorCmd.Flags().StringVar(&getWorkflowRunErrorRunId, "run-id", "", "This is a required parameter.The workflow run ID.")
	getWorkflowRunErrorCmd.MarkFlagRequired("run-id")

	mlCmd.AddCommand(getWorkflowRunLogCmd)

	var getWorkflowRunLogBuildId string
	getWorkflowRunLogCmd.Flags().StringVar(&getWorkflowRunLogBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	getWorkflowRunLogCmd.MarkFlagRequired("build-id")
	var getWorkflowRunLogId string
	getWorkflowRunLogCmd.Flags().StringVar(&getWorkflowRunLogId, "id", "", "This is a required parameter.The workflow ID.")
	getWorkflowRunLogCmd.MarkFlagRequired("id")
	var getWorkflowRunLogRunId string
	getWorkflowRunLogCmd.Flags().StringVar(&getWorkflowRunLogRunId, "run-id", "", "This is a required parameter.The workflow run ID.")
	getWorkflowRunLogCmd.MarkFlagRequired("run-id")

	mlCmd.AddCommand(getWorkflowStreamDeploymentCmd)

	var getWorkflowStreamDeploymentBuildId string
	getWorkflowStreamDeploymentCmd.Flags().StringVar(&getWorkflowStreamDeploymentBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	getWorkflowStreamDeploymentCmd.MarkFlagRequired("build-id")
	var getWorkflowStreamDeploymentId string
	getWorkflowStreamDeploymentCmd.Flags().StringVar(&getWorkflowStreamDeploymentId, "id", "", "This is a required parameter.The workflow ID.")
	getWorkflowStreamDeploymentCmd.MarkFlagRequired("id")
	var getWorkflowStreamDeploymentStreamDeploymentId string
	getWorkflowStreamDeploymentCmd.Flags().StringVar(&getWorkflowStreamDeploymentStreamDeploymentId, "stream-deployment-id", "", "This is a required parameter.The workflow streaming deployment ID.")
	getWorkflowStreamDeploymentCmd.MarkFlagRequired("stream-deployment-id")

	mlCmd.AddCommand(listWorkflowBuildsCmd)

	var listWorkflowBuildsId string
	listWorkflowBuildsCmd.Flags().StringVar(&listWorkflowBuildsId, "id", "", "This is a required parameter.The workflow ID.")
	listWorkflowBuildsCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(listWorkflowDeploymentsCmd)

	var listWorkflowDeploymentsBuildId string
	listWorkflowDeploymentsCmd.Flags().StringVar(&listWorkflowDeploymentsBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	listWorkflowDeploymentsCmd.MarkFlagRequired("build-id")
	var listWorkflowDeploymentsId string
	listWorkflowDeploymentsCmd.Flags().StringVar(&listWorkflowDeploymentsId, "id", "", "This is a required parameter.The workflow ID.")
	listWorkflowDeploymentsCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(listWorkflowRunsCmd)

	var listWorkflowRunsBuildId string
	listWorkflowRunsCmd.Flags().StringVar(&listWorkflowRunsBuildId, "build-id", "", "This is a required parameter.The workflow build ID.")
	listWorkflowRunsCmd.MarkFlagRequired("build-id")
	var listWorkflowRunsId string
	listWorkflowRunsCmd.Flags().StringVar(&listWorkflowRunsId, "id", "", "This is a required parameter.The workflow ID.")
	listWorkflowRunsCmd.MarkFlagRequired("id")

	mlCmd.AddCommand(listWorkflowsCmd)

}