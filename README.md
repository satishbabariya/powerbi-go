# Power BI REST APIs for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/satishbabariya/powerbi-go.svg)](https://pkg.go.dev/github.com/satishbabariya/powerbi-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/satishbabariya/powerbi-go)](https://goreportcard.com/report/github.com/satishbabariya/powerbi-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

The Power BI REST APIs provide service endpoints for embedding, user resources management, administration and governance.

For more information about Power BI REST APIs, see [Power BI REST APIs overview](https://docs.microsoft.com/rest/api/power-bi/).

## Power BI API Library

The `powerbi-go` library enables you to work with Power BI REST APIs in your Go application.

## Installation

```bash
go get github.com/satishbabariya/powerbi-go
```

## Creating the Power BI Client

The `PowerBIClient` can be created using either **Azure credentials** or a **Microsoft Entra Access Token**.

### Using Azure Identity (Recommended)

Authenticate using Azure Identity with service principal credentials:

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/satishbabariya/powerbi-go/powerbi"
)

func main() {
    ctx := context.Background()
    
    // Using Service Principal
    cred, err := azidentity.NewClientSecretCredential(
        "tenant-id",
        "client-id", 
        "client-secret",
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    client, err := powerbi.NewClient(cred, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    // Use the client
    datasets, err := client.Datasets.GetDatasets(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, ds := range datasets.Value {
        fmt.Printf("%s | %s\n", *ds.ID, *ds.Name)
    }
}
```

### Using Access Token

Alternatively, authenticate using a Microsoft Entra Access Token:

```go
package main

import (
    "context"
    "log"
    
    "github.com/satishbabariya/powerbi-go/powerbi"
)

func main() {
    ctx := context.Background()
    
    client, err := powerbi.NewClientWithToken("your-access-token", nil)
    if err != nil {
        log.Fatal(err)
    }
    
    // Use the client
    reports, err := client.Reports.GetReports(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Examples

### Get Datasets and Reports in a Workspace

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/satishbabariya/powerbi-go/powerbi"
)

func main() {
    ctx := context.Background()
    
    cred, err := azidentity.NewDefaultAzureCredential(nil)
    if err != nil {
        log.Fatal(err)
    }
    
    client, err := powerbi.NewClient(cred, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    groupID := "your-workspace-id"
    
    // Get datasets
    fmt.Println("\n*** DATASETS ***")
    datasets, err := client.Datasets.GetDatasetsInGroup(ctx, groupID, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, ds := range datasets.Value {
        fmt.Printf("%s | %s\n", *ds.ID, *ds.Name)
    }
    
    // Get reports
    fmt.Println("\n*** REPORTS ***")
    reports, err := client.Reports.GetReportsInGroup(ctx, groupID, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, rpt := range reports.Value {
        fmt.Printf("%s | %s | DatasetID = %s\n", *rpt.ID, *rpt.Name, *rpt.DatasetID)
    }
}
```

### Creating an Embed Token

Embed tokens are used to provide access to Power BI artifacts like reports and datasets to embed into an application.

```go
package main

import (
    "context"
    "log"
    
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/satishbabariya/powerbi-go/powerbi"
)

func main() {
    ctx := context.Background()
    
    cred, err := azidentity.NewDefaultAzureCredential(nil)
    if err != nil {
        log.Fatal(err)
    }
    
    client, err := powerbi.NewClient(cred, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    // Create embed token request
    reportID := "report-id"
    datasetID := "dataset-id"
    
    request := powerbi.GenerateTokenRequestV2{
        Datasets: []powerbi.GenerateTokenRequestV2Dataset{
            {ID: &datasetID},
        },
        Reports: []powerbi.GenerateTokenRequestV2Report{
            {ID: &reportID},
        },
    }
    
    embedToken, err := client.EmbedToken.GenerateToken(ctx, request, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Embed Token: %s\n", *embedToken.Token)
}
```

### Get Reports As Admin

Returns a list of reports for the organization. The caller must have administrator rights.

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/satishbabariya/powerbi-go/powerbi"
)

func main() {
    ctx := context.Background()
    
    cred, err := azidentity.NewDefaultAzureCredential(nil)
    if err != nil {
        log.Fatal(err)
    }
    
    client, err := powerbi.NewClient(cred, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("\n*** REPORTS (Admin) ***")
    reports, err := client.Admin.GetReportsAsAdmin(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, rpt := range reports.Value {
        fmt.Printf("%s | %s\n", *rpt.ID, *rpt.Name)
    }
}
```

## Features

The SDK provides complete coverage of Power BI REST APIs including:

- **Reports**: Get, create, update, delete, export, and clone reports
- **Datasets**: Manage datasets, refresh data, update parameters
- **Dashboards**: Get dashboards and tiles
- **Workspaces (Groups)**: Manage workspaces and workspace users
- **Embed Tokens**: Generate embed tokens for reports and datasets
- **Admin**: Administrative operations for organization-wide management
- **Gateways**: Manage on-premises data gateways
- **Dataflows**: Work with dataflows
- **Capacities**: Manage Power BI capacities
- **Apps**: Manage Power BI apps
- **Imports**: Import PBIX files
- **Tiles**: Get and manage dashboard tiles
- **Users**: Manage workspace users and permissions
- **Pipelines**: Manage deployment pipelines
- **Available Features**: Query available features
- **Profiles**: Manage service principal profiles
- **Scorecards**: Work with scorecards and goals
- **Information Protection**: Sensitivity labels
- **Template Apps**: Manage template apps
- **Workspace Info**: Get workspace information

## Additional Links

- [Power BI REST APIs](https://docs.microsoft.com/rest/api/power-bi/)
- [AAD Application registration](https://docs.microsoft.com/power-bi/developer/embedded/register-app?tabs=customers%2CAzure#register-an-azure-ad-app)
- [Power BI Embedded Analytics Playground](https://playground.powerbi.com)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

