// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"net/http"
	"net/url"

	"github.com/tmc/nvcf-go/internal/apijson"
	"github.com/tmc/nvcf-go/internal/apiquery"
	"github.com/tmc/nvcf-go/internal/param"
	"github.com/tmc/nvcf-go/internal/requestconfig"
	"github.com/tmc/nvcf-go/option"
)

// FunctionIDService contains methods and other services that help with interacting
// with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionIDService] method instead.
type FunctionIDService struct {
	Options []option.RequestOption
}

// NewFunctionIDService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionIDService(opts ...option.RequestOption) (r *FunctionIDService) {
	r = &FunctionIDService{}
	r.Options = opts
	return
}

// Lists ids of all the functions in the authenticated NVIDIA Cloud Account.
// Requires either a bearer token or an api-key with 'list_functions' or
// 'list_functions_details' scopes in the HTTP Authorization header.
func (r *FunctionIDService) List(ctx context.Context, query FunctionIDListParams, opts ...option.RequestOption) (res *ListFunctionIDsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/functions/ids"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response body containing list of function ids in an account
type ListFunctionIDsResponse struct {
	// List of function ids
	FunctionIDs []string                    `json:"functionIds,required" format:"uuid"`
	JSON        listFunctionIDsResponseJSON `json:"-"`
}

// listFunctionIDsResponseJSON contains the JSON metadata for the struct
// [ListFunctionIDsResponse]
type listFunctionIDsResponseJSON struct {
	FunctionIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListFunctionIDsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listFunctionIDsResponseJSON) RawJSON() string {
	return r.raw
}

type FunctionIDListParams struct {
	// Query param 'visibility' indicates the kind of functions to be included in the
	// response.
	Visibility param.Field[[]FunctionIDListParamsVisibility] `query:"visibility"`
}

// URLQuery serializes [FunctionIDListParams]'s query parameters as `url.Values`.
func (r FunctionIDListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FunctionIDListParamsVisibility string

const (
	FunctionIDListParamsVisibilityAuthorized FunctionIDListParamsVisibility = "authorized"
	FunctionIDListParamsVisibilityPrivate    FunctionIDListParamsVisibility = "private"
	FunctionIDListParamsVisibilityPublic     FunctionIDListParamsVisibility = "public"
)

func (r FunctionIDListParamsVisibility) IsKnown() bool {
	switch r {
	case FunctionIDListParamsVisibilityAuthorized, FunctionIDListParamsVisibilityPrivate, FunctionIDListParamsVisibilityPublic:
		return true
	}
	return false
}
