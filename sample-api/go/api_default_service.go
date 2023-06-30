/*
 * Matter replacement sample PoC API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// ConfigsDelete - Delete configurations
func (s *DefaultApiService) ConfigsDelete(ctx context.Context, policy string, clientId string, policyVersion string) (ImplResponse, error) {
	// TODO - update ConfigsDelete with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Metadata{}) or use other options such as http.Ok ...
	//return Response(200, Metadata{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ConfigsDelete method not implemented")
}

// ConfigsGet - Get configurations
func (s *DefaultApiService) ConfigsGet(ctx context.Context, clientId string, policy string, policyVersion string) (ImplResponse, error) {
	// TODO - update ConfigsGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ConfigsGet200Response{}) or use other options such as http.Ok ...
	//return Response(200, ConfigsGet200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ConfigsGet method not implemented")
}

// ExclusionsDelete - Delete exclusions
func (s *DefaultApiService) ExclusionsDelete(ctx context.Context, policy string, clientId string) (ImplResponse, error) {
	for k := range ExcludedPolicies {
		if ExcludedPolicies[k] {
			delete(ExcludedPolicies, k)
		}
	}

	return Response(http.StatusOK, Metadata{
		TransactionId: "01H448Y0XEBQ73YG7WXXVFWSNX",
	}), nil
}

// ExclusionsPost - Exempt a policy for a given client
func (s *DefaultApiService) ExclusionsPost(ctx context.Context, exclusionsPostRequest ExclusionsPostRequest) (ImplResponse, error) {
	ExcludedPolicies[exclusionsPostRequest.Policy] = true

	return Response(http.StatusCreated, Metadata{
		TransactionId: "01H448Y0XEBQ73YG7WXXVFWSNX",
	}), nil
}

// InfoGet - Get info
func (s *DefaultApiService) InfoGet(ctx context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, Info{
		Version: "1.0.0",
		Metadata: Metadata{
			TransactionId: "01H448Y0XEBQ73YG7WXXVFWSNX",
		},
	}), nil
}

var ExcludedPolicies = make(map[string]bool)

var AllPolicies = []Policy{
	{
		Name:     "CKV_AWS_93",
		Version:  "1.0.0",
		Location: "s3://example-bucket/CKV_AWS_93.json",
	},
	{
		Name:     "CKV_AWS_20",
		Version:  "1.0.0",
		Location: "s3://example-bucket/CKV_AWS_20.json",
	},
	{
		Name:     "CKV_AWS_19",
		Version:  "1.0.0",
		Location: "s3://example-bucket/CKV_AWS_19.json",
	},
	{
		Name:     "CKV_AWS_57",
		Version:  "1.0.0",
		Location: "s3://example-bucket/CKV_AWS_57.json",
	},
	{
		Name:     "CKV_AWS_145",
		Version:  "1.0.0",
		Location: "s3://example-bucket/CKV_AWS_145.json",
	},
	{
		Name:     "CKV2_AWS_61",
		Version:  "1.0.0",
		Location: "s3://example-bucket/CKV2_AWS_61.json",
	},
	{
		Name:     "CKV2_AWS_6",
		Version:  "1.0.0",
		Location: "s3://example-bucket/CKV2_AWS_6.json",
	},
	{
		Name:     "CKV_AWS_21",
		Version:  "1.0.0",
		Location: "s3://example-bucket/CKV_AWS_21.json",
	},
	{
		Name:     "CKV_AWS_144",
		Version:  "1.0.0",
		Location: "s3://example-bucket/CKV_AWS_144.json",
	},
	{
		Name:     "CKV_AWS_18",
		Version:  "1.0.0",
		Location: "s3://example-bucket/CKV_AWS_18.json",
	},
	{
		Name:     "CKV2_AWS_62",
		Version:  "1.0.0",
		Location: "s3://example-bucket/CKV2_AWS_62.json",
	}}

func PoliciesDelta() []Policy {
	policies := []Policy{}
	for _, v := range AllPolicies {
		if !ExcludedPolicies[v.Name] {
			policies = append(policies, v)
		}
	}
	return policies
}

// PoliciesGet - Get policies
func (s *DefaultApiService) PoliciesGet(ctx context.Context, clientId string, continuationToken string) (ImplResponse, error) {
	return Response(http.StatusOK, PoliciesGet200Response{
		Policies: PoliciesDelta(),
		Metadata: Metadata{
			TransactionId: "01H448Y0XEBQ73YG7WXXVFWSNX",
		},
	}), nil
}

// PoliciesPost - Add a policy
func (s *DefaultApiService) PoliciesPost(ctx context.Context, policiesPostRequest PoliciesPostRequest) (ImplResponse, error) {
	AllPolicies = append(AllPolicies, Policy{
		Name:     policiesPostRequest.Policy,
		Version:  policiesPostRequest.PolicyVersion,
		Location: fmt.Sprintf("s3://example-bucket/%s.json", strings.ToLower(policiesPostRequest.Policy)),
	})

	return Response(http.StatusCreated, Metadata{
		TransactionId: "01H44AN4JEY1G1WNMXQ0D14QZ4",
	}), nil
}
