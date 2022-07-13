/*
 * Kafka Instance API
 *
 * API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs
 *
 * API version: 0.12.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// TopicSettings The settings that are applicable to this topic. This includes partitions, configuration information, and number of replicas.
type TopicSettings struct {
	// Number of partitions for this topic. If not specified, the default for new topics is `1`. Number of partitions may not be reduced when updating existing topics
	NumPartitions *int32 `json:"numPartitions,omitempty"`
	// Topic configuration entries.
	Config *[]ConfigEntry `json:"config,omitempty"`
}

// NewTopicSettings instantiates a new TopicSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicSettings() *TopicSettings {
	this := TopicSettings{}
	return &this
}

// NewTopicSettingsWithDefaults instantiates a new TopicSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicSettingsWithDefaults() *TopicSettings {
	this := TopicSettings{}
	return &this
}

// GetNumPartitions returns the NumPartitions field value if set, zero value otherwise.
func (o *TopicSettings) GetNumPartitions() int32 {
	if o == nil || o.NumPartitions == nil {
		var ret int32
		return ret
	}
	return *o.NumPartitions
}

// GetNumPartitionsOk returns a tuple with the NumPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicSettings) GetNumPartitionsOk() (*int32, bool) {
	if o == nil || o.NumPartitions == nil {
		return nil, false
	}
	return o.NumPartitions, true
}

// HasNumPartitions returns a boolean if a field has been set.
func (o *TopicSettings) HasNumPartitions() bool {
	if o != nil && o.NumPartitions != nil {
		return true
	}

	return false
}

// SetNumPartitions gets a reference to the given int32 and assigns it to the NumPartitions field.
func (o *TopicSettings) SetNumPartitions(v int32) {
	o.NumPartitions = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *TopicSettings) GetConfig() []ConfigEntry {
	if o == nil || o.Config == nil {
		var ret []ConfigEntry
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicSettings) GetConfigOk() (*[]ConfigEntry, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *TopicSettings) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given []ConfigEntry and assigns it to the Config field.
func (o *TopicSettings) SetConfig(v []ConfigEntry) {
	o.Config = &v
}

func (o TopicSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumPartitions != nil {
		toSerialize["numPartitions"] = o.NumPartitions
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableTopicSettings struct {
	value *TopicSettings
	isSet bool
}

func (v NullableTopicSettings) Get() *TopicSettings {
	return v.value
}

func (v *NullableTopicSettings) Set(val *TopicSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicSettings(val *TopicSettings) *NullableTopicSettings {
	return &NullableTopicSettings{value: val, isSet: true}
}

func (v NullableTopicSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


