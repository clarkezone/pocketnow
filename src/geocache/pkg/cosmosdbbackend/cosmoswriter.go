package cosmosdbbackend

import (
	"fmt"
	"os"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"

	"github.com/clarkezone/geocache/pkg/geocacheservice"
)

const (
	dbname        = "pocketnow"
	containername = "geocache"
	partitionkey  = "/partitionid"
)

type CosmosDBWriter struct {
	cdal *CosmosDAL
}

func NewCosmosDBWriter() (geocacheservice.MessageProcessor, error) {
	cd := &CosmosDBWriter{}
	// TODO replace this code with VIPER
	endpoint := os.Getenv("AZURE_COSMOS_ENDPOINT")
	if endpoint == "" {
		return nil, fmt.Errorf("AZURE_COSMOS_ENDPOINT could not be found")
	}
	key := os.Getenv("AZURE_COSMOS_KEY")
	if key == "" {
		return nil, fmt.Errorf("AZURE_COSMOS_KEY could not be found")
	}

	cosmosdal, err := CreateCosmosDAL(endpoint, key)
	cd.cdal = cosmosdal
	return cd, err
}

func (c *CosmosDBWriter) Process(msg geocacheservice.Message) {
	for _, m := range msg.Locations.Locations {
		dao := GetThing(m)
		err := c.cdal.CreateItem(dbname, containername, "1", dao)
		if err != nil {
			clarkezoneLog.Debugf("Error writing to CosmosDB: %v", err)
		}

	}
}
