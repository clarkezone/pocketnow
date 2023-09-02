using System.Text.Json;
using System.Diagnostics;

namespace pocketnow
{

    class GeoService
    {
        const string baseurl = "https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/reverseGeocode";
        const string temp = "";
        Dictionary<(float, float), Root?> cityCach = new Dictionary<(float, float), Root?>();
        HttpClient http = new HttpClient();

        public async Task<Root?> AddressFromPoint(float lat, float lon)
        {
	    Console.WriteLine($"Lookup city with {lat}, {lon}");
            var city = await LookupCityCache(lat, lon);
            return city;
        }

        private async Task<Root?> LookupCity(float lat, float lon)
        {
	    var requestString = baseurl + "?location=" + lat + "," + lon + "&outSR=4326&f=json";
	    Console.WriteLine($"LookupCity: {requestString}");
            var result = await http.GetAsync(requestString);
	    var resultstring = await result.Content.ReadAsStringAsync();
	    Console.WriteLine($"Result: {resultstring}");
	    var options = new JsonSerializerOptions {IncludeFields = true};
	    var add = JsonSerializer.Deserialize<Root?>(resultstring,options);
	    Console.WriteLine($"Manual Parsed: {add}");
            return add;
        }

        private async Task<Root?> LookupCityCache(float lat, float lon)
        {
            Root? city = null;
            //TODO need to cache city more losely, not for every point
            //TODO contain a max size
	    Console.WriteLine($"Cache Count: {cityCach.Count}");
            if (cityCach.ContainsKey((lat, lon)))
            {
		Console.WriteLine($"Cache hit for {lat}, {lon}");
                city = cityCach[(lat, lon)];
            }
            else
            {
		Console.WriteLine($"Cache miss for {lat}, {lon}");
                city = await LookupCity(lat, lon);
                if (city != null)
                {
                    cityCach[(lat, lon)] = city;
                }
            }
	    Console.WriteLine($"City from cache: {city}");
            return city;
        }
    }
}
