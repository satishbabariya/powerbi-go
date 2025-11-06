package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/satishbabariya/powerbi-go/powerbi"
)

func main() {
	ctx := context.Background()

	// Get credentials from environment variables
	tenantID := os.Getenv("AZURE_TENANT_ID")
	clientID := os.Getenv("AZURE_CLIENT_ID")
	clientSecret := os.Getenv("AZURE_CLIENT_SECRET")

	// IDs from environment
	workspaceID := os.Getenv("POWERBI_WORKSPACE_ID")
	datasetID := os.Getenv("POWERBI_DATASET_ID")

	if workspaceID == "" || datasetID == "" {
		log.Fatal("Please set POWERBI_WORKSPACE_ID and POWERBI_DATASET_ID environment variables")
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

	// Get dataset details
	fmt.Println("=== Dataset Details ===")
	dataset, err := client.Datasets.GetDatasetInGroup(ctx, workspaceID, datasetID)
	if err != nil {
		log.Fatalf("Failed to get dataset: %v", err)
	}

	fmt.Printf("Name: %s\n", powerbi.StringValue(dataset.Name))
	fmt.Printf("ID: %s\n", powerbi.StringValue(dataset.ID))
	fmt.Printf("Is Refreshable: %v\n", powerbi.BoolValue(dataset.IsRefreshable))
	fmt.Printf("Configured By: %s\n", powerbi.StringValue(dataset.ConfiguredBy))

	// Get refresh history
	fmt.Println("\n=== Recent Refresh History ===")
	top := 5
	history, err := client.Datasets.GetRefreshHistoryInGroup(ctx, workspaceID, datasetID, &top)
	if err != nil {
		log.Fatalf("Failed to get refresh history: %v", err)
	}

	if len(history.Value) == 0 {
		fmt.Println("No refresh history found")
	} else {
		for i, entry := range history.Value {
			fmt.Printf("%d. Status: %s, Type: %s\n", 
				i+1,
				powerbi.StringValue(entry.Status),
				powerbi.StringValue(entry.RefreshType))
			fmt.Printf("   Start: %s\n", powerbi.StringValue(entry.StartTime))
			fmt.Printf("   End: %s\n", powerbi.StringValue(entry.EndTime))
		}
	}

	// Trigger a refresh (only if the dataset is refreshable)
	if powerbi.BoolValue(dataset.IsRefreshable) {
		fmt.Println("\n=== Triggering Dataset Refresh ===")
		
		notifyOption := "MailOnFailure"
		refreshRequest := &powerbi.RefreshRequest{
			NotifyOption: &notifyOption,
		}

		err = client.Datasets.RefreshDatasetInGroup(ctx, workspaceID, datasetID, refreshRequest)
		if err != nil {
			log.Fatalf("Failed to trigger refresh: %v", err)
		}

		fmt.Println("Refresh triggered successfully!")
		fmt.Println("The refresh is now running in the background.")
		
		// Wait a bit and check the status
		fmt.Println("\nWaiting 5 seconds before checking status...")
		time.Sleep(5 * time.Second)

		// Get updated refresh history
		fmt.Println("\n=== Updated Refresh History ===")
		history, err = client.Datasets.GetRefreshHistoryInGroup(ctx, workspaceID, datasetID, &top)
		if err != nil {
			log.Fatalf("Failed to get refresh history: %v", err)
		}

		if len(history.Value) > 0 {
			latestRefresh := history.Value[0]
			fmt.Printf("Latest Refresh Status: %s\n", powerbi.StringValue(latestRefresh.Status))
			fmt.Printf("Type: %s\n", powerbi.StringValue(latestRefresh.RefreshType))
			fmt.Printf("Start Time: %s\n", powerbi.StringValue(latestRefresh.StartTime))
			if latestRefresh.EndTime != nil {
				fmt.Printf("End Time: %s\n", powerbi.StringValue(latestRefresh.EndTime))
			}
		}
	} else {
		fmt.Println("\nDataset is not refreshable. Skipping refresh operation.")
	}

	fmt.Println("\n=== Done ===")
}

