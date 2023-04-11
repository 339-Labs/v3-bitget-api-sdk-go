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

// MarginBatchOrdersRequest struct for MarginBatchOrdersRequest
type MarginBatchOrdersRequest struct {
	ChannelApiCode *string              `json:"channelApiCode,omitempty"`
	Ip             *string              `json:"ip,omitempty"`
	OrderList      []MarginOrderRequest `json:"orderList,omitempty"`
	// symbol
	Symbol               string `json:"symbol"`
	AdditionalProperties map[string]interface{}
}

type _MarginBatchOrdersRequest MarginBatchOrdersRequest

// NewMarginBatchOrdersRequest instantiates a new MarginBatchOrdersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginBatchOrdersRequest(symbol string) *MarginBatchOrdersRequest {
	this := MarginBatchOrdersRequest{}
	this.Symbol = symbol
	return &this
}

// NewMarginBatchOrdersRequestWithDefaults instantiates a new MarginBatchOrdersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginBatchOrdersRequestWithDefaults() *MarginBatchOrdersRequest {
	this := MarginBatchOrdersRequest{}
	return &this
}

// GetChannelApiCode returns the ChannelApiCode field value if set, zero value otherwise.
func (o *MarginBatchOrdersRequest) GetChannelApiCode() string {
	if o == nil || isNil(o.ChannelApiCode) {
		var ret string
		return ret
	}
	return *o.ChannelApiCode
}

// GetChannelApiCodeOk returns a tuple with the ChannelApiCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBatchOrdersRequest) GetChannelApiCodeOk() (*string, bool) {
	if o == nil || isNil(o.ChannelApiCode) {
		return nil, false
	}
	return o.ChannelApiCode, true
}

// HasChannelApiCode returns a boolean if a field has been set.
func (o *MarginBatchOrdersRequest) HasChannelApiCode() bool {
	if o != nil && !isNil(o.ChannelApiCode) {
		return true
	}

	return false
}

// SetChannelApiCode gets a reference to the given string and assigns it to the ChannelApiCode field.
func (o *MarginBatchOrdersRequest) SetChannelApiCode(v string) {
	o.ChannelApiCode = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *MarginBatchOrdersRequest) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBatchOrdersRequest) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *MarginBatchOrdersRequest) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *MarginBatchOrdersRequest) SetIp(v string) {
	o.Ip = &v
}

// GetOrderList returns the OrderList field value if set, zero value otherwise.
func (o *MarginBatchOrdersRequest) GetOrderList() []MarginOrderRequest {
	if o == nil || isNil(o.OrderList) {
		var ret []MarginOrderRequest
		return ret
	}
	return o.OrderList
}

// GetOrderListOk returns a tuple with the OrderList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBatchOrdersRequest) GetOrderListOk() ([]MarginOrderRequest, bool) {
	if o == nil || isNil(o.OrderList) {
		return nil, false
	}
	return o.OrderList, true
}

// HasOrderList returns a boolean if a field has been set.
func (o *MarginBatchOrdersRequest) HasOrderList() bool {
	if o != nil && !isNil(o.OrderList) {
		return true
	}

	return false
}

// SetOrderList gets a reference to the given []MarginOrderRequest and assigns it to the OrderList field.
func (o *MarginBatchOrdersRequest) SetOrderList(v []MarginOrderRequest) {
	o.OrderList = v
}

// GetSymbol returns the Symbol field value
func (o *MarginBatchOrdersRequest) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *MarginBatchOrdersRequest) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *MarginBatchOrdersRequest) SetSymbol(v string) {
	o.Symbol = v
}

func (o MarginBatchOrdersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ChannelApiCode) {
		toSerialize["channelApiCode"] = o.ChannelApiCode
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.OrderList) {
		toSerialize["orderList"] = o.OrderList
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginBatchOrdersRequest) UnmarshalJSON(bytes []byte) (err error) {
	varMarginBatchOrdersRequest := _MarginBatchOrdersRequest{}

	if err = json.Unmarshal(bytes, &varMarginBatchOrdersRequest); err == nil {
		*o = MarginBatchOrdersRequest(varMarginBatchOrdersRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "channelApiCode")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "orderList")
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginBatchOrdersRequest struct {
	value *MarginBatchOrdersRequest
	isSet bool
}

func (v NullableMarginBatchOrdersRequest) Get() *MarginBatchOrdersRequest {
	return v.value
}

func (v *NullableMarginBatchOrdersRequest) Set(val *MarginBatchOrdersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginBatchOrdersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginBatchOrdersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginBatchOrdersRequest(val *MarginBatchOrdersRequest) *NullableMarginBatchOrdersRequest {
	return &NullableMarginBatchOrdersRequest{value: val, isSet: true}
}

func (v NullableMarginBatchOrdersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginBatchOrdersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
