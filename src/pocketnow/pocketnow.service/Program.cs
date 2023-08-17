using pocketnow;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
var app = builder.Build();

app.UseHttpsRedirection();

app.MapGeoQueries();

app.Run();
