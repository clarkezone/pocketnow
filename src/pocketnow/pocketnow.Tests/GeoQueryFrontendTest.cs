namespace pocketnow.Tests;

public class QueryFrontendTests
{
    [Fact]
	public async Task TestRootEndpoint()
	{
	    //await using var application = new WebApplicationFactory<Program>();
	    await using var application = new GeoServiceApplication();
	    using var client = application.CreateClient();

		var start = DateTime.Now;

		var pa = new GeoQueryParams() {QueryStart = start};

	    var response = await client.PostAsJsonAsync("/geoquery", pa);
	  
	    Assert.Equal(pa.QueryStart, start);
	}

}
