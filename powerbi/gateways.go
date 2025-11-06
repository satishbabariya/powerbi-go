package powerbi

import (
	"context"
	"fmt"
)

// GatewaysClient handles operations for Power BI gateways
type GatewaysClient struct {
	client *Client
}

// GetGateways returns a list of gateways
func (c *GatewaysClient) GetGateways(ctx context.Context) (*Gateways, error) {
	path := "/gateways"

	var result Gateways
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGateway returns the specified gateway
func (c *GatewaysClient) GetGateway(ctx context.Context, gatewayID string) (*Gateway, error) {
	if gatewayID == "" {
		return nil, fmt.Errorf("gatewayID cannot be empty")
	}

	path := fmt.Sprintf("/gateways/%s", gatewayID)

	var result Gateway
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGatewayDatasources returns a list of datasources from the specified gateway
func (c *GatewaysClient) GetGatewayDatasources(ctx context.Context, gatewayID string) (*GatewayDatasources, error) {
	if gatewayID == "" {
		return nil, fmt.Errorf("gatewayID cannot be empty")
	}

	path := fmt.Sprintf("/gateways/%s/datasources", gatewayID)

	var result GatewayDatasources
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGatewayDatasource returns the specified datasource from the specified gateway
func (c *GatewaysClient) GetGatewayDatasource(ctx context.Context, gatewayID, datasourceID string) (*GatewayDatasource, error) {
	if gatewayID == "" {
		return nil, fmt.Errorf("gatewayID cannot be empty")
	}
	if datasourceID == "" {
		return nil, fmt.Errorf("datasourceID cannot be empty")
	}

	path := fmt.Sprintf("/gateways/%s/datasources/%s", gatewayID, datasourceID)

	var result GatewayDatasource
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateDatasource creates a new datasource on the specified gateway
func (c *GatewaysClient) CreateDatasource(ctx context.Context, gatewayID string, request CreateDatasourceRequest) (*GatewayDatasource, error) {
	if gatewayID == "" {
		return nil, fmt.Errorf("gatewayID cannot be empty")
	}

	path := fmt.Sprintf("/gateways/%s/datasources", gatewayID)

	var result GatewayDatasource
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteDatasource deletes the specified datasource from the specified gateway
func (c *GatewaysClient) DeleteDatasource(ctx context.Context, gatewayID, datasourceID string) error {
	if gatewayID == "" {
		return fmt.Errorf("gatewayID cannot be empty")
	}
	if datasourceID == "" {
		return fmt.Errorf("datasourceID cannot be empty")
	}

	path := fmt.Sprintf("/gateways/%s/datasources/%s", gatewayID, datasourceID)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// UpdateDatasource updates the specified datasource credentials
func (c *GatewaysClient) UpdateDatasource(ctx context.Context, gatewayID, datasourceID string, request UpdateDatasourceRequest) error {
	if gatewayID == "" {
		return fmt.Errorf("gatewayID cannot be empty")
	}
	if datasourceID == "" {
		return fmt.Errorf("datasourceID cannot be empty")
	}

	path := fmt.Sprintf("/gateways/%s/datasources/%s", gatewayID, datasourceID)
	return c.client.doRequest(ctx, "PATCH", path, request, nil)
}

// GatewayDatasources represents a list of gateway datasources
type GatewayDatasources struct {
	ODataContext string               `json:"@odata.context,omitempty"`
	Value        []GatewayDatasource  `json:"value"`
}

// GatewayDatasource represents a gateway datasource
type GatewayDatasource struct {
	ID                 *string            `json:"id,omitempty"`
	GatewayID          *string            `json:"gatewayId,omitempty"`
	DatasourceName     *string            `json:"datasourceName,omitempty"`
	DatasourceType     *string            `json:"datasourceType,omitempty"`
	ConnectionDetails  *string            `json:"connectionDetails,omitempty"`
	CredentialType     *string            `json:"credentialType,omitempty"`
	CredentialDetails  *CredentialDetails `json:"credentialDetails,omitempty"`
}

// CredentialDetails represents credential details
type CredentialDetails struct {
	CredentialType    *string `json:"credentialType,omitempty"`
	Credentials       *string `json:"credentials,omitempty"`
	EncryptedConnection *string `json:"encryptedConnection,omitempty"`
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty"`
	PrivacyLevel      *string `json:"privacyLevel,omitempty"`
}

// CreateDatasourceRequest represents a request to create a datasource
type CreateDatasourceRequest struct {
	DatasourceName    *string            `json:"datasourceName,omitempty"`
	DatasourceType    *string            `json:"datasourceType,omitempty"`
	ConnectionDetails *string            `json:"connectionDetails,omitempty"`
	CredentialDetails *CredentialDetails `json:"credentialDetails,omitempty"`
}

// UpdateDatasourceRequest represents a request to update a datasource
type UpdateDatasourceRequest struct {
	CredentialDetails *CredentialDetails `json:"credentialDetails,omitempty"`
}

