using PocketBase;
using System.Diagnostics;
using System.Net;
using System.Text.Json;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
var app = builder.Build();

app.UseHttpsRedirection();
var pbhost = Environment.GetEnvironmentVariable("HOST") ?? string.Empty;
var un = Environment.GetEnvironmentVariable("UN") ?? string.Empty;
var pw = Environment.GetEnvironmentVariable("PW") ?? string.Empty;

Console.WriteLine($"Host: {pbhost}");
Console.WriteLine($"Username: {un}");
Console.WriteLine($"Password: {pw}");

Authorizor a = new Authorizor(pbhost, un, pw);
var result = await a.Authorize();
Console.WriteLine($"Result from auth {result}");
var client = a.GetClient();

app.MapGet("/geocode", async () =>
{
    if (client == null)
    {
        return Results.Unauthorized();
    }
    var root = await client.GetRecords<currentsitrep>("currentsitrep");
    if (root != null && root.items.Length > 0)
    {
        var item = root.items[0];
        HttpClient http = new HttpClient();
        var lat = item.lat;
        var lon = item.lon;
        var result = await http.GetAsync("https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/reverseGeocode?location=" + lat + "," + lon + "&outSR=4326&f=json");
        var add = await result.Content.ReadFromJsonAsync<Root>();
        Debug.WriteLine(add?.address.City);
        var r = new Returned()
        {
            City = add?.address.City??"",
            PhoneStatus = item.batterystate,
            Batterylevel = item.batterylevel
        };
        return Results.Json(r);
    }
    return Results.NotFound();
});

app.Run();

internal record WeatherForecast(DateTime Date, int TemperatureC, string? Summary)
{
    public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);
}

public class Returned
{
    public string City { get; set; }
    public string PhoneStatus { get; set; }

    public float Batterylevel { get; set; }
}

public class AddressRec
{
    public string Match_addr { get; set; }
    public string LongLabel { get; set; }
    public string ShortLabel { get; set; }
    public string Addr_type { get; set; }
    public string Type { get; set; }
    public string PlaceName { get; set; }
    public string AddNum { get; set; }
    public string Address { get; set; }
    public string Block { get; set; }
    public string Sector { get; set; }
    public string Neighborhood { get; set; }
    public string District { get; set; }
    public string City { get; set; }
    public string MetroArea { get; set; }
    public string Subregion { get; set; }
    public string Region { get; set; }
    public string RegionAbbr { get; set; }
    public string Territory { get; set; }
    public string Postal { get; set; }
    public string PostalExt { get; set; }
    public string CntryName { get; set; }
    public string CountryCode { get; set; }
}

public class Location
{
    public double x { get; set; }
    public double y { get; set; }
    public SpatialReference spatialReference { get; set; }
}

public class Root
{
    public AddressRec address { get; set; }
    public Location location { get; set; }
}

public class SpatialReference
{
    public int wkid { get; set; }
    public int latestWkid { get; set; }
}
public class currentsitrep
{
    public int page { get; set; }
    public int perPage { get; set; }
    public int totalItems { get; set; }
    public int totalPages { get; set; }
    public Item[] items { get; set; }
}

public class Item
{
    public int altitude { get; set; }
    public float batterylevel { get; set; }
    public string batterystate { get; set; }
    public string collectionId { get; set; }
    public string collectionName { get; set; }
    public string created { get; set; }
    public string id { get; set; }
    public float lat { get; set; }
    public float lon { get; set; }
    public string timestamp { get; set; }
    public string updated { get; set; }
    public string wifi { get; set; }
}