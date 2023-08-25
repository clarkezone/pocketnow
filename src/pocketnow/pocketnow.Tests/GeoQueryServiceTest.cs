namespace pocketnow.Tests;
using System.Diagnostics;

public class GeoServiceTest
{
    [Fact (Skip = "Integration")]
    public void TestCosmosEnvVars()
    {
        var cosmosendpoint = Environment.GetEnvironmentVariable("COSMOSDB_URL") ?? string.Empty;
        var cosmoskey = Environment.GetEnvironmentVariable("COSMOSDB_KEY") ?? string.Empty;

        Debug.WriteLine("" + cosmosendpoint);

        Assert.NotEmpty(cosmosendpoint);
        Assert.NotEmpty(cosmoskey);
    }

    // [Fact (Skip = "Integration")]
    [Fact ]
    public async Task TestQueryGeoLog()
    {
        var cosmosendpoint = Environment.GetEnvironmentVariable("COSMOSDB_URL") ?? string.Empty;
        var cosmoskey = Environment.GetEnvironmentVariable("COSMOSDB_KEY") ?? string.Empty;
        pocketnow.CosmosQueryService cosmosQueryService = new ();
        var container = cosmosQueryService.Connect(cosmosendpoint, cosmoskey);
        var result = await cosmosQueryService.GetGeoLog(container, DateTime.UtcNow-TimeSpan.FromHours(1), DateTime.UtcNow);
        Console.WriteLine($"Found georecords:{result.Count()}");

        //TODO: verify result contains correct numbers
    }
}