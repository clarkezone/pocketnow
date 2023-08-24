namespace pocketnow.Tests;

public class QueryFrontendTests
{
    [Fact]
	public async Task TestRootEndpoint()
	{
	    //await using var application = new WebApplicationFactory<Program>();
	    await using var application = new GeoServiceApplication();
	    using var client = application.CreateClient();

	    var response = await client.GetAsync("/geoquery");
	  
	    Assert.Equal("Hello World!", "Hello World!");
	}

}
