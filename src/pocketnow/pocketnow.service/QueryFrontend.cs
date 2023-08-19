namespace pocketnow
{

    public static class QueryFrontend
    {
        public static RouteGroupBuilder MapGeoQueries(this IEndpointRouteBuilder routes)
        {
            var group = routes.MapGroup("/geoquery");
            group.MapGet("/", async (HttpRequest req) =>
                    {
		    // TODO: consider caching the setup
			Console.WriteLine($"from: {req.Query["from"]}");
			Console.WriteLine($"to: {req.Query["to"]}");
			// TODO validate from, to via unittest
			// TODO: get CosmosQueryService from services
//                        pocketnow.CosmosQueryService cosmosQueryService = new();
//                        var mydb = cosmosQueryService.Connect(cosmosendpoint, cosmoskey);
//                        return await cosmosQueryService.GetGeoLog(mydb);
                    });
            return group;
        }

    }
}
