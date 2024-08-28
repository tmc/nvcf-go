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

// QueueFunctionVersionService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueFunctionVersionService] method instead.
type QueueFunctionVersionService struct {
	Options []option.RequestOption
}

// NewQueueFunctionVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewQueueFunctionVersionService(opts ...option.RequestOption) (r *QueueFunctionVersionService) {
	r = &QueueFunctionVersionService{}
	r.Options = opts
	return
}

// Provides details of all the queues associated with the specified function. If a
// function has multiple versions and they are all deployed, then the response
// includes details of all the queues. If the specified function is public, then
// Account Admin cannot perform this operation. Requires a bearer token or an
// api-key with 'queue_details' scope in the HTTP Authorization header.
func (r *QueueFunctionVersionService) List(ctx context.Context, functionID string, versionID string, opts ...option.RequestOption) (res *shared.GetQueuesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required versionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/queues/functions/%s/versions/%s", functionID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
