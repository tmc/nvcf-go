// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/stainless-sdks/nvcf-go/option"
)

// FunctionDeploymentFunctionService contains methods and other services that help
// with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionDeploymentFunctionService] method instead.
type FunctionDeploymentFunctionService struct {
	Options  []option.RequestOption
	Versions *FunctionDeploymentFunctionVersionService
}

// NewFunctionDeploymentFunctionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFunctionDeploymentFunctionService(opts ...option.RequestOption) (r *FunctionDeploymentFunctionService) {
	r = &FunctionDeploymentFunctionService{}
	r.Options = opts
	r.Versions = NewFunctionDeploymentFunctionVersionService(opts...)
	return
}
