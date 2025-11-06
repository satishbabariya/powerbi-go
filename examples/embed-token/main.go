package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/satishbabariya/powerbi-go/powerbi"
)

func main() {
	ctx := context.Background()

	// Get credentials from environment variables
	tenantID := os.Getenv("AZURE_TENANT_ID")
	clientID := os.Getenv("AZURE_CLIENT_ID")
	clientSecret := os.Getenv("AZURE_CLIENT_SECRET")

	// IDs from environment or hardcode for testing
	reportID := os.Getenv("POWERBI_REPORT_ID")
	datasetID := os.Getenv("POWERBI_DATASET_ID")
	workspaceID := os.Getenv("POWERBI_WORKSPACE_ID")

	if reportID == "" || datasetID == "" || workspaceID == "" {
		log.Fatal("Please set POWERBI_REPORT_ID, POWERBI_DATASET_ID, and POWERBI_WORKSPACE_ID environment variables")
	}

	// Create credentials using service principal
	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		log.Fatalf("Failed to create credentials: %v", err)
	}

	// Create Power BI client
	client, err := powerbi.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Println("=== Generating Embed Token ===")

	// Create embed token request
	request := powerbi.GenerateTokenRequestV2{
		Datasets: []powerbi.GenerateTokenRequestV2Dataset{
			{
				ID: &datasetID,
			},
		},
		Reports: []powerbi.GenerateTokenRequestV2Report{
			{
				ID:          &reportID,
				AllowEdit:   powerbi.Bool(false),
				AllowSaveAs: powerbi.Bool(false),
			},
		},
		TargetWorkspaces: []powerbi.GenerateTokenRequestV2TargetWorkspace{
			{
				ID: &workspaceID,
			},
		},
	}

	// Generate embed token
	embedToken, err := client.EmbedToken.GenerateToken(ctx, request, nil)
	if err != nil {
		log.Fatalf("Failed to generate embed token: %v", err)
	}

	fmt.Printf("Token: %s\n", powerbi.StringValue(embedToken.Token)[:50]+"...")
	fmt.Printf("Token ID: %s\n", powerbi.StringValue(embedToken.TokenID))
	if embedToken.Expiration != nil {
		fmt.Printf("Expiration: %s\n", embedToken.Expiration.String())
	}

	// Get report details for embedding
	report, err := client.Reports.GetReportInGroup(ctx, workspaceID, reportID)
	if err != nil {
		log.Fatalf("Failed to get report: %v", err)
	}

	fmt.Printf("\nReport Details:\n")
	fmt.Printf("  Name: %s\n", powerbi.StringValue(report.Name))
	fmt.Printf("  Embed URL: %s\n", powerbi.StringValue(report.EmbedURL))
	fmt.Printf("  Web URL: %s\n", powerbi.StringValue(report.WebURL))

	fmt.Println("\n=== Done ===")
	fmt.Println("\nYou can now use this token to embed the report in your application.")
	fmt.Println("The embed URL and token should be used in your web application's Power BI embed code.")
}

