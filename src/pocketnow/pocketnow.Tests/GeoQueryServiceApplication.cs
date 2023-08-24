namespace pocketnow.Tests;


class TestGeoQueryService : IGeoQueryService {
   public async Task<IEnumerable<GeoLogEntry>> GetGeoLog(DateTime start, DateTime end) {
    return new List<GeoLogEntry> {};
   }
}

public class GeoServiceApplication : WebApplicationFactory<Program> {
protected override IHost CreateHost(IHostBuilder builder)
    {
        //var myAppSettings = builder.Configuration.Get<MyAppSettings>();
            
        builder.ConfigureServices(services =>
        {
            Debug.WriteLine("");
            services.AddScoped<IGeoQueryService, TestGeoQueryService>(x => new TestGeoQueryService());
        });

        return base.CreateHost(builder);
    }

}