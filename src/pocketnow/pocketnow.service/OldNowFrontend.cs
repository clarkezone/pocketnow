using System.Diagnostics;
using Grpc.Net.Client;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace pocketnow
{
    public static class OldFrontend
    {
        public static RouteGroupBuilder MapGeoGeoService(this IEndpointRouteBuilder routes)
        {
            var group = routes.MapGroup("/");
            group.MapGet("/", async () =>
                {
                    var serviceurl = Environment.GetEnvironmentVariable("SERVICEURL") ?? string.Empty;
                    Console.WriteLine($"SERVICEURL: {serviceurl}");

                    var geoService = new GeoService();
                    var channel = GrpcChannel.ForAddress(serviceurl);
                    GrpcGeoCacheService.GeoCacheService.GeoCacheServiceClient wlient = new GrpcGeoCacheService.GeoCacheService.GeoCacheServiceClient(channel);
                    try
                    {
                        var last = await wlient.GetLastLocationAsync(new GrpcGeoCacheService.Empty());
			var x = (float)last.Geometry.Coordinates[0];
			var y = (float)last.Geometry.Coordinates[1];
			Console.WriteLine($"Got point: {last}, x:{x} y:{y}");

                        var address = await geoService.AddressFromPoint(x, y);
                        Returned r = new Returned(
                            address?.address?.City ?? "",
                            address?.address?.Neighborhood ?? "",
                            address?.address?.CountryCode ?? "",
                            address?.address?.MetroArea ?? "",
                            address?.address?.Postal ?? "",
                            last.Properties?.BatteryState ?? "",
                            last.Properties?.Wifi ?? "",
                            (float)(last.Properties?.BatteryLevel ?? 0.0),
                            last.Properties?.Timestamp.ToDateTime() ?? DateTime.MinValue) ;
                        return Results.Json(r);
                    }
                    catch (Exception ex)
                    {
                        Console.WriteLine(ex.Message);
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

    public record Returned(
        string City,
        string Neighborhood,
        string MetroArea,
        string PhoneStatus,
        string Postal,
        string Country,
        string Wifi,
        float Batterylevel,
        DateTime TimeStamp){
    };

    public record AddressRec(string Match_addr,
                          string LongLabel,
                          string ShortLabel,
                          string Addr_type,
                          string Type,
                          string PlaceName,
                          string AddNum,
                          string Address,
                          string Block,
                          string Sector,
                          string Neighborhood,
                          string District,
                          string City,
                          string MetroArea,
                          string Subregion,
                          string Region,
                          string RegionAbbr,
                          string Territory,
                          string Postal,
                          string PostalExt,
                          string CntryName,
                          string CountryCode) {

			  }

    public record Location(
        double x,
        double y,
        SpatialReference spatialReference) {};

    public record Root(AddressRec address, Location location){};

    public record SpatialReference(int wk, int lw){};

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
