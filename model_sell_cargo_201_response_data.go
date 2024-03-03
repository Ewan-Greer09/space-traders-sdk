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

// checks if the SellCargo201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SellCargo201ResponseData{}

// SellCargo201ResponseData struct for SellCargo201ResponseData
type SellCargo201ResponseData struct {
	Agent Agent `json:"agent"`
	Cargo ShipCargo `json:"cargo"`
	Transaction MarketTransaction `json:"transaction"`
}

type _SellCargo201ResponseData SellCargo201ResponseData

// NewSellCargo201ResponseData instantiates a new SellCargo201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSellCargo201ResponseData(agent Agent, cargo ShipCargo, transaction MarketTransaction) *SellCargo201ResponseData {
	this := SellCargo201ResponseData{}
	this.Agent = agent
	this.Cargo = cargo
	this.Transaction = transaction
	return &this
}

// NewSellCargo201ResponseDataWithDefaults instantiates a new SellCargo201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSellCargo201ResponseDataWithDefaults() *SellCargo201ResponseData {
	this := SellCargo201ResponseData{}
	return &this
}

// GetAgent returns the Agent field value
func (o *SellCargo201ResponseData) GetAgent() Agent {
	if o == nil {
		var ret Agent
		return ret
	}

	return o.Agent
}

// GetAgentOk returns a tuple with the Agent field value
// and a boolean to check if the value has been set.
func (o *SellCargo201ResponseData) GetAgentOk() (*Agent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent, true
}

// SetAgent sets field value
func (o *SellCargo201ResponseData) SetAgent(v Agent) {
	o.Agent = v
}

// GetCargo returns the Cargo field value
func (o *SellCargo201ResponseData) GetCargo() ShipCargo {
	if o == nil {
		var ret ShipCargo
		return ret
	}

	return o.Cargo
}

// GetCargoOk returns a tuple with the Cargo field value
// and a boolean to check if the value has been set.
func (o *SellCargo201ResponseData) GetCargoOk() (*ShipCargo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cargo, true
}

// SetCargo sets field value
func (o *SellCargo201ResponseData) SetCargo(v ShipCargo) {
	o.Cargo = v
}

// GetTransaction returns the Transaction field value
func (o *SellCargo201ResponseData) GetTransaction() MarketTransaction {
	if o == nil {
		var ret MarketTransaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *SellCargo201ResponseData) GetTransactionOk() (*MarketTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *SellCargo201ResponseData) SetTransaction(v MarketTransaction) {
	o.Transaction = v
}

func (o SellCargo201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SellCargo201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent"] = o.Agent
	toSerialize["cargo"] = o.Cargo
	toSerialize["transaction"] = o.Transaction
	return toSerialize, nil
}

func (o *SellCargo201ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agent",
		"cargo",
		"transaction",
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

	varSellCargo201ResponseData := _SellCargo201ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSellCargo201ResponseData)

	if err != nil {
		return err
	}

	*o = SellCargo201ResponseData(varSellCargo201ResponseData)

	return err
}

type NullableSellCargo201ResponseData struct {
	value *SellCargo201ResponseData
	isSet bool
}

func (v NullableSellCargo201ResponseData) Get() *SellCargo201ResponseData {
	return v.value
}

func (v *NullableSellCargo201ResponseData) Set(val *SellCargo201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSellCargo201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSellCargo201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSellCargo201ResponseData(val *SellCargo201ResponseData) *NullableSellCargo201ResponseData {
	return &NullableSellCargo201ResponseData{value: val, isSet: true}
}

func (v NullableSellCargo201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSellCargo201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


