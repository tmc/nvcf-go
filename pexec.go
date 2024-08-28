// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/stainless-sdks/nvcf-go/option"
)

// PexecService contains methods and other services that help with interacting with
// the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPexecService] method instead.
type PexecService struct {
	Options []option.RequestOption
	Status  *PexecStatusService
}

// NewPexecService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPexecService(opts ...option.RequestOption) (r *PexecService) {
	r = &PexecService{}
	r.Options = opts
	r.Status = NewPexecStatusService(opts...)
	return
}
