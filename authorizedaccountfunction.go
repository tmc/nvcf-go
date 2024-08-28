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

// Lists NVIDIA Cloud Account IDs that are authorized to invoke any version of the
// specified function. The response includes an array showing authorized accounts
// for each version. Individual versions of a function can have their own
// authorized accounts. So, each object in the array can have different authorized
// accounts listed. Access to this functionality mandates the inclusion of a bearer
// token with the 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizedAccountFunctionService) Get(ctx context.Context, functionID string, opts ...option.RequestOption) (res *ListAuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes all the extra NVIDIA Cloud Accounts that were authorized to invoke the
// function and all its versions. If a function version has its own set of
// authorized accounts, those are not deleted. If the specified function is public,
// then Account Admin cannot perform this operation. Access to this functionality
// mandates the inclusion of a bearer token with the 'authorize_clients' scope in
// the HTTP Authorization header
func (r *AuthorizedAccountFunctionService) Delete(ctx context.Context, functionID string, opts ...option.RequestOption) (res *shared.AuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Authorizes additional NVIDIA Cloud Accounts to invoke any version of the
// specified function. By default, a function belongs to the NVIDIA Cloud Account
// that created it, and the credentials used for function invocation must reference
// the same NVIDIA Cloud Account. Upon invocation of this endpoint, any existing
// authorized accounts will be overwritten by the newly specified authorized
// accounts. Access to this functionality mandates the inclusion of a bearer token
// with the 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizedAccountFunctionService) Authorize(ctx context.Context, functionID string, body AuthorizedAccountFunctionAuthorizeParams, opts ...option.RequestOption) (res *shared.AuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Parties authorized to invoke function
type ListAuthorizedPartiesResponse struct {
	// Functions with authorized parties and other details
	Functions []ListAuthorizedPartiesResponseFunction `json:"functions,required"`
	JSON      listAuthorizedPartiesResponseJSON       `json:"-"`
}

// listAuthorizedPartiesResponseJSON contains the JSON metadata for the struct
// [ListAuthorizedPartiesResponse]
type listAuthorizedPartiesResponseJSON struct {
	Functions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListAuthorizedPartiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listAuthorizedPartiesResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a function with authorized accounts
type ListAuthorizedPartiesResponseFunction struct {
	// Function id
	ID string `json:"id,required" format:"uuid"`
	// NVIDIA Cloud Account Id
	NcaID string `json:"ncaId,required"`
	// Authorized parties allowed to invoke the function
	AuthorizedParties []ListAuthorizedPartiesResponseFunctionsAuthorizedParty `json:"authorizedParties"`
	// Function version id
	VersionID string                                    `json:"versionId" format:"uuid"`
	JSON      listAuthorizedPartiesResponseFunctionJSON `json:"-"`
}

// listAuthorizedPartiesResponseFunctionJSON contains the JSON metadata for the
// struct [ListAuthorizedPartiesResponseFunction]
type listAuthorizedPartiesResponseFunctionJSON struct {
	ID                apijson.Field
	NcaID             apijson.Field
	AuthorizedParties apijson.Field
	VersionID         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ListAuthorizedPartiesResponseFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listAuthorizedPartiesResponseFunctionJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an authorized party.
type ListAuthorizedPartiesResponseFunctionsAuthorizedParty struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID string `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID string                                                    `json:"clientId"`
	JSON     listAuthorizedPartiesResponseFunctionsAuthorizedPartyJSON `json:"-"`
}

// listAuthorizedPartiesResponseFunctionsAuthorizedPartyJSON contains the JSON
// metadata for the struct [ListAuthorizedPartiesResponseFunctionsAuthorizedParty]
type listAuthorizedPartiesResponseFunctionsAuthorizedPartyJSON struct {
	NcaID       apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListAuthorizedPartiesResponseFunctionsAuthorizedParty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listAuthorizedPartiesResponseFunctionsAuthorizedPartyJSON) RawJSON() string {
	return r.raw
}

type AuthorizedAccountFunctionAuthorizeParams struct {
	// Parties authorized to invoke function
	AuthorizedParties param.Field[[]AuthorizedAccountFunctionAuthorizeParamsAuthorizedParty] `json:"authorizedParties,required"`
}

func (r AuthorizedAccountFunctionAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an authorized party.
type AuthorizedAccountFunctionAuthorizeParamsAuthorizedParty struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID param.Field[string] `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID param.Field[string] `json:"clientId"`
}

func (r AuthorizedAccountFunctionAuthorizeParamsAuthorizedParty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
