using pocketnow;

var builder = WebApplication.CreateBuilder(args);

//app.UseHttpsRedirection();

var pbhost = Environment.GetEnvironmentVariable("HOST") ?? string.Empty;
var cosmosendpoint = Environment.GetEnvironmentVariable("COSMOSDB_URL") ?? string.Empty;
var cosmoskey = Environment.GetEnvironmentVariable("COSMOSDB_KEY") ?? string.Empty;
var serviceurl = Environment.GetEnvironmentVariable("SERVICEURL") ?? string.Empty;

Console.WriteLine($"COSMOSDB_URL: {cosmosendpoint}");
Console.WriteLine($"COSMOSDB_KEY: {cosmoskey}");
builder.Services.AddControllers();
builder.Services.AddScoped<IGeoQueryService, GeoQueryService>(x => new GeoQueryService(cosmosendpoint, cosmoskey));
var app = builder.Build();
app.MapGeoQueries();
app.MapGeoGeoService(); //old frontend

app.Run();

public partial class Program { }
