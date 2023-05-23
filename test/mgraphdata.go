package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Agridence/arp-sdk-go-v1"
)

func main() {

    ctx := context.Background()
    // ctx = context.WithValue(ctx, openapiclient.ContextServerIndex, "0")
    // ctx = context.WithValue(ctx, openapiclient.ContextServerVariables, map[string]string{"environment": "api-testing",})

    configuration := openapiclient.NewConfiguration()
    configuration.Debug = true
    apiClient := openapiclient.NewAPIClient(configuration)
    // fmt.Println(configuration.Host)
    ctx = context.WithValue(ctx, openapiclient.ContextAPIKeys, map[string]openapiclient.APIKey{
	"ApiKeyAuth": openapiclient.APIKey{Key: "96a68b0635ee03a22e6399ce188c5da869bcdea1", Prefix: "Token",},
	"UserAgent": openapiclient.APIKey{Key: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15", Prefix: "",},
	})
    // start verify context
       if auth, ok := ctx.Value(openapiclient.ContextAPIKeys).(map[string]openapiclient.APIKey); ok {
                        if apiKey, ok := auth["ApiKeyAuth"]; ok {
                                var key string
                                if apiKey.Prefix != "" {
                                        key = apiKey.Prefix + " " + apiKey.Key
                                } else {
                                        key = apiKey.Key
                                }
             			fmt.Println("key = " + key)
                        }
                }
    // end verify context
    resp, r, err := apiClient.GraphDataApi.GraphData(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphDataApi.GraphData`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    fmt.Fprintf(os.Stdout, "Response from `GraphDataApi.GraphData`: %v\n", (*resp.Prices)["2023-03-01"]["SIR20"])
}
