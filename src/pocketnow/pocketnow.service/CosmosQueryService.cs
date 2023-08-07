
namespace PocketBase
{
    public class CosmosQueryService
    {
        public Connect()
        {
            using CosmosClient client = new(
                accountEndpoint: Environment.GetEnvironmentVariable("COSMOS_ENDPOINT")!,
                authKeyOrResourceToken: Environment.GetEnvironmentVariable("COSMOS_KEY")!

            );

        }
    }