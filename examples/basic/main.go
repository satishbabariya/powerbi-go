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

	// List workspaces
	fmt.Println("\n=== Workspaces ===")
	groups, err := client.Groups.GetGroups(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to get groups: %v", err)
	}

	for _, group := range groups.Value {
		fmt.Printf("ID: %s, Name: %s\n",
			powerbi.StringValue(group.ID),
			powerbi.StringValue(group.Name))
	}

	// If we have workspaces, get reports from the first one
	if len(groups.Value) > 0 {
		groupID := powerbi.StringValue(groups.Value[0].ID)

		fmt.Printf("\n=== Reports in workspace: %s ===\n", powerbi.StringValue(groups.Value[0].Name))
		reports, err := client.Reports.GetReportsInGroup(ctx, groupID, nil)
		if err != nil {
			log.Fatalf("Failed to get reports: %v", err)
		}

		for _, report := range reports.Value {
			fmt.Printf("ID: %s, Name: %s\n",
				powerbi.StringValue(report.ID),
				powerbi.StringValue(report.Name))
		}

		fmt.Printf("\n=== Datasets in workspace: %s ===\n", powerbi.StringValue(groups.Value[0].Name))
		datasets, err := client.Datasets.GetDatasetsInGroup(ctx, groupID, nil)
		if err != nil {
			log.Fatalf("Failed to get datasets: %v", err)
		}

		for _, dataset := range datasets.Value {
			fmt.Printf("ID: %s, Name: %s, Refreshable: %v\n",
				powerbi.StringValue(dataset.ID),
				powerbi.StringValue(dataset.Name),
				powerbi.BoolValue(dataset.IsRefreshable))
		}
	}

	fmt.Println("\n=== Done ===")
}
