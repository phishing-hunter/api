# PostObservationAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Expire** | Pointer to **int32** |  | [optional] 
**ApiConfig** | Pointer to [**PostObservationAddRequestApiConfig**](PostObservationAddRequestApiConfig.md) |  | [optional] 

## Methods

### NewPostObservationAddRequest

`func NewPostObservationAddRequest() *PostObservationAddRequest`

NewPostObservationAddRequest instantiates a new PostObservationAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostObservationAddRequestWithDefaults

`func NewPostObservationAddRequestWithDefaults() *PostObservationAddRequest`

NewPostObservationAddRequestWithDefaults instantiates a new PostObservationAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PostObservationAddRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PostObservationAddRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PostObservationAddRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PostObservationAddRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetInterval

`func (o *PostObservationAddRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PostObservationAddRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PostObservationAddRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PostObservationAddRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPeriod

`func (o *PostObservationAddRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *PostObservationAddRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *PostObservationAddRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *PostObservationAddRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetComment

`func (o *PostObservationAddRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PostObservationAddRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PostObservationAddRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PostObservationAddRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpire

`func (o *PostObservationAddRequest) GetExpire() int32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *PostObservationAddRequest) GetExpireOk() (*int32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *PostObservationAddRequest) SetExpire(v int32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *PostObservationAddRequest) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetApiConfig

`func (o *PostObservationAddRequest) GetApiConfig() PostObservationAddRequestApiConfig`

GetApiConfig returns the ApiConfig field if non-nil, zero value otherwise.

### GetApiConfigOk

`func (o *PostObservationAddRequest) GetApiConfigOk() (*PostObservationAddRequestApiConfig, bool)`

GetApiConfigOk returns a tuple with the ApiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiConfig

`func (o *PostObservationAddRequest) SetApiConfig(v PostObservationAddRequestApiConfig)`

SetApiConfig sets ApiConfig field to given value.

### HasApiConfig

`func (o *PostObservationAddRequest) HasApiConfig() bool`

HasApiConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


