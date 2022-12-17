/*
Phishing Hunter API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package phishing_hunter

import (
	"encoding/json"
)

// checks if the GetUsersInfo200ResponseQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUsersInfo200ResponseQuota{}

// GetUsersInfo200ResponseQuota struct for GetUsersInfo200ResponseQuota
type GetUsersInfo200ResponseQuota struct {
	DomainKeywords *int32 `json:"domain_keywords,omitempty"`
	Observations *int32 `json:"observations,omitempty"`
	Scanner *int32 `json:"scanner,omitempty"`
}

// NewGetUsersInfo200ResponseQuota instantiates a new GetUsersInfo200ResponseQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsersInfo200ResponseQuota() *GetUsersInfo200ResponseQuota {
	this := GetUsersInfo200ResponseQuota{}
	return &this
}

// NewGetUsersInfo200ResponseQuotaWithDefaults instantiates a new GetUsersInfo200ResponseQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUsersInfo200ResponseQuotaWithDefaults() *GetUsersInfo200ResponseQuota {
	this := GetUsersInfo200ResponseQuota{}
	return &this
}

// GetDomainKeywords returns the DomainKeywords field value if set, zero value otherwise.
func (o *GetUsersInfo200ResponseQuota) GetDomainKeywords() int32 {
	if o == nil || isNil(o.DomainKeywords) {
		var ret int32
		return ret
	}
	return *o.DomainKeywords
}

// GetDomainKeywordsOk returns a tuple with the DomainKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersInfo200ResponseQuota) GetDomainKeywordsOk() (*int32, bool) {
	if o == nil || isNil(o.DomainKeywords) {
		return nil, false
	}
	return o.DomainKeywords, true
}

// HasDomainKeywords returns a boolean if a field has been set.
func (o *GetUsersInfo200ResponseQuota) HasDomainKeywords() bool {
	if o != nil && !isNil(o.DomainKeywords) {
		return true
	}

	return false
}

// SetDomainKeywords gets a reference to the given int32 and assigns it to the DomainKeywords field.
func (o *GetUsersInfo200ResponseQuota) SetDomainKeywords(v int32) {
	o.DomainKeywords = &v
}

// GetObservations returns the Observations field value if set, zero value otherwise.
func (o *GetUsersInfo200ResponseQuota) GetObservations() int32 {
	if o == nil || isNil(o.Observations) {
		var ret int32
		return ret
	}
	return *o.Observations
}

// GetObservationsOk returns a tuple with the Observations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersInfo200ResponseQuota) GetObservationsOk() (*int32, bool) {
	if o == nil || isNil(o.Observations) {
		return nil, false
	}
	return o.Observations, true
}

// HasObservations returns a boolean if a field has been set.
func (o *GetUsersInfo200ResponseQuota) HasObservations() bool {
	if o != nil && !isNil(o.Observations) {
		return true
	}

	return false
}

// SetObservations gets a reference to the given int32 and assigns it to the Observations field.
func (o *GetUsersInfo200ResponseQuota) SetObservations(v int32) {
	o.Observations = &v
}

// GetScanner returns the Scanner field value if set, zero value otherwise.
func (o *GetUsersInfo200ResponseQuota) GetScanner() int32 {
	if o == nil || isNil(o.Scanner) {
		var ret int32
		return ret
	}
	return *o.Scanner
}

// GetScannerOk returns a tuple with the Scanner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersInfo200ResponseQuota) GetScannerOk() (*int32, bool) {
	if o == nil || isNil(o.Scanner) {
		return nil, false
	}
	return o.Scanner, true
}

// HasScanner returns a boolean if a field has been set.
func (o *GetUsersInfo200ResponseQuota) HasScanner() bool {
	if o != nil && !isNil(o.Scanner) {
		return true
	}

	return false
}

// SetScanner gets a reference to the given int32 and assigns it to the Scanner field.
func (o *GetUsersInfo200ResponseQuota) SetScanner(v int32) {
	o.Scanner = &v
}

func (o GetUsersInfo200ResponseQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUsersInfo200ResponseQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DomainKeywords) {
		toSerialize["domain_keywords"] = o.DomainKeywords
	}
	if !isNil(o.Observations) {
		toSerialize["observations"] = o.Observations
	}
	if !isNil(o.Scanner) {
		toSerialize["scanner"] = o.Scanner
	}
	return toSerialize, nil
}

type NullableGetUsersInfo200ResponseQuota struct {
	value *GetUsersInfo200ResponseQuota
	isSet bool
}

func (v NullableGetUsersInfo200ResponseQuota) Get() *GetUsersInfo200ResponseQuota {
	return v.value
}

func (v *NullableGetUsersInfo200ResponseQuota) Set(val *GetUsersInfo200ResponseQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsersInfo200ResponseQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsersInfo200ResponseQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsersInfo200ResponseQuota(val *GetUsersInfo200ResponseQuota) *NullableGetUsersInfo200ResponseQuota {
	return &NullableGetUsersInfo200ResponseQuota{value: val, isSet: true}
}

func (v NullableGetUsersInfo200ResponseQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsersInfo200ResponseQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


