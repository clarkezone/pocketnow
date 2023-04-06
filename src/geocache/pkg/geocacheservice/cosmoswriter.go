package geocacheservice

import (
	"fmt"
	"os"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"

	"github.com/clarkezone/geocache/pkg/cosmosdbbackend"
)

const (
	dbname        = "pocketnow"
	containername = "geocache"
	partitionkey  = "/partitionid"
)

type CosmosDBWriter struct {
	cdal *cosmosdbbackend.CosmosDAL
}

func NewCosmosDBWriter() (*CosmosDBWriter, error) {
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

	cosmosdal, err := cosmosdbbackend.CreateCosmosDAL(endpoint, key)
	cd.cdal = cosmosdal
	return cd, err
}

func (c *CosmosDBWriter) Process(msg Message) {
	for _, m := range msg.locations.Locations {
		dao := cosmosdbbackend.GetThing(m)
		err := c.cdal.CreateItem(dbname, containername, "1", dao)
		if err != nil {
			clarkezoneLog.Debugf("Error writing to CosmosDB: %v", err)
		}

	}
}
