namespace pocketnow_test;

public class TestImageService
{
    [Fact]
    public async void Test1()
    {
        ImageService s = new ImageService();
        var result = await s.GetLatestImage();
        System.Diagnostics.Debug.WriteLine(result);
        Assert.Equal(4, 4);
    }

      [Fact]
        public void WorkingTest()
        {
            Assert.Equal(4, Add(2, 2));
        }

        int Add(int x, int y)
        {
            return x + y;
        }
}