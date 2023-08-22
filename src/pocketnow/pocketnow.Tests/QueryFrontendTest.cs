using Microsoft.AspNetCore.Mvc.Testing;

namespace pocketnow.Tests;

public class QueryFrontendTests
{
    [Fact]
    public void TestDateConversion()
    {
        var startTime = "2023-08-17T00:45:59Z";
        var endTime = "2023-08-17T13:15:59Z";
        DateTime.Parse(startTime);
    }


    [Fact]
	public async Task TestRootEndpoint()
	{
	    await using var application = new WebApplicationFactory<Program>();
	    using var client = application.CreateClient();

	    var response = await client.GetStringAsync("/");
	  
	    Assert.Equal("Hello World!", response);
	}

}
