package powerbi

import (
	"context"
)

// EmbedTokenClient handles operations for Power BI embed tokens
type EmbedTokenClient struct {
	client *Client
}

// GenerateToken generates an embed token for multiple reports, datasets, and workspaces
func (c *EmbedTokenClient) GenerateToken(ctx context.Context, request GenerateTokenRequestV2, options *GenerateTokenOptions) (*EmbedToken, error) {
	path := "/GenerateToken"

	var result EmbedToken
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GenerateTokenOptions contains optional parameters for GenerateToken
type GenerateTokenOptions struct {
	// Additional options can be added here as needed
}

