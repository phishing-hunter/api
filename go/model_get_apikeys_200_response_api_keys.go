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

// checks if the GetApikeys200ResponseApiKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApikeys200ResponseApiKeys{}

// GetApikeys200ResponseApiKeys struct for GetApikeys200ResponseApiKeys
type GetApikeys200ResponseApiKeys struct {
	Urlscan *string `json:"urlscan,omitempty"`
	Shodan *string `json:"shodan,omitempty"`
	Vt *string `json:"vt,omitempty"`
}

// NewGetApikeys200ResponseApiKeys instantiates a new GetApikeys200ResponseApiKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApikeys200ResponseApiKeys() *GetApikeys200ResponseApiKeys {
	this := GetApikeys200ResponseApiKeys{}
	return &this
}

// NewGetApikeys200ResponseApiKeysWithDefaults instantiates a new GetApikeys200ResponseApiKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApikeys200ResponseApiKeysWithDefaults() *GetApikeys200ResponseApiKeys {
	this := GetApikeys200ResponseApiKeys{}
	return &this
}

// GetUrlscan returns the Urlscan field value if set, zero value otherwise.
func (o *GetApikeys200ResponseApiKeys) GetUrlscan() string {
	if o == nil || isNil(o.Urlscan) {
		var ret string
		return ret
	}
	return *o.Urlscan
}

// GetUrlscanOk returns a tuple with the Urlscan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApikeys200ResponseApiKeys) GetUrlscanOk() (*string, bool) {
	if o == nil || isNil(o.Urlscan) {
		return nil, false
	}
	return o.Urlscan, true
}

// HasUrlscan returns a boolean if a field has been set.
func (o *GetApikeys200ResponseApiKeys) HasUrlscan() bool {
	if o != nil && !isNil(o.Urlscan) {
		return true
	}

	return false
}

// SetUrlscan gets a reference to the given string and assigns it to the Urlscan field.
func (o *GetApikeys200ResponseApiKeys) SetUrlscan(v string) {
	o.Urlscan = &v
}

// GetShodan returns the Shodan field value if set, zero value otherwise.
func (o *GetApikeys200ResponseApiKeys) GetShodan() string {
	if o == nil || isNil(o.Shodan) {
		var ret string
		return ret
	}
	return *o.Shodan
}

// GetShodanOk returns a tuple with the Shodan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApikeys200ResponseApiKeys) GetShodanOk() (*string, bool) {
	if o == nil || isNil(o.Shodan) {
		return nil, false
	}
	return o.Shodan, true
}

// HasShodan returns a boolean if a field has been set.
func (o *GetApikeys200ResponseApiKeys) HasShodan() bool {
	if o != nil && !isNil(o.Shodan) {
		return true
	}

	return false
}

// SetShodan gets a reference to the given string and assigns it to the Shodan field.
func (o *GetApikeys200ResponseApiKeys) SetShodan(v string) {
	o.Shodan = &v
}

// GetVt returns the Vt field value if set, zero value otherwise.
func (o *GetApikeys200ResponseApiKeys) GetVt() string {
	if o == nil || isNil(o.Vt) {
		var ret string
		return ret
	}
	return *o.Vt
}

// GetVtOk returns a tuple with the Vt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApikeys200ResponseApiKeys) GetVtOk() (*string, bool) {
	if o == nil || isNil(o.Vt) {
		return nil, false
	}
	return o.Vt, true
}

// HasVt returns a boolean if a field has been set.
func (o *GetApikeys200ResponseApiKeys) HasVt() bool {
	if o != nil && !isNil(o.Vt) {
		return true
	}

	return false
}

// SetVt gets a reference to the given string and assigns it to the Vt field.
func (o *GetApikeys200ResponseApiKeys) SetVt(v string) {
	o.Vt = &v
}

func (o GetApikeys200ResponseApiKeys) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApikeys200ResponseApiKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Urlscan) {
		toSerialize["urlscan"] = o.Urlscan
	}
	if !isNil(o.Shodan) {
		toSerialize["shodan"] = o.Shodan
	}
	if !isNil(o.Vt) {
		toSerialize["vt"] = o.Vt
	}
	return toSerialize, nil
}

type NullableGetApikeys200ResponseApiKeys struct {
	value *GetApikeys200ResponseApiKeys
	isSet bool
}

func (v NullableGetApikeys200ResponseApiKeys) Get() *GetApikeys200ResponseApiKeys {
	return v.value
}

func (v *NullableGetApikeys200ResponseApiKeys) Set(val *GetApikeys200ResponseApiKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApikeys200ResponseApiKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApikeys200ResponseApiKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApikeys200ResponseApiKeys(val *GetApikeys200ResponseApiKeys) *NullableGetApikeys200ResponseApiKeys {
	return &NullableGetApikeys200ResponseApiKeys{value: val, isSet: true}
}

func (v NullableGetApikeys200ResponseApiKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApikeys200ResponseApiKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


