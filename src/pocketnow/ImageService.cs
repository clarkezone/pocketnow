using System.Diagnostics;
using OpenQA.Selenium;
using OpenQA.Selenium.Chrome;

public class ImageService
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
	using (var driver = new ChromeDriver(".")) {
		driver.Navigate().GoToUrl("https://lightroom.adobe.com/shares/f82b1d8aacc540628f3f978d99271b3e#");


	}
	return "ss";
        //var result = await http.GetAsync(baseurl);
        //var add = await result.Content.ReadAsStringAsync();

        //return add;
    }

}
