package cosmosdbbackend

import (
	"log"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/sirupsen/logrus"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"

	"github.com/clarkezone/geocache/internal"
)

// TestMain initizlie all tests
func TestMain(m *testing.M) {
	internal.SetupGitRoot()
	clarkezoneLog.Init(logrus.DebugLevel)
	code := m.Run()
	os.Exit(code)
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
