using System.Diagnostics;
using System.Net.Http.Json;

namespace PocketBase
{
    public class Authorizor
    {
        string email;
        string password;
        HttpClient client;
        Uri baseUrl;
        const string authUrl = "/api/admins/auth-with-password";
        Token? currentToken;

        public Authorizor(string rootUrl, string email, string password)
        {
            client = new HttpClient();
            baseUrl= new Uri(rootUrl);
            this.email= email;
            this.password= password;
        }

        public Client? GetClient()
        {
            if (currentToken != null)
            {
                return new Client(baseUrl, currentToken);
            }
            return null;
        }

        public async Task<bool> Authorize()
        {
            var authUri = new Uri(baseUrl, authUrl);
            Console.WriteLine($"Authenticating with {authUri}");
            try
            {
                var result = await client.PostAsJsonAsync<Identity>(authUri, new Identity(email, password));
                if (result.IsSuccessStatusCode) {
                    var token = await result.Content.ReadFromJsonAsync<authResponse>();
                    if (token != null) {
                    currentToken = new Token(token.Token, DateTime.Now + TimeSpan.FromHours(1));
                    }
                    return true;
                } else {
                    Console.WriteLine($"Failed to authenticate: {result.StatusCode}");
                }

            } catch (Exception ex)
            {
                Debug.WriteLine(ex.Message);
                Console.WriteLine($"Unhandled exception in auth {ex.Message}" );
            }

            return false;
        }

    }
 internal record Identity(string identity, string password) {}
    internal record authResponse(string Token) { }

    internal record Token(string token, DateTime validuntil) { }
}
