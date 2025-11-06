package powerbi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
)

// doRequest performs an HTTP request and handles the response
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	url := c.buildURL(path)
	
	req, err := c.newRequest(ctx, method, url)
	if err != nil {
		return err
	}

	// Add body if provided
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		if err := req.SetBody(streaming.NopCloser(bytes.NewReader(bodyBytes)), "application/json"); err != nil {
			return fmt.Errorf("failed to set request body: %w", err)
		}
	}

	// Send the request
	resp, err := c.pipeline.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check for errors
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return handleErrorResponse(resp)
	}

	// Parse response if result is provided
	if result != nil && resp.StatusCode != http.StatusNoContent {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		if len(bodyBytes) > 0 {
			if err := json.Unmarshal(bodyBytes, result); err != nil {
				return fmt.Errorf("failed to unmarshal response: %w", err)
			}
		}
	}

	return nil
}

// doRequestRaw performs an HTTP request and returns the raw response
func (c *Client) doRequestRaw(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	url := c.buildURL(path)
	
	req, err := c.newRequest(ctx, method, url)
	if err != nil {
		return nil, err
	}

	// Add body if provided
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		if err := req.SetBody(streaming.NopCloser(bytes.NewReader(bodyBytes)), "application/json"); err != nil {
			return nil, fmt.Errorf("failed to set request body: %w", err)
		}
	}

	// Send the request
	resp, err := c.pipeline.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// Check for errors
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		return nil, handleErrorResponse(resp)
	}

	return resp, nil
}

// handleErrorResponse handles error responses from the API
func handleErrorResponse(resp *http.Response) error {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HTTP %d: failed to read error response: %w", resp.StatusCode, err)
	}

	var errorResp ErrorResponse
	if err := json.Unmarshal(bodyBytes, &errorResp); err != nil {
		// If we can't parse the error response, return the raw body
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(bodyBytes))
	}

	if errorResp.Error != nil {
		return fmt.Errorf("HTTP %d: %s - %s", resp.StatusCode, errorResp.Error.Code, errorResp.Error.Message)
	}

	return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(bodyBytes))
}

// buildQueryParams builds query parameters from a map
func buildQueryParams(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}

	values := url.Values{}
	for k, v := range params {
		if v != "" {
			values.Add(k, v)
		}
	}

	query := values.Encode()
	if query != "" {
		return "?" + query
	}
	return ""
}

// buildPath constructs a path with optional path parameters and query parameters
func buildPath(base string, pathParams map[string]string, queryParams map[string]string) string {
	path := base
	
	// Replace path parameters
	for k, v := range pathParams {
		path = strings.ReplaceAll(path, "{"+k+"}", url.PathEscape(v))
	}

	// Add query parameters
	path += buildQueryParams(queryParams)

	return path
}

// String returns a pointer to a string
func String(s string) *string {
	return &s
}

// Bool returns a pointer to a bool
func Bool(b bool) *bool {
	return &b
}

// Int returns a pointer to an int
func Int(i int) *int {
	return &i
}

// StringValue returns the value of a string pointer or empty string if nil
func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// BoolValue returns the value of a bool pointer or false if nil
func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// IntValue returns the value of an int pointer or 0 if nil
func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}


