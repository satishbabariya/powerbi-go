package powerbi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
)

// uploadFile uploads a file using multipart/form-data
func (c *Client) uploadFile(ctx context.Context, path string, fileContent []byte, filename string, additionalFields map[string]string) error {
	// Create a buffer to write our multipart form data
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add the file
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := part.Write(fileContent); err != nil {
		return fmt.Errorf("failed to write file content: %w", err)
	}

	// Add additional fields
	for key, value := range additionalFields {
		if err := writer.WriteField(key, value); err != nil {
			return fmt.Errorf("failed to write field %s: %w", key, err)
		}
	}

	// Close the writer to finalize the multipart message
	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Build the full URL
	url := c.buildURL(path)

	// Create the HTTP request
	req, err := c.newRequest(ctx, "POST", url)
	if err != nil {
		return err
	}

	// Set the content type with boundary
	req.Raw().Header.Set("Content-Type", writer.FormDataContentType())

	// Set the body
	bodyReader := bytes.NewReader(buf.Bytes())
	if err := req.SetBody(streaming.NopCloser(bodyReader), writer.FormDataContentType()); err != nil {
		return fmt.Errorf("failed to set request body: %w", err)
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

	return nil
}

// uploadFileWithResponse uploads a file and returns the parsed response
func (c *Client) uploadFileWithResponse(ctx context.Context, path string, fileContent []byte, filename string, additionalFields map[string]string, result interface{}) error {
	// Create a buffer to write our multipart form data
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add the file
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := part.Write(fileContent); err != nil {
		return fmt.Errorf("failed to write file content: %w", err)
	}

	// Add additional fields
	for key, value := range additionalFields {
		if err := writer.WriteField(key, value); err != nil {
			return fmt.Errorf("failed to write field %s: %w", key, err)
		}
	}

	// Close the writer
	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Build the full URL
	url := c.buildURL(path)

	// Create the HTTP request
	req, err := c.newRequest(ctx, "POST", url)
	if err != nil {
		return err
	}

	// Set the content type
	req.Raw().Header.Set("Content-Type", writer.FormDataContentType())

	// Set the body
	bodyReader2 := bytes.NewReader(buf.Bytes())
	if err := req.SetBody(streaming.NopCloser(bodyReader2), writer.FormDataContentType()); err != nil {
		return fmt.Errorf("failed to set request body: %w", err)
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

