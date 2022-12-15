using System.Diagnostics;

class GeoService
{
    const string baseurl = "https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/reverseGeocode";
    Dictionary<(float, float), Root> cityCach = new Dictionary<(float, float), Root>();
    HttpClient http = new HttpClient();

    public async Task<Root?> AddressFromPoint(float lat, float lon)
    {
        var city = await LookupCityCache(lat, lon);
        return city;
    }

    private async Task<Root?> LookupCity(float lat, float lon)
    {
        var result = await http.GetAsync(baseurl + "?location=" + lat + "," + lon + "&outSR=4326&f=json");
        var add = await result.Content.ReadFromJsonAsync<Root>();
        Debug.WriteLine(add?.address.City);
        return add;
    }

    private async Task<Root?> LookupCityCache(float lat, float lon)
    {
        Root? city = null;
        //TODO need to cache city more losely, not for every point
        //TODO contain a max size
        if (cityCach.ContainsKey((lat, lon)))
        {
            city = cityCach[(lat, lon)];
        }
        else
        {
            city = await LookupCity(lat, lon);
            if (city != null)
            {
                cityCach[(lat, lon)] = city;
            }
        }
        return city;
    }
}
