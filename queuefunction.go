// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/tmc/nvcf-go/internal/requestconfig"
	"github.com/tmc/nvcf-go/option"
	"github.com/tmc/nvcf-go/shared"
)

// QueueFunctionService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueFunctionService] method instead.
type QueueFunctionService struct {
	Options  []option.RequestOption
	Versions *QueueFunctionVersionService
}

// NewQueueFunctionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQueueFunctionService(opts ...option.RequestOption) (r *QueueFunctionService) {
	r = &QueueFunctionService{}
	r.Options = opts
	r.Versions = NewQueueFunctionVersionService(opts...)
	return
}

// Provides details of all the queues associated with the specified function. If a
// function has multiple versions and they are all deployed, then the response
// includes details of all the queues. If the specified function is public, then
// Account Admin cannot perform this operation. Requires a bearer token or an
// api-key with 'queue_details' scope in the HTTP Authorization header.
func (r *QueueFunctionService) List(ctx context.Context, functionID string, opts ...option.RequestOption) (res *shared.GetQueuesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/queues/functions/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
