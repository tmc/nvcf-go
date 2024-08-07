// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/stainless-sdks/nvcf-go/option"
)

// QueueDetailService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueDetailService] method instead.
type QueueDetailService struct {
	Options []option.RequestOption
	Queues  *QueueDetailQueueService
}

// NewQueueDetailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQueueDetailService(opts ...option.RequestOption) (r *QueueDetailService) {
	r = &QueueDetailService{}
	r.Options = opts
	r.Queues = NewQueueDetailQueueService(opts...)
	return
}
