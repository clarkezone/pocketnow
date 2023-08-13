namespace pocketnow.Tests;
using System.Diagnostics;

public class UnitTest1
{
    [Fact]
    public void Test1()
    {
        var cosmosendpoint = Environment.GetEnvironmentVariable("COSMOSDB_URL") ?? string.Empty;
        var cosmoskey = Environment.GetEnvironmentVariable("COSMOSDB_KEY") ?? string.Empty;

        Debug.WriteLine("" + cosmosendpoint);

        Assert.NotEmpty(cosmosendpoint);
        Assert.NotEmpty(cosmoskey);
    }

    [Fact]
    public void Test2()
    {
        var cosmosendpoint = Environment.GetEnvironmentVariable("COSMOSDB_URL") ?? string.Empty;
        var cosmoskey = Environment.GetEnvironmentVariable("COSMOSDB_KEY") ?? string.Empty;
        pocketnow.CosmosQueryService cosmosQueryService = new ();
        var thing = cosmosQueryService.Connect(cosmosendpoint, cosmoskey);
        // cosmosQueryService.GetGeoLog(thing);
    }
}
