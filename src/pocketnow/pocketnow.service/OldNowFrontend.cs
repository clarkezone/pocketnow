using System.Diagnostics;
using Grpc.Net.Client;

namespace pocketnow
{
    public static class OldFrontend
    {
        public static RouteGroupBuilder MapGeoGeoService(this IEndpointRouteBuilder routes)
        {

            var group = routes.MapGroup("/oldfrontend");
            group.MapGet("/", async (CosmosQueryService db) =>
                {

                    var pbhost = Environment.GetEnvironmentVariable("HOST") ?? string.Empty;
                    var un = Environment.GetEnvironmentVariable("UN") ?? string.Empty;
                    var pw = Environment.GetEnvironmentVariable("PW") ?? string.Empty;
                    var serviceurl = Environment.GetEnvironmentVariable("SERVICEURL") ?? string.Empty;
                    Console.WriteLine($"Host: {pbhost}");
                    Console.WriteLine($"Username: {un}");
                    Console.WriteLine($"Password: {pw}");
                    Console.WriteLine($"SERVICEURL: {serviceurl}");

                    var geoService = new GeoService();
                    var channel = GrpcChannel.ForAddress(serviceurl);
                    GrpcGeoCacheService.GeoCacheService.GeoCacheServiceClient wlient = new GrpcGeoCacheService.GeoCacheService.GeoCacheServiceClient(channel);
                    try
                    {
                        var last = await wlient.GetLastLocationAsync(new GrpcGeoCacheService.Empty());

                        var address = await geoService.AddressFromPoint((float)last.Geometry.Coordinates[0], (float)last.Geometry.Coordinates[1]);
                        Returned r = new Returned()
                        {
                            City = address?.address?.City ?? "",
                            Neighborhood = address?.address?.Neighborhood ?? "",
                            Country = address?.address?.CountryCode ?? "",
                            MetroArea = address?.address?.MetroArea ?? "",
                            Postal = address?.address?.Postal ?? "",
                            PhoneStatus = last.Properties?.BatteryState ?? "",
                            Wifi = last.Properties?.Wifi ?? "",
                            Batterylevel = (float)(last.Properties?.BatteryLevel ?? 0.0),
                            TimeStamp = last.Properties?.Timestamp.ToDateTime() ?? DateTime.MinValue,
                        };
                        return Results.Json(r);
                    }
                    catch (Exception ex)
                    {
                        Debug.WriteLine(ex.Message);
                        return Results.NotFound();
                    }
                });
            return group;
        }
    }
    internal record WeatherForecast(DateTime Date, int TemperatureC, string? Summary)
    {
        public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);
    }

    public record Returned
    {
        public string City = "";
        public string Neighborhood = "";
        public string MetroArea = "";
        public string PhoneStatus = "";
        public string Postal = "";
        public string Country = "";
        public string Wifi = "";
        public float Batterylevel = 0;
        public string BatterylevelString = "";
        public DateTime TimeStamp = DateTime.MinValue;
    }

    public record AddressRec
    {
        public string Match_addr = "";
        public string LongLabel = "";
        public string ShortLabel = "";
        public string Addr_type = "";
        public string Type = "";
        public string PlaceName = "";
        public string AddNum = "";
        public string Address = "";
        public string Block = "";
        public string Sector = "";
        public string Neighborhood = "";
        public string District = "";
        public string City = "";
        public string MetroArea = "";
        public string Subregion = "";
        public string Region = "";
        public string RegionAbbr = "";
        public string Territory = "";
        public string Postal = "";
        public string PostalExt = "";
        public string CntryName = "";
        public string CountryCode = "";
    }

    public record Location
    {
        public double x = 0;
        public double y = 0;
        public SpatialReference spatialReference = new SpatialReference();
    }

    public record Root
    {
        public AddressRec address = new AddressRec();
        public Location location = new Location();
    }

    public record SpatialReference
    {
        public int wkid { get; set; }
        public int latestWkid { get; set; }
    }

    public record Item
    {
        public int altitude = 0;
        public float batterylevel = 0;
        public string batterystate = "";
        public string collectionId = "";
        public string collectionName = "";
        public string created = "";
        public string id = "";
        public float lat = 0;
        public float lon = 0;
        public string timestamp = "";
        public string updated = "";
        public string wifi = "";
    }
}