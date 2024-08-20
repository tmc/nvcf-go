// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/nvcf-go/internal/apijson"
	"github.com/stainless-sdks/nvcf-go/internal/requestconfig"
	"github.com/stainless-sdks/nvcf-go/option"
)

// QueueDetailQueuePositionService contains methods and other services that help
// with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueDetailQueuePositionService] method instead.
type QueueDetailQueuePositionService struct {
	Options []option.RequestOption
}

// NewQueueDetailQueuePositionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewQueueDetailQueuePositionService(opts ...option.RequestOption) (r *QueueDetailQueuePositionService) {
	r = &QueueDetailQueuePositionService{}
	r.Options = opts
	return
}

// Using the specified function invocation request id, returns the estimated
// position of the corresponding message up to 1000 in the queue. Requires a bearer
// token or an api-key with 'queue_details' scope in the HTTP Authorization header.
func (r *QueueDetailQueuePositionService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *QueueDetailQueuePositionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if requestID == "" {
		err = errors.New("missing required requestId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/queues/%s/position", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Request position in queue for invocation request
type QueueDetailQueuePositionGetResponse struct {
	// Function id
	FunctionID string `json:"functionId,required" format:"uuid"`
	// Function version id
	FunctionVersionID string `json:"functionVersionId,required" format:"uuid"`
	// Position of request in queue
	PositionInQueue int64                                   `json:"positionInQueue"`
	JSON            queueDetailQueuePositionGetResponseJSON `json:"-"`
}

// queueDetailQueuePositionGetResponseJSON contains the JSON metadata for the
// struct [QueueDetailQueuePositionGetResponse]
type queueDetailQueuePositionGetResponseJSON struct {
	FunctionID        apijson.Field
	FunctionVersionID apijson.Field
	PositionInQueue   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *QueueDetailQueuePositionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueDetailQueuePositionGetResponseJSON) RawJSON() string {
	return r.raw
}
