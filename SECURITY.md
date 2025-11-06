# Security Policy

## Supported Versions

We release patches for security vulnerabilities for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |

## Reporting a Vulnerability

If you discover a security vulnerability, please follow these steps:

### 1. Do Not Create a Public Issue

Please do not create a public GitHub issue for security vulnerabilities.

### 2. Contact Us Privately

Send an email to the maintainers with:

- A description of the vulnerability
- Steps to reproduce the issue
- Potential impact
- Any suggested fixes

### 3. Wait for a Response

We will acknowledge your email within 48 hours and will send a more detailed response within 7 days indicating the next steps.

### 4. Coordinated Disclosure

We believe in coordinated disclosure and will work with you to understand and address the issue before making it public.

## Security Best Practices

When using this SDK, please follow these security best practices:

### 1. Credential Management

**Never commit credentials to version control:**

```bash
# Use environment variables
export AZURE_CLIENT_SECRET="your-secret"

# Or use a .env file (add to .gitignore)
echo "AZURE_CLIENT_SECRET=your-secret" >> .env
```

**Use Azure Key Vault for production:**

```go
import "github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"

// Retrieve secrets from Key Vault
client, _ := azsecrets.NewClient(vaultURL, credential, nil)
secret, _ := client.GetSecret(ctx, "client-secret", "", nil)
```

### 2. Use Managed Identities

When running in Azure, use Managed Identity instead of service principals:

```go
cred, err := azidentity.NewManagedIdentityCredential(nil)
client, err := powerbi.NewClient(cred, nil)
```

### 3. Principle of Least Privilege

Grant only the minimum required permissions:

- Use workspace-specific permissions instead of admin access when possible
- Regularly audit service principal permissions
- Remove unused service principals

### 4. Token Handling

**Embed tokens should be short-lived:**

```go
request := powerbi.GenerateTokenRequestV2{
    LifetimeInMinutes: powerbi.Int(60), // 1 hour maximum
    // ...
}
```

**Never expose embed tokens in client-side code:**

```javascript
// ❌ Bad - token exposed in frontend
const embedToken = "eyJ0eXAiOiJKV...";

// ✅ Good - fetch token from secure backend
const response = await fetch('/api/embed-token');
const data = await response.json();
```

### 5. Network Security

**Use HTTPS for custom base URLs:**

```go
// ✅ Good
options := &powerbi.ClientOptions{
    BaseURL: "https://api.powerbi.com",
}

// ❌ Bad
options := &powerbi.ClientOptions{
    BaseURL: "http://api.powerbi.com", // Never use HTTP
}
```

### 6. Input Validation

Always validate user inputs before passing to API calls:

```go
func getReport(userInput string) (*powerbi.Report, error) {
    // Validate UUID format
    if !isValidUUID(userInput) {
        return nil, fmt.Errorf("invalid report ID")
    }
    
    return client.Reports.GetReport(ctx, userInput)
}
```

### 7. Error Handling

Don't expose sensitive information in error messages:

```go
// ❌ Bad - exposes credentials
log.Printf("Failed with secret %s: %v", secret, err)

// ✅ Good - no sensitive data
log.Printf("Failed to authenticate: %v", err)
```

### 8. Dependency Management

Regularly update dependencies:

```bash
go get -u ./...
go mod tidy
```

Use tools like Dependabot or Renovate to automate dependency updates.

### 9. Audit Logging

Monitor API usage using Power BI's audit logs:

```go
// Get activity events for compliance
events, err := client.Admin.GetActivityEvents(
    ctx,
    startTime,
    endTime,
    nil,
)
```

## Common Vulnerabilities

### Credential Leakage

**Problem:** Credentials committed to Git history

**Solution:**
```bash
# Remove from git history
git filter-branch --force --index-filter \
  'git rm --cached --ignore-unmatch .env' \
  --prune-empty --tag-name-filter cat -- --all

# Rotate compromised credentials immediately
```

### Insufficient Access Control

**Problem:** Over-privileged service principals

**Solution:**
- Use Azure RBAC
- Implement workspace-level permissions
- Regular access reviews

### Token Replay Attacks

**Problem:** Long-lived or reused tokens

**Solution:**
- Use short token lifetimes
- Implement token rotation
- Use effective identities for RLS

### Man-in-the-Middle (MITM)

**Problem:** Insecure connections

**Solution:**
- Always use HTTPS/TLS
- Verify certificates
- Use Azure Private Link for sensitive data

## Compliance

This SDK helps you maintain compliance with:

- **GDPR**: Data access controls and audit logs
- **SOC 2**: Security and availability controls
- **ISO 27001**: Information security management
- **HIPAA**: Healthcare data protection (when properly configured)

## Security Updates

We will notify users of security updates through:

1. GitHub Security Advisories
2. Release notes
3. Email notifications (for critical issues)

## Additional Resources

- [Power BI Security](https://docs.microsoft.com/power-bi/admin/service-admin-power-bi-security)
- [Azure Security Best Practices](https://docs.microsoft.com/azure/security/fundamentals/best-practices-and-patterns)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)

## Acknowledgments

We appreciate the security research community's efforts in responsibly disclosing vulnerabilities.

