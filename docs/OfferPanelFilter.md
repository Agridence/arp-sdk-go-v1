# OfferPanelFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Factories** | Pointer to [**[]OfferPanelFilterFactory**](OfferPanelFilterFactory.md) |  | [optional] [readonly] 
**DefaultFactory** | Pointer to **string** |  | [optional] [readonly] 
**Countries** | Pointer to [**[]OfferPanelFilterCountry**](OfferPanelFilterCountry.md) |  | [optional] [readonly] 
**DefaultCountry** | Pointer to **string** |  | [optional] [readonly] 
**Ports** | Pointer to [**[]OfferPanelFilterPort**](OfferPanelFilterPort.md) |  | [optional] [readonly] 
**DefaultPort** | Pointer to **string** |  | [optional] [readonly] 
**Grades** | Pointer to [**[]OfferPanelFilterGrade**](OfferPanelFilterGrade.md) |  | [optional] [readonly] 
**DefaultGrade** | Pointer to **string** |  | [optional] [readonly] 
**Packings** | Pointer to [**[]OfferPanelFilterPackings**](OfferPanelFilterPackings.md) |  | [optional] [readonly] 
**DefaultPacking** | Pointer to **string** |  | [optional] [readonly] 
**Payments** | Pointer to [**[]OfferPanelFilterPayments**](OfferPanelFilterPayments.md) |  | [optional] [readonly] 
**DefaultPayment** | Pointer to **string** |  | [optional] [readonly] 
**Shippings** | Pointer to [**[]OfferPanelFilterShippings**](OfferPanelFilterShippings.md) |  | [optional] [readonly] 
**DefaultShipping** | Pointer to **string** |  | [optional] [readonly] 
**Destinations** | Pointer to [**[]OfferPanelFilterDestinations**](OfferPanelFilterDestinations.md) |  | [optional] [readonly] 
**DefaultDestination** | Pointer to **string** |  | [optional] [readonly] 
**Months** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewOfferPanelFilter

`func NewOfferPanelFilter(name string, type_ string, ) *OfferPanelFilter`

NewOfferPanelFilter instantiates a new OfferPanelFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferPanelFilterWithDefaults

`func NewOfferPanelFilterWithDefaults() *OfferPanelFilter`

NewOfferPanelFilterWithDefaults instantiates a new OfferPanelFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OfferPanelFilter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OfferPanelFilter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OfferPanelFilter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OfferPanelFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OfferPanelFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OfferPanelFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OfferPanelFilter) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *OfferPanelFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OfferPanelFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OfferPanelFilter) SetType(v string)`

SetType sets Type field to given value.


### GetFactories

`func (o *OfferPanelFilter) GetFactories() []OfferPanelFilterFactory`

GetFactories returns the Factories field if non-nil, zero value otherwise.

### GetFactoriesOk

`func (o *OfferPanelFilter) GetFactoriesOk() (*[]OfferPanelFilterFactory, bool)`

GetFactoriesOk returns a tuple with the Factories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactories

`func (o *OfferPanelFilter) SetFactories(v []OfferPanelFilterFactory)`

SetFactories sets Factories field to given value.

### HasFactories

`func (o *OfferPanelFilter) HasFactories() bool`

HasFactories returns a boolean if a field has been set.

### GetDefaultFactory

`func (o *OfferPanelFilter) GetDefaultFactory() string`

GetDefaultFactory returns the DefaultFactory field if non-nil, zero value otherwise.

### GetDefaultFactoryOk

`func (o *OfferPanelFilter) GetDefaultFactoryOk() (*string, bool)`

GetDefaultFactoryOk returns a tuple with the DefaultFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFactory

`func (o *OfferPanelFilter) SetDefaultFactory(v string)`

SetDefaultFactory sets DefaultFactory field to given value.

### HasDefaultFactory

`func (o *OfferPanelFilter) HasDefaultFactory() bool`

HasDefaultFactory returns a boolean if a field has been set.

### GetCountries

`func (o *OfferPanelFilter) GetCountries() []OfferPanelFilterCountry`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *OfferPanelFilter) GetCountriesOk() (*[]OfferPanelFilterCountry, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *OfferPanelFilter) SetCountries(v []OfferPanelFilterCountry)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *OfferPanelFilter) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetDefaultCountry

`func (o *OfferPanelFilter) GetDefaultCountry() string`

GetDefaultCountry returns the DefaultCountry field if non-nil, zero value otherwise.

### GetDefaultCountryOk

`func (o *OfferPanelFilter) GetDefaultCountryOk() (*string, bool)`

GetDefaultCountryOk returns a tuple with the DefaultCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCountry

`func (o *OfferPanelFilter) SetDefaultCountry(v string)`

SetDefaultCountry sets DefaultCountry field to given value.

### HasDefaultCountry

`func (o *OfferPanelFilter) HasDefaultCountry() bool`

HasDefaultCountry returns a boolean if a field has been set.

### GetPorts

`func (o *OfferPanelFilter) GetPorts() []OfferPanelFilterPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OfferPanelFilter) GetPortsOk() (*[]OfferPanelFilterPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OfferPanelFilter) SetPorts(v []OfferPanelFilterPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OfferPanelFilter) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetDefaultPort

