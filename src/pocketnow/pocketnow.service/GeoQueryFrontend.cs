namespace pocketnow
{

    public static class GeoQueryFrontend
    {
        public static RouteGroupBuilder MapGeoQueries(this IEndpointRouteBuilder routes)
        {
            var group = routes.MapGroup("/geoquery");
            group.MapGet("/", async (HttpRequest req, IGeoQueryService dep) =>
                    {
                        // TODO: consider caching the setup
                        Console.WriteLine($"from: {req.Query["from"]}");
                        Console.WriteLine($"to: {req.Query["to"]}");
                        Console.WriteLine("Dependency is null=" + dep == null);
                        return Results.Json(await dep.GetGeoLog());
                    });
            return group;
        }

    }
}
