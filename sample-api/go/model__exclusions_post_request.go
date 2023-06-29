/*
 * Matter replacement sample PoC API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ExclusionsPostRequest struct {

	ClientId string `json:"clientId,omitempty"`

	Policy string `json:"policy,omitempty"`
}

// AssertExclusionsPostRequestRequired checks if the required fields are not zero-ed
func AssertExclusionsPostRequestRequired(obj ExclusionsPostRequest) error {
	return nil
}

// AssertRecurseExclusionsPostRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ExclusionsPostRequest (e.g. [][]ExclusionsPostRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseExclusionsPostRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aExclusionsPostRequest, ok := obj.(ExclusionsPostRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertExclusionsPostRequestRequired(aExclusionsPostRequest)
	})
}
