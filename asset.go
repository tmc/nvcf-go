// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/nvcf-go/internal/apijson"
	"github.com/stainless-sdks/nvcf-go/internal/param"
	"github.com/stainless-sdks/nvcf-go/internal/requestconfig"
	"github.com/stainless-sdks/nvcf-go/option"
)

// AssetService contains methods and other services that help with interacting with
// the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetService] method instead.
type AssetService struct {
	Options []option.RequestOption
}

// NewAssetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAssetService(opts ...option.RequestOption) (r *AssetService) {
	r = &AssetService{}
	r.Options = opts
	return
}

// Creates a unique id representing an asset and a pre-signed URL to upload the
// asset artifact to AWS S3 bucket for the NVIDIA Cloud Account. Requires either a
// bearer token or an api-key with 'invoke_function' scope in the HTTP
// Authorization header.
func (r *AssetService) New(ctx context.Context, body AssetNewParams, opts ...option.RequestOption) (res *CreateAssetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/assets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns details for the specified asset-id belonging to the current NVIDIA Cloud
// Account. Requires either a bearer token or an api-key with 'invoke_function'
// scope in the HTTP Authorization header.
func (r *AssetService) Get(ctx context.Context, assetID string, opts ...option.RequestOption) (res *AssetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List assets owned by the current NVIDIA Cloud Account. Requires either a bearer
// token or an api-key with invoke_function scope in the HTTP Authorization header.
func (r *AssetService) List(ctx context.Context, opts ...option.RequestOption) (res *ListAssetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/assets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes asset belonging to the current NVIDIA Cloud Account. Requires either a
// bearer token or an api-key with 'invoke_function' scope in the HTTP
// Authorization header.
func (r *AssetService) Delete(ctx context.Context, assetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AssetResponse struct {
	// Data Transfer Object(DTO) representing an asset
	Asset AssetResponseAsset `json:"asset"`
	JSON  assetResponseJSON  `json:"-"`
}

// assetResponseJSON contains the JSON metadata for the struct [AssetResponse]
type assetResponseJSON struct {
	Asset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an asset
type AssetResponseAsset struct {
	// Asset id
	AssetID string `json:"assetId" format:"uuid"`
	// Content-type specified when creating the asset
	ContentType string `json:"contentType"`
	// Description specified when creating the asset
	Description string                 `json:"description"`
	JSON        assetResponseAssetJSON `json:"-"`
}

// assetResponseAssetJSON contains the JSON metadata for the struct
// [AssetResponseAsset]
type assetResponseAssetJSON struct {
	AssetID     apijson.Field
	ContentType apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetResponseAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetResponseAssetJSON) RawJSON() string {
	return r.raw
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

// Response body containing list of assets of the current nca id
type ListAssetsResponse struct {
	// List of assets uploaded for the nca id
	Assets []ListAssetsResponseAsset `json:"assets"`
	JSON   listAssetsResponseJSON    `json:"-"`
}

// listAssetsResponseJSON contains the JSON metadata for the struct
// [ListAssetsResponse]
type listAssetsResponseJSON struct {
	Assets      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListAssetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listAssetsResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an asset
type ListAssetsResponseAsset struct {
	// Asset id
	AssetID string `json:"assetId" format:"uuid"`
	// Content-type specified when creating the asset
	ContentType string `json:"contentType"`
	// Description specified when creating the asset
	Description string                      `json:"description"`
	JSON        listAssetsResponseAssetJSON `json:"-"`
}

// listAssetsResponseAssetJSON contains the JSON metadata for the struct
// [ListAssetsResponseAsset]
type listAssetsResponseAssetJSON struct {
	AssetID     apijson.Field
	ContentType apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListAssetsResponseAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listAssetsResponseAssetJSON) RawJSON() string {
	return r.raw
}

type AssetNewParams struct {
	// Content type of the asset such image/png, image/jpeg, etc.
	ContentType param.Field[string] `json:"contentType,required"`
	// Asset description
	Description param.Field[string] `json:"description,required"`
}

func (r AssetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
