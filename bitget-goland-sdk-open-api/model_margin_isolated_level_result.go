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

// MarginIsolatedLevelResult struct for MarginIsolatedLevelResult
type MarginIsolatedLevelResult struct {
	BaseCoin                 *string `json:"baseCoin,omitempty"`
	BaseMaxBorrowableAmount  *string `json:"baseMaxBorrowableAmount,omitempty"`
	InitRate                 *string `json:"initRate,omitempty"`
	Leverage                 *string `json:"leverage,omitempty"`
	MaintainMarginRate       *string `json:"maintainMarginRate,omitempty"`
	QuoteCoin                *string `json:"quoteCoin,omitempty"`
	QuoteMaxBorrowableAmount *string `json:"quoteMaxBorrowableAmount,omitempty"`
	Symbol                   *string `json:"symbol,omitempty"`
	Tier                     *string `json:"tier,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _MarginIsolatedLevelResult MarginIsolatedLevelResult

// NewMarginIsolatedLevelResult instantiates a new MarginIsolatedLevelResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginIsolatedLevelResult() *MarginIsolatedLevelResult {
	this := MarginIsolatedLevelResult{}
	return &this
}

// NewMarginIsolatedLevelResultWithDefaults instantiates a new MarginIsolatedLevelResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginIsolatedLevelResultWithDefaults() *MarginIsolatedLevelResult {
	this := MarginIsolatedLevelResult{}
	return &this
}

// GetBaseCoin returns the BaseCoin field value if set, zero value otherwise.
func (o *MarginIsolatedLevelResult) GetBaseCoin() string {
	if o == nil || isNil(o.BaseCoin) {
		var ret string
		return ret
	}
	return *o.BaseCoin
}

// GetBaseCoinOk returns a tuple with the BaseCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLevelResult) GetBaseCoinOk() (*string, bool) {
	if o == nil || isNil(o.BaseCoin) {
		return nil, false
	}
	return o.BaseCoin, true
}

// HasBaseCoin returns a boolean if a field has been set.
func (o *MarginIsolatedLevelResult) HasBaseCoin() bool {
	if o != nil && !isNil(o.BaseCoin) {
		return true
	}

	return false
}

// SetBaseCoin gets a reference to the given string and assigns it to the BaseCoin field.
func (o *MarginIsolatedLevelResult) SetBaseCoin(v string) {
	o.BaseCoin = &v
}

// GetBaseMaxBorrowableAmount returns the BaseMaxBorrowableAmount field value if set, zero value otherwise.
func (o *MarginIsolatedLevelResult) GetBaseMaxBorrowableAmount() string {
	if o == nil || isNil(o.BaseMaxBorrowableAmount) {
		var ret string
		return ret
	}
	return *o.BaseMaxBorrowableAmount
}

// GetBaseMaxBorrowableAmountOk returns a tuple with the BaseMaxBorrowableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLevelResult) GetBaseMaxBorrowableAmountOk() (*string, bool) {
	if o == nil || isNil(o.BaseMaxBorrowableAmount) {
		return nil, false
	}
	return o.BaseMaxBorrowableAmount, true
}

// HasBaseMaxBorrowableAmount returns a boolean if a field has been set.
func (o *MarginIsolatedLevelResult) HasBaseMaxBorrowableAmount() bool {
	if o != nil && !isNil(o.BaseMaxBorrowableAmount) {
		return true
	}

	return false
}

// SetBaseMaxBorrowableAmount gets a reference to the given string and assigns it to the BaseMaxBorrowableAmount field.
func (o *MarginIsolatedLevelResult) SetBaseMaxBorrowableAmount(v string) {
	o.BaseMaxBorrowableAmount = &v
}

// GetInitRate returns the InitRate field value if set, zero value otherwise.
func (o *MarginIsolatedLevelResult) GetInitRate() string {
	if o == nil || isNil(o.InitRate) {
		var ret string
		return ret
	}
	return *o.InitRate
}

// GetInitRateOk returns a tuple with the InitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLevelResult) GetInitRateOk() (*string, bool) {
	if o == nil || isNil(o.InitRate) {
		return nil, false
	}
	return o.InitRate, true
}

// HasInitRate returns a boolean if a field has been set.
func (o *MarginIsolatedLevelResult) HasInitRate() bool {
	if o != nil && !isNil(o.InitRate) {
		return true
	}

	return false
}

// SetInitRate gets a reference to the given string and assigns it to the InitRate field.
func (o *MarginIsolatedLevelResult) SetInitRate(v string) {
	o.InitRate = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *MarginIsolatedLevelResult) GetLeverage() string {
	if o == nil || isNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLevelResult) GetLeverageOk() (*string, bool) {
	if o == nil || isNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *MarginIsolatedLevelResult) HasLeverage() bool {
	if o != nil && !isNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *MarginIsolatedLevelResult) SetLeverage(v string) {
	o.Leverage = &v
}

// GetMaintainMarginRate returns the MaintainMarginRate field value if set, zero value otherwise.
func (o *MarginIsolatedLevelResult) GetMaintainMarginRate() string {
	if o == nil || isNil(o.MaintainMarginRate) {
		var ret string
		return ret
	}
	return *o.MaintainMarginRate
}

// GetMaintainMarginRateOk returns a tuple with the MaintainMarginRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLevelResult) GetMaintainMarginRateOk() (*string, bool) {
	if o == nil || isNil(o.MaintainMarginRate) {
		return nil, false
	}
	return o.MaintainMarginRate, true
}

// HasMaintainMarginRate returns a boolean if a field has been set.
func (o *MarginIsolatedLevelResult) HasMaintainMarginRate() bool {
	if o != nil && !isNil(o.MaintainMarginRate) {
		return true
	}

	return false
}

// SetMaintainMarginRate gets a reference to the given string and assigns it to the MaintainMarginRate field.
func (o *MarginIsolatedLevelResult) SetMaintainMarginRate(v string) {
	o.MaintainMarginRate = &v
}

// GetQuoteCoin returns the QuoteCoin field value if set, zero value otherwise.
func (o *MarginIsolatedLevelResult) GetQuoteCoin() string {
	if o == nil || isNil(o.QuoteCoin) {
		var ret string
		return ret
	}
	return *o.QuoteCoin
}

// GetQuoteCoinOk returns a tuple with the QuoteCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLevelResult) GetQuoteCoinOk() (*string, bool) {
	if o == nil || isNil(o.QuoteCoin) {
		return nil, false
	}
	return o.QuoteCoin, true
}

// HasQuoteCoin returns a boolean if a field has been set.
func (o *MarginIsolatedLevelResult) HasQuoteCoin() bool {
	if o != nil && !isNil(o.QuoteCoin) {
		return true
	}

	return false
}

// SetQuoteCoin gets a reference to the given string and assigns it to the QuoteCoin field.
func (o *MarginIsolatedLevelResult) SetQuoteCoin(v string) {
	o.QuoteCoin = &v
}

// GetQuoteMaxBorrowableAmount returns the QuoteMaxBorrowableAmount field value if set, zero value otherwise.
func (o *MarginIsolatedLevelResult) GetQuoteMaxBorrowableAmount() string {
	if o == nil || isNil(o.QuoteMaxBorrowableAmount) {
		var ret string
		return ret
	}
	return *o.QuoteMaxBorrowableAmount
}

// GetQuoteMaxBorrowableAmountOk returns a tuple with the QuoteMaxBorrowableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLevelResult) GetQuoteMaxBorrowableAmountOk() (*string, bool) {
	if o == nil || isNil(o.QuoteMaxBorrowableAmount) {
		return nil, false
	}
	return o.QuoteMaxBorrowableAmount, true
}

// HasQuoteMaxBorrowableAmount returns a boolean if a field has been set.
func (o *MarginIsolatedLevelResult) HasQuoteMaxBorrowableAmount() bool {
	if o != nil && !isNil(o.QuoteMaxBorrowableAmount) {
		return true
	}

	return false
}

// SetQuoteMaxBorrowableAmount gets a reference to the given string and assigns it to the QuoteMaxBorrowableAmount field.
func (o *MarginIsolatedLevelResult) SetQuoteMaxBorrowableAmount(v string) {
	o.QuoteMaxBorrowableAmount = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginIsolatedLevelResult) GetSymbol() string {
	if o == nil || isNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLevelResult) GetSymbolOk() (*string, bool) {
	if o == nil || isNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginIsolatedLevelResult) HasSymbol() bool {
	if o != nil && !isNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginIsolatedLevelResult) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *MarginIsolatedLevelResult) GetTier() string {
	if o == nil || isNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLevelResult) GetTierOk() (*string, bool) {
	if o == nil || isNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *MarginIsolatedLevelResult) HasTier() bool {
	if o != nil && !isNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *MarginIsolatedLevelResult) SetTier(v string) {
	o.Tier = &v
}

func (o MarginIsolatedLevelResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BaseCoin) {
		toSerialize["baseCoin"] = o.BaseCoin
	}
	if !isNil(o.BaseMaxBorrowableAmount) {
		toSerialize["baseMaxBorrowableAmount"] = o.BaseMaxBorrowableAmount
	}
	if !isNil(o.InitRate) {
		toSerialize["initRate"] = o.InitRate
	}
	if !isNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !isNil(o.MaintainMarginRate) {
		toSerialize["maintainMarginRate"] = o.MaintainMarginRate
	}
	if !isNil(o.QuoteCoin) {
		toSerialize["quoteCoin"] = o.QuoteCoin
	}
	if !isNil(o.QuoteMaxBorrowableAmount) {
		toSerialize["quoteMaxBorrowableAmount"] = o.QuoteMaxBorrowableAmount
	}
	if !isNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !isNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginIsolatedLevelResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginIsolatedLevelResult := _MarginIsolatedLevelResult{}

	if err = json.Unmarshal(bytes, &varMarginIsolatedLevelResult); err == nil {
		*o = MarginIsolatedLevelResult(varMarginIsolatedLevelResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "baseCoin")
		delete(additionalProperties, "baseMaxBorrowableAmount")
		delete(additionalProperties, "initRate")
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "maintainMarginRate")
		delete(additionalProperties, "quoteCoin")
		delete(additionalProperties, "quoteMaxBorrowableAmount")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "tier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginIsolatedLevelResult struct {
	value *MarginIsolatedLevelResult
	isSet bool
}

func (v NullableMarginIsolatedLevelResult) Get() *MarginIsolatedLevelResult {
	return v.value
}

func (v *NullableMarginIsolatedLevelResult) Set(val *MarginIsolatedLevelResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginIsolatedLevelResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginIsolatedLevelResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginIsolatedLevelResult(val *MarginIsolatedLevelResult) *NullableMarginIsolatedLevelResult {
	return &NullableMarginIsolatedLevelResult{value: val, isSet: true}
}

func (v NullableMarginIsolatedLevelResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginIsolatedLevelResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
