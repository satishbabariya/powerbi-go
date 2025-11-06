# Power BI Go SDK Examples

This directory contains comprehensive examples demonstrating how to use the Power BI Go SDK.

## Prerequisites

Before running these examples, you need:

1. **Azure Active Directory App Registration**
   - Go to [Azure Portal](https://portal.azure.com)
   - Register a new application in Azure Active Directory
   - Grant necessary Power BI API permissions
   - Create a client secret

2. **Environment Variables**
   Set the following environment variables:
   ```bash
   export AZURE_TENANT_ID="your-tenant-id"
   export AZURE_CLIENT_ID="your-client-id"
   export AZURE_CLIENT_SECRET="your-client-secret"
   ```

3. **For specific examples, also set:**
   ```bash
   export POWERBI_WORKSPACE_ID="your-workspace-id"
   export POWERBI_REPORT_ID="your-report-id"
   export POWERBI_DATASET_ID="your-dataset-id"
   ```

## Examples

### 1. Basic Usage (`basic/`)

Demonstrates basic operations:
- Authenticating with Azure credentials
- Listing workspaces
- Getting reports and datasets from a workspace

**Run:**
```bash
cd basic
go run main.go
```

### 2. Embed Token Generation (`embed-token/`)

Shows how to:
- Generate embed tokens for reports and datasets
- Get embed URLs for reports
- Configure embed token parameters

**Run:**
```bash
cd embed-token
go run main.go
```

**Required environment variables:**
- `POWERBI_REPORT_ID`
- `POWERBI_DATASET_ID`
- `POWERBI_WORKSPACE_ID`

### 3. Admin Operations (`admin-operations/`)

Demonstrates administrative operations (requires admin privileges):
- Getting all reports in the organization
- Getting all datasets in the organization
- Getting all workspaces in the organization
- Getting capacity information

**Run:**
```bash
cd admin-operations
go run main.go
```

**Note:** The service principal must have Power BI admin rights.

### 4. Dataset Refresh (`dataset-refresh/`)

Shows how to:
- Get dataset details
- View refresh history
- Trigger dataset refresh
- Monitor refresh status

**Run:**
```bash
cd dataset-refresh
go run main.go
```

**Required environment variables:**
- `POWERBI_WORKSPACE_ID`
- `POWERBI_DATASET_ID`

### 5. Workspace Management (`workspace-management/`)

Demonstrates workspace operations:
- Listing all workspaces
- Creating a new workspace
- Getting workspace users
- Adding users to a workspace
- Updating workspace properties
- Deleting a workspace

**Run:**
```bash
cd workspace-management
go run main.go
```

## Authentication Methods

The SDK supports multiple authentication methods:

### 1. Service Principal (Client Secret)

```go
cred, err := azidentity.NewClientSecretCredential(
    tenantID,
    clientID,
    clientSecret,
    nil,
)
client, err := powerbi.NewClient(cred, nil)
```

### 2. Service Principal (Certificate)

```go
certData, err := os.ReadFile("path/to/cert.pem")
certs, key, err := azidentity.ParseCertificates(certData, nil)

cred, err := azidentity.NewClientCertificateCredential(
    tenantID,
    clientID,
    certs,
    key,
    nil,
)
client, err := powerbi.NewClient(cred, nil)
```

### 3. Managed Identity

```go
cred, err := azidentity.NewManagedIdentityCredential(nil)
client, err := powerbi.NewClient(cred, nil)
```

### 4. Static Access Token

```go
token := "your-access-token"
client, err := powerbi.NewClientWithToken(token, nil)
```

### 5. Default Azure Credential (tries multiple methods)

```go
cred, err := azidentity.NewDefaultAzureCredential(nil)
client, err := powerbi.NewClient(cred, nil)
```

## Common Operations

### Working with Workspaces

```go
// List workspaces
groups, err := client.Groups.GetGroups(ctx, nil)

// Get specific workspace
group, err := client.Groups.GetGroup(ctx, workspaceID)

// Create workspace
createReq := powerbi.CreateGroupRequest{
    Name: powerbi.String("My Workspace"),
}
group, err := client.Groups.CreateGroup(ctx, createReq)
```

### Working with Reports

```go
// List reports in workspace
reports, err := client.Reports.GetReportsInGroup(ctx, workspaceID, nil)

// Get specific report
report, err := client.Reports.GetReportInGroup(ctx, workspaceID, reportID)

// Clone report
cloneReq := powerbi.CloneReportRequest{
    Name:              powerbi.String("Cloned Report"),
    TargetWorkspaceID: powerbi.String(targetWorkspaceID),
}
newReport, err := client.Reports.CloneReportInGroup(ctx, workspaceID, reportID, cloneReq)
```

### Working with Datasets

```go
// List datasets
datasets, err := client.Datasets.GetDatasetsInGroup(ctx, workspaceID, nil)

// Trigger refresh
refreshReq := &powerbi.RefreshRequest{
    NotifyOption: powerbi.String("MailOnFailure"),
}
err := client.Datasets.RefreshDatasetInGroup(ctx, workspaceID, datasetID, refreshReq)

// Get refresh history
history, err := client.Datasets.GetRefreshHistoryInGroup(ctx, workspaceID, datasetID, nil)
```

### Generating Embed Tokens

```go
request := powerbi.GenerateTokenRequestV2{
    Datasets: []powerbi.GenerateTokenRequestV2Dataset{
        {ID: &datasetID},
    },
    Reports: []powerbi.GenerateTokenRequestV2Report{
        {
            ID:          &reportID,
            AllowEdit:   powerbi.Bool(false),
            AllowSaveAs: powerbi.Bool(false),
        },
    },
    TargetWorkspaces: []powerbi.GenerateTokenRequestV2TargetWorkspace{
        {ID: &workspaceID},
    },
}

embedToken, err := client.EmbedToken.GenerateToken(ctx, request, nil)
```

## Error Handling

Always check for errors and handle them appropriately:

```go
report, err := client.Reports.GetReport(ctx, reportID)
if err != nil {
    log.Fatalf("Failed to get report: %v", err)
}

// Use report...
fmt.Printf("Report: %s\n", powerbi.StringValue(report.Name))
```

## Helper Functions

The SDK provides helper functions for working with pointer types:

```go
// Create pointers
name := powerbi.String("My Report")
enabled := powerbi.Bool(true)
count := powerbi.Int(5)

// Get values (returns zero value if nil)
reportName := powerbi.StringValue(report.Name)  // returns "" if nil
isRefreshable := powerbi.BoolValue(dataset.IsRefreshable)  // returns false if nil
capacity := powerbi.IntValue(capacity.MaxParallelism)  // returns 0 if nil
```

## Additional Resources

- [Power BI REST API Documentation](https://docs.microsoft.com/rest/api/power-bi/)
- [Azure Identity SDK](https://github.com/Azure/azure-sdk-for-go/tree/main/sdk/azidentity)
- [Power BI Embedded Documentation](https://docs.microsoft.com/power-bi/developer/embedded/)
- [Register an Azure AD App](https://docs.microsoft.com/power-bi/developer/embedded/register-app)

## Troubleshooting

### Authentication Issues

If you get authentication errors:
1. Verify your credentials are correct
2. Check that your app has the necessary Power BI API permissions
3. Ensure the service principal has been granted access to your workspace

### Permission Errors

If you get "403 Forbidden" errors:
1. Verify your service principal has been added to the workspace
2. Check that the appropriate permissions have been granted
3. For admin operations, ensure your app has admin consent

### Rate Limiting

Power BI APIs have rate limits. If you hit them:
1. Implement exponential backoff
2. Cache results when possible
3. Use batch operations when available

## Contributing

Feel free to contribute more examples! Please follow the existing structure and include:
1. Clear comments explaining what the code does
2. Error handling
3. Documentation in this README

