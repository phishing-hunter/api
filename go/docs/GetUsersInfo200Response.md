# GetUsersInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quota** | Pointer to [**GetUsersInfo200ResponseQuota**](GetUsersInfo200ResponseQuota.md) |  | [optional] 
**UserPlan** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUsersInfo200Response

`func NewGetUsersInfo200Response() *GetUsersInfo200Response`

NewGetUsersInfo200Response instantiates a new GetUsersInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsersInfo200ResponseWithDefaults

`func NewGetUsersInfo200ResponseWithDefaults() *GetUsersInfo200Response`

NewGetUsersInfo200ResponseWithDefaults instantiates a new GetUsersInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuota

`func (o *GetUsersInfo200Response) GetQuota() GetUsersInfo200ResponseQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *GetUsersInfo200Response) GetQuotaOk() (*GetUsersInfo200ResponseQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *GetUsersInfo200Response) SetQuota(v GetUsersInfo200ResponseQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *GetUsersInfo200Response) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetUserPlan

`func (o *GetUsersInfo200Response) GetUserPlan() string`

GetUserPlan returns the UserPlan field if non-nil, zero value otherwise.

### GetUserPlanOk

`func (o *GetUsersInfo200Response) GetUserPlanOk() (*string, bool)`

GetUserPlanOk returns a tuple with the UserPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPlan

`func (o *GetUsersInfo200Response) SetUserPlan(v string)`

SetUserPlan sets UserPlan field to given value.

### HasUserPlan

`func (o *GetUsersInfo200Response) HasUserPlan() bool`

HasUserPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


