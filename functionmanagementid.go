// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/tmc/nvcf-go/internal/apijson"
	"github.com/tmc/nvcf-go/option"
)

// FunctionManagementIDService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionManagementIDService] method instead.
type FunctionManagementIDService struct {
	Options []option.RequestOption
}

// NewFunctionManagementIDService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFunctionManagementIDService(opts ...option.RequestOption) (r *FunctionManagementIDService) {
	r = &FunctionManagementIDService{}
	r.Options = opts
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
