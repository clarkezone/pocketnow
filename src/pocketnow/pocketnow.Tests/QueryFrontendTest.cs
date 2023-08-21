namespace pocketnow.Tests;
using System.Diagnostics;

public class QueryFrontendTests
{
    [Fact]
    public void TestDateConversion()
    {
        var startTime = "2023-08-17T00:45:59Z";
        var endTime = "2023-08-17T13:15:59Z";
        DateTime.Parse(startTime);
    }
}