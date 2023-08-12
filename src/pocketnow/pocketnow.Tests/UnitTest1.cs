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

        pocketnow.CosmosQueryService cosmosQueryService = new ();
        cosmosQueryService.Connect("",
         "");
    }
    [Fact]
    public void Test2()
    {
        pocketnow.CosmosQueryService cosmosQueryService = new ();
        var thing = cosmosQueryService.Connect("",
         "");
         cosmosQueryService.GetGeoLog(thing);
    }
}
