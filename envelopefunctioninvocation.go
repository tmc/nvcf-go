// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/tmc/nvcf-go/option"
)

// EnvelopeFunctionInvocationService contains methods and other services that help
// with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvelopeFunctionInvocationService] method instead.
type EnvelopeFunctionInvocationService struct {
	Options   []option.RequestOption
	Functions *EnvelopeFunctionInvocationFunctionService
}

// NewEnvelopeFunctionInvocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEnvelopeFunctionInvocationService(opts ...option.RequestOption) (r *EnvelopeFunctionInvocationService) {
	r = &EnvelopeFunctionInvocationService{}
	r.Options = opts
	r.Functions = NewEnvelopeFunctionInvocationFunctionService(opts...)
	return
}
