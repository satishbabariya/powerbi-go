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

	// Gateway and datasource IDs
	gatewayID := os.Getenv("POWERBI_GATEWAY_ID")
	datasourceName := os.Getenv("POWERBI_DATASOURCE_NAME")

	if gatewayID == "" || datasourceName == "" {
		log.Fatal("Please set POWERBI_GATEWAY_ID and POWERBI_DATASOURCE_NAME environment variables")
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

	fmt.Println("=== Gateway Credentials Management ===")

	// Get the gateway details
	gateway, err := client.Gateways.GetGateway(ctx, gatewayID)
	if err != nil {
		log.Fatalf("Failed to get gateway: %v", err)
	}

	fmt.Printf("Gateway: %s\n", powerbi.StringValue(gateway.Name))
	fmt.Printf("Status: %s\n", powerbi.StringValue(gateway.GatewayStatus))

	// Create credentials for a datasource
	fmt.Println("\n=== Creating Datasource with Encrypted Credentials ===")

	// Create basic credentials
	credentials := &powerbi.BasicCredentials{
		Username: "myusername",
		Password: "mypassword",
	}

	// Create encryptor using the gateway's public key
	encryptor, err := powerbi.NewAsymmetricKeyEncryptor(gateway.PublicKey)
	if err != nil {
		log.Fatalf("Failed to create encryptor: %v", err)
	}

	// Create credential details with encryption
	credDetails, err := powerbi.NewGatewayCredentialDetails(
		credentials,
		powerbi.PrivacyLevelNone,
		powerbi.EncryptedConnectionEncrypted,
		encryptor,
	)
	if err != nil {
		log.Fatalf("Failed to create credential details: %v", err)
	}

	fmt.Printf("Credential Type: %s\n", string(*credDetails.CredentialType))
	fmt.Printf("Encryption Algorithm: %s\n", string(*credDetails.EncryptionAlgorithm))
	fmt.Printf("Privacy Level: %s\n", string(*credDetails.PrivacyLevel))

	// Example of using other credential types
	fmt.Println("\n=== Other Credential Types ===")

	// Windows credentials
	windowsCreds := &powerbi.WindowsCredentials{
		Username: "DOMAIN\\user",
		Password: "password",
	}
	fmt.Printf("Windows Credentials Type: %s\n", string(windowsCreds.GetCredentialType()))

	// OAuth2 credentials
	oauth2Creds := &powerbi.OAuth2Credentials{
		AccessToken: "token",
	}
	fmt.Printf("OAuth2 Credentials Type: %s\n", string(oauth2Creds.GetCredentialType()))

	// Key credentials
	keyCreds := &powerbi.KeyCredentials{
		Key: "api-key",
	}
	fmt.Printf("Key Credentials Type: %s\n", string(keyCreds.GetCredentialType()))

	// Anonymous credentials
	anonCreds := &powerbi.AnonymousCredentials{}
	fmt.Printf("Anonymous Credentials Type: %s\n", string(anonCreds.GetCredentialType()))

	// Using caller AAD identity
	fmt.Println("\n=== Using Caller AAD Identity ===")
	aadCreds := powerbi.NewGatewayCredentialDetailsWithCallerAAD(
		powerbi.PrivacyLevelNone,
		powerbi.EncryptedConnectionEncrypted,
	)
	fmt.Printf("Use Caller AAD Identity: %v\n", powerbi.BoolValue(aadCreds.UseCallerAADIdentity))

	fmt.Println("\n=== Done ===")
	fmt.Println("\nNote: This example shows how to create encrypted credentials.")
	fmt.Println("To actually create a datasource, use client.Gateways.CreateDatasource()")
}

