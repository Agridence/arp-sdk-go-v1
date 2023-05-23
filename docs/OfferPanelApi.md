# \OfferPanelApi

All URIs are relative to *https://api-platform-testing.agridence.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfferPanelFilterList**](OfferPanelApi.md#OfferPanelFilterList) | **Get** /offer-panel/filter/ | 



## OfferPanelFilterList

> OfferPanelFilterList200Response OfferPanelFilterList(ctx).Page(page).PageSize(pageSize).Execute()





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
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfferPanelApi.OfferPanelFilterList(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfferPanelApi.OfferPanelFilterList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfferPanelFilterList`: OfferPanelFilterList200Response
    fmt.Fprintf(os.Stdout, "Response from `OfferPanelApi.OfferPanelFilterList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOfferPanelFilterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**OfferPanelFilterList200Response**](OfferPanelFilterList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

