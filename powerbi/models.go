package powerbi

import "time"

// Common response types

// ODataResponse is the base response type for OData responses
type ODataResponse struct {
	ODataContext string `json:"@odata.context,omitempty"`
	ODataCount   *int   `json:"@odata.count,omitempty"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Error *ErrorDetail `json:"error,omitempty"`
}

// ErrorDetail contains error details
type ErrorDetail struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Target  string `json:"target,omitempty"`
}

// Report represents a Power BI report
type Report struct {
	ID                       *string    `json:"id,omitempty"`
	ReportType               *string    `json:"reportType,omitempty"`
	Name                     *string    `json:"name,omitempty"`
	WebURL                   *string    `json:"webUrl,omitempty"`
	EmbedURL                 *string    `json:"embedUrl,omitempty"`
	IsFromPbix               *bool      `json:"isFromPbix,omitempty"`
	IsOwnedByMe              *bool      `json:"isOwnedByMe,omitempty"`
	DatasetID                *string    `json:"datasetId,omitempty"`
	DatasetWorkspaceID       *string    `json:"datasetWorkspaceId,omitempty"`
	Users                    []User     `json:"users,omitempty"`
	SubscriptionEnabled      *bool      `json:"subscriptionEnabled,omitempty"`
	ModifiedBy               *string    `json:"modifiedBy,omitempty"`
	ModifiedDateTime         *time.Time `json:"modifiedDateTime,omitempty"`
	CreatedDateTime          *time.Time `json:"createdDateTime,omitempty"`
	AppID                    *string    `json:"appId,omitempty"`
	Description              *string    `json:"description,omitempty"`
	AutoSave                 *bool      `json:"autoSave,omitempty"`
	ReportPages              []Page     `json:"reportPages,omitempty"`
	SensitivityLabel         *SensitivityLabel `json:"sensitivityLabel,omitempty"`
}

// Reports represents a list of reports
type Reports struct {
	ODataContext string   `json:"@odata.context,omitempty"`
	Value        []Report `json:"value"`
}

// Dataset represents a Power BI dataset
type Dataset struct {
	ID                                *string    `json:"id,omitempty"`
	Name                              *string    `json:"name,omitempty"`
	AddRowsAPIEnabled                 *bool      `json:"addRowsAPIEnabled,omitempty"`
	ConfiguredBy                      *string    `json:"configuredBy,omitempty"`
	IsRefreshable                     *bool      `json:"isRefreshable,omitempty"`
	IsEffectiveIdentityRequired       *bool      `json:"isEffectiveIdentityRequired,omitempty"`
	IsEffectiveIdentityRolesRequired  *bool      `json:"isEffectiveIdentityRolesRequired,omitempty"`
	IsOnPremGatewayRequired           *bool      `json:"isOnPremGatewayRequired,omitempty"`
	TargetStorageMode                 *string    `json:"targetStorageMode,omitempty"`
	CreatedDate                       *time.Time `json:"createdDate,omitempty"`
	CreateReportEmbedURL              *string    `json:"createReportEmbedURL,omitempty"`
	QnaEmbedURL                       *string    `json:"qnaEmbedURL,omitempty"`
	UpstreamDatasets                  []Dataset  `json:"upstreamDatasets,omitempty"`
	Users                             []User     `json:"users,omitempty"`
	WebURL                            *string    `json:"webUrl,omitempty"`
	Description                       *string    `json:"description,omitempty"`
	ContentProviderType               *string    `json:"contentProviderType,omitempty"`
	AutoSyncReadOnlyReplicas          *bool      `json:"autoSyncReadOnlyReplicas,omitempty"`
	QueryScaleOutSettings             *QueryScaleOutSettings `json:"queryScaleOutSettings,omitempty"`
	SensitivityLabel                  *SensitivityLabel `json:"sensitivityLabel,omitempty"`
}

// Datasets represents a list of datasets
type Datasets struct {
	ODataContext string    `json:"@odata.context,omitempty"`
	Value        []Dataset `json:"value"`
}

// Dashboard represents a Power BI dashboard
type Dashboard struct {
	ID                   *string    `json:"id,omitempty"`
	DisplayName          *string    `json:"displayName,omitempty"`
	EmbedURL             *string    `json:"embedUrl,omitempty"`
	IsReadOnly           *bool      `json:"isReadOnly,omitempty"`
	WebURL               *string    `json:"webUrl,omitempty"`
	Tiles                []Tile     `json:"tiles,omitempty"`
	Users                []User     `json:"users,omitempty"`
	SubscriptionEnabled  *bool      `json:"subscriptionEnabled,omitempty"`
	AppID                *string    `json:"appId,omitempty"`
	SensitivityLabel     *SensitivityLabel `json:"sensitivityLabel,omitempty"`
}

// Dashboards represents a list of dashboards
type Dashboards struct {
	ODataContext string      `json:"@odata.context,omitempty"`
	Value        []Dashboard `json:"value"`
}

// Tile represents a dashboard tile
type Tile struct {
	ID               *string `json:"id,omitempty"`
	Title            *string `json:"title,omitempty"`
	SubTitle         *string `json:"subTitle,omitempty"`
	EmbedURL         *string `json:"embedUrl,omitempty"`
	RowSpan          *int    `json:"rowSpan,omitempty"`
	ColSpan          *int    `json:"colSpan,omitempty"`
	ReportID         *string `json:"reportId,omitempty"`
	DatasetID        *string `json:"datasetId,omitempty"`
}

// Group represents a Power BI workspace (group)
type Group struct {
	ID                         *string    `json:"id,omitempty"`
	IsReadOnly                 *bool      `json:"isReadOnly,omitempty"`
	IsOnDedicatedCapacity      *bool      `json:"isOnDedicatedCapacity,omitempty"`
	CapacityID                 *string    `json:"capacityId,omitempty"`
	DefaultDatasetStorageFormat *string   `json:"defaultDatasetStorageFormat,omitempty"`
	Type                       *string    `json:"type,omitempty"`
	Name                       *string    `json:"name,omitempty"`
	Description                *string    `json:"description,omitempty"`
	CapacityMigrationStatus    *string    `json:"capacityMigrationStatus,omitempty"`
	State                      *string    `json:"state,omitempty"`
	HasWorkspaceLevelSettings  *bool      `json:"hasWorkspaceLevelSettings,omitempty"`
	PipelineID                 *string    `json:"pipelineId,omitempty"`
	Users                      []GroupUser `json:"users,omitempty"`
	LogAnalyticsWorkspace      *LogAnalyticsWorkspace `json:"logAnalyticsWorkspace,omitempty"`
}

// Groups represents a list of groups
type Groups struct {
	ODataContext string  `json:"@odata.context,omitempty"`
	Value        []Group `json:"value"`
}

// GroupUser represents a user in a workspace
type GroupUser struct {
	EmailAddress       *string `json:"emailAddress,omitempty"`
	DisplayName        *string `json:"displayName,omitempty"`
	GroupUserAccessRight *string `json:"groupUserAccessRight,omitempty"`
	Identifier         *string `json:"identifier,omitempty"`
	PrincipalType      *string `json:"principalType,omitempty"`
	Profile            *ServicePrincipalProfile `json:"profile,omitempty"`
}

// User represents a user
type User struct {
	EmailAddress       *string `json:"emailAddress,omitempty"`
	DisplayName        *string `json:"displayName,omitempty"`
	Identifier         *string `json:"identifier,omitempty"`
	PrincipalType      *string `json:"principalType,omitempty"`
	UserType           *string `json:"userType,omitempty"`
	GraphID            *string `json:"graphId,omitempty"`
	Profile            *ServicePrincipalProfile `json:"profile,omitempty"`
}

// Page represents a report page
type Page struct {
	Name        *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Order       *int    `json:"order,omitempty"`
}

// EmbedToken represents an embed token
type EmbedToken struct {
	Token      *string    `json:"token,omitempty"`
	TokenID    *string    `json:"tokenId,omitempty"`
	Expiration *time.Time `json:"expiration,omitempty"`
}

// GenerateTokenRequestV2 represents a request to generate an embed token
type GenerateTokenRequestV2 struct {
	Datasets         []GenerateTokenRequestV2Dataset         `json:"datasets,omitempty"`
	Reports          []GenerateTokenRequestV2Report          `json:"reports,omitempty"`
	TargetWorkspaces []GenerateTokenRequestV2TargetWorkspace `json:"targetWorkspaces,omitempty"`
	Identities       []EffectiveIdentity                     `json:"identities,omitempty"`
	LifetimeInMinutes *int                                   `json:"lifetimeInMinutes,omitempty"`
}

// GenerateTokenRequestV2Dataset represents a dataset in an embed token request
type GenerateTokenRequestV2Dataset struct {
	ID         *string   `json:"id,omitempty"`
	XMLAPermissions *string `json:"xmlaPermissions,omitempty"`
}

// GenerateTokenRequestV2Report represents a report in an embed token request
type GenerateTokenRequestV2Report struct {
	ID               *string   `json:"id,omitempty"`
	AllowEdit        *bool     `json:"allowEdit,omitempty"`
	AllowSaveAs      *bool     `json:"allowSaveAs,omitempty"`
}

// GenerateTokenRequestV2TargetWorkspace represents a target workspace in an embed token request
type GenerateTokenRequestV2TargetWorkspace struct {
	ID *string `json:"id,omitempty"`
}

// EffectiveIdentity represents an effective identity
type EffectiveIdentity struct {
	Username          *string   `json:"username,omitempty"`
	Roles             []string  `json:"roles,omitempty"`
	CustomData        *string   `json:"customData,omitempty"`
	Datasets          []string  `json:"datasets,omitempty"`
	IdentityBlob      *string   `json:"identityBlob,omitempty"`
	ReportID          *string   `json:"reportId,omitempty"`
	TargetWorkspaceID *string   `json:"targetWorkspaceId,omitempty"`
}

// SensitivityLabel represents a sensitivity label
type SensitivityLabel struct {
	ID     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
}

// QueryScaleOutSettings represents query scale-out settings
type QueryScaleOutSettings struct {
	MaxReadOnlyReplicas *int  `json:"maxReadOnlyReplicas,omitempty"`
	AutoSyncReadOnlyReplicas *bool `json:"autoSyncReadOnlyReplicas,omitempty"`
}

// ServicePrincipalProfile represents a service principal profile
type ServicePrincipalProfile struct {
	ID          *string `json:"id,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}

// LogAnalyticsWorkspace represents a Log Analytics workspace
type LogAnalyticsWorkspace struct {
	ID                 *string `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	ResourceGroup      *string `json:"resourceGroup,omitempty"`
	SubscriptionID     *string `json:"subscriptionId,omitempty"`
}

// Import represents a Power BI import
type Import struct {
	ID               *string    `json:"id,omitempty"`
	ImportState      *string    `json:"importState,omitempty"`
	CreatedDateTime  *time.Time `json:"createdDateTime,omitempty"`
	UpdatedDateTime  *time.Time `json:"updatedDateTime,omitempty"`
	Name             *string    `json:"name,omitempty"`
	ConnectionType   *string    `json:"connectionType,omitempty"`
	Source           *string    `json:"source,omitempty"`
	Datasets         []Dataset  `json:"datasets,omitempty"`
	Reports          []Report   `json:"reports,omitempty"`
}

// Imports represents a list of imports
type Imports struct {
	ODataContext string   `json:"@odata.context,omitempty"`
	Value        []Import `json:"value"`
}

// Gateway represents an on-premises data gateway
type Gateway struct {
	ID                         *string            `json:"id,omitempty"`
	Name                       *string            `json:"name,omitempty"`
	Type                       *string            `json:"type,omitempty"`
	PublicKey                  *GatewayPublicKey  `json:"publicKey,omitempty"`
	GatewayAnnotation          *string            `json:"gatewayAnnotation,omitempty"`
	GatewayStatus              *string            `json:"gatewayStatus,omitempty"`
	ExternalIPAddress          *string            `json:"externalIpAddress,omitempty"`
	VersionedGatewayVersion    *string            `json:"versionedGatewayVersion,omitempty"`
}

// GatewayPublicKey represents a gateway public key
type GatewayPublicKey struct {
	Exponent *string `json:"exponent,omitempty"`
	Modulus  *string `json:"modulus,omitempty"`
}

// Gateways represents a list of gateways
type Gateways struct {
	ODataContext string    `json:"@odata.context,omitempty"`
	Value        []Gateway `json:"value"`
}

// Dataflow represents a Power BI dataflow
type Dataflow struct {
	ObjectID            *string    `json:"objectId,omitempty"`
	Name                *string    `json:"name,omitempty"`
	Description         *string    `json:"description,omitempty"`
	ConfiguredBy        *string    `json:"configuredBy,omitempty"`
	ModifiedBy          *string    `json:"modifiedBy,omitempty"`
	ModifiedDateTime    *time.Time `json:"modifiedDateTime,omitempty"`
	Users               []User     `json:"users,omitempty"`
}

// Dataflows represents a list of dataflows
type Dataflows struct {
	ODataContext string     `json:"@odata.context,omitempty"`
	Value        []Dataflow `json:"value"`
}

// Capacity represents a Power BI capacity
type Capacity struct {
	ID               *string           `json:"id,omitempty"`
	DisplayName      *string           `json:"displayName,omitempty"`
	Admins           []CapacityAdmin   `json:"admins,omitempty"`
	SKU              *string           `json:"sku,omitempty"`
	State            *string           `json:"state,omitempty"`
	CapacityUserAccessRight *string    `json:"capacityUserAccessRight,omitempty"`
	Region           *string           `json:"region,omitempty"`
}

// CapacityAdmin represents a capacity admin
type CapacityAdmin struct {
	EmailAddress  *string `json:"emailAddress,omitempty"`
	DisplayName   *string `json:"displayName,omitempty"`
	Identifier    *string `json:"identifier,omitempty"`
	PrincipalType *string `json:"principalType,omitempty"`
}

// Capacities represents a list of capacities
type Capacities struct {
	ODataContext string     `json:"@odata.context,omitempty"`
	Value        []Capacity `json:"value"`
}

// App represents a Power BI app
type App struct {
	ID               *string `json:"id,omitempty"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	PublishedBy      *string `json:"publishedBy,omitempty"`
	LastUpdate       *time.Time `json:"lastUpdate,omitempty"`
	WorkspaceID      *string `json:"workspaceId,omitempty"`
}

// Apps represents a list of apps
type Apps struct {
	ODataContext string `json:"@odata.context,omitempty"`
	Value        []App  `json:"value"`
}

// RefreshRequest represents a dataset refresh request
type RefreshRequest struct {
	NotifyOption          *string `json:"notifyOption,omitempty"`
	RetryCount            *int    `json:"retryCount,omitempty"`
	Type                  *string `json:"type,omitempty"`
	CommitMode            *string `json:"commitMode,omitempty"`
	MaxParallelism        *int    `json:"maxParallelism,omitempty"`
	Objects               []RefreshObject `json:"objects,omitempty"`
	ApplyRefreshPolicy    *bool   `json:"applyRefreshPolicy,omitempty"`
}

// RefreshObject represents an object to refresh
type RefreshObject struct {
	Table     *string `json:"table,omitempty"`
	Partition *string `json:"partition,omitempty"`
}

// CloneReportRequest represents a request to clone a report
type CloneReportRequest struct {
	Name            *string `json:"name,omitempty"`
	TargetModelID   *string `json:"targetModelId,omitempty"`
	TargetWorkspaceID *string `json:"targetWorkspaceId,omitempty"`
}

// ExportToFileRequest represents a request to export a report to file
type ExportToFileRequest struct {
	Format       *string           `json:"format,omitempty"`
	PowerBIReportConfiguration *PowerBIReportConfiguration `json:"powerBIReportConfiguration,omitempty"`
}

// PowerBIReportConfiguration represents report export configuration
type PowerBIReportConfiguration struct {
	Pages             []PageExportConfiguration `json:"pages,omitempty"`
	DefaultBookmark   *BookmarkExportConfiguration `json:"defaultBookmark,omitempty"`
	Identities        []EffectiveIdentity `json:"identities,omitempty"`
	ReportLevelFilters []ExportFilter `json:"reportLevelFilters,omitempty"`
}

// PageExportConfiguration represents page export configuration
type PageExportConfiguration struct {
	PageName          *string        `json:"pageName,omitempty"`
	VisualName        *string        `json:"visualName,omitempty"`
	Bookmark          *BookmarkExportConfiguration `json:"bookmark,omitempty"`
	PageLevelFilters  []ExportFilter `json:"pageLevelFilters,omitempty"`
}

// BookmarkExportConfiguration represents bookmark export configuration
type BookmarkExportConfiguration struct {
	Name  *string `json:"name,omitempty"`
	State *string `json:"state,omitempty"`
}

// ExportFilter represents an export filter
type ExportFilter struct {
	Filter *string `json:"filter,omitempty"`
}

