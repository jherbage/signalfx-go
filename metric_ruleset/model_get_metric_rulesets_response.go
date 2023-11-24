/*
Metric Ruleset API

Metric ruleset API 

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metric_ruleset

import (
	"encoding/json"
)

// GetMetricRulesetsResponse Retrieving metric rulesets by query endpoint response body 
type GetMetricRulesetsResponse struct {
	// Number of matched metric rulesets 
	Count *int32 `json:"count,omitempty"`
	// List of metric rulesets 
	Results []MetricRuleset `json:"results,omitempty"`
}

// NewGetMetricRulesetsResponse instantiates a new GetMetricRulesetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMetricRulesetsResponse() *GetMetricRulesetsResponse {
	this := GetMetricRulesetsResponse{}
	return &this
}

// NewGetMetricRulesetsResponseWithDefaults instantiates a new GetMetricRulesetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMetricRulesetsResponseWithDefaults() *GetMetricRulesetsResponse {
	this := GetMetricRulesetsResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetMetricRulesetsResponse) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetricRulesetsResponse) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetMetricRulesetsResponse) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetMetricRulesetsResponse) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetMetricRulesetsResponse) GetResults() []MetricRuleset {
	if o == nil || isNil(o.Results) {
		var ret []MetricRuleset
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetricRulesetsResponse) GetResultsOk() ([]MetricRuleset, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetMetricRulesetsResponse) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MetricRuleset and assigns it to the Results field.
func (o *GetMetricRulesetsResponse) SetResults(v []MetricRuleset) {
	o.Results = v
}

func (o GetMetricRulesetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableGetMetricRulesetsResponse struct {
	value *GetMetricRulesetsResponse
	isSet bool
}

func (v NullableGetMetricRulesetsResponse) Get() *GetMetricRulesetsResponse {
	return v.value
}

func (v *NullableGetMetricRulesetsResponse) Set(val *GetMetricRulesetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMetricRulesetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMetricRulesetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMetricRulesetsResponse(val *GetMetricRulesetsResponse) *NullableGetMetricRulesetsResponse {
	return &NullableGetMetricRulesetsResponse{value: val, isSet: true}
}

func (v NullableGetMetricRulesetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMetricRulesetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


