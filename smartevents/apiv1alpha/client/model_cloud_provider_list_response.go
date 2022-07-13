/*
 * Red Hat Openshift SmartEvents Fleet Manager
 *
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * API version: 0.0.1
 * Contact: openbridge-dev@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsclient

import (
	"encoding/json"
)

// CloudProviderListResponse struct for CloudProviderListResponse
type CloudProviderListResponse struct {
	Kind *string `json:"kind,omitempty"`
	Items *[]CloudProviderResponse `json:"items,omitempty"`
	Page *int64 `json:"page,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Total *int64 `json:"total,omitempty"`
}

// NewCloudProviderListResponse instantiates a new CloudProviderListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderListResponse() *CloudProviderListResponse {
	this := CloudProviderListResponse{}
	return &this
}

// NewCloudProviderListResponseWithDefaults instantiates a new CloudProviderListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderListResponseWithDefaults() *CloudProviderListResponse {
	this := CloudProviderListResponse{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CloudProviderListResponse) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderListResponse) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CloudProviderListResponse) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CloudProviderListResponse) SetKind(v string) {
	o.Kind = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *CloudProviderListResponse) GetItems() []CloudProviderResponse {
	if o == nil || o.Items == nil {
		var ret []CloudProviderResponse
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderListResponse) GetItemsOk() (*[]CloudProviderResponse, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *CloudProviderListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []CloudProviderResponse and assigns it to the Items field.
func (o *CloudProviderListResponse) SetItems(v []CloudProviderResponse) {
	o.Items = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *CloudProviderListResponse) GetPage() int64 {
	if o == nil || o.Page == nil {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderListResponse) GetPageOk() (*int64, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *CloudProviderListResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *CloudProviderListResponse) SetPage(v int64) {
	o.Page = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CloudProviderListResponse) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderListResponse) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CloudProviderListResponse) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *CloudProviderListResponse) SetSize(v int64) {
	o.Size = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CloudProviderListResponse) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderListResponse) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CloudProviderListResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *CloudProviderListResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o CloudProviderListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableCloudProviderListResponse struct {
	value *CloudProviderListResponse
	isSet bool
}

func (v NullableCloudProviderListResponse) Get() *CloudProviderListResponse {
	return v.value
}

func (v *NullableCloudProviderListResponse) Set(val *CloudProviderListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderListResponse(val *CloudProviderListResponse) *NullableCloudProviderListResponse {
	return &NullableCloudProviderListResponse{value: val, isSet: true}
}

func (v NullableCloudProviderListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


