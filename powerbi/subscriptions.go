package powerbi

import (
	"context"
	"fmt"
)

// SubscriptionsClient handles operations for Power BI subscriptions
type SubscriptionsClient struct {
	client *Client
}

// Subscription represents a Power BI subscription
type Subscription struct {
	ID              *string              `json:"id,omitempty"`
	Title           *string              `json:"title,omitempty"`
	ArtifactID      *string              `json:"artifactId,omitempty"`
	ArtifactDisplayName *string          `json:"artifactDisplayName,omitempty"`
	SubArtifactDisplayName *string       `json:"subArtifactDisplayName,omitempty"`
	Frequency       *string              `json:"frequency,omitempty"` // Daily, Weekly, Monthly
	StartDate       *string              `json:"startDate,omitempty"`
	EndDate         *string              `json:"endDate,omitempty"`
	LinkToContent   *bool                `json:"linkToContent,omitempty"`
	PreviewImage    *bool                `json:"previewImage,omitempty"`
	AttachmentFormat *string             `json:"attachmentFormat,omitempty"` // PDF, PPTX, PNG
	Users           []SubscriptionUser   `json:"users,omitempty"`
	Schedule        *SubscriptionSchedule `json:"schedule,omitempty"`
	IsEnabled       *bool                `json:"isEnabled,omitempty"`
}

// SubscriptionUser represents a user in a subscription
type SubscriptionUser struct {
	EmailAddress  *string `json:"emailAddress,omitempty"`
	DisplayName   *string `json:"displayName,omitempty"`
	Identifier    *string `json:"identifier,omitempty"`
	PrincipalType *string `json:"principalType,omitempty"`
}

// SubscriptionSchedule represents the schedule for a subscription
type SubscriptionSchedule struct {
	Days      []string `json:"days,omitempty"`      // Monday, Tuesday, etc.
	Time      *string  `json:"time,omitempty"`      // HH:MM format
	TimeZone  *string  `json:"timeZone,omitempty"`  // IANA time zone
}

// CreateSubscriptionRequest represents a request to create a subscription
type CreateSubscriptionRequest struct {
	Title           *string              `json:"title,omitempty"`
	ArtifactID      *string              `json:"artifactId,omitempty"`
	Frequency       *string              `json:"frequency,omitempty"`
	StartDate       *string              `json:"startDate,omitempty"`
	EndDate         *string              `json:"endDate,omitempty"`
	LinkToContent   *bool                `json:"linkToContent,omitempty"`
	PreviewImage    *bool                `json:"previewImage,omitempty"`
	AttachmentFormat *string             `json:"attachmentFormat,omitempty"`
	Users           []SubscriptionUser   `json:"users,omitempty"`
	Schedule        *SubscriptionSchedule `json:"schedule,omitempty"`
}

// Subscriptions represents a list of subscriptions
type Subscriptions struct {
	ODataContext string         `json:"@odata.context,omitempty"`
	Value        []Subscription `json:"value"`
}

// GetSubscriptions returns a list of subscriptions for the current user
func (c *SubscriptionsClient) GetSubscriptions(ctx context.Context) (*Subscriptions, error) {
	path := "/myorg/subscriptions"

	var result Subscriptions
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetSubscription returns the specified subscription
func (c *SubscriptionsClient) GetSubscription(ctx context.Context, subscriptionID string) (*Subscription, error) {
	if subscriptionID == "" {
		return nil, fmt.Errorf("subscriptionID cannot be empty")
	}

	path := fmt.Sprintf("/myorg/subscriptions/%s", subscriptionID)

	var result Subscription
	if err := c.client.doRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateSubscription creates a new subscription
func (c *SubscriptionsClient) CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error) {
	path := "/myorg/subscriptions"

	var result Subscription
	if err := c.client.doRequest(ctx, "POST", path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateSubscription updates an existing subscription
func (c *SubscriptionsClient) UpdateSubscription(ctx context.Context, subscriptionID string, request CreateSubscriptionRequest) error {
	if subscriptionID == "" {
		return fmt.Errorf("subscriptionID cannot be empty")
	}

	path := fmt.Sprintf("/myorg/subscriptions/%s", subscriptionID)
	return c.client.doRequest(ctx, "PATCH", path, request, nil)
}

// DeleteSubscription deletes the specified subscription
func (c *SubscriptionsClient) DeleteSubscription(ctx context.Context, subscriptionID string) error {
	if subscriptionID == "" {
		return fmt.Errorf("subscriptionID cannot be empty")
	}

	path := fmt.Sprintf("/myorg/subscriptions/%s", subscriptionID)
	return c.client.doRequest(ctx, "DELETE", path, nil, nil)
}

// EnableSubscription enables the specified subscription
func (c *SubscriptionsClient) EnableSubscription(ctx context.Context, subscriptionID string) error {
	if subscriptionID == "" {
		return fmt.Errorf("subscriptionID cannot be empty")
	}

	path := fmt.Sprintf("/myorg/subscriptions/%s/enable", subscriptionID)
	return c.client.doRequest(ctx, "POST", path, nil, nil)
}

// DisableSubscription disables the specified subscription
func (c *SubscriptionsClient) DisableSubscription(ctx context.Context, subscriptionID string) error {
	if subscriptionID == "" {
		return fmt.Errorf("subscriptionID cannot be empty")
	}

	path := fmt.Sprintf("/myorg/subscriptions/%s/disable", subscriptionID)
	return c.client.doRequest(ctx, "POST", path, nil, nil)
}

