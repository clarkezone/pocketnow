package cosmosdbbackend

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/google/uuid"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"

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
	clarkezoneLog.Debugf("CreateCosmosDAL: endopint: %v, key: %v", endpoint, key)
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

func GetThing(lo *geocacheservice.Location) (ds DAOSample) {

	coords := lo.Geometry.Coordinates

	item := DAOSample{
		uuid.New().String(),
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

func (dal *CosmosDAL) ReadItem(databaseName, containerName, partitionKey, itemId string) error {
	//	databaseName = "adventureworks"
	//	containerName = "customer"
	//	partitionKey = "1"
	//	itemId = "1"

	// Create container client
	containerClient, err := dal.client.NewContainer(databaseName, containerName)
	if err != nil {
		return fmt.Errorf("failed to create a container client: %s", err)
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
		return fmt.Errorf("failed to create a container client: %s", err)
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

type PointStructFull struct {
	ID            string    `json:"id"`
	PartitionID   string    `json:"partitionid"`
	BatteryLevel  float64   `json:"BatteryLevel"`
	Altitude      int       `json:"Altitude"`
	Lat           float64   `json:"Lat"`
	Lon           float64   `json:"Lon"`
	BatteryState  string    `json:"BatteryState"`
	Timestamp     time.Time `json:"Timestamp"`
	RID           string    `json:"_rid"`
	Self          string    `json:"_self"`
	Etag          string    `json:"_etag"`
	Attachments   string    `json:"_attachments"`
	TimestampUnix int       `json:"_ts"`
}

func (p PointStructFull) GetLat() float64 {
	return p.Lat
}

func (p PointStructFull) GetLon() float64 {
	return p.Lon
}

func (p PointStructFull) GetTimestamp() time.Time {
	return p.Timestamp
}

// TODO return a slice of PointStruct
// TODO write a unittest
// TODO test to take slice of PointStruct and convert to GPX
// TODO make it work with a supplied date range
func (dal *CosmosDAL) Query(databaseName, containerName, partitionKey, sql string, ctx context.Context) ([]PointStructFull, error) {
	pk := azcosmos.NewPartitionKeyString(partitionKey)
	containerClient, err := dal.client.NewContainer(databaseName, containerName)
	if err != nil {
		return nil, fmt.Errorf("failed to create a container client: %s", err)
	}
	queryPager := containerClient.NewQueryItemsPager(sql, pk, nil)
	var resturnval []PointStructFull
	for queryPager.More() {
		queryResponse, err := queryPager.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Query returned %d items\n", len(queryResponse.Items))
		for _, item := range queryResponse.Items {
			var point PointStructFull
			err := json.Unmarshal(item, &point)
			if err != nil {
				return nil, err
			}
			tmp := point.Lat
			point.Lon = point.Lat
			point.Lat = tmp
			resturnval = append(resturnval, point)
			//var itemResponseBody map[string]interface{}
			//json.Unmarshal(item, &itemResponseBody)
		}

	}
	return resturnval, nil
}
