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

// MarginIsolatedLimitRequest struct for MarginIsolatedLimitRequest
type MarginIsolatedLimitRequest struct {
	// borrowAmount
	BorrowAmount string `json:"borrowAmount"`
	// coin
	Coin string `json:"coin"`
	// symbol
	Symbol               string `json:"symbol"`
	AdditionalProperties map[string]interface{}
}

type _MarginIsolatedLimitRequest MarginIsolatedLimitRequest

// NewMarginIsolatedLimitRequest instantiates a new MarginIsolatedLimitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginIsolatedLimitRequest(borrowAmount string, coin string, symbol string) *MarginIsolatedLimitRequest {
	this := MarginIsolatedLimitRequest{}
	this.BorrowAmount = borrowAmount
	this.Coin = coin
	this.Symbol = symbol
	return &this
}

// NewMarginIsolatedLimitRequestWithDefaults instantiates a new MarginIsolatedLimitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginIsolatedLimitRequestWithDefaults() *MarginIsolatedLimitRequest {
	this := MarginIsolatedLimitRequest{}
	return &this
}

// GetBorrowAmount returns the BorrowAmount field value
func (o *MarginIsolatedLimitRequest) GetBorrowAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BorrowAmount
}

// GetBorrowAmountOk returns a tuple with the BorrowAmount field value
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLimitRequest) GetBorrowAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BorrowAmount, true
}

// SetBorrowAmount sets field value
func (o *MarginIsolatedLimitRequest) SetBorrowAmount(v string) {
	o.BorrowAmount = v
}

// GetCoin returns the Coin field value
func (o *MarginIsolatedLimitRequest) GetCoin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coin
}

// GetCoinOk returns a tuple with the Coin field value
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLimitRequest) GetCoinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coin, true
}

// SetCoin sets field value
func (o *MarginIsolatedLimitRequest) SetCoin(v string) {
	o.Coin = v
}

// GetSymbol returns the Symbol field value
func (o *MarginIsolatedLimitRequest) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLimitRequest) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *MarginIsolatedLimitRequest) SetSymbol(v string) {
	o.Symbol = v
}

func (o MarginIsolatedLimitRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["borrowAmount"] = o.BorrowAmount
	}
	if true {
		toSerialize["coin"] = o.Coin
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginIsolatedLimitRequest) UnmarshalJSON(bytes []byte) (err error) {
	varMarginIsolatedLimitRequest := _MarginIsolatedLimitRequest{}

	if err = json.Unmarshal(bytes, &varMarginIsolatedLimitRequest); err == nil {
		*o = MarginIsolatedLimitRequest(varMarginIsolatedLimitRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "borrowAmount")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginIsolatedLimitRequest struct {
	value *MarginIsolatedLimitRequest
	isSet bool
}

func (v NullableMarginIsolatedLimitRequest) Get() *MarginIsolatedLimitRequest {
	return v.value
}

func (v *NullableMarginIsolatedLimitRequest) Set(val *MarginIsolatedLimitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginIsolatedLimitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginIsolatedLimitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginIsolatedLimitRequest(val *MarginIsolatedLimitRequest) *NullableMarginIsolatedLimitRequest {
	return &NullableMarginIsolatedLimitRequest{value: val, isSet: true}
}

func (v NullableMarginIsolatedLimitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginIsolatedLimitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
