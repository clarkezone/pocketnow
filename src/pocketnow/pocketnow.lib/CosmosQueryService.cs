using System.Security.Cryptography;
using Microsoft.Azure.Cosmos;

namespace pocketnow
{
    public class Product
    {
        public string ID { get; set; }
        public float Lat { get; set; }
    }

    public interface IMyDependency
    {
        public Task<IEnumerable<Product>> GetGeoLog();
    }

    public class MyDependency : IMyDependency
    {
        public MyDependency(string url, string key)
        {
            CosmosUrl = url;
            CosmosKey = key;
            _queryService = new();
        }

        public string CosmosUrl { get; set; }
        public string CosmosKey { get; set; }

        CosmosQueryService _queryService;
        Container? _container;

        public async Task<IEnumerable<Product>> GetGeoLog()
        {
            _container = _queryService.Connect(CosmosUrl, CosmosKey);
            return await _queryService.GetGeoLog(_container);
        }
    }

    /*
	public static class MyConfigServiceCollectionExtensions
    {

        public static IServiceCollection AddMyDependencyGroup(
             this IServiceCollection services, string url, string key)
        {
            services.AddScoped<MyDependency>();

            return services;
        }
    }
    */

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

        public async Task<IEnumerable<Product>> GetGeoLog(Container container)
        {
            //TODO validate strings are valid times
            var startTime = "2023-08-17T00:45:59Z";
            var endTime = "2023-08-17T13:15:59Z";
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
