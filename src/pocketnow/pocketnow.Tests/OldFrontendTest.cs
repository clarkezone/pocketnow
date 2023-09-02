using System.Text.Json;
namespace pocketnow.Tests;

public class QueryOldFrontendTests
{
    [Fact]
	public void TestOldFrontend()
	{
		var parse = """
		{"address":{"Match_addr":"16632 NE 46th St, Redmond, Washington, 98052","LongLabel":"16632 NE 46th St, Redmond, WA, 98052, USA","ShortLabel":"16632 NE 46th St","Addr_type":"PointAddress","Type":"","PlaceName":"","AddNum":"16632","Address":"16632 NE 46th St","Block":"","Sector":"","Neighborhood":"","District":"","City":"Redmond","MetroArea":"","Subregion":"King County","Region":"Washington","RegionAbbr":"WA","Territory":"","Postal":"98052","PostalExt":"5441","CntryName":"United States","CountryCode":"USA"},"location":{"x":-122.11786204116319,"y":47.650423016236374,"spatialReference":{"wkid":4326,"latestWkid":4326}}}
	""";
	    var options = new JsonSerializerOptions {IncludeFields = true};
	    var parsed = JsonSerializer.Deserialize<Root>(parse);
	    Console.WriteLine($"Manual Parsed: {parsed}");
	    Assert.Equal(parsed.location.x, -122.11786204116319);
	}

}