`func (o *OfferPanelFilter) GetDefaultPort() string`

GetDefaultPort returns the DefaultPort field if non-nil, zero value otherwise.

### GetDefaultPortOk

`func (o *OfferPanelFilter) GetDefaultPortOk() (*string, bool)`

GetDefaultPortOk returns a tuple with the DefaultPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPort

`func (o *OfferPanelFilter) SetDefaultPort(v string)`

SetDefaultPort sets DefaultPort field to given value.

### HasDefaultPort

`func (o *OfferPanelFilter) HasDefaultPort() bool`

HasDefaultPort returns a boolean if a field has been set.

### GetGrades

`func (o *OfferPanelFilter) GetGrades() []OfferPanelFilterGrade`

GetGrades returns the Grades field if non-nil, zero value otherwise.

### GetGradesOk

`func (o *OfferPanelFilter) GetGradesOk() (*[]OfferPanelFilterGrade, bool)`

GetGradesOk returns a tuple with the Grades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrades

`func (o *OfferPanelFilter) SetGrades(v []OfferPanelFilterGrade)`

SetGrades sets Grades field to given value.

### HasGrades

`func (o *OfferPanelFilter) HasGrades() bool`

HasGrades returns a boolean if a field has been set.

### GetDefaultGrade

`func (o *OfferPanelFilter) GetDefaultGrade() string`

GetDefaultGrade returns the DefaultGrade field if non-nil, zero value otherwise.

### GetDefaultGradeOk

`func (o *OfferPanelFilter) GetDefaultGradeOk() (*string, bool)`

GetDefaultGradeOk returns a tuple with the DefaultGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGrade

`func (o *OfferPanelFilter) SetDefaultGrade(v string)`

SetDefaultGrade sets DefaultGrade field to given value.

### HasDefaultGrade

`func (o *OfferPanelFilter) HasDefaultGrade() bool`

HasDefaultGrade returns a boolean if a field has been set.

### GetPackings

`func (o *OfferPanelFilter) GetPackings() []OfferPanelFilterPackings`

GetPackings returns the Packings field if non-nil, zero value otherwise.

### GetPackingsOk

`func (o *OfferPanelFilter) GetPackingsOk() (*[]OfferPanelFilterPackings, bool)`

GetPackingsOk returns a tuple with the Packings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackings

`func (o *OfferPanelFilter) SetPackings(v []OfferPanelFilterPackings)`

SetPackings sets Packings field to given value.

### HasPackings

`func (o *OfferPanelFilter) HasPackings() bool`

HasPackings returns a boolean if a field has been set.

### GetDefaultPacking

`func (o *OfferPanelFilter) GetDefaultPacking() string`

GetDefaultPacking returns the DefaultPacking field if non-nil, zero value otherwise.

### GetDefaultPackingOk

