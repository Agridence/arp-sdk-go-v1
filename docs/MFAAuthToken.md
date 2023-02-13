# MFAAuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**Token** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMFAAuthToken

`func NewMFAAuthToken(email string, password string, ) *MFAAuthToken`

NewMFAAuthToken instantiates a new MFAAuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAAuthTokenWithDefaults

`func NewMFAAuthTokenWithDefaults() *MFAAuthToken`

NewMFAAuthTokenWithDefaults instantiates a new MFAAuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *MFAAuthToken) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MFAAuthToken) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MFAAuthToken) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *MFAAuthToken) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MFAAuthToken) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MFAAuthToken) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetToken

`func (o *MFAAuthToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MFAAuthToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MFAAuthToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MFAAuthToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


