package powerbi

import (
	"context"
	"fmt"
)

// PushDatasetsClient handles operations for Power BI push datasets (streaming datasets)
type PushDatasetsClient struct {
	client *Client
}

// Table represents a table in a push dataset
type Table struct {
	Name    *string  `json:"name,omitempty"`
	Columns []Column `json:"columns,omitempty"`
	Rows    []Row    `json:"rows,omitempty"`
	Measures []Measure `json:"measures,omitempty"`
}

// Column represents a column in a table
type Column struct {
	Name     *string `json:"name,omitempty"`
	DataType *string `json:"dataType,omitempty"`
}

// Row represents a row of data
type Row map[string]interface{}

// Measure represents a measure in a table
type Measure struct {
	Name       *string `json:"name,omitempty"`
	Expression *string `json:"expression,omitempty"`
}

// CreatePushDatasetRequest represents a request to create a push dataset
type CreatePushDatasetRequest struct {
	Name                 *string  `json:"name,omitempty"`
	DefaultMode          *string  `json:"defaultMode,omitempty"` // Push, Streaming, PushStreaming
	Tables               []Table  `json:"tables,omitempty"`
	DefaultRetentionPolicy *string `json:"defaultRetentionPolicy,omitempty"`
}

// PostRowsRequest represents a request to add rows to a table
type PostRowsRequest struct {
	Rows []Row `json:"rows"`
}

// CreatePushDataset creates a new push dataset in "My Workspace"
func (c *PushDatasetsClient) CreatePushDataset(ctx context.Context, request CreatePushDatasetRequest) (*Dataset, error) {
	path := "/datasets"
	
	var result Dataset
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreatePushDatasetInGroup creates a new push dataset in the specified workspace
func (c *PushDatasetsClient) CreatePushDatasetInGroup(ctx context.Context, groupID string, request CreatePushDatasetRequest) (*Dataset, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets", groupID)
	
	var result Dataset
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// PostRows adds new data rows to the specified table within the specified dataset from "My Workspace"
func (c *PushDatasetsClient) PostRows(ctx context.Context, datasetID, tableName string, rows []Row) error {
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}
	if tableName == "" {
		return fmt.Errorf("tableName cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/tables/%s/rows", datasetID, tableName)
	
	request := PostRowsRequest{Rows: rows}
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// PostRowsInGroup adds new data rows to the specified table within the specified dataset from the specified workspace
func (c *PushDatasetsClient) PostRowsInGroup(ctx context.Context, groupID, datasetID, tableName string, rows []Row) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}
	if tableName == "" {
		return fmt.Errorf("tableName cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/tables/%s/rows", groupID, datasetID, tableName)
	
	request := PostRowsRequest{Rows: rows}
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// DeleteRows deletes all rows from the specified table within the specified dataset from "My Workspace"
func (c *PushDatasetsClient) DeleteRows(ctx context.Context, datasetID, tableName string) error {
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}
	if tableName == "" {
		return fmt.Errorf("tableName cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/tables/%s/rows", datasetID, tableName)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// DeleteRowsInGroup deletes all rows from the specified table within the specified dataset from the specified workspace
func (c *PushDatasetsClient) DeleteRowsInGroup(ctx context.Context, groupID, datasetID, tableName string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}
	if tableName == "" {
		return fmt.Errorf("tableName cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/tables/%s/rows", groupID, datasetID, tableName)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// GetTables returns a list of tables within the specified dataset from "My Workspace"
func (c *PushDatasetsClient) GetTables(ctx context.Context, datasetID string) ([]Table, error) {
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/tables", datasetID)

	var result struct {
		Value []Table `json:"value"`
	}
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}

// GetTablesInGroup returns a list of tables within the specified dataset from the specified workspace
func (c *PushDatasetsClient) GetTablesInGroup(ctx context.Context, groupID, datasetID string) ([]Table, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return nil, fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/tables", groupID, datasetID)

	var result struct {
		Value []Table `json:"value"`
	}
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}

// PutTable updates the metadata and schema for the specified table within the specified dataset from "My Workspace"
func (c *PushDatasetsClient) PutTable(ctx context.Context, datasetID, tableName string, table Table) error {
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}
	if tableName == "" {
		return fmt.Errorf("tableName cannot be empty")
	}

	path := fmt.Sprintf("/datasets/%s/tables/%s", datasetID, tableName)
	return c.client.doRequest(ctx, "PUT", path, table, nil)
}

// PutTableInGroup updates the metadata and schema for the specified table within the specified dataset from the specified workspace
func (c *PushDatasetsClient) PutTableInGroup(ctx context.Context, groupID, datasetID, tableName string, table Table) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}
	if tableName == "" {
		return fmt.Errorf("tableName cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/datasets/%s/tables/%s", groupID, datasetID, tableName)
	return c.client.doRequest(ctx, "PUT", path, table, nil)
}

