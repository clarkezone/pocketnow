using System.ComponentModel.DataAnnotations;
using MiniValidation;

namespace pocketnow
{

    public class GeoQueryParams {
	[Required, DataType(DataType.DateTime)]
        public DateTime QueryStart { get; set; }

	[Required, DataType(DataType.DateTime)]
        public DateTime QueryEnd { get; set; }
    }

    public static class GeoQueryFrontend
    {
        public static RouteGroupBuilder MapGeoQueries(this IEndpointRouteBuilder routes)
        {
            var group = routes.MapGroup("/geoquery");
            group.MapPost("/", async (GeoQueryParams pa, IGeoQueryService dep) =>
                    {
		    	if (!MiniValidator.TryValidate(pa, out var errors))
				{
			return Results.ValidationProblem(errors);
			}

                        Console.WriteLine(string.Format("{0:yyyy-MM-ddTHH:mm:ss.FFFZ}", pa.QueryStart));
                        Console.WriteLine(string.Format("{0:yyyy-MM-ddTHH:mm:ss.FFFZ}", pa.QueryEnd));
//                        Console.WriteLine($"to: {req.Query["to"]}");
                        Console.WriteLine("Dependency is null=" + dep == null);
                        return Results.Json(await dep.GetGeoLog(pa.QueryStart, pa.QueryEnd));
                    });
            return group;
        }

    }
}
