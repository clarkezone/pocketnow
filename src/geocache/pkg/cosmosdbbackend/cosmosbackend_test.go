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
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/sirupsen/logrus"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/clarkezone/geocache/internal"
	"github.com/clarkezone/geocache/pkg/geocacheservice"
)

// TestMain initizlie all tests
func TestMain(m *testing.M) {
	internal.SetupGitRoot()
	clarkezoneLog.Init(logrus.DebugLevel)
	code := m.Run()
	os.Exit(code)
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
			result := getThing(tc.input)
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

	var databaseName = "adventureworks"
	var containerName = "customer"
	var partitionKey = "/customerId"

	item := struct {
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

	cred, err := azcosmos.NewKeyCredential(key)
	if err != nil {
		log.Fatal("Failed to create a credential: ", err)
	}

	// Create a CosmosDB client
	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)
	if err != nil {
		log.Fatal("Failed to create Azure Cosmos DB db client: ", err)
	}

	err = createDatabase(client, databaseName)
	if err != nil {
		log.Printf("createDatabase failed: %s\n", err)
	}

	err = createContainer(client, databaseName, containerName, partitionKey)
	if err != nil {
		log.Printf("createContainer failed: %s\n", err)
	}

	err = createItem(client, databaseName, containerName, item.CustomerId, item)
	if err != nil {
		log.Printf("createItem failed: %s\n", err)
	}

	//	err = readItem(client, databaseName, containerName, item.CustomerId, item.ID)
	//	if err != nil {
	//		log.Printf("readItem failed: %s\n", err)
	//	}

	//	err = deleteItem(client, databaseName, containerName, item.CustomerId, item.ID)
	//	if err != nil {
	//		log.Printf("deleteItem failed: %s\n", err)
	//	}

}
