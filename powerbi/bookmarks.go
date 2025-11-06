package powerbi

import (
	"context"
	"fmt"
)

// BookmarksClient handles operations for Power BI bookmarks
type BookmarksClient struct {
	client *Client
}

// Bookmark represents a Power BI bookmark
type Bookmark struct {
	Name        *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	State       *string `json:"state,omitempty"`
	ReportID    *string `json:"reportId,omitempty"`
}

// Bookmarks represents a list of bookmarks
type Bookmarks struct {
	Value []Bookmark `json:"value"`
}

// GetBookmarks returns a list of bookmarks from the specified report in "My Workspace"
func (c *BookmarksClient) GetBookmarks(ctx context.Context, reportID string) (*Bookmarks, error) {
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/reports/%s/bookmarks", reportID)

	var result Bookmarks
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetBookmarksInGroup returns a list of bookmarks from the specified report in the specified workspace
func (c *BookmarksClient) GetBookmarksInGroup(ctx context.Context, groupID, reportID string) (*Bookmarks, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/reports/%s/bookmarks", groupID, reportID)

	var result Bookmarks
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CaptureBookmark captures the current state of a report as a bookmark
func (c *BookmarksClient) CaptureBookmark(ctx context.Context, reportID string, request CaptureBookmarkRequest) (*Bookmark, error) {
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/reports/%s/bookmarks/capture", reportID)

	var result Bookmark
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CaptureBookmarkInGroup captures the current state of a report as a bookmark in the specified workspace
func (c *BookmarksClient) CaptureBookmarkInGroup(ctx context.Context, groupID, reportID string, request CaptureBookmarkRequest) (*Bookmark, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/reports/%s/bookmarks/capture", groupID, reportID)

	var result Bookmark
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CaptureBookmarkRequest represents a request to capture a bookmark
type CaptureBookmarkRequest struct {
	Name        *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}

