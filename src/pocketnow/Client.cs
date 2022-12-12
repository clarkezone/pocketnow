using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Text;
using System.Threading.Tasks;

namespace PocketBase
{
    public class Client
    {
        private HttpClient client;
        private Uri baseUrl;
        private Token currentToken;
        internal Client(Uri baseUrl, Token validtoken) {
            client = new HttpClient();
            client.DefaultRequestHeaders.Accept.Add(new MediaTypeWithQualityHeaderValue("application/json"));
            this.baseUrl = baseUrl;
            currentToken = validtoken;
        }

        public async Task<T?> GetRecords<T>(string tableName)
        {
            //SetAuthHeader();
            var callUrl = GetUrl(baseUrl, tableName);
            HttpRequestMessage rm =new HttpRequestMessage(HttpMethod.Get, callUrl);
            //TODO handle do auth and token validity
            rm.Headers.Add("Authorization", currentToken.token);
            var result = await client.SendAsync(rm);
            if (result.IsSuccessStatusCode)
            {
//                var temp = result.Content.ReadAsStringAsync();
                var decoded = result.Content.ReadFromJsonAsync<T>();
                return decoded.Result;
            }
            return default(T);
        }

        private Uri GetUrl(Uri baseUrl, string tableName)
        {
            // string interpolation
            string thing = "api/collections/"+tableName+"/records";
            return new Uri(baseUrl, thing);
        }

        public void SetAuthHeader()
        {
            throw new NotImplementedException();
        }
    }
}
