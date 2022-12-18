namespace pocketnow_test;

public class TestImageService
{
    [Fact]
    public async void Test1()
    {
        ImageService s = new ImageService();
        var result = await s.GetLatestImage();
        System.Diagnostics.Debug.WriteLine(result);
    }
}