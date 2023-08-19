using pocketnow;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
var app = builder.Build();


//app.UseHttpsRedirection();

                        var pbhost = Environment.GetEnvironmentVariable("HOST") ?? string.Empty;
                        var cosmosendpoint = Environment.GetEnvironmentVariable("COSMOSDB_URL") ?? string.Empty;
                        var cosmoskey = Environment.GetEnvironmentVariable("COSMOSDB_KEY") ?? string.Empty;
                        var serviceurl = Environment.GetEnvironmentVariable("SERVICEURL") ?? string.Empty;

                        Console.WriteLine($"COSMOSDB_URL: {cosmosendpoint}");
                        Console.WriteLine($"COSMOSDB_KEY: {cosmoskey}");
//builder.Service.Add()
app.MapGeoQueries();

app.Run();
