// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/tmc/nvcf-go/internal/apijson"
	"github.com/tmc/nvcf-go/internal/param"
	"github.com/tmc/nvcf-go/internal/requestconfig"
	"github.com/tmc/nvcf-go/option"
	"github.com/tmc/nvcf-go/shared"
)

// FunctionVersionService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionVersionService] method instead.
type FunctionVersionService struct {
	Options []option.RequestOption
}

// NewFunctionVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionVersionService(opts ...option.RequestOption) (r *FunctionVersionService) {
	r = &FunctionVersionService{}
	r.Options = opts
	return
}

// Creates a version of the specified function within the authenticated NVIDIA
// Cloud Account. Requires a bearer token with 'register_function' scope in the
// HTTP Authorization header.
func (r *FunctionVersionService) New(ctx context.Context, functionID string, body FunctionVersionNewParams, opts ...option.RequestOption) (res *CreateFunctionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/functions/%s/versions", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed information of the specified function version in the
// authenticated NVIDIA Cloud Account. Requires either a bearer token or an api-key
// with 'list_functions' or 'list_functions_details' scopes in the HTTP
// Authorization header.
func (r *FunctionVersionService) Get(ctx context.Context, functionID string, functionVersionID string, opts ...option.RequestOption) (res *shared.FunctionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists details of all the versions of the specified function in the authenticated
// NVIDIA Cloud Account. Requires either a bearer token or an api-key with
// 'list_functions' or 'list_functions_details' scopes in the HTTP Authorization
// header.
func (r *FunctionVersionService) List(ctx context.Context, functionID string, opts ...option.RequestOption) (res *ListFunctionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/functions/%s/versions", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes the specified function version in the authenticated NVIDIA Cloud
// Account. Requires a bearer token with 'delete_function' scope in the HTTP
// Authorization header. If the function version is public, then Account Admin
// cannot delete the function.
func (r *FunctionVersionService) Delete(ctx context.Context, functionID string, functionVersionID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type FunctionVersionNewParams struct {
	// Entrypoint for invoking the container to process a request
	InferenceURL param.Field[string] `json:"inferenceUrl,required" format:"uri"`
	// Function name must start with lowercase/uppercase/digit and can only contain
	// lowercase, uppercase, digit, hyphen, and underscore characters
	Name param.Field[string] `json:"name,required"`
	// Invocation request body format
	APIBodyFormat param.Field[FunctionVersionNewParamsAPIBodyFormat] `json:"apiBodyFormat"`
	// Args to be passed when launching the container
	ContainerArgs param.Field[string] `json:"containerArgs"`
	// Environment settings for launching the container
	ContainerEnvironment param.Field[[]FunctionVersionNewParamsContainerEnvironment] `json:"containerEnvironment"`
	// Optional custom container image
	ContainerImage param.Field[string] `json:"containerImage" format:"uri"`
	// Optional function/version description
	Description param.Field[string] `json:"description"`
	// Optional function type, used to indicate a STREAMING function. Defaults to
	// DEFAULT.
	FunctionType param.Field[FunctionVersionNewParamsFunctionType] `json:"functionType"`
	// Data Transfer Object(DTO) representing a function ne
	Health param.Field[FunctionVersionNewParamsHealth] `json:"health"`
	// Health endpoint for the container or the helmChart
	HealthUri param.Field[string] `json:"healthUri" format:"uri"`
	// Optional Helm Chart
	HelmChart param.Field[string] `json:"helmChart" format:"uri"`
	// Helm Chart Service Name is required when helmChart property is specified
	HelmChartServiceName param.Field[string] `json:"helmChartServiceName"`
	// Optional port number where the inference listener is running. Defaults to 8000
	// for Triton.
	InferencePort param.Field[int64] `json:"inferencePort"`
	// Optional set of models
	Models param.Field[[]FunctionVersionNewParamsModel] `json:"models"`
	// Optional set of resources
	Resources param.Field[[]FunctionVersionNewParamsResource] `json:"resources"`
	// Optional secrets
	Secrets param.Field[[]FunctionVersionNewParamsSecret] `json:"secrets"`
	// Optional set of tags - could be empty. Provided by user
	Tags param.Field[[]string] `json:"tags"`
}

func (r FunctionVersionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Invocation request body format
type FunctionVersionNewParamsAPIBodyFormat string

const (
	FunctionVersionNewParamsAPIBodyFormatPredictV2 FunctionVersionNewParamsAPIBodyFormat = "PREDICT_V2"
	FunctionVersionNewParamsAPIBodyFormatCustom    FunctionVersionNewParamsAPIBodyFormat = "CUSTOM"
)

func (r FunctionVersionNewParamsAPIBodyFormat) IsKnown() bool {
	switch r {
	case FunctionVersionNewParamsAPIBodyFormatPredictV2, FunctionVersionNewParamsAPIBodyFormatCustom:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a container environment entry
type FunctionVersionNewParamsContainerEnvironment struct {
	// Container environment key
	Key param.Field[string] `json:"key,required"`
	// Container environment value
	Value param.Field[string] `json:"value,required"`
}

func (r FunctionVersionNewParamsContainerEnvironment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional function type, used to indicate a STREAMING function. Defaults to
// DEFAULT.
type FunctionVersionNewParamsFunctionType string

const (
	FunctionVersionNewParamsFunctionTypeDefault   FunctionVersionNewParamsFunctionType = "DEFAULT"
	FunctionVersionNewParamsFunctionTypeStreaming FunctionVersionNewParamsFunctionType = "STREAMING"
)

func (r FunctionVersionNewParamsFunctionType) IsKnown() bool {
	switch r {
	case FunctionVersionNewParamsFunctionTypeDefault, FunctionVersionNewParamsFunctionTypeStreaming:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a function ne
type FunctionVersionNewParamsHealth struct {
	// Expected return status code considered as successful.
	ExpectedStatusCode param.Field[int64] `json:"expectedStatusCode,required"`
	// Port number where the health listener is running
	Port param.Field[int64] `json:"port,required"`
	// HTTP/gPRC protocol type for health endpoint
	Protocol param.Field[FunctionVersionNewParamsHealthProtocol] `json:"protocol,required"`
	// ISO 8601 duration string in PnDTnHnMn.nS format
	Timeout param.Field[string] `json:"timeout,required" format:"PnDTnHnMn.nS"`
	// Health endpoint for the container or the helmChart
	Uri param.Field[string] `json:"uri,required" format:"uri"`
}

func (r FunctionVersionNewParamsHealth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP/gPRC protocol type for health endpoint
type FunctionVersionNewParamsHealthProtocol string

const (
	FunctionVersionNewParamsHealthProtocolHTTP FunctionVersionNewParamsHealthProtocol = "HTTP"
	FunctionVersionNewParamsHealthProtocolGRpc FunctionVersionNewParamsHealthProtocol = "gRPC"
)

func (r FunctionVersionNewParamsHealthProtocol) IsKnown() bool {
	switch r {
	case FunctionVersionNewParamsHealthProtocolHTTP, FunctionVersionNewParamsHealthProtocolGRpc:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing an artifact
type FunctionVersionNewParamsModel struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r FunctionVersionNewParamsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an artifact
type FunctionVersionNewParamsResource struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r FunctionVersionNewParamsResource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing secret name/value pair
type FunctionVersionNewParamsSecret struct {
	// Secret name
	Name param.Field[string] `json:"name,required"`
	// Secret value
	Value param.Field[string] `json:"value,required"`
}

func (r FunctionVersionNewParamsSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
