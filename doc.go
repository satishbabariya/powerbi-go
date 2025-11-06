/*
Package powerbi provides a Go SDK for the Power BI REST APIs.

The Power BI REST APIs provide service endpoints for embedding, user resources
management, administration and governance.

# Installation

	go get github.com/satishbabariya/powerbi-go

# Quick Start

Create a client using Azure credentials:

	import (
		"context"
		"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
		"github.com/satishbabariya/powerbi-go/powerbi"
	)

	func main() {
		ctx := context.Background()

		// Create credentials
		cred, err := azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
			log.Fatal(err)
		}

		// Create client
		client, err := powerbi.NewClient(cred, nil)
		if err != nil {
			log.Fatal(err)
		}

		// Use the client
		reports, err := client.Reports.GetReports(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

# Authentication

The SDK supports multiple authentication methods:

Service Principal with Client Secret:

	cred, err := azidentity.NewClientSecretCredential(
		tenantID, clientID, clientSecret, nil)
	client, err := powerbi.NewClient(cred, nil)

Service Principal with Certificate:

	certData, _ := os.ReadFile("cert.pem")
	certs, key, _ := azidentity.ParseCertificates(certData, nil)
	cred, err := azidentity.NewClientCertificateCredential(
		tenantID, clientID, certs, key, nil)
	client, err := powerbi.NewClient(cred, nil)

Managed Identity:

	cred, err := azidentity.NewManagedIdentityCredential(nil)
	client, err := powerbi.NewClient(cred, nil)

Static Access Token:

	client, err := powerbi.NewClientWithToken(token, nil)

# Features

The SDK provides comprehensive coverage of Power BI REST APIs:

  - Reports: Get, create, update, delete, export, clone reports
  - Datasets: Manage datasets, refresh data, update parameters
  - Dashboards: Get dashboards and tiles
  - Workspaces (Groups): Manage workspaces and users
  - Embed Tokens: Generate tokens for embedding
  - Admin: Organization-wide administrative operations
  - Gateways: Manage on-premises data gateways
  - Dataflows: Work with dataflows
  - Capacities: Manage Power BI capacities
  - And many more...

# Examples

List reports in a workspace:

	reports, err := client.Reports.GetReportsInGroup(ctx, workspaceID, nil)
	for _, report := range reports.Value {
		fmt.Printf("%s: %s\n",
			powerbi.StringValue(report.Name),
			powerbi.StringValue(report.ID))
	}

Refresh a dataset:

	refreshReq := &powerbi.RefreshRequest{
		NotifyOption: powerbi.String("MailOnFailure"),
	}
	err := client.Datasets.RefreshDatasetInGroup(
		ctx, workspaceID, datasetID, refreshReq)

Generate an embed token:

	request := powerbi.GenerateTokenRequestV2{
		Reports: []powerbi.GenerateTokenRequestV2Report{
			{ID: &reportID},
		},
		Datasets: []powerbi.GenerateTokenRequestV2Dataset{
			{ID: &datasetID},
		},
	}
	token, err := client.EmbedToken.GenerateToken(ctx, request, nil)

Create a workspace:

	req := powerbi.CreateGroupRequest{
		Name: powerbi.String("My Workspace"),
	}
	workspace, err := client.Groups.CreateGroup(ctx, req)

# Error Handling

All methods return errors that should be checked:

	report, err := client.Reports.GetReport(ctx, reportID)
	if err != nil {
		// Handle error
		log.Fatalf("Failed to get report: %v", err)
	}

# Helper Functions

The SDK provides helper functions for working with pointer types:

	// Create pointers
	name := powerbi.String("My Report")
	enabled := powerbi.Bool(true)
	count := powerbi.Int(5)

	// Get values (safe, returns zero value if nil)
	reportName := powerbi.StringValue(report.Name)
	isEnabled := powerbi.BoolValue(feature.Enabled)
	maxCount := powerbi.IntValue(config.MaxCount)

# Additional Resources

  - Power BI REST API: https://docs.microsoft.com/rest/api/power-bi/
  - Azure Identity: https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity
  - Examples: https://github.com/satishbabariya/powerbi-go/tree/main/examples

*/
package powerbi

