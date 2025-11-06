package powerbi

import (
	"context"
	"fmt"
)

// AdminClient handles administrative operations for Power BI
type AdminClient struct {
	client *Client
}

// GetReportsAsAdminOptions contains optional parameters for GetReportsAsAdmin
type GetReportsAsAdminOptions struct {
	// Filter OData filter query
	Filter *string
	// Top OData top query - maximum number of items to return
	Top *int
	// Skip OData skip query - number of items to skip
	Skip *int
}

// GetReportsAsAdmin returns a list of reports for the organization
func (c *AdminClient) GetReportsAsAdmin(ctx context.Context, options *GetReportsAsAdminOptions) (*AdminReports, error) {
	path := "/admin/reports"
	
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

	var result AdminReports
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDatasetsAsAdmin returns a list of datasets for the organization
func (c *AdminClient) GetDatasetsAsAdmin(ctx context.Context, options *GetReportsAsAdminOptions) (*AdminDatasets, error) {
	path := "/admin/datasets"
	
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

	var result AdminDatasets
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDashboardsAsAdmin returns a list of dashboards for the organization
func (c *AdminClient) GetDashboardsAsAdmin(ctx context.Context, options *GetReportsAsAdminOptions) (*AdminDashboards, error) {
	path := "/admin/dashboards"
	
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

	var result AdminDashboards
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGroupsAsAdmin returns a list of workspaces for the organization
func (c *AdminClient) GetGroupsAsAdmin(ctx context.Context, options *GetGroupsAsAdminOptions) (*AdminGroups, error) {
	path := "/admin/groups"
	
	if options != nil {
		queryParams := make(map[string]string)
		if options.Expand != nil {
			queryParams["$expand"] = *options.Expand
		}
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

	var result AdminGroups
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGroupAsAdmin returns the specified workspace
func (c *AdminClient) GetGroupAsAdmin(ctx context.Context, groupID string, expand *string) (*Group, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/admin/groups/%s", groupID)
	
	if expand != nil {
		queryParams := map[string]string{
			"$expand": *expand,
		}
		path += buildQueryParams(queryParams)
	}

	var result Group
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGroupUsersAsAdmin returns a list of users that have access to the specified workspace
func (c *AdminClient) GetGroupUsersAsAdmin(ctx context.Context, groupID string) ([]GroupUser, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/admin/groups/%s/users", groupID)

	var result struct {
		Value []GroupUser `json:"value"`
	}
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}

// AddGroupUserAsAdmin grants user permissions to the specified workspace
func (c *AdminClient) AddGroupUserAsAdmin(ctx context.Context, groupID string, request GroupUserAccessRight) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/admin/groups/%s/users", groupID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// DeleteGroupUserAsAdmin removes user permissions from the specified workspace
func (c *AdminClient) DeleteGroupUserAsAdmin(ctx context.Context, groupID, userID string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if userID == "" {
		return fmt.Errorf("userID cannot be empty")
	}

	path := fmt.Sprintf("/admin/groups/%s/users/%s", groupID, userID)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// RestoreGroupAsAdmin restores a deleted workspace
func (c *AdminClient) RestoreGroupAsAdmin(ctx context.Context, groupID string, request RestoreGroupRequest) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/admin/groups/%s/restore", groupID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// GetCapacitiesAsAdmin returns a list of capacities for the organization
func (c *AdminClient) GetCapacitiesAsAdmin(ctx context.Context, expand *string) (*Capacities, error) {
	path := "/admin/capacities"
	
	if expand != nil {
		queryParams := map[string]string{
			"$expand": *expand,
		}
		path += buildQueryParams(queryParams)
	}

	var result Capacities
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetRefreshableForCapacity returns a list of refreshables for the specified capacity
func (c *AdminClient) GetRefreshableForCapacity(ctx context.Context, capacityID string, options *GetRefreshableOptions) (*Refreshables, error) {
	if capacityID == "" {
		return nil, fmt.Errorf("capacityID cannot be empty")
	}

	path := fmt.Sprintf("/admin/capacities/%s/refreshables", capacityID)
	
	if options != nil {
		queryParams := make(map[string]string)
		if options.Top != nil {
			queryParams["$top"] = fmt.Sprintf("%d", *options.Top)
		}
		if options.Skip != nil {
			queryParams["$skip"] = fmt.Sprintf("%d", *options.Skip)
		}
		if options.Expand != nil {
			queryParams["$expand"] = *options.Expand
		}
		if options.Filter != nil {
			queryParams["$filter"] = *options.Filter
		}
		path += buildQueryParams(queryParams)
	}

	var result Refreshables
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetActivityEvents returns a list of audit activity events
func (c *AdminClient) GetActivityEvents(ctx context.Context, startDateTime, endDateTime string, options *GetActivityEventsOptions) (*ActivityEvents, error) {
	if startDateTime == "" {
		return nil, fmt.Errorf("startDateTime cannot be empty")
	}
	if endDateTime == "" {
		return nil, fmt.Errorf("endDateTime cannot be empty")
	}

	path := "/admin/activityevents"
	
	queryParams := map[string]string{
		"startDateTime": startDateTime,
		"endDateTime":   endDateTime,
	}
	
	if options != nil {
		if options.Filter != nil {
			queryParams["$filter"] = *options.Filter
		}
		if options.ContinuationToken != nil {
			queryParams["continuationToken"] = *options.ContinuationToken
		}
	}
	
	path += buildQueryParams(queryParams)

	var result ActivityEvents
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGroupsAsAdminOptions contains optional parameters for GetGroupsAsAdmin
type GetGroupsAsAdminOptions struct {
	Expand *string
	Filter *string
	Top    *int
	Skip   *int
}

// GetRefreshableOptions contains optional parameters for GetRefreshableForCapacity
type GetRefreshableOptions struct {
	Top    *int
	Skip   *int
	Expand *string
	Filter *string
}

// GetActivityEventsOptions contains optional parameters for GetActivityEvents
type GetActivityEventsOptions struct {
	Filter            *string
	ContinuationToken *string
}

// AdminReports represents a list of reports (admin view)
type AdminReports struct {
	ODataContext string        `json:"@odata.context,omitempty"`
	Value        []AdminReport `json:"value"`
}

// AdminReport represents a report (admin view)
type AdminReport struct {
	ID                       *string    `json:"id,omitempty"`
	ReportType               *string    `json:"reportType,omitempty"`
	Name                     *string    `json:"name,omitempty"`
	WebURL                   *string    `json:"webUrl,omitempty"`
	EmbedURL                 *string    `json:"embedUrl,omitempty"`
	DatasetID                *string    `json:"datasetId,omitempty"`
	DatasetWorkspaceID       *string    `json:"datasetWorkspaceId,omitempty"`
	CreatedBy                *string    `json:"createdBy,omitempty"`
	CreatedDateTime          *string    `json:"createdDateTime,omitempty"`
	ModifiedBy               *string    `json:"modifiedBy,omitempty"`
	ModifiedDateTime         *string    `json:"modifiedDateTime,omitempty"`
	WorkspaceID              *string    `json:"workspaceId,omitempty"`
	Description              *string    `json:"description,omitempty"`
	Users                    []User     `json:"users,omitempty"`
}

// AdminDatasets represents a list of datasets (admin view)
type AdminDatasets struct {
	ODataContext string         `json:"@odata.context,omitempty"`
	Value        []AdminDataset `json:"value"`
}

// AdminDataset represents a dataset (admin view)
type AdminDataset struct {
	ID                                *string `json:"id,omitempty"`
	Name                              *string `json:"name,omitempty"`
	ConfiguredBy                      *string `json:"configuredBy,omitempty"`
	CreatedDate                       *string `json:"createdDate,omitempty"`
	ContentProviderType               *string `json:"contentProviderType,omitempty"`
	Description                       *string `json:"description,omitempty"`
	WorkspaceID                       *string `json:"workspaceId,omitempty"`
	IsRefreshable                     *bool   `json:"isRefreshable,omitempty"`
	IsEffectiveIdentityRequired       *bool   `json:"isEffectiveIdentityRequired,omitempty"`
	IsEffectiveIdentityRolesRequired  *bool   `json:"isEffectiveIdentityRolesRequired,omitempty"`
	TargetStorageMode                 *string `json:"targetStorageMode,omitempty"`
	Users                             []User  `json:"users,omitempty"`
}

// AdminDashboards represents a list of dashboards (admin view)
type AdminDashboards struct {
	ODataContext string           `json:"@odata.context,omitempty"`
	Value        []AdminDashboard `json:"value"`
}

// AdminDashboard represents a dashboard (admin view)
type AdminDashboard struct {
	ID               *string `json:"id,omitempty"`
	DisplayName      *string `json:"displayName,omitempty"`
	IsReadOnly       *bool   `json:"isReadOnly,omitempty"`
	WorkspaceID      *string `json:"workspaceId,omitempty"`
	Users            []User  `json:"users,omitempty"`
}

// AdminGroups represents a list of workspaces (admin view)
type AdminGroups struct {
	ODataContext string  `json:"@odata.context,omitempty"`
	Value        []Group `json:"value"`
}

// Refreshables represents a list of refreshables
type Refreshables struct {
	ODataContext string        `json:"@odata.context,omitempty"`
	Value        []Refreshable `json:"value"`
}

// Refreshable represents a refreshable item
type Refreshable struct {
	ID               *string `json:"id,omitempty"`
	Name             *string `json:"name,omitempty"`
	Kind             *string `json:"kind,omitempty"`
	StartTime        *string `json:"startTime,omitempty"`
	EndTime          *string `json:"endTime,omitempty"`
	RefreshCount     *int    `json:"refreshCount,omitempty"`
	RefreshFailures  *int    `json:"refreshFailures,omitempty"`
	AverageTime      *float64 `json:"averageTime,omitempty"`
	MedianTime       *float64 `json:"medianTime,omitempty"`
	RefreshSchedule  *RefreshSchedule `json:"refreshSchedule,omitempty"`
	ConfiguredBy     []string `json:"configuredBy,omitempty"`
	Capacity         *CapacityReference `json:"capacity,omitempty"`
}

// RefreshSchedule represents a refresh schedule
type RefreshSchedule struct {
	Days      []string `json:"days,omitempty"`
	Times     []string `json:"times,omitempty"`
	Enabled   *bool    `json:"enabled,omitempty"`
	LocalTimeZoneID *string `json:"localTimeZoneId,omitempty"`
}

// CapacityReference represents a capacity reference
type CapacityReference struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ActivityEvents represents a list of activity events
type ActivityEvents struct {
	ActivityEventEntities []ActivityEvent `json:"activityEventEntities,omitempty"`
	ContinuationURI       *string         `json:"continuationUri,omitempty"`
	ContinuationToken     *string         `json:"continuationToken,omitempty"`
}

// ActivityEvent represents an activity event
type ActivityEvent struct {
	ID                      *string `json:"Id,omitempty"`
	RecordType              *int    `json:"RecordType,omitempty"`
	CreationTime            *string `json:"CreationTime,omitempty"`
	Operation               *string `json:"Operation,omitempty"`
	OrganizationID          *string `json:"OrganizationId,omitempty"`
	UserType                *int    `json:"UserType,omitempty"`
	UserKey                 *string `json:"UserKey,omitempty"`
	Workload                *string `json:"Workload,omitempty"`
	UserID                  *string `json:"UserId,omitempty"`
	ClientIP                *string `json:"ClientIP,omitempty"`
	UserAgent               *string `json:"UserAgent,omitempty"`
	Activity                *string `json:"Activity,omitempty"`
	ItemName                *string `json:"ItemName,omitempty"`
	WorkspaceID             *string `json:"WorkspaceId,omitempty"`
	WorkspaceName           *string `json:"WorkspaceName,omitempty"`
	DatasetID               *string `json:"DatasetId,omitempty"`
	DatasetName             *string `json:"DatasetName,omitempty"`
	ReportID                *string `json:"ReportId,omitempty"`
	ReportName              *string `json:"ReportName,omitempty"`
	CapacityID              *string `json:"CapacityId,omitempty"`
	CapacityName            *string `json:"CapacityName,omitempty"`
}

