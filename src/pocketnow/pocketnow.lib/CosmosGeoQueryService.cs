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

        public async Task<IEnumerable<GeoLogEntry>> GetGeoLog(Container container, DateTime start, DateTime end)
        {
            //TODO validate strings are valid times
//            var startTime = "2023-08-17T00:45:59Z";
//            var endTime = "2023-08-17T13:15:59Z";
            //ISO 8601 UTC
            var startTime = string.Format("{0:yyyy-MM-ddTHH:mm:ss.FFFZ}", start);
            var endTime = string.Format("{0:yyyy-MM-ddTHH:mm:ss.FFFZ}", end);
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
                    var tmp = item.Lat;
                    item.Lat = item.Lon;
                    item.Lon = tmp;
                    logEntries.Add(item);
                }
            }
            return logEntries;
        }

    }
}
