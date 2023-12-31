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
	"net/http"
)



// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface { 
	ConfigsDelete(http.ResponseWriter, *http.Request)
	ConfigsGet(http.ResponseWriter, *http.Request)
	ExclusionsDelete(http.ResponseWriter, *http.Request)
	ExclusionsPost(http.ResponseWriter, *http.Request)
	InfoGet(http.ResponseWriter, *http.Request)
	PoliciesGet(http.ResponseWriter, *http.Request)
	PoliciesPost(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface { 
	ConfigsDelete(context.Context, string, string, string) (ImplResponse, error)
	ConfigsGet(context.Context, string, string, string) (ImplResponse, error)
	ExclusionsDelete(context.Context, string, string) (ImplResponse, error)
	ExclusionsPost(context.Context, ExclusionsPostRequest) (ImplResponse, error)
	InfoGet(context.Context) (ImplResponse, error)
	PoliciesGet(context.Context, string, string) (ImplResponse, error)
	PoliciesPost(context.Context, PoliciesPostRequest) (ImplResponse, error)
}
