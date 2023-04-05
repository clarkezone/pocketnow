package cosmosdbbackend

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/clarkezone/geocache/pkg/geocacheservice"
)

type DAOSample struct {
	ID           string `json:"id"`
	PartitionID  string `json:"partitionid"`
	BatteryLevel float64
	Altitude     int32
	Lat          float64
	Lon          float64
	BatteryState string
	Timestamp    string
}

type CosmosDAL struct {
	endpoint string
	key      string
	client   *azcosmos.Client
}

func CreateCosmosDAL(endpoint, key string) (*CosmosDAL, error) {
	cred, err := azcosmos.NewKeyCredential(key)
	if err != nil {
		return nil, err
	}

	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)
	if err != nil {
		return nil, err
	}
	return &CosmosDAL{
		endpoint: endpoint,
		key:      key,
		client:   client,
	}, nil
}

func getThing(lo *geocacheservice.Location) (ds DAOSample) {

	coords := lo.Geometry.Coordinates

	item := DAOSample{
		"1",
		"1",
		lo.Properties.BatteryLevel,
		lo.Properties.Altitude,
		coords[0],
		coords[1],
		lo.Properties.BatteryState,
		lo.Properties.Timestamp.AsTime().Format(time.RFC3339),
	}
	return item
}

func (dal *CosmosDAL) CreateDatabase(databaseName string) error {
	databaseProperties := azcosmos.DatabaseProperties{ID: databaseName}

	ctx := context.TODO()
	_, err := dal.client.CreateDatabase(ctx, databaseProperties, nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (dal *CosmosDAL) CreateContainer(databaseName, containerName, partitionKey string) error {
	databaseClient, err := dal.client.NewDatabase(databaseName) // returns a struct that represents a database
	if err != nil {
		log.Fatal("Failed to create a database client:", err)
	}

	containerProperties := azcosmos.ContainerProperties{
		ID: containerName,
		PartitionKeyDefinition: azcosmos.PartitionKeyDefinition{
			Paths: []string{partitionKey},
		},
	}

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

func (dal *CosmosDAL) CreateItem(databaseName, containerName, partitionKey string, item any) error {
	containerClient, err := dal.client.NewContainer(databaseName, containerName)
	if err != nil {
		return fmt.Errorf("failed to create a container client: %s", err)
	}

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
