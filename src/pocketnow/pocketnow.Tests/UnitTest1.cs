namespace pocketnow.Tests;
using System.Diagnostics;

public class UnitTest1
{
    [Fact]
    public void TestCosmosEnvVars()
    {
        var cosmosendpoint = Environment.GetEnvironmentVariable("COSMOSDB_URL") ?? string.Empty;
        var cosmoskey = Environment.GetEnvironmentVariable("COSMOSDB_KEY") ?? string.Empty;

        Debug.WriteLine("" + cosmosendpoint);

        Assert.NotEmpty(cosmosendpoint);
        Assert.NotEmpty(cosmoskey);
    }

    [Fact]
    public void TestQueryGeoLog()
    {
        var cosmosendpoint = Environment.GetEnvironmentVariable("COSMOSDB_URL") ?? string.Empty;
        var cosmoskey = Environment.GetEnvironmentVariable("COSMOSDB_KEY") ?? string.Empty;
        pocketnow.CosmosQueryService cosmosQueryService = new ();
        var container = cosmosQueryService.Connect(cosmosendpoint, cosmoskey);
        cosmosQueryService.GetGeoLog(container);
    }
}
