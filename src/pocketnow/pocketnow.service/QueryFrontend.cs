namespace pocketnow
{
    public static class QueryFrontend
    {
        public static RouteGroupBuilder MapGeoQueries(this IEndpointRouteBuilder routes)
        {
            var group = routes.MapGroup("/geoquery");
            group.MapGet("/", async (CosmosQueryService db) =>
                    {
                        var pbhost = Environment.GetEnvironmentVariable("HOST") ?? string.Empty;
                        var cosmosendpoint = Environment.GetEnvironmentVariable("COSMOS_URL") ?? string.Empty;
                        var cosmoskey = Environment.GetEnvironmentVariable("COSMOS_KEY") ?? string.Empty;
                        var serviceurl = Environment.GetEnvironmentVariable("SERVICEURL") ?? string.Empty;

                        Console.WriteLine($"Username: {cosmosendpoint}");
                        Console.WriteLine($"Password: {cosmoskey}");
                        pocketnow.CosmosQueryService cosmosQueryService = new();
                        var mydb = cosmosQueryService.Connect(cosmosendpoint, cosmoskey);
                        return await cosmosQueryService.GetGeoLog(mydb);
                    });
            return group;
        }

// This is not used
        private static void Configure()
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
        }
    }
}
