/*
 * Matter replacement sample PoC API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PoliciesPostRequest struct {

	ClientId string `json:"clientId,omitempty"`

	Policy string `json:"policy,omitempty"`

	PolicyVersion string `json:"policyVersion,omitempty"`
}

// AssertPoliciesPostRequestRequired checks if the required fields are not zero-ed
func AssertPoliciesPostRequestRequired(obj PoliciesPostRequest) error {
	return nil
}

// AssertRecursePoliciesPostRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PoliciesPostRequest (e.g. [][]PoliciesPostRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePoliciesPostRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPoliciesPostRequest, ok := obj.(PoliciesPostRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPoliciesPostRequestRequired(aPoliciesPostRequest)
	})
}
