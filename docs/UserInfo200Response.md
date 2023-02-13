# UserInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **map[string]interface{}** |  | 
**Name** | **string** |  | 
**SgxtsrSubscription** | **NullableString** |  | 

## Methods

### NewUserInfo200Response

`func NewUserInfo200Response(company map[string]interface{}, name string, sgxtsrSubscription NullableString, ) *UserInfo200Response`

NewUserInfo200Response instantiates a new UserInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfo200ResponseWithDefaults

`func NewUserInfo200ResponseWithDefaults() *UserInfo200Response`

NewUserInfo200ResponseWithDefaults instantiates a new UserInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *UserInfo200Response) GetCompany() map[string]interface{}`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UserInfo200Response) GetCompanyOk() (*map[string]interface{}, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UserInfo200Response) SetCompany(v map[string]interface{})`

SetCompany sets Company field to given value.


### GetName

`func (o *UserInfo200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserInfo200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserInfo200Response) SetName(v string)`

SetName sets Name field to given value.


### GetSgxtsrSubscription

`func (o *UserInfo200Response) GetSgxtsrSubscription() string`

GetSgxtsrSubscription returns the SgxtsrSubscription field if non-nil, zero value otherwise.

### GetSgxtsrSubscriptionOk

`func (o *UserInfo200Response) GetSgxtsrSubscriptionOk() (*string, bool)`

GetSgxtsrSubscriptionOk returns a tuple with the SgxtsrSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxtsrSubscription

`func (o *UserInfo200Response) SetSgxtsrSubscription(v string)`

SetSgxtsrSubscription sets SgxtsrSubscription field to given value.


### SetSgxtsrSubscriptionNil

`func (o *UserInfo200Response) SetSgxtsrSubscriptionNil(b bool)`

 SetSgxtsrSubscriptionNil sets the value for SgxtsrSubscription to be an explicit nil

### UnsetSgxtsrSubscription
`func (o *UserInfo200Response) UnsetSgxtsrSubscription()`

UnsetSgxtsrSubscription ensures that no value is present for SgxtsrSubscription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


