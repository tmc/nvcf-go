// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/tmc/nvcf-go/option"
)

// UserSecretManagementService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSecretManagementService] method instead.
type UserSecretManagementService struct {
	Options   []option.RequestOption
	Functions *UserSecretManagementFunctionService
}

// NewUserSecretManagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserSecretManagementService(opts ...option.RequestOption) (r *UserSecretManagementService) {
	r = &UserSecretManagementService{}
	r.Options = opts
	r.Functions = NewUserSecretManagementFunctionService(opts...)
	return
}
