package powerbi

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
)

// CredentialType represents the type of credential
type CredentialType string

const (
	CredentialTypeBasic     CredentialType = "Basic"
	CredentialTypeWindows   CredentialType = "Windows"
	CredentialTypeOAuth2    CredentialType = "OAuth2"
	CredentialTypeKey       CredentialType = "Key"
	CredentialTypeAnonymous CredentialType = "Anonymous"
)

// PrivacyLevel represents the privacy level for credentials
type PrivacyLevel string

const (
	PrivacyLevelNone         PrivacyLevel = "None"
	PrivacyLevelPublic       PrivacyLevel = "Public"
	PrivacyLevelOrganizational PrivacyLevel = "Organizational"
	PrivacyLevelPrivate      PrivacyLevel = "Private"
)

// EncryptedConnection represents whether the connection is encrypted
type EncryptedConnection string

const (
	EncryptedConnectionEncrypted   EncryptedConnection = "Encrypted"
	EncryptedConnectionNotEncrypted EncryptedConnection = "NotEncrypted"
)

// EncryptionAlgorithm represents the encryption algorithm used
type EncryptionAlgorithm string

const (
	EncryptionAlgorithmNone    EncryptionAlgorithm = "None"
	EncryptionAlgorithmRSAOAEP EncryptionAlgorithm = "RSA-OAEP"
)

// CredentialsBase is the base interface for all credential types
type CredentialsBase interface {
	GetCredentialData() map[string]string
	GetCredentialType() CredentialType
}

// BasicCredentials represents username and password credentials
type BasicCredentials struct {
	Username string
	Password string
}

// GetCredentialData returns the credential data as a map
func (c *BasicCredentials) GetCredentialData() map[string]string {
	return map[string]string{
		"username": c.Username,
		"password": c.Password,
	}
}

// GetCredentialType returns the credential type
func (c *BasicCredentials) GetCredentialType() CredentialType {
	return CredentialTypeBasic
}

// WindowsCredentials represents Windows authentication credentials
type WindowsCredentials struct {
	Username string
	Password string
}

// GetCredentialData returns the credential data as a map
func (c *WindowsCredentials) GetCredentialData() map[string]string {
	return map[string]string{
		"username": c.Username,
		"password": c.Password,
	}
}

// GetCredentialType returns the credential type
func (c *WindowsCredentials) GetCredentialType() CredentialType {
	return CredentialTypeWindows
}

// OAuth2Credentials represents OAuth2 credentials
type OAuth2Credentials struct {
	AccessToken string
}

// GetCredentialData returns the credential data as a map
func (c *OAuth2Credentials) GetCredentialData() map[string]string {
	return map[string]string{
		"accessToken": c.AccessToken,
	}
}

// GetCredentialType returns the credential type
func (c *OAuth2Credentials) GetCredentialType() CredentialType {
	return CredentialTypeOAuth2
}

// KeyCredentials represents key-based credentials
type KeyCredentials struct {
	Key string
}

// GetCredentialData returns the credential data as a map
func (c *KeyCredentials) GetCredentialData() map[string]string {
	return map[string]string{
		"key": c.Key,
	}
}

// GetCredentialType returns the credential type
func (c *KeyCredentials) GetCredentialType() CredentialType {
	return CredentialTypeKey
}

// AnonymousCredentials represents anonymous credentials
type AnonymousCredentials struct{}

// GetCredentialData returns the credential data as a map
func (c *AnonymousCredentials) GetCredentialData() map[string]string {
	return map[string]string{}
}

// GetCredentialType returns the credential type
func (c *AnonymousCredentials) GetCredentialType() CredentialType {
	return CredentialTypeAnonymous
}

// CredentialsEncryptor is an interface for encrypting credentials
type CredentialsEncryptor interface {
	EncryptCredentials(plainText string) (string, error)
}

// AsymmetricKeyEncryptor encrypts credentials using RSA encryption
type AsymmetricKeyEncryptor struct {
	publicKey *GatewayPublicKey
}

// NewAsymmetricKeyEncryptor creates a new asymmetric key encryptor
func NewAsymmetricKeyEncryptor(publicKey *GatewayPublicKey) (*AsymmetricKeyEncryptor, error) {
	if publicKey == nil {
		return nil, fmt.Errorf("publicKey cannot be nil")
	}
	if publicKey.Exponent == nil || *publicKey.Exponent == "" {
		return nil, fmt.Errorf("publicKey.Exponent cannot be empty")
	}
	if publicKey.Modulus == nil || *publicKey.Modulus == "" {
		return nil, fmt.Errorf("publicKey.Modulus cannot be empty")
	}

	return &AsymmetricKeyEncryptor{publicKey: publicKey}, nil
}

