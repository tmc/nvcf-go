// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/tmc/nvcf-go/internal/apijson"
	"github.com/tmc/nvcf-go/internal/param"
	"github.com/tmc/nvcf-go/internal/requestconfig"
	"github.com/tmc/nvcf-go/option"
	"github.com/tmc/nvcf-go/shared"
)

// AuthorizationFunctionVersionService contains methods and other services that
// help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthorizationFunctionVersionService] method instead.
type AuthorizationFunctionVersionService struct {
	Options []option.RequestOption
}

// NewAuthorizationFunctionVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAuthorizationFunctionVersionService(opts ...option.RequestOption) (r *AuthorizationFunctionVersionService) {
	r = &AuthorizationFunctionVersionService{}
	r.Options = opts
	return
}

// Adds the specified NVIDIA Cloud Account to the set of authorized accounts that
// can invoke the specified function version. If the specified function version
// does not have any existing inheritable authorized accounts, it results in a
// response with status 404. If the specified account is already in the set of
// existing authorized accounts that are directly associated with the function
// version, it results in a response wit status code 409. If a function is public,
// then Account Admin cannot perform this operation. Access to this functionality
// mandates the inclusion of a bearer token with the 'authorize_clients' scope in
// the HTTP Authorization header
func (r *AuthorizationFunctionVersionService) Add(ctx context.Context, functionID string, functionVersionID string, body AuthorizationFunctionVersionAddParams, opts ...option.RequestOption) (res *shared.AuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/versions/%s/add", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Removes the specified NVIDIA Cloud Account from the set of authorized accounts
// that are directly associated with specified function version. If the specified
// function version does not have any of its own(not inherited) authorized
// accounts, it results in a response with status 404. Also, if the specified
// authorized account is not in the set of existing authorized parties that are
// directly associated with the specified function version, it results in a
// response with status code 400. If the specified function version is public, then
// Account Admin cannot perform this operation. Access to this functionality
// mandates the inclusion of a bearer token with the 'authorize_clients' scope in
// the HTTP Authorization header
func (r *AuthorizationFunctionVersionService) Remove(ctx context.Context, functionID string, functionVersionID string, body AuthorizationFunctionVersionRemoveParams, opts ...option.RequestOption) (res *shared.AuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/versions/%s/remove", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AuthorizationFunctionVersionAddParams struct {
	// Data Transfer Object(DTO) representing an authorized party.
	AuthorizedParty param.Field[AuthorizationFunctionVersionAddParamsAuthorizedParty] `json:"authorizedParty,required"`
}

func (r AuthorizationFunctionVersionAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an authorized party.
type AuthorizationFunctionVersionAddParamsAuthorizedParty struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID param.Field[string] `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID param.Field[string] `json:"clientId"`
}

func (r AuthorizationFunctionVersionAddParamsAuthorizedParty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthorizationFunctionVersionRemoveParams struct {
	// Data Transfer Object(DTO) representing an authorized party.
	AuthorizedParty param.Field[AuthorizationFunctionVersionRemoveParamsAuthorizedParty] `json:"authorizedParty,required"`
}

func (r AuthorizationFunctionVersionRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an authorized party.
type AuthorizationFunctionVersionRemoveParamsAuthorizedParty struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID param.Field[string] `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID param.Field[string] `json:"clientId"`
}

func (r AuthorizationFunctionVersionRemoveParamsAuthorizedParty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
