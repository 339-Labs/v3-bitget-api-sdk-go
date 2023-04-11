/*
Bitget Open API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MarginCancelOrderFailureResult struct for MarginCancelOrderFailureResult
type MarginCancelOrderFailureResult struct {
	ClientOid            *string `json:"clientOid,omitempty"`
	ErrorMsg             *string `json:"errorMsg,omitempty"`
	OrderId              *string `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginCancelOrderFailureResult MarginCancelOrderFailureResult

// NewMarginCancelOrderFailureResult instantiates a new MarginCancelOrderFailureResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCancelOrderFailureResult() *MarginCancelOrderFailureResult {
	this := MarginCancelOrderFailureResult{}
	return &this
}

// NewMarginCancelOrderFailureResultWithDefaults instantiates a new MarginCancelOrderFailureResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCancelOrderFailureResultWithDefaults() *MarginCancelOrderFailureResult {
	this := MarginCancelOrderFailureResult{}
	return &this
}

// GetClientOid returns the ClientOid field value if set, zero value otherwise.
func (o *MarginCancelOrderFailureResult) GetClientOid() string {
	if o == nil || isNil(o.ClientOid) {
		var ret string
		return ret
	}
	return *o.ClientOid
}

// GetClientOidOk returns a tuple with the ClientOid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCancelOrderFailureResult) GetClientOidOk() (*string, bool) {
	if o == nil || isNil(o.ClientOid) {
		return nil, false
	}
	return o.ClientOid, true
}

// HasClientOid returns a boolean if a field has been set.
func (o *MarginCancelOrderFailureResult) HasClientOid() bool {
	if o != nil && !isNil(o.ClientOid) {
		return true
	}

	return false
}

// SetClientOid gets a reference to the given string and assigns it to the ClientOid field.
func (o *MarginCancelOrderFailureResult) SetClientOid(v string) {
	o.ClientOid = &v
}

// GetErrorMsg returns the ErrorMsg field value if set, zero value otherwise.
func (o *MarginCancelOrderFailureResult) GetErrorMsg() string {
	if o == nil || isNil(o.ErrorMsg) {
		var ret string
		return ret
	}
	return *o.ErrorMsg
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCancelOrderFailureResult) GetErrorMsgOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMsg) {
		return nil, false
	}
	return o.ErrorMsg, true
}

// HasErrorMsg returns a boolean if a field has been set.
func (o *MarginCancelOrderFailureResult) HasErrorMsg() bool {
	if o != nil && !isNil(o.ErrorMsg) {
		return true
	}

	return false
}

// SetErrorMsg gets a reference to the given string and assigns it to the ErrorMsg field.
func (o *MarginCancelOrderFailureResult) SetErrorMsg(v string) {
	o.ErrorMsg = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MarginCancelOrderFailureResult) GetOrderId() string {
	if o == nil || isNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCancelOrderFailureResult) GetOrderIdOk() (*string, bool) {
	if o == nil || isNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MarginCancelOrderFailureResult) HasOrderId() bool {
	if o != nil && !isNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *MarginCancelOrderFailureResult) SetOrderId(v string) {
	o.OrderId = &v
}

func (o MarginCancelOrderFailureResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientOid) {
		toSerialize["clientOid"] = o.ClientOid
	}
	if !isNil(o.ErrorMsg) {
		toSerialize["errorMsg"] = o.ErrorMsg
	}
	if !isNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginCancelOrderFailureResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginCancelOrderFailureResult := _MarginCancelOrderFailureResult{}

	if err = json.Unmarshal(bytes, &varMarginCancelOrderFailureResult); err == nil {
		*o = MarginCancelOrderFailureResult(varMarginCancelOrderFailureResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "clientOid")
		delete(additionalProperties, "errorMsg")
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginCancelOrderFailureResult struct {
	value *MarginCancelOrderFailureResult
	isSet bool
}

func (v NullableMarginCancelOrderFailureResult) Get() *MarginCancelOrderFailureResult {
	return v.value
}

func (v *NullableMarginCancelOrderFailureResult) Set(val *MarginCancelOrderFailureResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCancelOrderFailureResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCancelOrderFailureResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCancelOrderFailureResult(val *MarginCancelOrderFailureResult) *NullableMarginCancelOrderFailureResult {
	return &NullableMarginCancelOrderFailureResult{value: val, isSet: true}
}

func (v NullableMarginCancelOrderFailureResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCancelOrderFailureResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
