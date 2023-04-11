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

// MarginIsolatedMaxBorrowResult struct for MarginIsolatedMaxBorrowResult
type MarginIsolatedMaxBorrowResult struct {
	Coin                 *string `json:"coin,omitempty"`
	MaxBorrowableAmount  *string `json:"maxBorrowableAmount,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginIsolatedMaxBorrowResult MarginIsolatedMaxBorrowResult

// NewMarginIsolatedMaxBorrowResult instantiates a new MarginIsolatedMaxBorrowResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginIsolatedMaxBorrowResult() *MarginIsolatedMaxBorrowResult {
	this := MarginIsolatedMaxBorrowResult{}
	return &this
}

// NewMarginIsolatedMaxBorrowResultWithDefaults instantiates a new MarginIsolatedMaxBorrowResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginIsolatedMaxBorrowResultWithDefaults() *MarginIsolatedMaxBorrowResult {
	this := MarginIsolatedMaxBorrowResult{}
	return &this
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *MarginIsolatedMaxBorrowResult) GetCoin() string {
	if o == nil || isNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedMaxBorrowResult) GetCoinOk() (*string, bool) {
	if o == nil || isNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *MarginIsolatedMaxBorrowResult) HasCoin() bool {
	if o != nil && !isNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *MarginIsolatedMaxBorrowResult) SetCoin(v string) {
	o.Coin = &v
}

// GetMaxBorrowableAmount returns the MaxBorrowableAmount field value if set, zero value otherwise.
func (o *MarginIsolatedMaxBorrowResult) GetMaxBorrowableAmount() string {
	if o == nil || isNil(o.MaxBorrowableAmount) {
		var ret string
		return ret
	}
	return *o.MaxBorrowableAmount
}

// GetMaxBorrowableAmountOk returns a tuple with the MaxBorrowableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedMaxBorrowResult) GetMaxBorrowableAmountOk() (*string, bool) {
	if o == nil || isNil(o.MaxBorrowableAmount) {
		return nil, false
	}
	return o.MaxBorrowableAmount, true
}

// HasMaxBorrowableAmount returns a boolean if a field has been set.
func (o *MarginIsolatedMaxBorrowResult) HasMaxBorrowableAmount() bool {
	if o != nil && !isNil(o.MaxBorrowableAmount) {
		return true
	}

	return false
}

// SetMaxBorrowableAmount gets a reference to the given string and assigns it to the MaxBorrowableAmount field.
func (o *MarginIsolatedMaxBorrowResult) SetMaxBorrowableAmount(v string) {
	o.MaxBorrowableAmount = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginIsolatedMaxBorrowResult) GetSymbol() string {
	if o == nil || isNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedMaxBorrowResult) GetSymbolOk() (*string, bool) {
	if o == nil || isNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginIsolatedMaxBorrowResult) HasSymbol() bool {
	if o != nil && !isNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginIsolatedMaxBorrowResult) SetSymbol(v string) {
	o.Symbol = &v
}

func (o MarginIsolatedMaxBorrowResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !isNil(o.MaxBorrowableAmount) {
		toSerialize["maxBorrowableAmount"] = o.MaxBorrowableAmount
	}
	if !isNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginIsolatedMaxBorrowResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginIsolatedMaxBorrowResult := _MarginIsolatedMaxBorrowResult{}

	if err = json.Unmarshal(bytes, &varMarginIsolatedMaxBorrowResult); err == nil {
		*o = MarginIsolatedMaxBorrowResult(varMarginIsolatedMaxBorrowResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "coin")
		delete(additionalProperties, "maxBorrowableAmount")
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginIsolatedMaxBorrowResult struct {
	value *MarginIsolatedMaxBorrowResult
	isSet bool
}

func (v NullableMarginIsolatedMaxBorrowResult) Get() *MarginIsolatedMaxBorrowResult {
	return v.value
}

func (v *NullableMarginIsolatedMaxBorrowResult) Set(val *MarginIsolatedMaxBorrowResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginIsolatedMaxBorrowResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginIsolatedMaxBorrowResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginIsolatedMaxBorrowResult(val *MarginIsolatedMaxBorrowResult) *NullableMarginIsolatedMaxBorrowResult {
	return &NullableMarginIsolatedMaxBorrowResult{value: val, isSet: true}
}

func (v NullableMarginIsolatedMaxBorrowResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginIsolatedMaxBorrowResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
