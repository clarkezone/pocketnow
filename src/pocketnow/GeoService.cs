using System.Diagnostics;

class GeoService
{
    const string baseurl = "https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/reverseGeocode";
    Dictionary<(float, float), string> cityCach = new Dictionary<(float, float), string>();
    HttpClient http = new HttpClient();

    public async Task<string> AddressFromPoint(Item item, float lat, float lon)
    {
        var city = await LookupCityCache(lat, lon);
        return city;
    }

    private async Task<string?> LookupCity(float lat, float lon)
    {
        var result = await http.GetAsync(baseurl + "?location=" + lat + "," + lon + "&outSR=4326&f=json");
        var add = await result.Content.ReadFromJsonAsync<Root>();
        Debug.WriteLine(add?.address.City);
        return add?.address.City;
    }

    private async Task<string> LookupCityCache(float lat, float lon)
    {
        string city = string.Empty;
        //TODO need to cache city more losely, not for every point
        //TODO contain a max size
        if (cityCach.ContainsKey((lat, lon)))
        {
            city = cityCach[(lat, lon)];
        }
        else
        {
            city = await LookupCity(lat, lon)??"";
            cityCach[(lat,lon)] = city;
        }
        return city;
    }
}