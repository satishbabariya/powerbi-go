package powerbi

import (
	"context"
	"fmt"
)

// DatasetsClient handles operations for Power BI datasets
type DatasetsClient struct {
	client *Client
}

// GetDatasetsOptions contains optional parameters for GetDatasets
type GetDatasetsOptions struct {
	// Filter OData filter query
	Filter *string
	// Top OData top query - maximum number of items to return
	Top *int
	// Skip OData skip query - number of items to skip
	Skip *int
}

// GetDatasets returns a list of datasets from "My Workspace"
func (c *DatasetsClient) GetDatasets(ctx context.Context, options *GetDatasetsOptions) (*Datasets, error) {
	path := "/datasets"
	
	if options != nil {
		queryParams := make(map[string]string)
		if options.Filter != nil {
			queryParams["$filter"] = *options.Filter
		}
		if options.Top != nil {
			queryParams["$top"] = fmt.Sprintf("%d", *options.Top)
		}
		if options.Skip != nil {
			queryParams["$skip"] = fmt.Sprintf("%d", *options.Skip)
		}
		path += buildQueryParams(queryParams)
	}

	var result Datasets
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDatasetsInGroup returns a list of datasets from the specified workspace
func (c *DatasetsClient) GetDatasetsInGroup(ctx context.Context, groupID string, options *GetDatasetsOptions) (*Datasets, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets", groupID)
	
	if options != nil {
		queryParams := make(map[string]string)
		if options.Filter != nil {
			queryParams["$filter"] = *options.Filter
		}
		if options.Top != nil {
			queryParams["$top"] = fmt.Sprintf("%d", *options.Top)
		}
		if options.Skip != nil {
			queryParams["$skip"] = fmt.Sprintf("%d", *options.Skip)
		}
		path += buildQueryParams(queryParams)
	}

	var result Datasets
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDataset returns the specified dataset from "My Workspace"
func (c *DatasetsClient) GetDataset(ctx context.Context, datasetID string) (*Dataset, error) {
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s", datasetID)

	var result Dataset
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDatasetInGroup returns the specified dataset from the specified workspace
func (c *DatasetsClient) GetDatasetInGroup(ctx context.Context, groupID, datasetID string) (*Dataset, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s", groupID, datasetID)

	var result Dataset
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteDataset deletes the specified dataset from "My Workspace"
func (c *DatasetsClient) DeleteDataset(ctx context.Context, datasetID string) error {
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s", datasetID)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// DeleteDatasetInGroup deletes the specified dataset from the specified workspace
func (c *DatasetsClient) DeleteDatasetInGroup(ctx context.Context, groupID, datasetID string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s", groupID, datasetID)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// RefreshDataset triggers a refresh for the specified dataset from "My Workspace"
func (c *DatasetsClient) RefreshDataset(ctx context.Context, datasetID string, request *RefreshRequest) error {
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/refreshes", datasetID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// RefreshDatasetInGroup triggers a refresh for the specified dataset from the specified workspace
func (c *DatasetsClient) RefreshDatasetInGroup(ctx context.Context, groupID, datasetID string, request *RefreshRequest) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/refreshes", groupID, datasetID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// GetRefreshHistory returns the refresh history for the specified dataset from "My Workspace"
func (c *DatasetsClient) GetRefreshHistory(ctx context.Context, datasetID string, top *int) (*RefreshHistory, error) {
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/refreshes", datasetID)
	
	if top != nil {
		queryParams := map[string]string{
			"$top": fmt.Sprintf("%d", *top),
		}
		path += buildQueryParams(queryParams)
	}

	var result RefreshHistory
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetRefreshHistoryInGroup returns the refresh history for the specified dataset from the specified workspace
func (c *DatasetsClient) GetRefreshHistoryInGroup(ctx context.Context, groupID, datasetID string, top *int) (*RefreshHistory, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/refreshes", groupID, datasetID)
	
	if top != nil {
		queryParams := map[string]string{
			"$top": fmt.Sprintf("%d", *top),
		}
		path += buildQueryParams(queryParams)
	}

	var result RefreshHistory
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// TakeOverDataset takes over the specified dataset
func (c *DatasetsClient) TakeOverDataset(ctx context.Context, datasetID string) error {
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/Default.TakeOver", datasetID)
	return c.client.doRequest(ctx, "POST", path, nil, nil)
}

// TakeOverDatasetInGroup takes over the specified dataset in the specified workspace
func (c *DatasetsClient) TakeOverDatasetInGroup(ctx context.Context, groupID, datasetID string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/Default.TakeOver", groupID, datasetID)
	return c.client.doRequest(ctx, "POST", path, nil, nil)
}

// GetDatasources returns a list of datasources for the specified dataset
func (c *DatasetsClient) GetDatasources(ctx context.Context, datasetID string) (*Datasources, error) {
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/datasources", datasetID)

	var result Datasources
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDatasourcesInGroup returns a list of datasources for the specified dataset in the specified workspace
func (c *DatasetsClient) GetDatasourcesInGroup(ctx context.Context, groupID, datasetID string) (*Datasources, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/datasources", groupID, datasetID)

	var result Datasources
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateDatasources updates the datasources of the specified dataset
func (c *DatasetsClient) UpdateDatasources(ctx context.Context, datasetID string, request UpdateDatasourcesRequest) error {
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/Default.UpdateDatasources", datasetID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// UpdateDatasourcesInGroup updates the datasources of the specified dataset in the specified workspace
func (c *DatasetsClient) UpdateDatasourcesInGroup(ctx context.Context, groupID, datasetID string, request UpdateDatasourcesRequest) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/Default.UpdateDatasources", groupID, datasetID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// BindToGateway binds the specified dataset to the specified gateway
func (c *DatasetsClient) BindToGateway(ctx context.Context, datasetID, gatewayID string) error {
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}
	if gatewayID == "" {
		return fmt.Errorf("gatewayID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/Default.BindToGateway", datasetID)
	
	body := map[string]string{
		"gatewayObjectId": gatewayID,
	}

	return c.client.doRequest(ctx, "POST", path, body, nil)
}

// BindToGatewayInGroup binds the specified dataset to the specified gateway in the specified workspace
func (c *DatasetsClient) BindToGatewayInGroup(ctx context.Context, groupID, datasetID, gatewayID string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}
	if gatewayID == "" {
		return fmt.Errorf("gatewayID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/Default.BindToGateway", groupID, datasetID)
	
	body := map[string]string{
		"gatewayObjectId": gatewayID,
	}

	return c.client.doRequest(ctx, "POST", path, body, nil)
}

// RefreshHistory represents the refresh history for a dataset
type RefreshHistory struct {
	Value []RefreshHistoryEntry `json:"value"`
}

// RefreshHistoryEntry represents a single refresh history entry
type RefreshHistoryEntry struct {
	RefreshType      *string `json:"refreshType,omitempty"`
	StartTime        *string `json:"startTime,omitempty"`
	EndTime          *string `json:"endTime,omitempty"`
	Status           *string `json:"status,omitempty"`
	RequestID        *string `json:"requestId,omitempty"`
	ServiceExceptionJSON *string `json:"serviceExceptionJson,omitempty"`
}

// Datasources represents a list of datasources
type Datasources struct {
	Value []Datasource `json:"value"`
}

// Datasource represents a datasource
type Datasource struct {
	Name              *string            `json:"name,omitempty"`
	ConnectionString  *string            `json:"connectionString,omitempty"`
	DatasourceType    *string            `json:"datasourceType,omitempty"`
	DatasourceID      *string            `json:"datasourceId,omitempty"`
	GatewayID         *string            `json:"gatewayId,omitempty"`
	ConnectionDetails *ConnectionDetails `json:"connectionDetails,omitempty"`
}

// ConnectionDetails represents connection details for a datasource
type ConnectionDetails struct {
	Server   *string `json:"server,omitempty"`
	Database *string `json:"database,omitempty"`
	URL      *string `json:"url,omitempty"`
	Path     *string `json:"path,omitempty"`
}

// UpdateDatasourcesRequest represents a request to update datasources
type UpdateDatasourcesRequest struct {
	UpdateDetails []UpdateDatasourceDetails `json:"updateDetails"`
}

// UpdateDatasourceDetails represents details for updating a datasource
type UpdateDatasourceDetails struct {
	DatasourceSelector *DatasourceSelector `json:"datasourceSelector,omitempty"`
	ConnectionDetails  *ConnectionDetails  `json:"connectionDetails,omitempty"`
}

// DatasourceSelector represents a datasource selector
type DatasourceSelector struct {
	DatasourceType    *string            `json:"datasourceType,omitempty"`
	ConnectionDetails *ConnectionDetails `json:"connectionDetails,omitempty"`
}

// GetParameters returns a list of parameters for the specified dataset
func (c *DatasetsClient) GetParameters(ctx context.Context, datasetID string) (*DatasetParameters, error) {
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/parameters", datasetID)

	var result DatasetParameters
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetParametersInGroup returns a list of parameters for the specified dataset in the specified workspace
func (c *DatasetsClient) GetParametersInGroup(ctx context.Context, groupID, datasetID string) (*DatasetParameters, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/parameters", groupID, datasetID)

	var result DatasetParameters
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateParameters updates the parameters for the specified dataset
func (c *DatasetsClient) UpdateParameters(ctx context.Context, datasetID string, request UpdateParametersRequest) error {
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/Default.UpdateParameters", datasetID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// UpdateParametersInGroup updates the parameters for the specified dataset in the specified workspace
func (c *DatasetsClient) UpdateParametersInGroup(ctx context.Context, groupID, datasetID string, request UpdateParametersRequest) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/Default.UpdateParameters", groupID, datasetID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// DatasetParameters represents dataset parameters
type DatasetParameters struct {
	Value []DatasetParameter `json:"value"`
}

// DatasetParameter represents a single dataset parameter
type DatasetParameter struct {
	Name          *string `json:"name,omitempty"`
	Type          *string `json:"type,omitempty"`
	CurrentValue  *string `json:"currentValue,omitempty"`
	SuggestedValues []string `json:"suggestedValues,omitempty"`
	IsRequired    *bool   `json:"isRequired,omitempty"`
}

// UpdateParametersRequest represents a request to update dataset parameters
type UpdateParametersRequest struct {
	UpdateDetails []UpdateParameterDetails `json:"updateDetails"`
}

// UpdateParameterDetails represents details for updating a parameter
type UpdateParameterDetails struct {
	Name     *string `json:"name,omitempty"`
	NewValue *string `json:"newValue,omitempty"`
}

// ExecuteQueries executes DAX queries against the specified dataset
func (c *DatasetsClient) ExecuteQueries(ctx context.Context, datasetID string, request ExecuteQueriesRequest) (*ExecuteQueriesResponse, error) {
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/executeQueries", datasetID)

	var result ExecuteQueriesResponse
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ExecuteQueriesInGroup executes DAX queries against the specified dataset in the specified workspace
func (c *DatasetsClient) ExecuteQueriesInGroup(ctx context.Context, groupID, datasetID string, request ExecuteQueriesRequest) (*ExecuteQueriesResponse, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/executeQueries", groupID, datasetID)

	var result ExecuteQueriesResponse
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ExecuteQueriesRequest represents a request to execute queries
type ExecuteQueriesRequest struct {
	Queries               []DatasetQuery `json:"queries"`
	SerializerSettings    *SerializerSettings `json:"serializerSettings,omitempty"`
	ImpersonatedUserName  *string `json:"impersonatedUserName,omitempty"`
}

// DatasetQuery represents a DAX query
type DatasetQuery struct {
	Query *string `json:"query,omitempty"`
}

// SerializerSettings represents serializer settings for query results
type SerializerSettings struct {
	IncludeNulls *bool `json:"includeNulls,omitempty"`
}

// ExecuteQueriesResponse represents the response from executing queries
type ExecuteQueriesResponse struct {
	Results []QueryResult `json:"results,omitempty"`
	Error   *QueryError   `json:"error,omitempty"`
}

// QueryResult represents a single query result
type QueryResult struct {
	Tables []QueryTable `json:"tables,omitempty"`
}

// QueryTable represents a table in a query result
type QueryTable struct {
	Rows []map[string]interface{} `json:"rows,omitempty"`
}

// QueryError represents an error from query execution
type QueryError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

