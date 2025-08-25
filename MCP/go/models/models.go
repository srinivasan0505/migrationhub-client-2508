package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DeleteProgressUpdateStreamRequest represents the DeleteProgressUpdateStreamRequest schema from the OpenAPI specification
type DeleteProgressUpdateStreamRequest struct {
	Dryrun interface{} `json:"DryRun,omitempty"`
	Progressupdatestreamname interface{} `json:"ProgressUpdateStreamName"`
}

// ListMigrationTasksRequest represents the ListMigrationTasksRequest schema from the OpenAPI specification
type ListMigrationTasksRequest struct {
	Resourcename interface{} `json:"ResourceName,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ResourceAttribute represents the ResourceAttribute schema from the OpenAPI specification
type ResourceAttribute struct {
	TypeField interface{} `json:"Type"`
	Value interface{} `json:"Value"`
}

// AssociateCreatedArtifactRequestCreatedArtifact represents the AssociateCreatedArtifactRequestCreatedArtifact schema from the OpenAPI specification
type AssociateCreatedArtifactRequestCreatedArtifact struct {
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name"`
}

// MigrationTask represents the MigrationTask schema from the OpenAPI specification
type MigrationTask struct {
	Task MigrationTaskTask `json:"Task,omitempty"`
	Updatedatetime interface{} `json:"UpdateDateTime,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName,omitempty"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream,omitempty"`
	Resourceattributelist interface{} `json:"ResourceAttributeList,omitempty"`
}

