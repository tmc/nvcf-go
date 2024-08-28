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

// AuthorizedAccountFunctionVersionService contains methods and other services that
// help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthorizedAccountFunctionVersionService] method instead.
type AuthorizedAccountFunctionVersionService struct {
	Options []option.RequestOption
}

// NewAuthorizedAccountFunctionVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAuthorizedAccountFunctionVersionService(opts ...option.RequestOption) (r *AuthorizedAccountFunctionVersionService) {
	r = &AuthorizedAccountFunctionVersionService{}
	r.Options = opts
	return
}

// Gets NVIDIA Cloud Account IDs that are authorized to invoke specified function
// version. Response includes authorized accounts that were added specifically to
// the function version and the inherited authorized accounts that were added at
// the function level. Access to this functionality mandates the inclusion of a
// bearer token with the 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizedAccountFunctionVersionService) Get(ctx context.Context, functionID string, functionVersionID string, opts ...option.RequestOption) (res *shared.AuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes all the authorized accounts that are directly associated with the
// specified function version. Authorized parties that are inherited by the
// function version are not deleted. If the specified function version is public,
// then Account Admin cannot perform this operation. Access to this functionality
// mandates the inclusion of a bearer token with the 'authorize_clients' scope in
// the HTTP Authorization header
func (r *AuthorizedAccountFunctionVersionService) Delete(ctx context.Context, functionID string, functionVersionID string, opts ...option.RequestOption) (res *shared.AuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Authorizes additional NVIDIA Cloud Accounts to invoke a specific function
// version. By default, a function belongs to the NVIDIA Cloud Account that created
// it, and the credentials used for function invocation must reference the same
// NVIDIA Cloud Account. Upon invocation of this endpoint, any existing authorized
// accounts will be overwritten by the newly specified authorized accounts. Access
// to this functionality mandates the inclusion of a bearer token with the
// 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizedAccountFunctionVersionService) Authorize(ctx context.Context, functionID string, functionVersionID string, body AuthorizedAccountFunctionVersionAuthorizeParams, opts ...option.RequestOption) (res *shared.AuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AuthorizedAccountFunctionVersionAuthorizeParams struct {
	// Parties authorized to invoke function
	AuthorizedParties param.Field[[]AuthorizedAccountFunctionVersionAuthorizeParamsAuthorizedParty] `json:"authorizedParties,required"`
}

func (r AuthorizedAccountFunctionVersionAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an authorized party.
type AuthorizedAccountFunctionVersionAuthorizeParamsAuthorizedParty struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID param.Field[string] `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID param.Field[string] `json:"clientId"`
}

func (r AuthorizedAccountFunctionVersionAuthorizeParamsAuthorizedParty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
