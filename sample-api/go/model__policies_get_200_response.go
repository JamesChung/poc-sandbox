/*
 * Matter replacement sample PoC API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PoliciesGet200Response struct {

	Policies []Policy `json:"policies,omitempty"`

	Metadata Metadata `json:"metadata,omitempty"`
}

// AssertPoliciesGet200ResponseRequired checks if the required fields are not zero-ed
func AssertPoliciesGet200ResponseRequired(obj PoliciesGet200Response) error {
	for _, el := range obj.Policies {
		if err := AssertPolicyRequired(el); err != nil {
			return err
		}
	}
	if err := AssertMetadataRequired(obj.Metadata); err != nil {
		return err
	}
	return nil
}

// AssertRecursePoliciesGet200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PoliciesGet200Response (e.g. [][]PoliciesGet200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePoliciesGet200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPoliciesGet200Response, ok := obj.(PoliciesGet200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPoliciesGet200ResponseRequired(aPoliciesGet200Response)
	})
}
