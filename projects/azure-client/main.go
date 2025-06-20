package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

func main() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain credential: %v", err)
	}

	client, err := armresources.NewResourceGroupsClient("6bf5efd4-c0e6-4ac1-b690-4d8133396c9c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			log.Fatalf("failed to get next page: %v", err)
		}
		for _, rg := range page.Value {
			fmt.Println(*rg.Name)
		}
	}
}