// MigrationTaskSummary represents the MigrationTaskSummary schema from the OpenAPI specification
type MigrationTaskSummary struct {
	Migrationtaskname interface{} `json:"MigrationTaskName,omitempty"`
	Progresspercent interface{} `json:"ProgressPercent,omitempty"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusdetail interface{} `json:"StatusDetail,omitempty"`
	Updatedatetime interface{} `json:"UpdateDateTime,omitempty"`
}

// NotifyMigrationTaskStateRequestTask represents the NotifyMigrationTaskStateRequestTask schema from the OpenAPI specification
type NotifyMigrationTaskStateRequestTask struct {
	Statusdetail interface{} `json:"StatusDetail,omitempty"`
	Progresspercent interface{} `json:"ProgressPercent,omitempty"`
	Status interface{} `json:"Status"`
}

// ListApplicationStatesRequest represents the ListApplicationStatesRequest schema from the OpenAPI specification
type ListApplicationStatesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Applicationids interface{} `json:"ApplicationIds,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// CreatedArtifact represents the CreatedArtifact schema from the OpenAPI specification
type CreatedArtifact struct {
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name"`
}

// AssociateCreatedArtifactRequest represents the AssociateCreatedArtifactRequest schema from the OpenAPI specification
type AssociateCreatedArtifactRequest struct {
	Createdartifact AssociateCreatedArtifactRequestCreatedArtifact `json:"CreatedArtifact"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// PutResourceAttributesResult represents the PutResourceAttributesResult schema from the OpenAPI specification
type PutResourceAttributesResult struct {
}

// DescribeMigrationTaskResultMigrationTask represents the DescribeMigrationTaskResultMigrationTask schema from the OpenAPI specification
type DescribeMigrationTaskResultMigrationTask struct {
	Task MigrationTaskTask `json:"Task,omitempty"`
	Updatedatetime interface{} `json:"UpdateDateTime,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName,omitempty"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream,omitempty"`
	Resourceattributelist interface{} `json:"ResourceAttributeList,omitempty"`
}

// ListMigrationTasksResult represents the ListMigrationTasksResult schema from the OpenAPI specification
type ListMigrationTasksResult struct {
	Migrationtasksummarylist interface{} `json:"MigrationTaskSummaryList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DisassociateDiscoveredResourceResult represents the DisassociateDiscoveredResourceResult schema from the OpenAPI specification
type DisassociateDiscoveredResourceResult struct {
}

// NotifyMigrationTaskStateResult represents the NotifyMigrationTaskStateResult schema from the OpenAPI specification
type NotifyMigrationTaskStateResult struct {
}

// DisassociateCreatedArtifactResult represents the DisassociateCreatedArtifactResult schema from the OpenAPI specification
type DisassociateCreatedArtifactResult struct {
}

// PutResourceAttributesRequest represents the PutResourceAttributesRequest schema from the OpenAPI specification
type PutResourceAttributesRequest struct {
	Resourceattributelist interface{} `json:"ResourceAttributeList"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// AssociateCreatedArtifactResult represents the AssociateCreatedArtifactResult schema from the OpenAPI specification
type AssociateCreatedArtifactResult struct {
}

// CreateProgressUpdateStreamRequest represents the CreateProgressUpdateStreamRequest schema from the OpenAPI specification
type CreateProgressUpdateStreamRequest struct {
	Dryrun interface{} `json:"DryRun,omitempty"`
	Progressupdatestreamname interface{} `json:"ProgressUpdateStreamName"`
}

// ListProgressUpdateStreamsRequest represents the ListProgressUpdateStreamsRequest schema from the OpenAPI specification
type ListProgressUpdateStreamsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DisassociateDiscoveredResourceRequest represents the DisassociateDiscoveredResourceRequest schema from the OpenAPI specification
type DisassociateDiscoveredResourceRequest struct {
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
	Configurationid interface{} `json:"ConfigurationId"`
	Dryrun interface{} `json:"DryRun,omitempty"`
}

// ListDiscoveredResourcesRequest represents the ListDiscoveredResourcesRequest schema from the OpenAPI specification
type ListDiscoveredResourcesRequest struct {
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// AssociateDiscoveredResourceResult represents the AssociateDiscoveredResourceResult schema from the OpenAPI specification
type AssociateDiscoveredResourceResult struct {
}

// ImportMigrationTaskResult represents the ImportMigrationTaskResult schema from the OpenAPI specification
type ImportMigrationTaskResult struct {
}

// DescribeApplicationStateResult represents the DescribeApplicationStateResult schema from the OpenAPI specification
type DescribeApplicationStateResult struct {
	Applicationstatus interface{} `json:"ApplicationStatus,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
}

// ListDiscoveredResourcesResult represents the ListDiscoveredResourcesResult schema from the OpenAPI specification
type ListDiscoveredResourcesResult struct {
	Discoveredresourcelist interface{} `json:"DiscoveredResourceList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ApplicationState represents the ApplicationState schema from the OpenAPI specification
type ApplicationState struct {
	Applicationid interface{} `json:"ApplicationId,omitempty"`
	Applicationstatus interface{} `json:"ApplicationStatus,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
}

// ListProgressUpdateStreamsResult represents the ListProgressUpdateStreamsResult schema from the OpenAPI specification
type ListProgressUpdateStreamsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Progressupdatestreamsummarylist interface{} `json:"ProgressUpdateStreamSummaryList,omitempty"`
}

// NotifyApplicationStateResult represents the NotifyApplicationStateResult schema from the OpenAPI specification
type NotifyApplicationStateResult struct {
}

// NotifyApplicationStateRequest represents the NotifyApplicationStateRequest schema from the OpenAPI specification
type NotifyApplicationStateRequest struct {
	Applicationid interface{} `json:"ApplicationId"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Status interface{} `json:"Status"`
	Updatedatetime interface{} `json:"UpdateDateTime,omitempty"`
}

// DiscoveredResource represents the DiscoveredResource schema from the OpenAPI specification
type DiscoveredResource struct {
	Configurationid interface{} `json:"ConfigurationId"`
	Description interface{} `json:"Description,omitempty"`
}

// DescribeMigrationTaskResult represents the DescribeMigrationTaskResult schema from the OpenAPI specification
type DescribeMigrationTaskResult struct {
	Migrationtask DescribeMigrationTaskResultMigrationTask `json:"MigrationTask,omitempty"`
}

// DescribeApplicationStateRequest represents the DescribeApplicationStateRequest schema from the OpenAPI specification
type DescribeApplicationStateRequest struct {
	Applicationid interface{} `json:"ApplicationId"`
}

// ListCreatedArtifactsResult represents the ListCreatedArtifactsResult schema from the OpenAPI specification
type ListCreatedArtifactsResult struct {
	Createdartifactlist interface{} `json:"CreatedArtifactList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MigrationTaskTask represents the MigrationTaskTask schema from the OpenAPI specification
type MigrationTaskTask struct {
	Progresspercent interface{} `json:"ProgressPercent,omitempty"`
	Status interface{} `json:"Status"`
	Statusdetail interface{} `json:"StatusDetail,omitempty"`
}

// DisassociateCreatedArtifactRequest represents the DisassociateCreatedArtifactRequest schema from the OpenAPI specification
type DisassociateCreatedArtifactRequest struct {
	Createdartifactname interface{} `json:"CreatedArtifactName"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// ProgressUpdateStreamSummary represents the ProgressUpdateStreamSummary schema from the OpenAPI specification
type ProgressUpdateStreamSummary struct {
	Progressupdatestreamname interface{} `json:"ProgressUpdateStreamName,omitempty"`
}

// AssociateDiscoveredResourceRequestDiscoveredResource represents the AssociateDiscoveredResourceRequestDiscoveredResource schema from the OpenAPI specification
type AssociateDiscoveredResourceRequestDiscoveredResource struct {
	Configurationid interface{} `json:"ConfigurationId"`
	Description interface{} `json:"Description,omitempty"`
}

// DeleteProgressUpdateStreamResult represents the DeleteProgressUpdateStreamResult schema from the OpenAPI specification
type DeleteProgressUpdateStreamResult struct {
}

// Task represents the Task schema from the OpenAPI specification
type Task struct {
	Progresspercent interface{} `json:"ProgressPercent,omitempty"`
	Status interface{} `json:"Status"`
	Statusdetail interface{} `json:"StatusDetail,omitempty"`
}

// DescribeMigrationTaskRequest represents the DescribeMigrationTaskRequest schema from the OpenAPI specification
type DescribeMigrationTaskRequest struct {
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// AssociateDiscoveredResourceRequest represents the AssociateDiscoveredResourceRequest schema from the OpenAPI specification
type AssociateDiscoveredResourceRequest struct {
	Discoveredresource AssociateDiscoveredResourceRequestDiscoveredResource `json:"DiscoveredResource"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// NotifyMigrationTaskStateRequest represents the NotifyMigrationTaskStateRequest schema from the OpenAPI specification
type NotifyMigrationTaskStateRequest struct {
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Nextupdateseconds interface{} `json:"NextUpdateSeconds"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
	Task NotifyMigrationTaskStateRequestTask `json:"Task"`
	Updatedatetime interface{} `json:"UpdateDateTime"`
	Dryrun interface{} `json:"DryRun,omitempty"`
}

// ListCreatedArtifactsRequest represents the ListCreatedArtifactsRequest schema from the OpenAPI specification
type ListCreatedArtifactsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// ImportMigrationTaskRequest represents the ImportMigrationTaskRequest schema from the OpenAPI specification
type ImportMigrationTaskRequest struct {
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// ListApplicationStatesResult represents the ListApplicationStatesResult schema from the OpenAPI specification
type ListApplicationStatesResult struct {
	Applicationstatelist interface{} `json:"ApplicationStateList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateProgressUpdateStreamResult represents the CreateProgressUpdateStreamResult schema from the OpenAPI specification
type CreateProgressUpdateStreamResult struct {
}
