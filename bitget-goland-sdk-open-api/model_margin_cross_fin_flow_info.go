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

// MarginCrossFinFlowInfo struct for MarginCrossFinFlowInfo
type MarginCrossFinFlowInfo struct {
	Amount               *string `json:"amount,omitempty"`
	Balance              *string `json:"balance,omitempty"`
	Coin                 *string `json:"coin,omitempty"`
	Ctime                *string `json:"ctime,omitempty"`
	Fee                  *string `json:"fee,omitempty"`
	MarginId             *string `json:"marginId,omitempty"`
	MarginType           *string `json:"marginType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginCrossFinFlowInfo MarginCrossFinFlowInfo

// NewMarginCrossFinFlowInfo instantiates a new MarginCrossFinFlowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCrossFinFlowInfo() *MarginCrossFinFlowInfo {
	this := MarginCrossFinFlowInfo{}
	return &this
}

// NewMarginCrossFinFlowInfoWithDefaults instantiates a new MarginCrossFinFlowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCrossFinFlowInfoWithDefaults() *MarginCrossFinFlowInfo {
	this := MarginCrossFinFlowInfo{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MarginCrossFinFlowInfo) GetAmount() string {
	if o == nil || isNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossFinFlowInfo) GetAmountOk() (*string, bool) {
	if o == nil || isNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MarginCrossFinFlowInfo) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *MarginCrossFinFlowInfo) SetAmount(v string) {
	o.Amount = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *MarginCrossFinFlowInfo) GetBalance() string {
	if o == nil || isNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossFinFlowInfo) GetBalanceOk() (*string, bool) {
	if o == nil || isNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *MarginCrossFinFlowInfo) HasBalance() bool {
	if o != nil && !isNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *MarginCrossFinFlowInfo) SetBalance(v string) {
	o.Balance = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *MarginCrossFinFlowInfo) GetCoin() string {
	if o == nil || isNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossFinFlowInfo) GetCoinOk() (*string, bool) {
	if o == nil || isNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *MarginCrossFinFlowInfo) HasCoin() bool {
	if o != nil && !isNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *MarginCrossFinFlowInfo) SetCoin(v string) {
	o.Coin = &v
}

// GetCtime returns the Ctime field value if set, zero value otherwise.
func (o *MarginCrossFinFlowInfo) GetCtime() string {
	if o == nil || isNil(o.Ctime) {
		var ret string
		return ret
	}
	return *o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossFinFlowInfo) GetCtimeOk() (*string, bool) {
	if o == nil || isNil(o.Ctime) {
		return nil, false
	}
	return o.Ctime, true
}

// HasCtime returns a boolean if a field has been set.
func (o *MarginCrossFinFlowInfo) HasCtime() bool {
	if o != nil && !isNil(o.Ctime) {
		return true
	}

	return false
}

// SetCtime gets a reference to the given string and assigns it to the Ctime field.
func (o *MarginCrossFinFlowInfo) SetCtime(v string) {
	o.Ctime = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *MarginCrossFinFlowInfo) GetFee() string {
	if o == nil || isNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossFinFlowInfo) GetFeeOk() (*string, bool) {
	if o == nil || isNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *MarginCrossFinFlowInfo) HasFee() bool {
	if o != nil && !isNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *MarginCrossFinFlowInfo) SetFee(v string) {
	o.Fee = &v
}

// GetMarginId returns the MarginId field value if set, zero value otherwise.
func (o *MarginCrossFinFlowInfo) GetMarginId() string {
	if o == nil || isNil(o.MarginId) {
		var ret string
		return ret
	}
	return *o.MarginId
}

// GetMarginIdOk returns a tuple with the MarginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossFinFlowInfo) GetMarginIdOk() (*string, bool) {
	if o == nil || isNil(o.MarginId) {
		return nil, false
	}
	return o.MarginId, true
}

// HasMarginId returns a boolean if a field has been set.
func (o *MarginCrossFinFlowInfo) HasMarginId() bool {
	if o != nil && !isNil(o.MarginId) {
		return true
	}

	return false
}

// SetMarginId gets a reference to the given string and assigns it to the MarginId field.
func (o *MarginCrossFinFlowInfo) SetMarginId(v string) {
	o.MarginId = &v
}

// GetMarginType returns the MarginType field value if set, zero value otherwise.
func (o *MarginCrossFinFlowInfo) GetMarginType() string {
	if o == nil || isNil(o.MarginType) {
		var ret string
		return ret
	}
	return *o.MarginType
}

// GetMarginTypeOk returns a tuple with the MarginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossFinFlowInfo) GetMarginTypeOk() (*string, bool) {
	if o == nil || isNil(o.MarginType) {
		return nil, false
	}
	return o.MarginType, true
}

// HasMarginType returns a boolean if a field has been set.
func (o *MarginCrossFinFlowInfo) HasMarginType() bool {
	if o != nil && !isNil(o.MarginType) {
		return true
	}

	return false
}

// SetMarginType gets a reference to the given string and assigns it to the MarginType field.
func (o *MarginCrossFinFlowInfo) SetMarginType(v string) {
	o.MarginType = &v
}

func (o MarginCrossFinFlowInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !isNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !isNil(o.Ctime) {
		toSerialize["ctime"] = o.Ctime
	}
	if !isNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !isNil(o.MarginId) {
		toSerialize["marginId"] = o.MarginId
	}
	if !isNil(o.MarginType) {
		toSerialize["marginType"] = o.MarginType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginCrossFinFlowInfo) UnmarshalJSON(bytes []byte) (err error) {
	varMarginCrossFinFlowInfo := _MarginCrossFinFlowInfo{}

	if err = json.Unmarshal(bytes, &varMarginCrossFinFlowInfo); err == nil {
		*o = MarginCrossFinFlowInfo(varMarginCrossFinFlowInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "ctime")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "marginId")
		delete(additionalProperties, "marginType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginCrossFinFlowInfo struct {
	value *MarginCrossFinFlowInfo
	isSet bool
}

func (v NullableMarginCrossFinFlowInfo) Get() *MarginCrossFinFlowInfo {
	return v.value
}

func (v *NullableMarginCrossFinFlowInfo) Set(val *MarginCrossFinFlowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCrossFinFlowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCrossFinFlowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCrossFinFlowInfo(val *MarginCrossFinFlowInfo) *NullableMarginCrossFinFlowInfo {
	return &NullableMarginCrossFinFlowInfo{value: val, isSet: true}
}

func (v NullableMarginCrossFinFlowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCrossFinFlowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
