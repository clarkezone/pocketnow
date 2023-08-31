using pocketnow;

var builder = WebApplication.CreateBuilder(args);

//app.UseHttpsRedirection();

var pocketnowmode = Environment.GetEnvironmentVariable("NOWMODE") ?? string.Empty;
var cosmosendpoint = Environment.GetEnvironmentVariable("COSMOSDB_URL") ?? string.Empty;
var cosmoskey = Environment.GetEnvironmentVariable("COSMOSDB_KEY") ?? string.Empty;
var serviceurl = Environment.GetEnvironmentVariable("SERVICEURL") ?? string.Empty;

Console.WriteLine($"COSMOSDB_URL: {cosmosendpoint}");
Console.WriteLine($"COSMOSDB_KEY: {cosmoskey}");
Console.WriteLine($"SERVICEURL: {serviceurl}");
Console.WriteLine($"NOWMODE: {pocketnowmode}");
builder.Services.AddControllers();
if (string.IsNullOrEmpty(pocketnowmode) ||  pocketnowmode=="FALSE") {
	builder.Services.AddScoped<IGeoQueryService, GeoQueryService>(x => new GeoQueryService(cosmosendpoint, cosmoskey));
}

var app = builder.Build();
if (string.IsNullOrEmpty(pocketnowmode) ||  pocketnowmode=="FALSE") {
	Console.WriteLine($"Starting PocketNow mode");
	app.MapGeoGeoService(); //old frontend
} else {
	Console.WriteLine($"Starting GeoQuery mode");
	app.MapGeoQueries();
}

app.Run();

public partial class Program { }
