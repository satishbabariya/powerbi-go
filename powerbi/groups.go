package powerbi

import (
	"context"
	"fmt"
)

// GroupsClient handles operations for Power BI workspaces (groups)
type GroupsClient struct {
	client *Client
}

// GetGroupsOptions contains optional parameters for GetGroups
type GetGroupsOptions struct {
	// Filter OData filter query
	Filter *string
	// Top OData top query - maximum number of items to return
	Top *int
	// Skip OData skip query - number of items to skip
	Skip *int
}

// GetGroups returns a list of workspaces the user has access to
func (c *GroupsClient) GetGroups(ctx context.Context, options *GetGroupsOptions) (*Groups, error) {
	path := "/groups"
	
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

	var result Groups
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGroup returns a workspace by ID
func (c *GroupsClient) GetGroup(ctx context.Context, groupID string) (*Group, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s", groupID)

	var result Group
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateGroup creates a new workspace
func (c *GroupsClient) CreateGroup(ctx context.Context, request CreateGroupRequest) (*Group, error) {
	path := "/groups"

	var result Group
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteGroup deletes a workspace by ID
func (c *GroupsClient) DeleteGroup(ctx context.Context, groupID string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s", groupID)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// UpdateGroup updates a workspace
func (c *GroupsClient) UpdateGroup(ctx context.Context, groupID string, request UpdateGroupRequest) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s", groupID)
	return c.client.doRequest(ctx, "PATCH", path, request, nil)
}

// GetGroupUsers returns a list of users that have access to the specified workspace
func (c *GroupsClient) GetGroupUsers(ctx context.Context, groupID string) ([]GroupUser, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/users", groupID)

	var result struct {
		Value []GroupUser `json:"value"`
	}
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}

// AddGroupUser grants user permissions to the specified workspace
func (c *GroupsClient) AddGroupUser(ctx context.Context, groupID string, request GroupUserAccessRight) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/users", groupID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// DeleteGroupUser removes user permissions from the specified workspace
func (c *GroupsClient) DeleteGroupUser(ctx context.Context, groupID, userID string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if userID == "" {
		return fmt.Errorf("userID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/users/%s", groupID, userID)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// UpdateGroupUser updates user permissions for the specified workspace
func (c *GroupsClient) UpdateGroupUser(ctx context.Context, groupID, userID string, request GroupUserAccessRight) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if userID == "" {
		return fmt.Errorf("userID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/users/%s", groupID, userID)
	return c.client.doRequest(ctx, "PUT", path, request, nil)
}

// RestoreGroup restores a deleted workspace
func (c *GroupsClient) RestoreGroup(ctx context.Context, groupID string, request RestoreGroupRequest) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/restore", groupID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// AssignToCapacity assigns the specified workspace to the specified capacity
func (c *GroupsClient) AssignToCapacity(ctx context.Context, groupID string, request AssignToCapacityRequest) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/AssignToCapacity", groupID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// UnassignFromCapacity unassigns the specified workspace from capacity
func (c *GroupsClient) UnassignFromCapacity(ctx context.Context, groupID string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/UnassignFromCapacity", groupID)
	return c.client.doRequest(ctx, "POST", path, nil, nil)
}

// CreateGroupRequest represents a request to create a workspace
type CreateGroupRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateGroupRequest represents a request to update a workspace
type UpdateGroupRequest struct {
	Name                       *string `json:"name,omitempty"`
	Description                *string `json:"description,omitempty"`
	DefaultDatasetStorageFormat *string `json:"defaultDatasetStorageFormat,omitempty"`
}

// GroupUserAccessRight represents user access rights for a workspace
type GroupUserAccessRight struct {
	EmailAddress         *string `json:"emailAddress,omitempty"`
	Identifier           *string `json:"identifier,omitempty"`
	PrincipalType        *string `json:"principalType,omitempty"`
	GroupUserAccessRight *string `json:"groupUserAccessRight,omitempty"`
	DisplayName          *string `json:"displayName,omitempty"`
	Profile              *ServicePrincipalProfile `json:"profile,omitempty"`
}

// RestoreGroupRequest represents a request to restore a workspace
type RestoreGroupRequest struct {
	EmailAddress *string `json:"emailAddress,omitempty"`
	Name         *string `json:"name,omitempty"`
}

// AssignToCapacityRequest represents a request to assign a workspace to capacity
type AssignToCapacityRequest struct {
	CapacityID *string `json:"capacityId,omitempty"`
}

