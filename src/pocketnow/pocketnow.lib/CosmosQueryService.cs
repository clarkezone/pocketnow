
using Microsoft.Azure.Cosmos;

namespace pocketnow
{
    public class CosmosQueryService
    {
        public Container Connect(string Endpoint, string key)
        {
            CosmosClient client = new(
                accountEndpoint: Endpoint,
                authKeyOrResourceToken: key
            );
            //authKeyOrResourceToken: Environment.GetEnvironmentVariable("COSMOS_KEY")!
            var databaseName = "pocketnow";
            var containerName = "geocache";
            var partitionKey = "1";
            var db = client.GetDatabase(databaseName);
            return db.GetContainer(containerName);
        }

        /*
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
        */

        public record Product
        {
            public string ID = "";
            public float Lat = 0;
        }

        public async Task<IEnumerable<Product>> GetGeoLog(Container container)
        {
	//TODO validate strings are valid times
	    var startTime = "2023-05-13T14:45:59Z";
	    var endTime = "2023-05-13T15:15:59Z";
            var sql = "SELECT * FROM geocache c where c.Timestamp >= @starttime AND c.Timestamp <= @endtime";
		var qd = new QueryDefinition(query: sql);
		qd.WithParameter("@starttime", startTime);
		qd.WithParameter("@endtime", endTime);

            // Query multiple items from container
            using FeedIterator<Product> feed = container.GetItemQueryIterator<Product>(
                queryDefinition: qd
            );

            List<Product> products = new();

            // Iterate query result pages
            while (feed.HasMoreResults)
            {
                FeedResponse<Product> response = await feed.ReadNextAsync();

                // Iterate query results
                foreach (Product item in response)
                {
                    products.Add(item);
                }
            }
            return products;
        }

    }
}
