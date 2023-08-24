using System.Security.Cryptography;
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
        //
        // TODO: Rename and populate

        public async Task<IEnumerable<GeoLogEntry>> GetGeoLog(Container container)
        {
            //TODO validate strings are valid times
            var startTime = "2023-08-17T00:45:59Z";
            var endTime = "2023-08-17T13:15:59Z";
            var sql = "SELECT * FROM geocache c where c.Timestamp >= @starttime AND c.Timestamp <= @endtime";
            var qd = new QueryDefinition(query: sql);
            qd.WithParameter("@starttime", startTime);
            qd.WithParameter("@endtime", endTime);

            // Query multiple items from container
            using FeedIterator<GeoLogEntry> feed = container.GetItemQueryIterator<GeoLogEntry>(
                queryDefinition: qd
            );

            List<GeoLogEntry> logEntries = new();

            // Iterate query result pages
            while (feed.HasMoreResults)
            {
                FeedResponse<GeoLogEntry> response = await feed.ReadNextAsync();

                // Iterate query results
                foreach (GeoLogEntry item in response)
                {
                    logEntries.Add(item);
                }
            }
            return logEntries;
        }

    }
}
