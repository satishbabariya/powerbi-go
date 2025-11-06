package powerbi

import (
	"context"
	"fmt"
)

// DashboardsClient handles operations for Power BI dashboards
type DashboardsClient struct {
	client *Client
}

// GetDashboardsOptions contains optional parameters for GetDashboards
type GetDashboardsOptions struct {
	// Filter OData filter query
	Filter *string
	// Top OData top query - maximum number of items to return
	Top *int
	// Skip OData skip query - number of items to skip
	Skip *int
}

// GetDashboards returns a list of dashboards from "My Workspace"
func (c *DashboardsClient) GetDashboards(ctx context.Context, options *GetDashboardsOptions) (*Dashboards, error) {
	path := "/dashboards"
	
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

	var result Dashboards
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDashboardsInGroup returns a list of dashboards from the specified workspace
func (c *DashboardsClient) GetDashboardsInGroup(ctx context.Context, groupID string, options *GetDashboardsOptions) (*Dashboards, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/dashboards", groupID)
	
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

	var result Dashboards
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDashboard returns the specified dashboard from "My Workspace"
func (c *DashboardsClient) GetDashboard(ctx context.Context, dashboardID string) (*Dashboard, error) {
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboardID cannot be empty")
	}

	path := fmt.Sprintf("/dashboards/%s", dashboardID)

	var result Dashboard
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDashboardInGroup returns the specified dashboard from the specified workspace
func (c *DashboardsClient) GetDashboardInGroup(ctx context.Context, groupID, dashboardID string) (*Dashboard, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboardID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/dashboards/%s", groupID, dashboardID)

	var result Dashboard
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteDashboard deletes the specified dashboard from "My Workspace"
func (c *DashboardsClient) DeleteDashboard(ctx context.Context, dashboardID string) error {
	if dashboardID == "" {
		return fmt.Errorf("dashboardID cannot be empty")
	}

	path := fmt.Sprintf("/dashboards/%s", dashboardID)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// DeleteDashboardInGroup deletes the specified dashboard from the specified workspace
func (c *DashboardsClient) DeleteDashboardInGroup(ctx context.Context, groupID, dashboardID string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if dashboardID == "" {
		return fmt.Errorf("dashboardID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/dashboards/%s", groupID, dashboardID)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// GetTiles returns a list of tiles within the specified dashboard from "My Workspace"
func (c *DashboardsClient) GetTiles(ctx context.Context, dashboardID string) ([]Tile, error) {
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboardID cannot be empty")
	}

	path := fmt.Sprintf("/dashboards/%s/tiles", dashboardID)

	var result struct {
		Value []Tile `json:"value"`
	}
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}

// GetTilesInGroup returns a list of tiles within the specified dashboard from the specified workspace
func (c *DashboardsClient) GetTilesInGroup(ctx context.Context, groupID, dashboardID string) ([]Tile, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboardID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/dashboards/%s/tiles", groupID, dashboardID)

	var result struct {
		Value []Tile `json:"value"`
	}
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}

// GetTile returns the specified tile within the specified dashboard from "My Workspace"
func (c *DashboardsClient) GetTile(ctx context.Context, dashboardID, tileID string) (*Tile, error) {
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboardID cannot be empty")
	}
	if tileID == "" {
		return nil, fmt.Errorf("tileID cannot be empty")
	}

	path := fmt.Sprintf("/dashboards/%s/tiles/%s", dashboardID, tileID)

	var result Tile
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTileInGroup returns the specified tile within the specified dashboard from the specified workspace
func (c *DashboardsClient) GetTileInGroup(ctx context.Context, groupID, dashboardID, tileID string) (*Tile, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboardID cannot be empty")
	}
	if tileID == "" {
		return nil, fmt.Errorf("tileID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/dashboards/%s/tiles/%s", groupID, dashboardID, tileID)

	var result Tile
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CloneTile clones the specified tile from the specified dashboard
func (c *DashboardsClient) CloneTile(ctx context.Context, dashboardID, tileID string, request CloneTileRequest) (*Tile, error) {
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboardID cannot be empty")
	}
	if tileID == "" {
		return nil, fmt.Errorf("tileID cannot be empty")
	}

	path := fmt.Sprintf("/dashboards/%s/tiles/%s/Clone", dashboardID, tileID)

	var result Tile
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CloneTileInGroup clones the specified tile from the specified dashboard in the specified workspace
func (c *DashboardsClient) CloneTileInGroup(ctx context.Context, groupID, dashboardID, tileID string, request CloneTileRequest) (*Tile, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboardID cannot be empty")
	}
	if tileID == "" {
		return nil, fmt.Errorf("tileID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/dashboards/%s/tiles/%s/Clone", groupID, dashboardID, tileID)

	var result Tile
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CloneTileRequest represents a request to clone a tile
type CloneTileRequest struct {
	TargetDashboardID *string `json:"targetDashboardId,omitempty"`
	TargetWorkspaceID *string `json:"targetWorkspaceId,omitempty"`
	TargetModelID     *string `json:"targetModelId,omitempty"`
	PositionConflictAction *string `json:"positionConflictAction,omitempty"`
}

