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

	// List all workspaces
	fmt.Println("=== All Workspaces ===")
	groups, err := client.Groups.GetGroups(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to get groups: %v", err)
	}

	for i, group := range groups.Value {
		fmt.Printf("%d. ID: %s\n", i+1, powerbi.StringValue(group.ID))
		fmt.Printf("   Name: %s\n", powerbi.StringValue(group.Name))
		fmt.Printf("   Type: %s\n", powerbi.StringValue(group.Type))
		fmt.Printf("   Is On Dedicated Capacity: %v\n", powerbi.BoolValue(group.IsOnDedicatedCapacity))
		if group.CapacityID != nil {
			fmt.Printf("   Capacity ID: %s\n", powerbi.StringValue(group.CapacityID))
		}
		fmt.Println()
	}

	// Create a new workspace
	fmt.Println("=== Creating New Workspace ===")
	workspaceName := "SDK Test Workspace"
	description := "Created by Power BI Go SDK"
	
	createRequest := powerbi.CreateGroupRequest{
		Name:        &workspaceName,
		Description: &description,
	}

	newGroup, err := client.Groups.CreateGroup(ctx, createRequest)
	if err != nil {
		log.Printf("Failed to create workspace (this is normal if it already exists): %v", err)
	} else {
		fmt.Printf("Created workspace: %s (ID: %s)\n", 
			powerbi.StringValue(newGroup.Name), 
			powerbi.StringValue(newGroup.ID))
		
		workspaceID := powerbi.StringValue(newGroup.ID)

		// Get workspace users
		fmt.Println("\n=== Workspace Users ===")
		users, err := client.Groups.GetGroupUsers(ctx, workspaceID)
		if err != nil {
			log.Printf("Failed to get workspace users: %v", err)
		} else {
			for i, user := range users {
				fmt.Printf("%d. Email: %s\n", i+1, powerbi.StringValue(user.EmailAddress))
				fmt.Printf("   Display Name: %s\n", powerbi.StringValue(user.DisplayName))
				fmt.Printf("   Access Right: %s\n", powerbi.StringValue(user.GroupUserAccessRight))
				fmt.Printf("   Principal Type: %s\n", powerbi.StringValue(user.PrincipalType))
				fmt.Println()
			}
		}

		// Add a user to the workspace (uncomment and modify as needed)
		/*
		fmt.Println("=== Adding User to Workspace ===")
		userEmail := "user@example.com"
		accessRight := "Member"
		principalType := "User"
		
		addUserRequest := powerbi.GroupUserAccessRight{
			EmailAddress:         &userEmail,
			GroupUserAccessRight: &accessRight,
			PrincipalType:        &principalType,
		}

		err = client.Groups.AddGroupUser(ctx, workspaceID, addUserRequest)
		if err != nil {
			log.Printf("Failed to add user: %v", err)
		} else {
			fmt.Printf("Added user %s to workspace\n", userEmail)
		}
		*/

		// Update workspace
		fmt.Println("\n=== Updating Workspace ===")
		updatedDescription := "Updated by Power BI Go SDK"
		updateRequest := powerbi.UpdateGroupRequest{
			Description: &updatedDescription,
		}

		err = client.Groups.UpdateGroup(ctx, workspaceID, updateRequest)
		if err != nil {
			log.Printf("Failed to update workspace: %v", err)
		} else {
			fmt.Println("Workspace updated successfully")
		}

		// Get updated workspace details
		fmt.Println("\n=== Updated Workspace Details ===")
		updatedGroup, err := client.Groups.GetGroup(ctx, workspaceID)
		if err != nil {
			log.Printf("Failed to get workspace: %v", err)
		} else {
			fmt.Printf("Name: %s\n", powerbi.StringValue(updatedGroup.Name))
			fmt.Printf("Description: %s\n", powerbi.StringValue(updatedGroup.Description))
		}

		// Delete the workspace (uncomment if you want to clean up)
		/*
		fmt.Println("\n=== Deleting Workspace ===")
		err = client.Groups.DeleteGroup(ctx, workspaceID)
		if err != nil {
			log.Printf("Failed to delete workspace: %v", err)
		} else {
			fmt.Println("Workspace deleted successfully")
		}
		*/
	}

	fmt.Println("\n=== Done ===")
}

