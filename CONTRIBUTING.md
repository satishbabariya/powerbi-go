# Contributing to Power BI Go SDK

First off, thank you for considering contributing to the Power BI Go SDK! It's people like you that make this SDK better for everyone.

## Code of Conduct

This project adheres to a code of conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check existing issues to avoid duplicates. When creating a bug report, include as many details as possible:

- **Use a clear and descriptive title**
- **Describe the exact steps to reproduce the problem**
- **Provide specific examples** to demonstrate the steps
- **Describe the behavior you observed** and why it's a problem
- **Explain the behavior you expected to see**
- **Include code samples** and error messages

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion:

- **Use a clear and descriptive title**
- **Provide a detailed description** of the suggested enhancement
- **Explain why this enhancement would be useful**
- **List some examples** of where this would be used

### Pull Requests

1. Fork the repo and create your branch from `main`
2. If you've added code that should be tested, add tests
3. Ensure the test suite passes
4. Make sure your code follows the existing code style
5. Write a clear commit message

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git

### Setting Up Your Development Environment

1. **Fork and clone the repository**
   ```bash
   git clone https://github.com/satishbabariya/powerbi-go.git
   cd powerbi-go
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables for testing**
   ```bash
   export AZURE_TENANT_ID="your-tenant-id"
   export AZURE_CLIENT_ID="your-client-id"
   export AZURE_CLIENT_SECRET="your-client-secret"
   export POWERBI_WORKSPACE_ID="your-workspace-id"
   ```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for a specific package
go test ./powerbi

# Run a specific test
go test -run TestClientCreation ./powerbi
```

### Running Examples

```bash
cd examples/basic
go run main.go
```

## Code Style Guidelines

### General Guidelines

- Follow standard Go conventions and idioms
- Use `gofmt` to format your code
- Use `golint` to check for style issues
- Write clear, readable code with meaningful variable names
- Add comments for exported functions, types, and constants

### Naming Conventions

- Use `camelCase` for unexported names
- Use `PascalCase` for exported names
- Use descriptive names that clearly indicate purpose
- Avoid abbreviations unless they're widely understood

### Error Handling

- Always check and handle errors
- Provide context in error messages
- Use `fmt.Errorf` with `%w` to wrap errors when appropriate

Example:
```go
if err != nil {
    return fmt.Errorf("failed to get report: %w", err)
}
```

### Function Documentation

All exported functions should have documentation comments:

```go
// GetReport returns the specified report from "My Workspace".
// It returns an error if the reportID is empty or if the API call fails.
func (c *ReportsClient) GetReport(ctx context.Context, reportID string) (*Report, error) {
    // implementation
}
```

### Testing

- Write unit tests for new functionality
- Aim for high test coverage
- Use table-driven tests where appropriate
- Mock external dependencies

Example:
```go
func TestGetReport(t *testing.T) {
    tests := []struct {
        name      string
        reportID  string
        wantError bool
    }{
        {
            name:      "valid report ID",
            reportID:  "valid-id",
            wantError: false,
        },
        {
            name:      "empty report ID",
            reportID:  "",
            wantError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

## Project Structure

```
powerbi-go/
├── powerbi/              # Main SDK code
│   ├── client.go         # Main client implementation
│   ├── models.go         # Data models
│   ├── reports.go        # Reports client
│   ├── datasets.go       # Datasets client
│   ├── dashboards.go     # Dashboards client
│   ├── groups.go         # Groups (workspaces) client
│   ├── embedtoken.go     # Embed token client
│   ├── admin.go          # Admin operations client
│   ├── gateways.go       # Gateways client
│   ├── other_clients.go  # Other client implementations
│   └── utils.go          # Utility functions
├── examples/             # Example code
│   ├── basic/
│   ├── embed-token/
│   ├── admin-operations/
│   ├── dataset-refresh/
│   └── workspace-management/
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
├── README.md             # Main documentation
├── CONTRIBUTING.md       # This file
└── LICENSE               # License file
```

## Adding New Features

### Adding a New Client Method

1. Add the method to the appropriate client file (e.g., `reports.go` for report operations)
2. Follow the existing pattern for method signatures
3. Add proper error handling and validation
4. Document the method with a clear comment
5. Add tests for the new method
6. Update the examples if appropriate

Example:
```go
// GetReportSubscriptions returns a list of subscriptions for the specified report.
func (c *ReportsClient) GetReportSubscriptions(ctx context.Context, reportID string) (*Subscriptions, error) {
    if reportID == "" {
        return nil, fmt.Errorf("reportID cannot be empty")
    }

    path := fmt.Sprintf("/reports/%s/subscriptions", reportID)

    var result Subscriptions
    if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
        return nil, err
    }

    return &result, nil
}
```

### Adding a New Model

1. Add the model to `models.go`
2. Use pointer types for optional fields
3. Add appropriate JSON tags
4. Document the model and its fields

Example:
```go
// Subscription represents a Power BI subscription
type Subscription struct {
    ID          *string    `json:"id,omitempty"`
    Title       *string    `json:"title,omitempty"`
    ArtifactID  *string    `json:"artifactId,omitempty"`
    IsEnabled   *bool      `json:"isEnabled,omitempty"`
    CreatedBy   *string    `json:"createdBy,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
}
```

## Documentation

- Update README.md for major changes
- Add examples for new features
- Update API documentation comments
- Keep CHANGELOG.md updated

## Release Process

1. Update version in relevant files
2. Update CHANGELOG.md
3. Create a git tag
4. Push the tag to trigger release

## Questions?

Feel free to open an issue with your question or reach out to the maintainers.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

