// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetAppResponseData The definition of `GetAppResponseData` object.
type GetAppResponseData struct {
	// The definition of `GetAppResponseDataAttributes` object.
	Attributes GetAppResponseDataAttributes `json:"attributes"`
	// The `data` `id`.
	Id string `json:"id"`
	// The definition of `GetAppResponseDataType` object.
	Type GetAppResponseDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetAppResponseData instantiates a new GetAppResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetAppResponseData(attributes GetAppResponseDataAttributes, id string, typeVar GetAppResponseDataType) *GetAppResponseData {
	this := GetAppResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewGetAppResponseDataWithDefaults instantiates a new GetAppResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetAppResponseDataWithDefaults() *GetAppResponseData {
	this := GetAppResponseData{}
	var typeVar GetAppResponseDataType = GETAPPRESPONSEDATATYPE_APPDEFINITIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *GetAppResponseData) GetAttributes() GetAppResponseDataAttributes {
	if o == nil {
		var ret GetAppResponseDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetAppResponseData) GetAttributesOk() (*GetAppResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *GetAppResponseData) SetAttributes(v GetAppResponseDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *GetAppResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetAppResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *GetAppResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *GetAppResponseData) GetType() GetAppResponseDataType {
	if o == nil {
		var ret GetAppResponseDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetAppResponseData) GetTypeOk() (*GetAppResponseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GetAppResponseData) SetType(v GetAppResponseDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetAppResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetAppResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *GetAppResponseDataAttributes `json:"attributes"`
		Id         *string                       `json:"id"`
		Type       *GetAppResponseDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
