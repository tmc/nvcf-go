// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/nvcf-go/internal/requestconfig"
	"github.com/stainless-sdks/nvcf-go/option"
	"github.com/stainless-sdks/nvcf-go/shared"
)

// ExecStatusService contains methods and other services that help with interacting
// with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExecStatusService] method instead.
type ExecStatusService struct {
	Options []option.RequestOption
}

// NewExecStatusService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExecStatusService(opts ...option.RequestOption) (r *ExecStatusService) {
	r = &ExecStatusService{}
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
func (r *ExecStatusService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *shared.InvokeFunctionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if requestID == "" {
		err = errors.New("missing required requestId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/exec/status/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
