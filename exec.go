// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/stainless-sdks/nvcf-go/option"
)

// ExecService contains methods and other services that help with interacting with
// the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExecService] method instead.
type ExecService struct {
	Options []option.RequestOption
	Status  *ExecStatusService
}

// NewExecService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExecService(opts ...option.RequestOption) (r *ExecService) {
	r = &ExecService{}
	r.Options = opts
	r.Status = NewExecStatusService(opts...)
	return
}
