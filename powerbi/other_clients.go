package powerbi

import (
	"context"
	"fmt"
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

// GenerateTileTokenInGroup generates an embed token for the specified tile in the specified workspace
func (c *TilesClient) GenerateTileTokenInGroup(ctx context.Context, groupID, dashboardID, tileID string, request GenerateTokenRequest) (*EmbedToken, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboardID cannot be empty")
	}
	if tileID == "" {
		return nil, fmt.Errorf("tileID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/dashboards/%s/tiles/%s/GenerateToken", groupID, dashboardID, tileID)

	var result EmbedToken
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GenerateTileToken generates an embed token for the specified tile
func (c *TilesClient) GenerateTileToken(ctx context.Context, dashboardID, tileID string, request GenerateTokenRequest) (*EmbedToken, error) {
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboardID cannot be empty")
	}
	if tileID == "" {
		return nil, fmt.Errorf("tileID cannot be empty")
	}

	path := fmt.Sprintf("/dashboards/%s/tiles/%s/GenerateToken", dashboardID, tileID)

	var result EmbedToken
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GenerateTokenRequest represents a request to generate a token
type GenerateTokenRequest struct {
	AccessLevel *string `json:"accessLevel,omitempty"` // View, Edit, Create
	Identities  []EffectiveIdentity `json:"identities,omitempty"`
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

// GetDataflow returns the specified dataflow from the specified workspace
func (c *DataflowsClient) GetDataflow(ctx context.Context, groupID, dataflowID string) (*Dataflow, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if dataflowID == "" {
		return nil, fmt.Errorf("dataflowID cannot be empty")
	}

	var result Dataflow
	if err := c.client.doRequest(ctx, "GET", fmt.Sprintf("/groups/%s/dataflows/%s", groupID, dataflowID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDataflow deletes the specified dataflow from the specified workspace
func (c *DataflowsClient) DeleteDataflow(ctx context.Context, groupID, dataflowID string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if dataflowID == "" {
		return fmt.Errorf("dataflowID cannot be empty")
	}

	return c.client.doRequest(ctx, "DELETE", fmt.Sprintf("/groups/%s/dataflows/%s", groupID, dataflowID), nil, nil)
}

// RefreshDataflow triggers a refresh of the specified dataflow
func (c *DataflowsClient) RefreshDataflow(ctx context.Context, groupID, dataflowID string, request *DataflowRefreshRequest) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if dataflowID == "" {
		return fmt.Errorf("dataflowID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/dataflows/%s/refreshes", groupID, dataflowID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// GetDataflowTransactions returns the list of transactions for the specified dataflow
func (c *DataflowsClient) GetDataflowTransactions(ctx context.Context, groupID, dataflowID string) (*DataflowTransactions, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if dataflowID == "" {
		return nil, fmt.Errorf("dataflowID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/dataflows/%s/transactions", groupID, dataflowID)
	
	var result DataflowTransactions
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DataflowRefreshRequest represents a request to refresh a dataflow
type DataflowRefreshRequest struct {
	NotifyOption *string `json:"notifyOption,omitempty"` // MailOnFailure, MailOnCompletion, NoNotification
}

// DataflowTransactions represents a list of dataflow transactions
type DataflowTransactions struct {
	Value []DataflowTransaction `json:"value"`
}

// DataflowTransaction represents a dataflow transaction
type DataflowTransaction struct {
	ID               *string `json:"id,omitempty"`
	RefreshType      *string `json:"refreshType,omitempty"`
	StartTime        *string `json:"startTime,omitempty"`
	EndTime          *string `json:"endTime,omitempty"`
	Status           *string `json:"status,omitempty"`
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

// PostImportOptions contains options for importing a file
type PostImportOptions struct {
	DatasetDisplayName *string
	NameConflict       *string // CreateOrOverwrite, Abort, Overwrite, GenerateUniqueName
	SkipReport         *bool
}

// PostImportFile uploads and imports a PBIX file to My Workspace
func (c *ImportsClient) PostImportFile(ctx context.Context, file []byte, filename string, options *PostImportOptions) (*Import, error) {
	return c.postImportFileInternal(ctx, "", file, filename, options)
}

// PostImportFileInGroup uploads and imports a PBIX file to the specified workspace
func (c *ImportsClient) PostImportFileInGroup(ctx context.Context, groupID string, file []byte, filename string, options *PostImportOptions) (*Import, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	return c.postImportFileInternal(ctx, groupID, file, filename, options)
}

// postImportFileInternal is the internal implementation for file import
func (c *ImportsClient) postImportFileInternal(ctx context.Context, groupID string, file []byte, filename string, options *PostImportOptions) (*Import, error) {
	if len(file) == 0 {
		return nil, fmt.Errorf("file cannot be empty")
	}
	if filename == "" {
		return nil, fmt.Errorf("filename cannot be empty")
	}

	// Build the path
	var path string
	if groupID == "" {
		path = "/imports"
	} else {
		path = fmt.Sprintf("/groups/%s/imports", groupID)
	}

	// Add query parameters
	queryParams := make(map[string]string)
	if options != nil {
		if options.DatasetDisplayName != nil {
			queryParams["datasetDisplayName"] = *options.DatasetDisplayName
		}
		if options.NameConflict != nil {
			queryParams["nameConflict"] = *options.NameConflict
		}
		if options.SkipReport != nil {
			if *options.SkipReport {
				queryParams["skipReport"] = "true"
			} else {
				queryParams["skipReport"] = "false"
			}
		}
	}
	path += buildQueryParams(queryParams)

	// Use multipart upload
	var result Import
	additionalFields := make(map[string]string)
	if err := c.client.uploadFileWithResponse(ctx, path, file, filename, additionalFields, &result); err != nil {
		return nil, err
	}
	
	return &result, nil
}

