//go:build integration
// +build integration

// open settings json or remote settings json
// {
//"go.buildFlags": [
//    "-tags=unit,integration"
//],
//"go.buildTags": "-tags=unit,integration",
//"go.testTags": "-tags=unit,integration"
// }

package cosmosdbbackend

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/clarkezone/geocache/internal"
	"github.com/clarkezone/geocache/pkg/geocacheservice"
	"github.com/clarkezone/geocache/pkg/gpxconverter"
	// "github.com/google/uuid"
)

// TestMain initizlie all tests
func TestMain(m *testing.M) {
	internal.SetupGitRoot()
	clarkezoneLog.Init(logrus.DebugLevel)
	code := m.Run()
	os.Exit(code)
}

func TestReadQuery(t *testing.T) {

}

func TestGetThing(t *testing.T) {
	testCases := []struct {
		name     string
		input    *geocacheservice.Location
		expected DAOSample
	}{
		{
			name: "Test 1",
			input: &geocacheservice.Location{
				Geometry: &geocacheservice.Geometry{
					Type:        "Point",
					Coordinates: []float64{37.4219999, -122.0840575},
				},
				Properties: &geocacheservice.Properties{
					BatteryLevel: 75.0,
					Altitude:     1000,
					BatteryState: "Good",
					Timestamp:    timestamppb.New(time.Date(2023, 4, 4, 12, 0, 0, 0, time.UTC)),
				},
			},
			expected: DAOSample{
				ID:           "1",
				PartitionID:  "1",
				BatteryLevel: 75.0,
				Altitude:     1000,
				Lat:          37.4219999,
				Lon:          -122.0840575,
				BatteryState: "Good",
				Timestamp:    "2023-04-04T12:00:00Z",
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetThing(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %+v, got: %+v", tc.expected, result)
			}
		})
	}
}

func Test_CosmosBootstrap(t *testing.T) {
	endpoint := os.Getenv("AZURE_COSMOS_ENDPOINT")
	if endpoint == "" {
		log.Fatal("AZURE_COSMOS_ENDPOINT could not be found")
	}
	key := os.Getenv("AZURE_COSMOS_KEY")
	if key == "" {
		log.Fatal("AZURE_COSMOS_KEY could not be found")
	}

	cosmosdal, err := CreateCosmosDAL(endpoint, key)
	if err != nil {
		log.Fatal(err)
	}

	var databaseName = "integrationtest"
	var containerName = "geocacheintegrationtest"
	//var partitionKey = "/partitionid"

	// item2 := DAOSample{
	// 	ID:           uuid.New().String(),
	// 	PartitionID:  "1",
	// 	BatteryLevel: 75.0,
	// 	Altitude:     1000,
	// 	Lat:          37.4219999,
	// 	Lon:          -122.0840575,
	// 	BatteryState: "Good",
	// 	Timestamp:    "2023-04-04T12:00:00Z",
	// }

	// err = cosmosdal.CreateDatabase(databaseName)
	// if err != nil {
	// 	log.Fatalf("createDatabase failed: %s\n", err)
	// }

	// err = cosmosdal.CreateContainer(databaseName, containerName, partitionKey)
	// if err != nil {
	// 	log.Fatalf("createContainer failed: %s\n", err)
	// }

	// err = cosmosdal.CreateItem(databaseName, containerName, item2.PartitionID, item2)
	// if err != nil {
	// 	log.Fatalf("createItem failed: %s\n", err)
	// }

	// if err != nil {
	// 	log.Printf("createItem failed: %s\n", err)
	// }

	sql := "SELECT top 10 * FROM geocacheintegrationtest p order by p.Timestamp desc"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err = cosmosdal.Query(databaseName, containerName, "1", sql, ctx)

	//TODO verify returned items

	//	err = readItem(client, databaseName, containerName, item.CustomerId, item.ID)
	//	if err != nil {
	//		log.Printf("readItem failed: %s\n", err)
	//	}

	//	err = deleteItem(client, databaseName, containerName, item.CustomerId, item.ID)
	//	if err != nil {
	//		log.Printf("deleteItem failed: %s\n", err)
	//	}

	//TODO: delete container and database
}

func Test_CosmosQuery(t *testing.T) {
	endpoint := os.Getenv("AZURE_COSMOS_ENDPOINT")
	if endpoint == "" {
		log.Fatal("AZURE_COSMOS_ENDPOINT could not be found")
	}
	key := os.Getenv("AZURE_COSMOS_KEY")
	if key == "" {
		log.Fatal("AZURE_COSMOS_KEY could not be found")
	}

	cosmosdal, err := CreateCosmosDAL(endpoint, key)
	if err != nil {
		log.Fatal(err)
	}

	var databaseName = "integrationtest"
	var containerName = "geocacheintegrationtest"
	var partitionKey = "/partitionid"

	sql := "SELECT top 10 * FROM geocacheintegrationtest p order by p.Timestamp desc"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err = cosmosdal.Query(databaseName, containerName, partitionKey, sql, ctx)

	if err != nil {
		log.Printf("createItem failed: %s\n", err)
	}

}

func Test_CosmosQuerySFOPoints(t *testing.T) {
	endpoint := os.Getenv("AZURE_COSMOS_ENDPOINT")
	if endpoint == "" {
		log.Fatal("AZURE_COSMOS_ENDPOINT could not be found")
	}
	key := os.Getenv("AZURE_COSMOS_KEY")
	if key == "" {
		log.Fatal("AZURE_COSMOS_KEY could not be found")
	}

	cosmosdal, err := CreateCosmosDAL(endpoint, key)
	if err != nil {
		log.Fatal(err)
	}

	var databaseName = "pocketnow"
	var containerName = "geocache"
	var partitionKey = "1"

	sql := "SELECT * FROM geocache c where c.Timestamp >= '2023-05-31T14:00:00Z' AND c.Timestamp <= '2023-05-31T14:55:59Z'"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	returned, err := cosmosdal.Query(databaseName, containerName, partitionKey, sql, ctx)

	if err != nil {
		log.Fatalf("createItem failed: %s\n", err)
	}

	log.Printf("returned: %+v\n", len(returned))

	b := make([]gpxconverter.Point, len(returned))
	for i := range returned {
		b[i] = returned[i]
	}

	gpx, err := gpxconverter.ConvertToGPX(b)
	if err != nil {
		log.Fatalf("createItem failed: %s\n", err)
	}

	log.Printf("gpx: %+v\n", gpx)
}

func Test_Cosmosquery(t *testing.T) {

}
