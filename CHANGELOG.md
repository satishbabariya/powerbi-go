# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-01-01

### Added

- Initial release of Power BI Go SDK
- Complete implementation of Power BI REST API v1.0
- Core client with Azure Identity integration
- Authentication support:
  - Service Principal (Client Secret)
  - Service Principal (Certificate)
  - Managed Identity
  - Static Access Token
  - Default Azure Credential
- Client implementations:
  - **Reports**: Get, create, update, delete, export, clone, rebind reports
  - **Datasets**: Get, refresh, update, manage datasources
  - **Dashboards**: Get dashboards and tiles
  - **Groups (Workspaces)**: Manage workspaces and users
  - **Embed Token**: Generate embed tokens for embedding
  - **Admin**: Administrative operations for organization-wide management
  - **Gateways**: Manage on-premises data gateways
  - **Dataflows**: Work with dataflows
  - **Capacities**: Manage Power BI capacities
  - **Apps**: Manage Power BI apps
  - **Imports**: Import PBIX files
  - **Tiles**: Get and manage dashboard tiles
  - **Users**: Manage workspace users and permissions
  - **Pipelines**: Manage deployment pipelines
  - **Available Features**: Query available features
  - **Profiles**: Manage service principal profiles
  - **Scorecards**: Work with scorecards and goals
  - **Information Protection**: Sensitivity labels
  - **Template Apps**: Manage template apps
  - **Workspace Info**: Get workspace information
- Comprehensive examples:
  - Basic usage
  - Embed token generation
  - Admin operations
  - Dataset refresh
  - Workspace management
- Helper functions for working with pointer types
- Full documentation and usage guides
- MIT License

### Features by Category

#### Reports
- List reports in workspace or My Workspace
- Get individual report details
- Clone reports to same or different workspace
- Delete reports
- Export reports to file
- Rebind reports to different datasets
- Update report content
- Get report pages

#### Datasets
- List datasets in workspace or My Workspace
- Get individual dataset details
- Trigger dataset refresh
- Get refresh history
- Take over datasets
- Get and update datasources
- Bind to gateway

#### Dashboards
- List dashboards in workspace or My Workspace
- Get individual dashboard details
- Get tiles from dashboards
- Clone tiles
- Delete dashboards

#### Workspaces (Groups)
- List all workspaces
- Get individual workspace details
- Create new workspaces
- Update workspace properties
- Delete workspaces
- Manage workspace users (add, remove, update)
- Restore deleted workspaces
- Assign/unassign to/from capacity

#### Admin Operations
- Get all reports in organization
- Get all datasets in organization
- Get all dashboards in organization
- Get all workspaces in organization
- Manage workspace users as admin
- Get capacity information
- Get refreshables for capacity
- Get activity events (audit logs)

#### Gateways
- List gateways
- Get gateway details
- Manage gateway datasources (create, update, delete)

#### Embed Tokens
- Generate embed tokens for multiple reports, datasets, and workspaces
- Support for effective identities
- Configurable token lifetime

### Documentation
- Comprehensive README with examples
- Detailed API documentation
- Contributing guidelines
- Examples for common scenarios
- Authentication guide
- Error handling guide

### Development
- Go 1.21+ support
- Azure SDK for Go integration
- Modular client architecture
- Context-based operations
- Proper error handling and wrapping

## [Unreleased]

### Planned
- Unit tests
- Integration tests
- More examples
- Performance optimizations
- Additional helper utilities
- Retry policies
- Rate limiting handling
- Webhook support
- Batch operations

---

## Version History

- **1.0.0** - Initial release with complete API coverage

[1.0.0]: https://github.com/satishbabariya/powerbi-go/releases/tag/v1.0.0
[Unreleased]: https://github.com/satishbabariya/powerbi-go/compare/v1.0.0...HEAD

