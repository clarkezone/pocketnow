namespace pocketnow
{
    public class Foo
    {
        public void Configure()
        {
            var builder = WebApplication.CreateBuilder();
            builder.Services.AddControllers();
            var app = builder.Build();
            app.UseHttpsRedirection();
            var pbhost = Environment.GetEnvironmentVariable("HOST") ?? string.Empty;
            var cosmosendpoint = Environment.GetEnvironmentVariable("COSMOS_URL") ?? string.Empty;
            var cosmoskey = Environment.GetEnvironmentVariable("COSMOS_KEY") ?? string.Empty;
            var serviceurl = Environment.GetEnvironmentVariable("SERVICEURL") ?? string.Empty;

            Console.WriteLine($"Username: {cosmosendpoint}");
            Console.WriteLine($"Password: {cosmoskey}");

            pocketnow.CosmosQueryService cosmosQueryService = new();
            cosmosQueryService.Connect(cosmosendpoint, cosmoskey);

            app.MapGet("/", async () =>
            {
            });

        }
    }
}