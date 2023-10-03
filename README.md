# Go API client for arp_sdk_go_v1

Agridence

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import arp_sdk_go_v1 "github.com/Agridence/arp-sdk-go-v1"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), arp_sdk_go_v1.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), arp_sdk_go_v1.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), arp_sdk_go_v1.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), arp_sdk_go_v1.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api-platform-testing.agridence.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthApi* | [**AuthLoginCreate**](docs/AuthApi.md#authlogincreate) | **Post** /auth/login/ | Login to the Agridence API with email and password.
*AuthApi* | [**AuthLogoutCreate**](docs/AuthApi.md#authlogoutcreate) | **Post** /auth/logout/ | 
*GraphDataApi* | [**GraphData**](docs/GraphDataApi.md#graphdata) | **Get** /dashboard/graph-without-sicom | 
*OfferPanelApi* | [**OfferPanelFilterList**](docs/OfferPanelApi.md#offerpanelfilterlist) | **Get** /offer-panel/filter/ | 
*UserInfoApi* | [**UserInfo**](docs/UserInfoApi.md#userinfo) | **Get** /user/info | 


## Documentation For Models

 - [CompanyDetails](docs/CompanyDetails.md)
 - [GraphDataForSymbol](docs/GraphDataForSymbol.md)
 - [GraphDataForSymbolCategoryValueInner](docs/GraphDataForSymbolCategoryValueInner.md)
 - [MFAAuthToken](docs/MFAAuthToken.md)
 - [OfferPanelFilter](docs/OfferPanelFilter.md)
 - [OfferPanelFilterCountry](docs/OfferPanelFilterCountry.md)
 - [OfferPanelFilterDestinations](docs/OfferPanelFilterDestinations.md)
 - [OfferPanelFilterFactory](docs/OfferPanelFilterFactory.md)
 - [OfferPanelFilterGrade](docs/OfferPanelFilterGrade.md)
 - [OfferPanelFilterList200Response](docs/OfferPanelFilterList200Response.md)
 - [OfferPanelFilterPackings](docs/OfferPanelFilterPackings.md)
 - [OfferPanelFilterPayments](docs/OfferPanelFilterPayments.md)
 - [OfferPanelFilterPort](docs/OfferPanelFilterPort.md)
 - [OfferPanelFilterShippings](docs/OfferPanelFilterShippings.md)
 - [UserInfo200Response](docs/UserInfo200Response.md)


## Documentation For Authorization



### UserAgent

- **Type**: API key
- **API key parameter name**: User-Agent
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: User-Agent and passed in as the auth context for each request.


### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


### ClientType

- **Type**: API key
- **API key parameter name**: X-Client-Type
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Client-Type and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

it@agridence.com

