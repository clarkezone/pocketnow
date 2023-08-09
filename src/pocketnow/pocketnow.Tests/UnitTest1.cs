namespace pocketnow.Tests;

public class UnitTest1
{
    [Fact]
    public void Test1()
    {
        pocketnow.CosmosQueryService cosmosQueryService = new ();
        cosmosQueryService.Connect("https://pocketnow.documents.azure.com:443",
         "78FDsSHV7CNb0cfOnG7TvpFzjjFYhHsx5p0eolg8ZVweVFMfMQUcZGSNqLB72RNwHO9sKHzy7Mo0ACDbU3wLtQ==");
    }
    [Fact]
    public void Test2()
    {
        pocketnow.CosmosQueryService cosmosQueryService = new ();
        var thing = cosmosQueryService.Connect("https://pocketnow.documents.azure.com:443",
         "78FDsSHV7CNb0cfOnG7TvpFzjjFYhHsx5p0eolg8ZVweVFMfMQUcZGSNqLB72RNwHO9sKHzy7Mo0ACDbU3wLtQ==");
         cosmosQueryService.GetGeoLog(thing);
    }
}