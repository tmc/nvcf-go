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
)

// FunctionInvocationFunctionVersionService contains methods and other services
// that help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionInvocationFunctionVersionService] method instead.
type FunctionInvocationFunctionVersionService struct {
	Options []option.RequestOption
}

// NewFunctionInvocationFunctionVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFunctionInvocationFunctionVersionService(opts ...option.RequestOption) (r *FunctionInvocationFunctionVersionService) {
	r = &FunctionInvocationFunctionVersionService{}
	r.Options = opts
	return
}

// Invokes the specified function that was successfully deployed. If the version is
// not specified, any active function versions will handle the request. If the
// version is specified in the URI, then the request is exclusively processed by
// the designated version of the function. By default, this endpoint will block for
// 5 seconds. If the request is not fulfilled before the timeout, it's status is
// considered in-progress or pending and the response includes HTTP status code 202
// with an invocation request ID, indicating that the client should commence
// polling for the result using the invocation request ID. Access to this endpoint
// mandates inclusion of either a bearer token or an api-key with 'invoke_function'
// scope in the HTTP Authorization header. Additionally, this endpoint has the
// capability to provide updates on the progress of the request, contingent upon
// the workload's provision of such information. In-progress responses are returned
// in order. If no in-progress response is received during polling you will receive
// the most recent in-progress response. Only the first 256 unread in-progress
// messages are kept.
func (r *FunctionInvocationFunctionVersionService) Invoke(ctx context.Context, functionID string, versionID string, params FunctionInvocationFunctionVersionInvokeParams, opts ...option.RequestOption) (res *FunctionInvocationFunctionVersionInvokeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required versionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/pexec/functions/%s/versions/%s", functionID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type FunctionInvocationFunctionVersionInvokeResponse struct {
	Char     string                                              `json:"char"`
	Direct   bool                                                `json:"direct"`
	Double   float64                                             `json:"double"`
	Float    float64                                             `json:"float"`
	Int      int64                                               `json:"int"`
	Long     int64                                               `json:"long"`
	ReadOnly bool                                                `json:"readOnly"`
	Short    int64                                               `json:"short"`
	JSON     functionInvocationFunctionVersionInvokeResponseJSON `json:"-"`
}

// functionInvocationFunctionVersionInvokeResponseJSON contains the JSON metadata
// for the struct [FunctionInvocationFunctionVersionInvokeResponse]
type functionInvocationFunctionVersionInvokeResponseJSON struct {
	Char        apijson.Field
	Direct      apijson.Field
	Double      apijson.Field
	Float       apijson.Field
	Int         apijson.Field
	Long        apijson.Field
	ReadOnly    apijson.Field
	Short       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionInvocationFunctionVersionInvokeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionInvocationFunctionVersionInvokeResponseJSON) RawJSON() string {
	return r.raw
}

type FunctionInvocationFunctionVersionInvokeParams struct {
	Body                     interface{}           `json:"body,required"`
	NvcfInputAssetReferences param.Field[[]string] `header:"NVCF-INPUT-ASSET-REFERENCES"`
	NvcfPollSeconds          param.Field[int64]    `header:"NVCF-POLL-SECONDS"`
}

func (r FunctionInvocationFunctionVersionInvokeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
