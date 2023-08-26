using System.Data;
using System.Security.Cryptography;
using Microsoft.Azure.Cosmos;

namespace pocketnow
{
        /*
        type PointStructFull struct {
            ID            string    `json:"id"`
            PartitionID   string    `json:"partitionid"`
            BatteryLevel  float64   `json:"BatteryLevel"`
            Altitude      int       `json:"Altitude"`
            Lat           float64   `json:"Lat"`
            Lon           float64   `json:"Lon"`
            BatteryState  string    `json:"BatteryState"`
            Timestamp     time.Time `json:"Timestamp"`
            RID           string    `json:"_rid"`
            Self          string    `json:"_self"`
            Etag          string    `json:"_etag"`
            Attachments   string    `json:"_attachments"`
            TimestampUnix int       `json:"_ts"`
        }
        */
        //
        // TODO: Rename and populate

    public class GeoLogEntry
    {
        public string ID { get; set; }
        public float Lat { get; set; }
        public float Lon { get; set; }
        public string BatteryState { get; set; }
        public int Altitude { get; set; }
        public DateTime Timestamp { get; set; }
    }

    public interface IGeoQueryService
    {
        public Task<IEnumerable<GeoLogEntry>> GetGeoLog(DateTime start, DateTime end);
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

        public async Task<IEnumerable<GeoLogEntry>> GetGeoLog(DateTime start, DateTime end)
        {
            _container = _queryService.Connect(CosmosUrl, CosmosKey);
            return await _queryService.GetGeoLog(_container, start, end);
        }
    }
}
