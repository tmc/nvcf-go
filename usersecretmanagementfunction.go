// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/stainless-sdks/nvcf-go/option"
)

// UserSecretManagementFunctionService contains methods and other services that
// help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSecretManagementFunctionService] method instead.
type UserSecretManagementFunctionService struct {
	Options  []option.RequestOption
	Versions *UserSecretManagementFunctionVersionService
}

// NewUserSecretManagementFunctionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserSecretManagementFunctionService(opts ...option.RequestOption) (r *UserSecretManagementFunctionService) {
	r = &UserSecretManagementFunctionService{}
	r.Options = opts
	r.Versions = NewUserSecretManagementFunctionVersionService(opts...)
	return
}
