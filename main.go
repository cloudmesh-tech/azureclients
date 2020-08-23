package main

import (
	"fmt"
	"log"

	"github.com/Azure/go-autorest/autorest/to"
)

var (
	azureClients = &Clients{
		Credential: &Credential{},
		Location:   to.StringPtr("EastAsia"),
	}
)

func main() {
	fmt.Println("hi")

	azureClients.Credential.AuthorizeFromFile()
	rgClient := azureClients.ResourcesGroupsClient()

	something, _ := rgClient.List(azureClients.Credential.Ctx, "", nil)

	for _, j := range something.Values() {
		log.Println(*j.Name)
	}
}
