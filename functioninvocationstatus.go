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
)

// FunctionInvocationStatusService contains methods and other services that help
// with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionInvocationStatusService] method instead.
type FunctionInvocationStatusService struct {
	Options []option.RequestOption
}

// NewFunctionInvocationStatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFunctionInvocationStatusService(opts ...option.RequestOption) (r *FunctionInvocationStatusService) {
	r = &FunctionInvocationStatusService{}
	r.Options = opts
	return
}

// Retrieves the status of an in-progress or pending request using its unique
// invocation request ID. If the result is available, it will be included in the
// response, marking the request as fulfilled. Conversely, if the result is not yet
// available, the request is deemed pending. Access to this endpoint mandates
// inclusion of either a bearer token or an api-key with 'invoke_function' scope in
// the HTTP Authorization header. In-progress responses are returned in order. If
// no in-progress response is received during polling you will receive the most
// recent in-progress response. Only the first 256 unread in-progress messages are
// kept.
func (r *FunctionInvocationStatusService) Get(ctx context.Context, requestID string, query FunctionInvocationStatusGetParams, opts ...option.RequestOption) (res *FunctionInvocationStatusGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if requestID == "" {
		err = errors.New("missing required requestId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/pexec/status/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FunctionInvocationStatusGetResponse struct {
	Char     string                                  `json:"char"`
	Direct   bool                                    `json:"direct"`
	Double   float64                                 `json:"double"`
	Float    float64                                 `json:"float"`
	Int      int64                                   `json:"int"`
	Long     int64                                   `json:"long"`
	ReadOnly bool                                    `json:"readOnly"`
	Short    int64                                   `json:"short"`
	JSON     functionInvocationStatusGetResponseJSON `json:"-"`
}

// functionInvocationStatusGetResponseJSON contains the JSON metadata for the
// struct [FunctionInvocationStatusGetResponse]
type functionInvocationStatusGetResponseJSON struct {
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

func (r *FunctionInvocationStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionInvocationStatusGetResponseJSON) RawJSON() string {
	return r.raw
}

type FunctionInvocationStatusGetParams struct {
	NvcfPollSeconds param.Field[int64] `header:"NVCF-POLL-SECONDS"`
}
