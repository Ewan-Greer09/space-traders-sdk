/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Market type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Market{}

// Market 
type Market struct {
	// The symbol of the market. The symbol is the same as the waypoint where the market is located.
	Symbol string `json:"symbol"`
	// The list of goods that are exported from this market.
	Exports []TradeGood `json:"exports"`
	// The list of goods that are sought as imports in this market.
	Imports []TradeGood `json:"imports"`
	// The list of goods that are bought and sold between agents at this market.
	Exchange []TradeGood `json:"exchange"`
	// The list of recent transactions at this market. Visible only when a ship is present at the market.
	Transactions []MarketTransaction `json:"transactions,omitempty"`
	// The list of goods that are traded at this market. Visible only when a ship is present at the market.
	TradeGoods []MarketTradeGood `json:"tradeGoods,omitempty"`
}

type _Market Market

// NewMarket instantiates a new Market object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarket(symbol string, exports []TradeGood, imports []TradeGood, exchange []TradeGood) *Market {
	this := Market{}
	this.Symbol = symbol
	this.Exports = exports
	this.Imports = imports
	this.Exchange = exchange
	return &this
}

// NewMarketWithDefaults instantiates a new Market object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketWithDefaults() *Market {
	this := Market{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *Market) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *Market) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *Market) SetSymbol(v string) {
	o.Symbol = v
}

// GetExports returns the Exports field value
func (o *Market) GetExports() []TradeGood {
	if o == nil {
		var ret []TradeGood
		return ret
	}

	return o.Exports
}

// GetExportsOk returns a tuple with the Exports field value
// and a boolean to check if the value has been set.
func (o *Market) GetExportsOk() ([]TradeGood, bool) {
	if o == nil {
		return nil, false
	}
	return o.Exports, true
}

// SetExports sets field value
func (o *Market) SetExports(v []TradeGood) {
	o.Exports = v
}

// GetImports returns the Imports field value
func (o *Market) GetImports() []TradeGood {
	if o == nil {
		var ret []TradeGood
		return ret
	}

	return o.Imports
}

// GetImportsOk returns a tuple with the Imports field value
// and a boolean to check if the value has been set.
func (o *Market) GetImportsOk() ([]TradeGood, bool) {
	if o == nil {
		return nil, false
	}
	return o.Imports, true
}

// SetImports sets field value
func (o *Market) SetImports(v []TradeGood) {
	o.Imports = v
}

// GetExchange returns the Exchange field value
func (o *Market) GetExchange() []TradeGood {
	if o == nil {
		var ret []TradeGood
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *Market) GetExchangeOk() ([]TradeGood, bool) {
	if o == nil {
		return nil, false
	}
	return o.Exchange, true
}

// SetExchange sets field value
func (o *Market) SetExchange(v []TradeGood) {
	o.Exchange = v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *Market) GetTransactions() []MarketTransaction {
	if o == nil || IsNil(o.Transactions) {
		var ret []MarketTransaction
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Market) GetTransactionsOk() ([]MarketTransaction, bool) {
	if o == nil || IsNil(o.Transactions) {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *Market) HasTransactions() bool {
	if o != nil && !IsNil(o.Transactions) {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []MarketTransaction and assigns it to the Transactions field.
func (o *Market) SetTransactions(v []MarketTransaction) {
	o.Transactions = v
}

// GetTradeGoods returns the TradeGoods field value if set, zero value otherwise.
func (o *Market) GetTradeGoods() []MarketTradeGood {
	if o == nil || IsNil(o.TradeGoods) {
		var ret []MarketTradeGood
		return ret
	}
	return o.TradeGoods
}

// GetTradeGoodsOk returns a tuple with the TradeGoods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Market) GetTradeGoodsOk() ([]MarketTradeGood, bool) {
	if o == nil || IsNil(o.TradeGoods) {
		return nil, false
	}
	return o.TradeGoods, true
}

// HasTradeGoods returns a boolean if a field has been set.
func (o *Market) HasTradeGoods() bool {
	if o != nil && !IsNil(o.TradeGoods) {
		return true
	}

	return false
}

// SetTradeGoods gets a reference to the given []MarketTradeGood and assigns it to the TradeGoods field.
func (o *Market) SetTradeGoods(v []MarketTradeGood) {
	o.TradeGoods = v
}

func (o Market) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Market) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["exports"] = o.Exports
	toSerialize["imports"] = o.Imports
	toSerialize["exchange"] = o.Exchange
	if !IsNil(o.Transactions) {
		toSerialize["transactions"] = o.Transactions
	}
	if !IsNil(o.TradeGoods) {
		toSerialize["tradeGoods"] = o.TradeGoods
	}
	return toSerialize, nil
}

func (o *Market) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"exports",
		"imports",
		"exchange",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMarket := _Market{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarket)

	if err != nil {
		return err
	}

	*o = Market(varMarket)

	return err
}

type NullableMarket struct {
	value *Market
	isSet bool
}

func (v NullableMarket) Get() *Market {
	return v.value
}

func (v *NullableMarket) Set(val *Market) {
	v.value = val
	v.isSet = true
}

func (v NullableMarket) IsSet() bool {
	return v.isSet
}

func (v *NullableMarket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarket(val *Market) *NullableMarket {
	return &NullableMarket{value: val, isSet: true}
}

func (v NullableMarket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