`func (o *OfferPanelFilter) GetDefaultPackingOk() (*string, bool)`

GetDefaultPackingOk returns a tuple with the DefaultPacking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPacking

`func (o *OfferPanelFilter) SetDefaultPacking(v string)`

SetDefaultPacking sets DefaultPacking field to given value.

### HasDefaultPacking

`func (o *OfferPanelFilter) HasDefaultPacking() bool`

HasDefaultPacking returns a boolean if a field has been set.

### GetPayments

`func (o *OfferPanelFilter) GetPayments() []OfferPanelFilterPayments`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *OfferPanelFilter) GetPaymentsOk() (*[]OfferPanelFilterPayments, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *OfferPanelFilter) SetPayments(v []OfferPanelFilterPayments)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *OfferPanelFilter) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetDefaultPayment

`func (o *OfferPanelFilter) GetDefaultPayment() string`

GetDefaultPayment returns the DefaultPayment field if non-nil, zero value otherwise.

### GetDefaultPaymentOk

`func (o *OfferPanelFilter) GetDefaultPaymentOk() (*string, bool)`

GetDefaultPaymentOk returns a tuple with the DefaultPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPayment

`func (o *OfferPanelFilter) SetDefaultPayment(v string)`

SetDefaultPayment sets DefaultPayment field to given value.

### HasDefaultPayment

`func (o *OfferPanelFilter) HasDefaultPayment() bool`

HasDefaultPayment returns a boolean if a field has been set.

### GetShippings

`func (o *OfferPanelFilter) GetShippings() []OfferPanelFilterShippings`

GetShippings returns the Shippings field if non-nil, zero value otherwise.

### GetShippingsOk

`func (o *OfferPanelFilter) GetShippingsOk() (*[]OfferPanelFilterShippings, bool)`

GetShippingsOk returns a tuple with the Shippings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippings

`func (o *OfferPanelFilter) SetShippings(v []OfferPanelFilterShippings)`

SetShippings sets Shippings field to given value.

### HasShippings

`func (o *OfferPanelFilter) HasShippings() bool`

HasShippings returns a boolean if a field has been set.

### GetDefaultShipping

`func (o *OfferPanelFilter) GetDefaultShipping() string`

GetDefaultShipping returns the DefaultShipping field if non-nil, zero value otherwise.

### GetDefaultShippingOk

`func (o *OfferPanelFilter) GetDefaultShippingOk() (*string, bool)`

GetDefaultShippingOk returns a tuple with the DefaultShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShipping

`func (o *OfferPanelFilter) SetDefaultShipping(v string)`

SetDefaultShipping sets DefaultShipping field to given value.

### HasDefaultShipping

`func (o *OfferPanelFilter) HasDefaultShipping() bool`

HasDefaultShipping returns a boolean if a field has been set.

### GetDestinations

`func (o *OfferPanelFilter) GetDestinations() []OfferPanelFilterDestinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *OfferPanelFilter) GetDestinationsOk() (*[]OfferPanelFilterDestinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *OfferPanelFilter) SetDestinations(v []OfferPanelFilterDestinations)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *OfferPanelFilter) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetDefaultDestination

`func (o *OfferPanelFilter) GetDefaultDestination() string`

GetDefaultDestination returns the DefaultDestination field if non-nil, zero value otherwise.

### GetDefaultDestinationOk

`func (o *OfferPanelFilter) GetDefaultDestinationOk() (*string, bool)`

GetDefaultDestinationOk returns a tuple with the DefaultDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestination

`func (o *OfferPanelFilter) SetDefaultDestination(v string)`

SetDefaultDestination sets DefaultDestination field to given value.

### HasDefaultDestination

`func (o *OfferPanelFilter) HasDefaultDestination() bool`

HasDefaultDestination returns a boolean if a field has been set.

### GetMonths

`func (o *OfferPanelFilter) GetMonths() []string`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *OfferPanelFilter) GetMonthsOk() (*[]string, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *OfferPanelFilter) SetMonths(v []string)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *OfferPanelFilter) HasMonths() bool`

HasMonths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


