namespace pocketnow_test;

// RUn this with dotnet test -v n

public class TestImageService
{
    public TestImageService() {
System.Diagnostics.Trace.Listeners.Add(new System.Diagnostics.DefaultTraceListener());
    }

    [Fact]
    public async void Test1()
    {
        ImageService s = new ImageService();
        var result = await s.GetLatestImage();
        //Console.WriteLine("Result:" + result);
	Console.WriteLine("Contains 17dcfaa4c117197c91a2eb1eeb28fd8d:" + result.Contains("17dcfaa4c117197c91a2eb1eeb28fd8d".ToString()));
	var linkParser = new Regex(@"\b(?:https?://|www\.)\S+\b", RegexOptions.Compiled | RegexOptions.IgnoreCase);
	foreach(Match m in linkParser.Matches(result))
		    Console.WriteLine(m.Value);
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
