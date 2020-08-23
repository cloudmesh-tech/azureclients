package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

// Credential Contains Azure client details, an authorizer token and context
type Credential struct {
	ServicePrincipal *ServicePrincipal
	Authorizer       autorest.Authorizer
	Ctx              context.Context
}

// AuthorizeFromFile Authorizes the Azure API client from file and returns an AzureCredential struct
func (creds *Credential) AuthorizeFromFile() {

	authorizer, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		log.Fatalf("Failed to get OAuth config: %v", err)
	}

	servicePrincipal, err := readJSON(os.Getenv("AZURE_AUTH_LOCATION"))

	if err != nil {
		log.Fatalf("Failed to read JSON: %+v", err)
	}

	creds.ServicePrincipal = servicePrincipal
	creds.Authorizer = authorizer
	creds.Ctx = context.Background()
}

// readJSON Reads json and returns a map
func readJSON(path string) (*ServicePrincipal, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	contents := make(map[string]string)
	json.Unmarshal(data, &contents)

	spInfo := &ServicePrincipal{
		ClientID:                   contents["clientId"],
		ClientSecret:               contents["clientSecret"],
		SubscriptionID:             contents["subscriptionId"],
		TenantID:                   contents["tenantId"],
		ActiveDirectoryEndPointURL: contents["activeDirectoryEndpointUrl"],
		ManagementEndpointURL:      contents["managementEndpointUrl"],
	}

	return spInfo, nil
}
