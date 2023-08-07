
using Microsoft.Azure.Cosmos;

namespace pocketnow
{
    public class CosmosQueryService
    {
        public void Connect(string Endpoint, string key)
        {
            using CosmosClient client = new(
                accountEndpoint: Endpoint,
                authKeyOrResourceToken: key
            );
            //authKeyOrResourceToken: Environment.GetEnvironmentVariable("COSMOS_KEY")!
            var databaseName = "pocketnow";
            var containerName = "geocache";
            var partitionKey = "1";
            var db = client.GetDatabase(databaseName);
            db.GetContainer(containerName);

        }
    }
}