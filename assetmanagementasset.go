// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/nvcf-go/internal/apijson"
	"github.com/stainless-sdks/nvcf-go/internal/param"
	"github.com/stainless-sdks/nvcf-go/internal/requestconfig"
	"github.com/stainless-sdks/nvcf-go/option"
)

// AssetManagementAssetService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetManagementAssetService] method instead.
type AssetManagementAssetService struct {
	Options []option.RequestOption
}

// NewAssetManagementAssetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssetManagementAssetService(opts ...option.RequestOption) (r *AssetManagementAssetService) {
	r = &AssetManagementAssetService{}
	r.Options = opts
	return
}

// Creates a unique id representing an asset and a pre-signed URL to upload the
// asset artifact to AWS S3 bucket for the NVIDIA Cloud Account. Requires either a
// bearer token or an api-key with 'invoke_function' scope in the HTTP
// Authorization header.
func (r *AssetManagementAssetService) New(ctx context.Context, body AssetManagementAssetNewParams, opts ...option.RequestOption) (res *CreateAssetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/assets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response body containing asset-id and the corresponding pre-signed URL to upload
// an asset of specified content-type to AWS S3 bucket.
type CreateAssetResponse struct {
	// Unique id of the asset to be uploaded to AWS S3 bucket
	AssetID string `json:"assetId" format:"uuid"`
	// Content type of the asset such image/png, image/jpeg, etc.
	ContentType string `json:"contentType"`
	// Asset description to be used when uploading the asset
	Description string `json:"description"`
	// Pre-signed upload URL to upload asset
	UploadURL string                  `json:"uploadUrl" format:"url"`
	JSON      createAssetResponseJSON `json:"-"`
}

// createAssetResponseJSON contains the JSON metadata for the struct
// [CreateAssetResponse]
type createAssetResponseJSON struct {
	AssetID     apijson.Field
	ContentType apijson.Field
	Description apijson.Field
	UploadURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateAssetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createAssetResponseJSON) RawJSON() string {
	return r.raw
}

type AssetManagementAssetNewParams struct {
	// Content type of the asset such image/png, image/jpeg, etc.
	ContentType param.Field[string] `json:"contentType,required"`
	// Asset description
	Description param.Field[string] `json:"description,required"`
}

func (r AssetManagementAssetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
