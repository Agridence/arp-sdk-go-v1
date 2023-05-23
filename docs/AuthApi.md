# \AuthApi

All URIs are relative to *https://api-platform-testing.agridence.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthLoginCreate**](AuthApi.md#AuthLoginCreate) | **Post** /auth/login/ | Login to the Agridence API with email and password.
[**AuthLogoutCreate**](AuthApi.md#AuthLogoutCreate) | **Post** /auth/logout/ | 



## AuthLoginCreate

> MFAAuthToken AuthLoginCreate(ctx).Data(data).Execute()

Login to the Agridence API with email and password.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Agridence/arp-sdk-go-v1"
)

func main() {
    data := *openapiclient.NewMFAAuthToken("Email_example", "Password_example") // MFAAuthToken | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.AuthLoginCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthLoginCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthLoginCreate`: MFAAuthToken
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthLoginCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthLoginCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**MFAAuthToken**](MFAAuthToken.md) |  | 

### Return type

[**MFAAuthToken**](MFAAuthToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthLogoutCreate

> AuthLogoutCreate(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Agridence/arp-sdk-go-v1"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthApi.AuthLogoutCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthLogoutCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthLogoutCreateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

