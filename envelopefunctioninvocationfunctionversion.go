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

// EnvelopeFunctionInvocationFunctionVersionService contains methods and other
// services that help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvelopeFunctionInvocationFunctionVersionService] method instead.
type EnvelopeFunctionInvocationFunctionVersionService struct {
	Options []option.RequestOption
}

// NewEnvelopeFunctionInvocationFunctionVersionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewEnvelopeFunctionInvocationFunctionVersionService(opts ...option.RequestOption) (r *EnvelopeFunctionInvocationFunctionVersionService) {
	r = &EnvelopeFunctionInvocationFunctionVersionService{}
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
func (r *EnvelopeFunctionInvocationFunctionVersionService) Invoke(ctx context.Context, functionID string, versionID string, body EnvelopeFunctionInvocationFunctionVersionInvokeParams, opts ...option.RequestOption) (res *shared.InvokeFunctionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required versionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/exec/functions/%s/versions/%s", functionID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EnvelopeFunctionInvocationFunctionVersionInvokeParams struct {
	RequestBody param.Field[interface{}] `json:"requestBody,required"`
	// Data Transfer Object(DTO) representing header/address for Cloud Functions
	// processing.
	RequestHeader param.Field[EnvelopeFunctionInvocationFunctionVersionInvokeParamsRequestHeader] `json:"requestHeader"`
}

func (r EnvelopeFunctionInvocationFunctionVersionInvokeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing header/address for Cloud Functions
// processing.
type EnvelopeFunctionInvocationFunctionVersionInvokeParamsRequestHeader struct {
	// List of UUIDs corresponding to the uploaded assets to be used as input for
	// executing the task.
	InputAssetReferences param.Field[[]string] `json:"inputAssetReferences" format:"uuid"`
	// Metadata used for billing/metering purposes.
	MeteringData param.Field[[]EnvelopeFunctionInvocationFunctionVersionInvokeParamsRequestHeaderMeteringData] `json:"meteringData"`
	// Polling timeout duration.
	PollDurationSeconds param.Field[int64] `json:"pollDurationSeconds"`
}

func (r EnvelopeFunctionInvocationFunctionVersionInvokeParamsRequestHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing a billing/metering data entry
type EnvelopeFunctionInvocationFunctionVersionInvokeParamsRequestHeaderMeteringData struct {
	// Metering/Billing key
	Key param.Field[string] `json:"key,required"`
	// Metering/Billing value
	Value param.Field[string] `json:"value,required"`
}

func (r EnvelopeFunctionInvocationFunctionVersionInvokeParamsRequestHeaderMeteringData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
