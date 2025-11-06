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
	// Note: This service principal must have Power BI admin rights
	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		log.Fatalf("Failed to create credentials: %v", err)
	}

	// Create Power BI client
	client, err := powerbi.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Get all reports in the organization (admin operation)
	fmt.Println("\n=== All Reports in Organization (Admin) ===")
	top := 10
	reports, err := client.Admin.GetReportsAsAdmin(ctx, &powerbi.GetReportsAsAdminOptions{
		Top: &top,
	})
	if err != nil {
		log.Fatalf("Failed to get reports as admin: %v", err)
	}

	for i, report := range reports.Value {
		fmt.Printf("%d. ID: %s, Name: %s\n", 
			i+1,
			powerbi.StringValue(report.ID), 
			powerbi.StringValue(report.Name))
	}

	// Get all datasets in the organization (admin operation)
	fmt.Println("\n=== All Datasets in Organization (Admin) ===")
	datasets, err := client.Admin.GetDatasetsAsAdmin(ctx, &powerbi.GetReportsAsAdminOptions{
		Top: &top,
	})
	if err != nil {
		log.Fatalf("Failed to get datasets as admin: %v", err)
	}

	for i, dataset := range datasets.Value {
		fmt.Printf("%d. ID: %s, Name: %s, Refreshable: %v\n", 
			i+1,
			powerbi.StringValue(dataset.ID), 
			powerbi.StringValue(dataset.Name),
			powerbi.BoolValue(dataset.IsRefreshable))
	}

	// Get all workspaces in the organization (admin operation)
	fmt.Println("\n=== All Workspaces in Organization (Admin) ===")
	groups, err := client.Admin.GetGroupsAsAdmin(ctx, &powerbi.GetGroupsAsAdminOptions{
		Top: &top,
	})
	if err != nil {
		log.Fatalf("Failed to get groups as admin: %v", err)
	}

	for i, group := range groups.Value {
		fmt.Printf("%d. ID: %s, Name: %s, Type: %s\n", 
			i+1,
			powerbi.StringValue(group.ID), 
			powerbi.StringValue(group.Name),
			powerbi.StringValue(group.Type))
	}

	// Get capacities in the organization (admin operation)
	fmt.Println("\n=== Capacities in Organization (Admin) ===")
	capacities, err := client.Admin.GetCapacitiesAsAdmin(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to get capacities as admin: %v", err)
	}

	for i, capacity := range capacities.Value {
		fmt.Printf("%d. ID: %s, Name: %s, SKU: %s, State: %s\n", 
			i+1,
			powerbi.StringValue(capacity.ID), 
			powerbi.StringValue(capacity.DisplayName),
			powerbi.StringValue(capacity.SKU),
			powerbi.StringValue(capacity.State))
	}

	fmt.Println("\n=== Done ===")
}

