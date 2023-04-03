package cosmosdbbackend

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

var endpoint = "<azure_cosmos_uri>"
var key = "<azure_cosmos_primary_key"
var databasename = ""
var containername = ""

func auth() (*azcosmos.DatabaseClient, *azcosmos.ContainerClient) {

	cred, err := azcosmos.NewKeyCredential(key)
	if err != nil {
		log.Fatal("Failed to create a credential: ", err)
	}

	// Create a CosmosDB client
	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)
	if err != nil {
		log.Fatal("Failed to create Azure Cosmos DB client: ", err)
	}

	// Create database client
	databaseClient, err := client.NewDatabase("<databaseName>")
	if err != nil {
		log.Fatal("Failed to create database client:", err)
	}

	// Create container client
	containerClient, err := client.NewContainer("<databaseName>", "<containerName>")
	if err != nil {
		log.Fatal("Failed to create a container client:", err)
	}

	return databaseClient, containerClient
}

func createDatabase(client *azcosmos.Client, databaseName string) error {
	//	databaseName := "adventureworks"

	// sets the name of the database
	databaseProperties := azcosmos.DatabaseProperties{ID: databaseName}

	// creating the database
	ctx := context.TODO()
	_, err := client.CreateDatabase(ctx, databaseProperties, nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func createContainer(client *azcosmos.Client, databaseName, containerName, partitionKey string) error {
	//	databaseName = "adventureworks"
	//	containerName = "customer"
	//	partitionKey = "/customerId"

	databaseClient, err := client.NewDatabase(databaseName) // returns a struct that represents a database
	if err != nil {
		log.Fatal("Failed to create a database client:", err)
	}

	// Setting container properties
	containerProperties := azcosmos.ContainerProperties{
		ID: containerName,
		PartitionKeyDefinition: azcosmos.PartitionKeyDefinition{
			Paths: []string{partitionKey},
		},
	}

	// Setting container options
	throughputProperties := azcosmos.NewManualThroughputProperties(400) //defaults to 400 if not set
	options := &azcosmos.CreateContainerOptions{
		ThroughputProperties: &throughputProperties,
	}

	ctx := context.TODO()
	containerResponse, err := databaseClient.CreateContainer(ctx, containerProperties, options)
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("Container [%v] created. ActivityId %s\n", containerName, containerResponse.ActivityID)

	return nil
}

func createItem(client *azcosmos.Client, databaseName, containerName, partitionKey string, item any) error {
	//	databaseName = "adventureworks"
	//	containerName = "customer"
	//	partitionKey = "1"
	/*
		item = struct {
			ID           string `json:"id"`
			CustomerId   string `json:"customerId"`
			Title        string
			FirstName    string
			LastName     string
			EmailAddress string
			PhoneNumber  string
			CreationDate string
		}{
			ID:           "1",
			CustomerId:   "1",
			Title:        "Mr",
			FirstName:    "Luke",
			LastName:     "Hayes",
			EmailAddress: "luke12@adventure-works.com",
			PhoneNumber:  "879-555-0197",
		}
	*/
	// Create container client
	containerClient, err := client.NewContainer(databaseName, containerName)
	if err != nil {
		return fmt.Errorf("failed to create a container client: %s", err)
	}

	// Specifies the value of the partiton key
	pk := azcosmos.NewPartitionKeyString(partitionKey)

	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	// setting item options upon creating ie. consistency level
	itemOptions := azcosmos.ItemOptions{
		ConsistencyLevel: azcosmos.ConsistencyLevelSession.ToPtr(),
	}
	ctx := context.TODO()
	itemResponse, err := containerClient.CreateItem(ctx, pk, b, &itemOptions)

	if err != nil {
		return err
	}
	log.Printf("Status %d. Item %v created. ActivityId %s. Consuming %v Request Units.\n", itemResponse.RawResponse.StatusCode, pk, itemResponse.ActivityID, itemResponse.RequestCharge)

	return nil
}

func readItem(client *azcosmos.Client, databaseName, containerName, partitionKey, itemId string) error {
	//	databaseName = "adventureworks"
	//	containerName = "customer"
	//	partitionKey = "1"
	//	itemId = "1"

	// Create container client
	containerClient, err := client.NewContainer(databaseName, containerName)
	if err != nil {
		return fmt.Errorf("Failed to create a container client: %s", err)
	}

	// Specifies the value of the partiton key
	pk := azcosmos.NewPartitionKeyString(partitionKey)

	// Read an item
	ctx := context.TODO()
	itemResponse, err := containerClient.ReadItem(ctx, pk, itemId, nil)
	if err != nil {
		return err
	}

	itemResponseBody := struct {
		ID           string `json:"id"`
		CustomerId   string `json:"customerId"`
		Title        string
		FirstName    string
		LastName     string
		EmailAddress string
		PhoneNumber  string
		CreationDate string
	}{}

	err = json.Unmarshal(itemResponse.Value, &itemResponseBody)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(itemResponseBody, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("Read item with customerId %s\n", itemResponseBody.CustomerId)
	fmt.Printf("%s\n", b)

	log.Printf("Status %d. Item %v read. ActivityId %s. Consuming %v Request Units.\n", itemResponse.RawResponse.StatusCode, pk, itemResponse.ActivityID, itemResponse.RequestCharge)

	return nil
}

func deleteItem(client *azcosmos.Client, databaseName, containerName, partitionKey, itemId string) error {
	//	databaseName = "adventureworks"
	//	containerName = "customer"
	//	partitionKey = "1"
	//	itemId = "1"

	// Create container client
	containerClient, err := client.NewContainer(databaseName, containerName)
	if err != nil {
		return fmt.Errorf("Failed to create a container client: %s", err)
	}
	// Specifies the value of the partiton key
	pk := azcosmos.NewPartitionKeyString(partitionKey)

	// Delete an item
	ctx := context.TODO()
	res, err := containerClient.DeleteItem(ctx, pk, itemId, nil)
	if err != nil {
		return err
	}

	log.Printf("Status %d. Item %v deleted. ActivityId %s. Consuming %v Request Units.\n", res.RawResponse.StatusCode, pk, res.ActivityID, res.RequestCharge)

	return nil
}
