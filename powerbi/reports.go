package powerbi

import (
	"context"
	"fmt"
)

// ReportsClient handles operations for Power BI reports
type ReportsClient struct {
	client *Client
}

// GetReportsOptions contains optional parameters for GetReports
type GetReportsOptions struct {
	// Filter OData filter query
	Filter *string
	// Top OData top query - maximum number of items to return
	Top *int
	// Skip OData skip query - number of items to skip
	Skip *int
}

// GetReports returns a list of reports from "My Workspace"
func (c *ReportsClient) GetReports(ctx context.Context, options *GetReportsOptions) (*Reports, error) {
	path := "/reports"
	
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

	var result Reports
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetReportsInGroup returns a list of reports from the specified workspace
func (c *ReportsClient) GetReportsInGroup(ctx context.Context, groupID string, options *GetReportsOptions) (*Reports, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/reports", groupID)
	
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

	var result Reports
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetReport returns the specified report from "My Workspace"
func (c *ReportsClient) GetReport(ctx context.Context, reportID string) (*Report, error) {
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/reports/%s", reportID)

	var result Report
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetReportInGroup returns the specified report from the specified workspace
func (c *ReportsClient) GetReportInGroup(ctx context.Context, groupID, reportID string) (*Report, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/reports/%s", groupID, reportID)

	var result Report
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteReport deletes the specified report from "My Workspace"
func (c *ReportsClient) DeleteReport(ctx context.Context, reportID string) error {
	if reportID == "" {
		return fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/reports/%s", reportID)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// DeleteReportInGroup deletes the specified report from the specified workspace
func (c *ReportsClient) DeleteReportInGroup(ctx context.Context, groupID, reportID string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if reportID == "" {
		return fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/reports/%s", groupID, reportID)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// CloneReport clones the specified report from "My Workspace"
func (c *ReportsClient) CloneReport(ctx context.Context, reportID string, request CloneReportRequest) (*Report, error) {
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/reports/%s/Clone", reportID)

	var result Report
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CloneReportInGroup clones the specified report from the specified workspace
func (c *ReportsClient) CloneReportInGroup(ctx context.Context, groupID, reportID string, request CloneReportRequest) (*Report, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/reports/%s/Clone", groupID, reportID)

	var result Report
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ExportToFile exports the specified report to file
func (c *ReportsClient) ExportToFile(ctx context.Context, reportID string, request ExportToFileRequest) error {
	if reportID == "" {
		return fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/reports/%s/ExportTo", reportID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// ExportToFileInGroup exports the specified report from the specified workspace to file
func (c *ReportsClient) ExportToFileInGroup(ctx context.Context, groupID, reportID string, request ExportToFileRequest) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if reportID == "" {
		return fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/reports/%s/ExportTo", groupID, reportID)
	return c.client.doRequest(ctx, "POST", path, request, nil)
}

// UpdateReportContent updates the content of the specified report
func (c *ReportsClient) UpdateReportContent(ctx context.Context, reportID string, sourceReportID string, sourceType string) (*Report, error) {
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/reports/%s/UpdateReportContent", reportID)
	
	body := map[string]string{
		"sourceReport": sourceReportID,
		"sourceType":   sourceType,
	}

	var result Report
	if err := c.client.doRequest(ctx, "POST", path, body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateReportContentInGroup updates the content of the specified report in the specified workspace
func (c *ReportsClient) UpdateReportContentInGroup(ctx context.Context, groupID, reportID string, sourceReportID string, sourceType string) (*Report, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/reports/%s/UpdateReportContent", groupID, reportID)
	
	body := map[string]string{
		"sourceReport": sourceReportID,
		"sourceType":   sourceType,
	}

	var result Report
	if err := c.client.doRequest(ctx, "POST", path, body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// RebindReport rebinds the specified report to the specified dataset
func (c *ReportsClient) RebindReport(ctx context.Context, reportID, datasetID string) error {
	if reportID == "" {
		return fmt.Errorf("reportID cannot be empty")
	}
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/reports/%s/Rebind", reportID)
	
	body := map[string]string{
		"datasetId": datasetID,
	}

	return c.client.doRequest(ctx, "POST", path, body, nil)
}

// RebindReportInGroup rebinds the specified report to the specified dataset in the specified workspace
func (c *ReportsClient) RebindReportInGroup(ctx context.Context, groupID, reportID, datasetID string) error {
	if groupID == "" {
		return fmt.Errorf("groupID cannot be empty")
	}
	if reportID == "" {
		return fmt.Errorf("reportID cannot be empty")
	}
	if datasetID == "" {
		return fmt.Errorf("datasetID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/reports/%s/Rebind", groupID, reportID)
	
	body := map[string]string{
		"datasetId": datasetID,
	}

	return c.client.doRequest(ctx, "POST", path, body, nil)
}

// GetPages returns a list of pages within the specified report
func (c *ReportsClient) GetPages(ctx context.Context, reportID string) ([]Page, error) {
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/reports/%s/pages", reportID)

	var result struct {
		Value []Page `json:"value"`
	}
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}

// GetPagesInGroup returns a list of pages within the specified report from the specified workspace
func (c *ReportsClient) GetPagesInGroup(ctx context.Context, groupID, reportID string) ([]Page, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if reportID == "" {
		return nil, fmt.Errorf("reportID cannot be empty")
	}

	path := fmt.Sprintf("/groups/%s/reports/%s/pages", groupID, reportID)

	var result struct {
		Value []Page `json:"value"`
	}
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}

