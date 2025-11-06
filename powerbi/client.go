package powerbi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

const (
	// DefaultBaseURL is the default base URL for Power BI REST API
	DefaultBaseURL = "https://api.powerbi.com"
	// DefaultScope is the default scope for Power BI API
	DefaultScope = "https://analysis.windows.net/powerbi/api/.default"
	// APIVersion is the API version
	APIVersion = "v1.0"
)

// ClientOptions contains optional configuration for the PowerBIClient
type ClientOptions struct {
	// BaseURL is the base URL for the Power BI API
	// Default: https://api.powerbi.com
	BaseURL string

	// ServicePrincipalProfile is the service principal profile ID
	ServicePrincipalProfile *string

	// HTTPClient is the HTTP client to use for requests
	HTTPClient *http.Client

	// RetryOptions specifies the retry policy options
	RetryOptions policy.RetryOptions

	// Telemetry options
	Telemetry policy.TelemetryOptions
}

// Client is the main Power BI client
type Client struct {
	baseURL    string
	credential azcore.TokenCredential
	pipeline   runtime.Pipeline

	// Service clients
	Reports                 *ReportsClient
	Datasets                *DatasetsClient
	Dashboards              *DashboardsClient
	Groups                  *GroupsClient
	EmbedToken              *EmbedTokenClient
	Admin                   *AdminClient
	Gateways                *GatewaysClient
	Dataflows               *DataflowsClient
	Capacities              *CapacitiesClient
	Apps                    *AppsClient
	Imports                 *ImportsClient
	Tiles                   *TilesClient
	Users                   *UsersClient
	Pipelines               *PipelinesClient
	AvailableFeatures       *AvailableFeaturesClient
	Profiles                *ProfilesClient
	Scorecards              *ScorecardsClient
	Goals                   *GoalsClient
	GoalsStatusRules        *GoalsStatusRulesClient
	GoalValues              *GoalValuesClient
	GoalNotes               *GoalNotesClient
	TemplateApps            *TemplateAppsClient
	DataflowStorageAccounts *DataflowStorageAccountsClient
	WorkspaceInfo           *WorkspaceInfoClient
	WidelySharedArtifacts   *WidelySharedArtifactsClient
	InformationProtection   *InformationProtectionClient
	PushDatasets            *PushDatasetsClient
	Subscriptions           *SubscriptionsClient
	Bookmarks               *BookmarksClient
}

// NewClient creates a new Power BI client with Azure credentials
func NewClient(credential azcore.TokenCredential, options *ClientOptions) (*Client, error) {
	if credential == nil {
		return nil, fmt.Errorf("credential cannot be nil")
	}

	if options == nil {
		options = &ClientOptions{}
	}

	baseURL := options.BaseURL
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	// Create pipeline options
	pipelineOptions := runtime.PipelineOptions{
		PerRetry: []policy.Policy{},
	}

	// Add service principal profile header if specified
	if options.ServicePrincipalProfile != nil {
		pipelineOptions.PerRetry = append(pipelineOptions.PerRetry,
			&servicePrincipalProfilePolicy{profileID: *options.ServicePrincipalProfile})
	}

	// Create the pipeline
	pipeline := runtime.NewPipeline("powerbi-go", "1.0.0", pipelineOptions, &policy.ClientOptions{
		Transport:        options.HTTPClient,
		Retry:            options.RetryOptions,
		Telemetry:        options.Telemetry,
		PerCallPolicies:  []policy.Policy{},
		PerRetryPolicies: pipelineOptions.PerRetry,
	})

	client := &Client{
		baseURL:    baseURL,
		credential: credential,
		pipeline:   pipeline,
	}

	// Initialize all service clients
	client.Reports = &ReportsClient{client: client}
	client.Datasets = &DatasetsClient{client: client}
	client.Dashboards = &DashboardsClient{client: client}
	client.Groups = &GroupsClient{client: client}
	client.EmbedToken = &EmbedTokenClient{client: client}
	client.Admin = &AdminClient{client: client}
	client.Gateways = &GatewaysClient{client: client}
	client.Dataflows = &DataflowsClient{client: client}
	client.Capacities = &CapacitiesClient{client: client}
	client.Apps = &AppsClient{client: client}
	client.Imports = &ImportsClient{client: client}
	client.Tiles = &TilesClient{client: client}
	client.Users = &UsersClient{client: client}
	client.Pipelines = &PipelinesClient{client: client}
	client.AvailableFeatures = &AvailableFeaturesClient{client: client}
	client.Profiles = &ProfilesClient{client: client}
	client.Scorecards = &ScorecardsClient{client: client}
	client.Goals = &GoalsClient{client: client}
	client.GoalsStatusRules = &GoalsStatusRulesClient{client: client}
	client.GoalValues = &GoalValuesClient{client: client}
	client.GoalNotes = &GoalNotesClient{client: client}
	client.TemplateApps = &TemplateAppsClient{client: client}
	client.DataflowStorageAccounts = &DataflowStorageAccountsClient{client: client}
	client.WorkspaceInfo = &WorkspaceInfoClient{client: client}
	client.WidelySharedArtifacts = &WidelySharedArtifactsClient{client: client}
	client.InformationProtection = &InformationProtectionClient{client: client}
	client.PushDatasets = &PushDatasetsClient{client: client}
	client.Subscriptions = &SubscriptionsClient{client: client}
	client.Bookmarks = &BookmarksClient{client: client}

	return client, nil
}

// NewClientWithToken creates a new Power BI client with an access token
func NewClientWithToken(token string, options *ClientOptions) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("token cannot be empty")
	}

	credential := &staticTokenCredential{token: token}
	return NewClient(credential, options)
}

// staticTokenCredential implements azcore.TokenCredential for a static token
type staticTokenCredential struct {
	token string
}

// GetToken returns the static token
func (c *staticTokenCredential) GetToken(ctx context.Context, opts policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{
		Token:     c.token,
		ExpiresOn: time.Now().Add(1 * time.Hour), // Assume token is valid for 1 hour
	}, nil
}

// servicePrincipalProfilePolicy adds the X-PowerBI-Profile-Id header
type servicePrincipalProfilePolicy struct {
	profileID string
}

// Do implements the policy.Policy interface
func (p *servicePrincipalProfilePolicy) Do(req *policy.Request) (*http.Response, error) {
	req.Raw().Header.Set("X-PowerBI-Profile-Id", p.profileID)
	return req.Next()
}

// buildURL constructs a full URL from the base URL and path
func (c *Client) buildURL(path string) string {
	return fmt.Sprintf("%s/%s/myorg%s", c.baseURL, APIVersion, path)
}

// newRequest creates a new HTTP request with authentication
func (c *Client) newRequest(ctx context.Context, method, url string) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, method, url)
	if err != nil {
		return nil, err
	}

	// Get access token
	token, err := c.credential.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: []string{DefaultScope},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	// Add authorization header
	req.Raw().Header.Set("Authorization", "Bearer "+token.Token)
	req.Raw().Header.Set("Content-Type", "application/json")

	return req, nil
}
