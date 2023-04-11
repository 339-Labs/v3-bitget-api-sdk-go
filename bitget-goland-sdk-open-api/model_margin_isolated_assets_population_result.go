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

// MarginIsolatedAssetsPopulationResult struct for MarginIsolatedAssetsPopulationResult
type MarginIsolatedAssetsPopulationResult struct {
	Available            *string `json:"available,omitempty"`
	Borrow               *string `json:"borrow,omitempty"`
	Coin                 *string `json:"coin,omitempty"`
	Ctime                *string `json:"ctime,omitempty"`
	Frozen               *string `json:"frozen,omitempty"`
	Interest             *string `json:"interest,omitempty"`
	Net                  *string `json:"net,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	TotalAmount          *string `json:"totalAmount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginIsolatedAssetsPopulationResult MarginIsolatedAssetsPopulationResult

// NewMarginIsolatedAssetsPopulationResult instantiates a new MarginIsolatedAssetsPopulationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginIsolatedAssetsPopulationResult() *MarginIsolatedAssetsPopulationResult {
	this := MarginIsolatedAssetsPopulationResult{}
	return &this
}

// NewMarginIsolatedAssetsPopulationResultWithDefaults instantiates a new MarginIsolatedAssetsPopulationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginIsolatedAssetsPopulationResultWithDefaults() *MarginIsolatedAssetsPopulationResult {
	this := MarginIsolatedAssetsPopulationResult{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *MarginIsolatedAssetsPopulationResult) GetAvailable() string {
	if o == nil || isNil(o.Available) {
		var ret string
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedAssetsPopulationResult) GetAvailableOk() (*string, bool) {
	if o == nil || isNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *MarginIsolatedAssetsPopulationResult) HasAvailable() bool {
	if o != nil && !isNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given string and assigns it to the Available field.
func (o *MarginIsolatedAssetsPopulationResult) SetAvailable(v string) {
	o.Available = &v
}

// GetBorrow returns the Borrow field value if set, zero value otherwise.
func (o *MarginIsolatedAssetsPopulationResult) GetBorrow() string {
	if o == nil || isNil(o.Borrow) {
		var ret string
		return ret
	}
	return *o.Borrow
}

// GetBorrowOk returns a tuple with the Borrow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedAssetsPopulationResult) GetBorrowOk() (*string, bool) {
	if o == nil || isNil(o.Borrow) {
		return nil, false
	}
	return o.Borrow, true
}

// HasBorrow returns a boolean if a field has been set.
func (o *MarginIsolatedAssetsPopulationResult) HasBorrow() bool {
	if o != nil && !isNil(o.Borrow) {
		return true
	}

	return false
}

// SetBorrow gets a reference to the given string and assigns it to the Borrow field.
func (o *MarginIsolatedAssetsPopulationResult) SetBorrow(v string) {
	o.Borrow = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *MarginIsolatedAssetsPopulationResult) GetCoin() string {
	if o == nil || isNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedAssetsPopulationResult) GetCoinOk() (*string, bool) {
	if o == nil || isNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *MarginIsolatedAssetsPopulationResult) HasCoin() bool {
	if o != nil && !isNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *MarginIsolatedAssetsPopulationResult) SetCoin(v string) {
	o.Coin = &v
}

// GetCtime returns the Ctime field value if set, zero value otherwise.
func (o *MarginIsolatedAssetsPopulationResult) GetCtime() string {
	if o == nil || isNil(o.Ctime) {
		var ret string
		return ret
	}
	return *o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedAssetsPopulationResult) GetCtimeOk() (*string, bool) {
	if o == nil || isNil(o.Ctime) {
		return nil, false
	}
	return o.Ctime, true
}

// HasCtime returns a boolean if a field has been set.
func (o *MarginIsolatedAssetsPopulationResult) HasCtime() bool {
	if o != nil && !isNil(o.Ctime) {
		return true
	}

	return false
}

// SetCtime gets a reference to the given string and assigns it to the Ctime field.
func (o *MarginIsolatedAssetsPopulationResult) SetCtime(v string) {
	o.Ctime = &v
}

// GetFrozen returns the Frozen field value if set, zero value otherwise.
func (o *MarginIsolatedAssetsPopulationResult) GetFrozen() string {
	if o == nil || isNil(o.Frozen) {
		var ret string
		return ret
	}
	return *o.Frozen
}

// GetFrozenOk returns a tuple with the Frozen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedAssetsPopulationResult) GetFrozenOk() (*string, bool) {
	if o == nil || isNil(o.Frozen) {
		return nil, false
	}
	return o.Frozen, true
}

