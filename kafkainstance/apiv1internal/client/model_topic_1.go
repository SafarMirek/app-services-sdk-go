/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.11.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// Topic1 Kafka Topic (A feed where records are stored and published)
type Topic1 struct {

	// The name of the topic.
	Name *string `json:"name,omitempty"`

	IsInternal *bool `json:"isInternal,omitempty"`

	// Partitions for this topic.
	Partitions *[]Partition `json:"partitions,omitempty"`

	// Topic configuration entry.
	Config *[]ConfigEntry `json:"config,omitempty"`

}

// NewTopic1 instantiates a new Topic1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopic1() *Topic1 {
	this := Topic1{}
	return &this
}

// NewTopic1WithDefaults instantiates a new Topic1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopic1WithDefaults() *Topic1 {
	this := Topic1{}





	return &this
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *Topic1) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic1) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Topic1) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Topic1) SetName(v string) {
	o.Name = &v
}


// GetIsInternal returns the IsInternal field value if set, zero value otherwise.
func (o *Topic1) GetIsInternal() bool {
	if o == nil || o.IsInternal == nil {
		var ret bool
		return ret
	}
	return *o.IsInternal
}

// GetIsInternalOk returns a tuple with the IsInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic1) GetIsInternalOk() (*bool, bool) {
	if o == nil || o.IsInternal == nil {
		return nil, false
	}
	return o.IsInternal, true
}

// HasIsInternal returns a boolean if a field has been set.
func (o *Topic1) HasIsInternal() bool {
	if o != nil && o.IsInternal != nil {
		return true
	}

	return false
}

// SetIsInternal gets a reference to the given bool and assigns it to the IsInternal field.
func (o *Topic1) SetIsInternal(v bool) {
	o.IsInternal = &v
}


// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *Topic1) GetPartitions() []Partition {
	if o == nil || o.Partitions == nil {
		var ret []Partition
		return ret
	}
	return *o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic1) GetPartitionsOk() (*[]Partition, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *Topic1) HasPartitions() bool {
	if o != nil && o.Partitions != nil {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []Partition and assigns it to the Partitions field.
func (o *Topic1) SetPartitions(v []Partition) {
	o.Partitions = &v
}


// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Topic1) GetConfig() []ConfigEntry {
	if o == nil || o.Config == nil {
		var ret []ConfigEntry
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic1) GetConfigOk() (*[]ConfigEntry, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Topic1) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given []ConfigEntry and assigns it to the Config field.
func (o *Topic1) SetConfig(v []ConfigEntry) {
	o.Config = &v
}


func (o Topic1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.IsInternal != nil {
		toSerialize["isInternal"] = o.IsInternal
	}
    
	if o.Partitions != nil {
		toSerialize["partitions"] = o.Partitions
	}
    
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
    
	return json.Marshal(toSerialize)
}

type NullableTopic1 struct {
	value *Topic1
	isSet bool
}

func (v NullableTopic1) Get() *Topic1 {
	return v.value
}

func (v *NullableTopic1) Set(val *Topic1) {
	v.value = val
	v.isSet = true
}

func (v NullableTopic1) IsSet() bool {
	return v.isSet
}

func (v *NullableTopic1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopic1(val *Topic1) *NullableTopic1 {
	return &NullableTopic1{value: val, isSet: true}
}

func (v NullableTopic1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopic1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

