package main

// ServicePrincipal Contatins details about the service principal used to authenticate with Azure
type ServicePrincipal struct {
	ClientID                   string
	ClientSecret               string
	SubscriptionID             string
	TenantID                   string
	ActiveDirectoryEndPointURL string
	ManagementEndpointURL      string
}
