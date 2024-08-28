// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/nvcf-go/internal/apijson"
	"github.com/stainless-sdks/nvcf-go/internal/param"
	"github.com/stainless-sdks/nvcf-go/internal/requestconfig"
	"github.com/stainless-sdks/nvcf-go/option"
	"github.com/stainless-sdks/nvcf-go/shared"
)

// AuthorizationFunctionService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthorizationFunctionService] method instead.
type AuthorizationFunctionService struct {
	Options  []option.RequestOption
	Versions *AuthorizationFunctionVersionService
}

// NewAuthorizationFunctionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAuthorizationFunctionService(opts ...option.RequestOption) (r *AuthorizationFunctionService) {
	r = &AuthorizationFunctionService{}
	r.Options = opts
	r.Versions = NewAuthorizationFunctionVersionService(opts...)
	return
}

// Adds the specified NVIDIA Cloud Account to the set of authorized accounts that
// are can invoke all the versions of the specified function. If the specified
// function does not have any existing inheritable authorized accounts, it results
// in a response with status 404. If the specified account is already in the set of
// existing inheritable authorized accounts, it results in a response with status
// code 409. If a function is public, then Account Admin cannot perform this
// operation. Access to this functionality mandates the inclusion of a bearer token
// with the 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizationFunctionService) Add(ctx context.Context, functionID string, body AuthorizationFunctionAddParams, opts ...option.RequestOption) (res *shared.AuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/add", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Removes the specified NVIDIA Cloud Account from the set of authorized accounts
// that can invoke all the versions of the specified function. If the specified
// function does not have any existing inheritable authorized parties, it results
// in a response with status 404. Also, if the specified account is not in the
// existing set of inheritable authorized accounts, it results in a response with
// status 400. If the specified function is public, then Account Admin cannot
// perform this operation. Access to this functionality mandates the inclusion of a
// bearer token with the 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizationFunctionService) Remove(ctx context.Context, functionID string, body AuthorizationFunctionRemoveParams, opts ...option.RequestOption) (res *shared.AuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/remove", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AuthorizationFunctionAddParams struct {
	// Data Transfer Object(DTO) representing an authorized party.
	AuthorizedParty param.Field[AuthorizationFunctionAddParamsAuthorizedParty] `json:"authorizedParty,required"`
}

func (r AuthorizationFunctionAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an authorized party.
type AuthorizationFunctionAddParamsAuthorizedParty struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID param.Field[string] `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID param.Field[string] `json:"clientId"`
}

func (r AuthorizationFunctionAddParamsAuthorizedParty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthorizationFunctionRemoveParams struct {
	// Data Transfer Object(DTO) representing an authorized party.
	AuthorizedParty param.Field[AuthorizationFunctionRemoveParamsAuthorizedParty] `json:"authorizedParty,required"`
}

func (r AuthorizationFunctionRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an authorized party.
type AuthorizationFunctionRemoveParamsAuthorizedParty struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID param.Field[string] `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID param.Field[string] `json:"clientId"`
}

func (r AuthorizationFunctionRemoveParamsAuthorizedParty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
