# \UserInfoApi

All URIs are relative to *https://api-testing.heveaconnect.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserInfo**](UserInfoApi.md#UserInfo) | **Get** /user/info | 



## UserInfo

> UserInfo200Response UserInfo(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserInfoApi.UserInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInfoApi.UserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserInfo`: UserInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `UserInfoApi.UserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserInfoRequest struct via the builder pattern


### Return type

[**UserInfo200Response**](UserInfo200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

