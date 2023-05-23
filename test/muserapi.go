package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Agridence/arp-sdk-go-v1"
)

func main() {

    ctx := context.Background()
    ctx = context.WithValue(ctx, openapiclient.ContextServerIndex, 0)
    // ctx = context.WithValue(ctx, openapiclient.ContextServerVariables, map[string]string{"environment": "api-platform-testing",})

    configuration := openapiclient.NewConfiguration()
    configuration.Debug = true
    apiClient := openapiclient.NewAPIClient(configuration)
    // fmt.Println(configuration.Host)
    ctx = context.WithValue(ctx, openapiclient.ContextAPIKeys, map[string]openapiclient.APIKey{
	"ApiKeyAuth": openapiclient.APIKey{Key: "9fbf858d7f55f1dee1ffa9a43ffa1a0ca315b6c0", Prefix: "Token",},
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
    resp, r, err := apiClient.UserInfoApi.UserInfo(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInfoApi.UserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserInfo`: UserInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `UserInfoApi.UserInfo`: %v, %d\n", resp, resp.GetId())
}
