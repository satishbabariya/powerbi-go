package powerbi

import (
	"context"
)

// DataflowsClient handles operations for Power BI dataflows
type DataflowsClient struct {
	client *Client
}

// CapacitiesClient handles operations for Power BI capacities
type CapacitiesClient struct {
	client *Client
}

// AppsClient handles operations for Power BI apps
type AppsClient struct {
	client *Client
}

// ImportsClient handles operations for Power BI imports
type ImportsClient struct {
	client *Client
}

// TilesClient handles operations for Power BI tiles
type TilesClient struct {
	client *Client
}

// UsersClient handles operations for Power BI users
type UsersClient struct {
	client *Client
}

// PipelinesClient handles operations for Power BI pipelines
type PipelinesClient struct {
	client *Client
}

// AvailableFeaturesClient handles operations for Power BI available features
type AvailableFeaturesClient struct {
	client *Client
}

// ProfilesClient handles operations for Power BI profiles
type ProfilesClient struct {
	client *Client
}

// ScorecardsClient handles operations for Power BI scorecards
type ScorecardsClient struct {
	client *Client
}

// GoalsClient handles operations for Power BI goals
type GoalsClient struct {
	client *Client
}

// GoalsStatusRulesClient handles operations for Power BI goals status rules
type GoalsStatusRulesClient struct {
	client *Client
}

// GoalValuesClient handles operations for Power BI goal values
type GoalValuesClient struct {
	client *Client
}

// GoalNotesClient handles operations for Power BI goal notes
type GoalNotesClient struct {
	client *Client
}

// TemplateAppsClient handles operations for Power BI template apps
type TemplateAppsClient struct {
	client *Client
}

// DataflowStorageAccountsClient handles operations for Power BI dataflow storage accounts
type DataflowStorageAccountsClient struct {
	client *Client
}

// WorkspaceInfoClient handles operations for Power BI workspace info
type WorkspaceInfoClient struct {
	client *Client
}

// WidelySharedArtifactsClient handles operations for Power BI widely shared artifacts
type WidelySharedArtifactsClient struct {
	client *Client
}

// InformationProtectionClient handles operations for Power BI information protection
type InformationProtectionClient struct {
	client *Client
}

// GetDataflows returns a list of dataflows from the specified workspace
func (c *DataflowsClient) GetDataflows(ctx context.Context, groupID string) (*Dataflows, error) {
	var result Dataflows
	if err := c.client.doRequest(ctx, "GET", "/groups/"+groupID+"/dataflows", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCapacities returns a list of capacities
func (c *CapacitiesClient) GetCapacities(ctx context.Context) (*Capacities, error) {
	var result Capacities
	if err := c.client.doRequest(ctx, "GET", "/capacities", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApps returns a list of apps
func (c *AppsClient) GetApps(ctx context.Context) (*Apps, error) {
	var result Apps
	if err := c.client.doRequest(ctx, "GET", "/apps", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetImports returns a list of imports from "My Workspace"
func (c *ImportsClient) GetImports(ctx context.Context) (*Imports, error) {
	var result Imports
	if err := c.client.doRequest(ctx, "GET", "/imports", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetImportsInGroup returns a list of imports from the specified workspace
func (c *ImportsClient) GetImportsInGroup(ctx context.Context, groupID string) (*Imports, error) {
	var result Imports
	if err := c.client.doRequest(ctx, "GET", "/groups/"+groupID+"/imports", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

