namespace pocketnow.Tests;

public class UnitTest1
{
    [Fact]
    public void Test1()
    {
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