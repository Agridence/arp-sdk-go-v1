package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/HeveaConnect/arp-sdk-go-v1"
)

func main() {
    data := *openapiclient.NewMFAAuthToken("apiproxy@serviceaccount-testing-hevea-global.com", "Test12345") // MFAAuthToken | 

    configuration := openapiclient.NewConfiguration()
    configuration.Debug = true
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.AuthLoginCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthLoginCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthLoginCreate`: MFAAuthToken
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthLoginCreate`: %v, Token: %s\n", resp, *resp.Token)
}
