
using Microsoft.Azure.Cosmos;

namespace PocketBase
{
    public class CosmosQueryService
    {
        public void Connect()
        {
            using CosmosClient client = new(
                accountEndpoint: Environment.GetEnvironmentVariable("COSMOS_ENDPOINT")!,
                authKeyOrResourceToken: Environment.GetEnvironmentVariable("COSMOS_KEY")!

            );

        }
    }
}