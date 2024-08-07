// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/stainless-sdks/nvcf-go/option"
)

// QueueDetailQueueService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueDetailQueueService] method instead.
type QueueDetailQueueService struct {
	Options   []option.RequestOption
	Functions *QueueDetailQueueFunctionService
	Position  *QueueDetailQueuePositionService
}

// NewQueueDetailQueueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewQueueDetailQueueService(opts ...option.RequestOption) (r *QueueDetailQueueService) {
	r = &QueueDetailQueueService{}
	r.Options = opts
	r.Functions = NewQueueDetailQueueFunctionService(opts...)
	r.Position = NewQueueDetailQueuePositionService(opts...)
	return
}
