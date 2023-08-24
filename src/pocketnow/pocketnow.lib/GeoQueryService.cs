using System.Security.Cryptography;
using Microsoft.Azure.Cosmos;

namespace pocketnow
{
    public class GeoLogEntry
    {
        public string ID { get; set; }
        public float Lat { get; set; }
    }

    public interface IGeoQueryService
    {
        public Task<IEnumerable<GeoLogEntry>> GetGeoLog();
    }

    public class GeoQueryService : IGeoQueryService
    {
        public GeoQueryService(string url, string key)
        {
            CosmosUrl = url;
            CosmosKey = key;
            _queryService = new();
        }

        public string CosmosUrl { get; set; }
        public string CosmosKey { get; set; }

        CosmosQueryService _queryService;
        Container? _container;

        public async Task<IEnumerable<GeoLogEntry>> GetGeoLog()
        {
            _container = _queryService.Connect(CosmosUrl, CosmosKey);
            return await _queryService.GetGeoLog(_container);
        }
    }
}