// HasFrozen returns a boolean if a field has been set.
func (o *MarginIsolatedAssetsPopulationResult) HasFrozen() bool {
	if o != nil && !isNil(o.Frozen) {
		return true
	}

	return false
}

// SetFrozen gets a reference to the given string and assigns it to the Frozen field.
func (o *MarginIsolatedAssetsPopulationResult) SetFrozen(v string) {
	o.Frozen = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *MarginIsolatedAssetsPopulationResult) GetInterest() string {
	if o == nil || isNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedAssetsPopulationResult) GetInterestOk() (*string, bool) {
	if o == nil || isNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *MarginIsolatedAssetsPopulationResult) HasInterest() bool {
	if o != nil && !isNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *MarginIsolatedAssetsPopulationResult) SetInterest(v string) {
	o.Interest = &v
}

// GetNet returns the Net field value if set, zero value otherwise.
func (o *MarginIsolatedAssetsPopulationResult) GetNet() string {
	if o == nil || isNil(o.Net) {
		var ret string
		return ret
	}
	return *o.Net
}

// GetNetOk returns a tuple with the Net field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedAssetsPopulationResult) GetNetOk() (*string, bool) {
	if o == nil || isNil(o.Net) {
		return nil, false
	}
	return o.Net, true
}

// HasNet returns a boolean if a field has been set.
func (o *MarginIsolatedAssetsPopulationResult) HasNet() bool {
	if o != nil && !isNil(o.Net) {
		return true
	}

	return false
}

// SetNet gets a reference to the given string and assigns it to the Net field.
func (o *MarginIsolatedAssetsPopulationResult) SetNet(v string) {
	o.Net = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginIsolatedAssetsPopulationResult) GetSymbol() string {
	if o == nil || isNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedAssetsPopulationResult) GetSymbolOk() (*string, bool) {
	if o == nil || isNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginIsolatedAssetsPopulationResult) HasSymbol() bool {
	if o != nil && !isNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginIsolatedAssetsPopulationResult) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *MarginIsolatedAssetsPopulationResult) GetTotalAmount() string {
	if o == nil || isNil(o.TotalAmount) {
		var ret string
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedAssetsPopulationResult) GetTotalAmountOk() (*string, bool) {
	if o == nil || isNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *MarginIsolatedAssetsPopulationResult) HasTotalAmount() bool {
	if o != nil && !isNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given string and assigns it to the TotalAmount field.
func (o *MarginIsolatedAssetsPopulationResult) SetTotalAmount(v string) {
	o.TotalAmount = &v
}

func (o MarginIsolatedAssetsPopulationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !isNil(o.Borrow) {
		toSerialize["borrow"] = o.Borrow
	}
	if !isNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !isNil(o.Ctime) {
		toSerialize["ctime"] = o.Ctime
	}
	if !isNil(o.Frozen) {
		toSerialize["frozen"] = o.Frozen
	}
	if !isNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !isNil(o.Net) {
		toSerialize["net"] = o.Net
	}
	if !isNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !isNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginIsolatedAssetsPopulationResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginIsolatedAssetsPopulationResult := _MarginIsolatedAssetsPopulationResult{}

	if err = json.Unmarshal(bytes, &varMarginIsolatedAssetsPopulationResult); err == nil {
		*o = MarginIsolatedAssetsPopulationResult(varMarginIsolatedAssetsPopulationResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "available")
		delete(additionalProperties, "borrow")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "ctime")
		delete(additionalProperties, "frozen")
		delete(additionalProperties, "interest")
		delete(additionalProperties, "net")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "totalAmount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginIsolatedAssetsPopulationResult struct {
	value *MarginIsolatedAssetsPopulationResult
	isSet bool
}

func (v NullableMarginIsolatedAssetsPopulationResult) Get() *MarginIsolatedAssetsPopulationResult {
	return v.value
}

func (v *NullableMarginIsolatedAssetsPopulationResult) Set(val *MarginIsolatedAssetsPopulationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginIsolatedAssetsPopulationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginIsolatedAssetsPopulationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginIsolatedAssetsPopulationResult(val *MarginIsolatedAssetsPopulationResult) *NullableMarginIsolatedAssetsPopulationResult {
	return &NullableMarginIsolatedAssetsPopulationResult{value: val, isSet: true}
}

func (v NullableMarginIsolatedAssetsPopulationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginIsolatedAssetsPopulationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
