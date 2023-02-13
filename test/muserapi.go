package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/HeveaConnect/arp-sdk-go-v1"
)

func main() {

    ctx := context.Background()
    ctx = context.WithValue(ctx, openapiclient.ContextServerIndex, "0")
    ctx = context.WithValue(ctx, openapiclient.ContextServerVariables, map[string]string{"environment": "api-testing",})

    configuration := openapiclient.NewConfiguration()
    configuration.Debug = true
    configuration.AddDefaultHeader("Authorization", "Token 0b99fddd9b8268e157df0f4746855763b9936c84")
    apiClient := openapiclient.NewAPIClient(configuration)
    // fmt.Println(configuration.Host)
    resp, r, err := apiClient.UserInfoApi.UserInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInfoApi.UserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserInfo`: UserInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `UserInfoApi.UserInfo`: %v\n", resp)
}
