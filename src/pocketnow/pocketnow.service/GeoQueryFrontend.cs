namespace pocketnow
{

    public class GeoQueryParams {
        public DateTime QueryStart { get; set; }
    }

    public static class GeoQueryFrontend
    {
        public static RouteGroupBuilder MapGeoQueries(this IEndpointRouteBuilder routes)
        {
            var group = routes.MapGroup("/geoquery");
            group.MapPost("/", async (GeoQueryParams pa, IGeoQueryService dep) =>
                    {
                        // TODO: consider caching the setup
                        Console.WriteLine(pa.QueryStart);
//                        Console.WriteLine($"to: {req.Query["to"]}");
                        Console.WriteLine("Dependency is null=" + dep == null);
                        return Results.Json(await dep.GetGeoLog());
                    });
            return group;
        }

    }
}
