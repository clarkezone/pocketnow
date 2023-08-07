using System.Diagnostics;

class ImageService
{
    const string baseurl = "https://lightroom.adobe.com/shares/f82b1d8aacc540628f3f978d99271b3e";
    HttpClient http = new HttpClient();

    public async Task<String?> GetLatestImage()
    {
        var city = await LookupImage();
        return city;
    }

    private async Task<String?> LookupImage()
    {
        var result = await http.GetAsync(baseurl);
        var add = await result.Content.ReadAsStringAsync();

        return add;
    }

}