// EncryptCredentials encrypts credentials using RSA-OAEP
func (e *AsymmetricKeyEncryptor) EncryptCredentials(plainText string) (string, error) {
	if plainText == "" {
		return "", fmt.Errorf("plainText cannot be empty")
	}

	// Decode the modulus and exponent from base64
	modulusBytes, err := base64.StdEncoding.DecodeString(*e.publicKey.Modulus)
	if err != nil {
		return "", fmt.Errorf("failed to decode modulus: %w", err)
	}

	exponentBytes, err := base64.StdEncoding.DecodeString(*e.publicKey.Exponent)
	if err != nil {
		return "", fmt.Errorf("failed to decode exponent: %w", err)
	}

	// Create RSA public key
	pubKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(modulusBytes),
		E: int(new(big.Int).SetBytes(exponentBytes).Int64()),
	}

	// Encrypt using RSA-OAEP with SHA-1
	plainTextBytes := []byte(plainText)
	ciphertext, err := rsa.EncryptOAEP(
		nil, // Use default SHA-1
		rand.Reader,
		pubKey,
		plainTextBytes,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("encryption failed: %w", err)
	}

	// Return base64 encoded ciphertext
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// GatewayCredentialDetails represents credential details for gateway datasources
type GatewayCredentialDetails struct {
	CredentialType       *CredentialType       `json:"credentialType,omitempty"`
	Credentials          *string               `json:"credentials,omitempty"`
	EncryptedConnection  *EncryptedConnection  `json:"encryptedConnection,omitempty"`
	EncryptionAlgorithm  *EncryptionAlgorithm  `json:"encryptionAlgorithm,omitempty"`
	PrivacyLevel         *PrivacyLevel         `json:"privacyLevel,omitempty"`
	UseCallerAADIdentity *bool                 `json:"useCallerAADIdentity,omitempty"`
	UseEndUserOAuth2Credentials *bool         `json:"useEndUserOAuth2Credentials,omitempty"`
}

// NewGatewayCredentialDetails creates credential details for a gateway datasource
func NewGatewayCredentialDetails(
	credentials CredentialsBase,
	privacyLevel PrivacyLevel,
	encryptedConnection EncryptedConnection,
	encryptor CredentialsEncryptor,
) (*GatewayCredentialDetails, error) {
	// Create the credential request
	type nameValuePair struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	type credentialsRequest struct {
		CredentialData []nameValuePair `json:"credentialData"`
	}

	credData := credentials.GetCredentialData()
	pairs := make([]nameValuePair, 0, len(credData))
	for k, v := range credData {
		pairs = append(pairs, nameValuePair{Name: k, Value: v})
	}

	request := credentialsRequest{CredentialData: pairs}

	// Serialize to JSON
	credJSON, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal credentials: %w", err)
	}

	credString := string(credJSON)

	// Encrypt if encryptor provided
	var encryptionAlgorithm EncryptionAlgorithm
	if encryptor != nil {
		encrypted, err := encryptor.EncryptCredentials(credString)
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt credentials: %w", err)
		}
		credString = encrypted
		encryptionAlgorithm = EncryptionAlgorithmRSAOAEP
	} else {
		encryptionAlgorithm = EncryptionAlgorithmNone
	}

	credType := credentials.GetCredentialType()

	return &GatewayCredentialDetails{
		CredentialType:      &credType,
		Credentials:         &credString,
		EncryptedConnection: &encryptedConnection,
		EncryptionAlgorithm: &encryptionAlgorithm,
		PrivacyLevel:        &privacyLevel,
		UseCallerAADIdentity: Bool(false),
	}, nil
}

// NewGatewayCredentialDetailsWithCallerAAD creates credential details using caller AAD identity
func NewGatewayCredentialDetailsWithCallerAAD(
	privacyLevel PrivacyLevel,
	encryptedConnection EncryptedConnection,
) *GatewayCredentialDetails {
	credType := CredentialTypeOAuth2
	encAlg := EncryptionAlgorithmNone

	return &GatewayCredentialDetails{
		CredentialType:       &credType,
		EncryptedConnection:  &encryptedConnection,
		EncryptionAlgorithm:  &encAlg,
		PrivacyLevel:         &privacyLevel,
		UseCallerAADIdentity: Bool(true),
	}
}

