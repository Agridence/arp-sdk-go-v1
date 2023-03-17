# \GraphDataApi

All URIs are relative to *https://api-testing.heveaconnect.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphData**](GraphDataApi.md#GraphData) | **Get** /dashboard/graph-without-sicom | 



## GraphData

> GraphDataForSymbol GraphData(ctx).Execute()





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
    resp, r, err := apiClient.GraphDataApi.GraphData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphDataApi.GraphData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphData`: GraphDataForSymbol
    fmt.Fprintf(os.Stdout, "Response from `GraphDataApi.GraphData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGraphDataRequest struct via the builder pattern


### Return type

[**GraphDataForSymbol**](GraphDataForSymbol.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [UserAgent](../README.md#UserAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

