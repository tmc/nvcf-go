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

// AuthorizedAccountFunctionService contains methods and other services that help
// with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthorizedAccountFunctionService] method instead.
type AuthorizedAccountFunctionService struct {
	Options  []option.RequestOption
	Versions *AuthorizedAccountFunctionVersionService
}

// NewAuthorizedAccountFunctionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAuthorizedAccountFunctionService(opts ...option.RequestOption) (r *AuthorizedAccountFunctionService) {
	r = &AuthorizedAccountFunctionService{}
	r.Options = opts
	r.Versions = NewAuthorizedAccountFunctionVersionService(opts...)
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
func (r *AuthorizedAccountFunctionService) Add(ctx context.Context, functionID string, body AuthorizedAccountFunctionAddParams, opts ...option.RequestOption) (res *shared.AuthorizedPartiesResponse, err error) {
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
func (r *AuthorizedAccountFunctionService) Remove(ctx context.Context, functionID string, body AuthorizedAccountFunctionRemoveParams, opts ...option.RequestOption) (res *shared.AuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/remove", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AuthorizedAccountFunctionAddParams struct {
	// Data Transfer Object(DTO) representing an authorized party.
	AuthorizedParty param.Field[AuthorizedAccountFunctionAddParamsAuthorizedParty] `json:"authorizedParty,required"`
}

func (r AuthorizedAccountFunctionAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an authorized party.
type AuthorizedAccountFunctionAddParamsAuthorizedParty struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID param.Field[string] `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID param.Field[string] `json:"clientId"`
}

func (r AuthorizedAccountFunctionAddParamsAuthorizedParty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthorizedAccountFunctionRemoveParams struct {
	// Data Transfer Object(DTO) representing an authorized party.
	AuthorizedParty param.Field[AuthorizedAccountFunctionRemoveParamsAuthorizedParty] `json:"authorizedParty,required"`
}

func (r AuthorizedAccountFunctionRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an authorized party.
type AuthorizedAccountFunctionRemoveParamsAuthorizedParty struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID param.Field[string] `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID param.Field[string] `json:"clientId"`
}

func (r AuthorizedAccountFunctionRemoveParamsAuthorizedParty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
