// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/nvcf-go/internal/apijson"
	"github.com/stainless-sdks/nvcf-go/internal/apiquery"
	"github.com/stainless-sdks/nvcf-go/internal/param"
	"github.com/stainless-sdks/nvcf-go/internal/requestconfig"
	"github.com/stainless-sdks/nvcf-go/option"
	"github.com/stainless-sdks/nvcf-go/shared"
)

// FunctionDeploymentFunctionVersionService contains methods and other services
// that help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionDeploymentFunctionVersionService] method instead.
type FunctionDeploymentFunctionVersionService struct {
	Options []option.RequestOption
}

// NewFunctionDeploymentFunctionVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFunctionDeploymentFunctionVersionService(opts ...option.RequestOption) (r *FunctionDeploymentFunctionVersionService) {
	r = &FunctionDeploymentFunctionVersionService{}
	r.Options = opts
	return
}

// Initiates deployment for the specified function version. Upon invocation of this
// endpoint, the function's status transitions to 'DEPLOYING'. If the specified
// function version is public, then Account Admin cannot perform this operation.
// Access to this endpoint mandates a bearer token with 'deploy_function' scope in
// the HTTP Authorization header.
func (r *FunctionDeploymentFunctionVersionService) New(ctx context.Context, functionID string, functionVersionID string, body FunctionDeploymentFunctionVersionNewParams, opts ...option.RequestOption) (res *DeploymentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/deployments/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Allows Account Admins to retrieve the deployment details of the specified
// function version. Access to this endpoint mandates a bearer token with
// 'deploy_function' scope in the HTTP Authorization header.
func (r *FunctionDeploymentFunctionVersionService) Get(ctx context.Context, functionID string, functionVersionID string, opts ...option.RequestOption) (res *DeploymentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/deployments/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the deployment specs of the specified function version. It's important
// to note that GPU type and backend configurations cannot be modified through this
// endpoint. If the specified function is public, then Account Admin cannot perform
// this operation. Access to this endpoint mandates a bearer token with
// 'deploy_function' scope in the HTTP Authorization header.
func (r *FunctionDeploymentFunctionVersionService) Update(ctx context.Context, functionID string, functionVersionID string, body FunctionDeploymentFunctionVersionUpdateParams, opts ...option.RequestOption) (res *DeploymentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/deployments/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes the deployment associated with the specified function. Upon deletion,
// any active instances will be terminated, and the function's status will
// transition to 'INACTIVE'. To undeploy a function version gracefully, specify
// 'graceful=true' query parameter, allowing current tasks to complete before
// terminating the instances. If the specified function version is public, then
// Account Admin cannot perform this operation. Access to this endpoint mandates a
// bearer token with 'deploy_function' scope in the HTTP Authorization header.
func (r *FunctionDeploymentFunctionVersionService) Delete(ctx context.Context, functionID string, functionVersionID string, body FunctionDeploymentFunctionVersionDeleteParams, opts ...option.RequestOption) (res *shared.FunctionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/deployments/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Function Deployment Response
type DeploymentResponse struct {
	// Function deployment response
	Deployment DeploymentResponseDeployment `json:"deployment,required"`
	JSON       deploymentResponseJSON       `json:"-"`
}

// deploymentResponseJSON contains the JSON metadata for the struct
// [DeploymentResponse]
type deploymentResponseJSON struct {
	Deployment  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentResponseJSON) RawJSON() string {
	return r.raw
}

// Function deployment response
type DeploymentResponseDeployment struct {
	// Function deployment creation timestamp
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Function deployment details
	DeploymentSpecifications []DeploymentResponseDeploymentDeploymentSpecification `json:"deploymentSpecifications,required"`
	// Function id
	FunctionID string `json:"functionId,required" format:"uuid"`
	// Function name
	FunctionName string `json:"functionName,required"`
	// Function status
	FunctionStatus DeploymentResponseDeploymentFunctionStatus `json:"functionStatus,required"`
	// Function version id
	FunctionVersionID string `json:"functionVersionId,required" format:"uuid"`
	// NVIDIA Cloud Account Id
	NcaID string `json:"ncaId,required"`
	// Health info for a deployment specification is included only if there are any
	// issues/errors.
	HealthInfo []DeploymentResponseDeploymentHealthInfo `json:"healthInfo"`
	// Deprecated Request Queue URL
	RequestQueueURL string                           `json:"requestQueueUrl"`
	JSON            deploymentResponseDeploymentJSON `json:"-"`
}

// deploymentResponseDeploymentJSON contains the JSON metadata for the struct
// [DeploymentResponseDeployment]
type deploymentResponseDeploymentJSON struct {
	CreatedAt                apijson.Field
	DeploymentSpecifications apijson.Field
	FunctionID               apijson.Field
	FunctionName             apijson.Field
	FunctionStatus           apijson.Field
	FunctionVersionID        apijson.Field
	NcaID                    apijson.Field
	HealthInfo               apijson.Field
	RequestQueueURL          apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *DeploymentResponseDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentResponseDeploymentJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing GPU specification.
type DeploymentResponseDeploymentDeploymentSpecification struct {
	// GPU name from the cluster
	GPU string `json:"gpu,required"`
	// Instance type, based on GPU, assigned to a Worker
	InstanceType string `json:"instanceType,required"`
	// Maximum number of spot instances for the deployment
	MaxInstances int64 `json:"maxInstances,required"`
	// Minimum number of spot instances for the deployment
	MinInstances int64 `json:"minInstances,required"`
	// Specific attributes capabilities to deploy functions.
	Attributes []string `json:"attributes"`
	// List of availability-zones(or clusters) in the cluster group
	AvailabilityZones []string `json:"availabilityZones"`
	// Backend/CSP where the GPU powered instance will be launched
	Backend string `json:"backend"`
	// Specific clusters within spot instance or worker node powered by the selected
	// instance-type to deploy function.
	Clusters      []string    `json:"clusters"`
	Configuration interface{} `json:"configuration"`
	// Max request concurrency between 1 (default) and 1024.
	MaxRequestConcurrency int64 `json:"maxRequestConcurrency"`
	// Preferred order of deployment if there are several gpu specs.
	PreferredOrder int64 `json:"preferredOrder"`
	// List of regions allowed to deploy. The instance or worker node will be in one of
	// the specified geographical regions.
	Regions []string                                                `json:"regions"`
	JSON    deploymentResponseDeploymentDeploymentSpecificationJSON `json:"-"`
}

// deploymentResponseDeploymentDeploymentSpecificationJSON contains the JSON
// metadata for the struct [DeploymentResponseDeploymentDeploymentSpecification]
type deploymentResponseDeploymentDeploymentSpecificationJSON struct {
	GPU                   apijson.Field
	InstanceType          apijson.Field
	MaxInstances          apijson.Field
	MinInstances          apijson.Field
	Attributes            apijson.Field
	AvailabilityZones     apijson.Field
	Backend               apijson.Field
	Clusters              apijson.Field
	Configuration         apijson.Field
	MaxRequestConcurrency apijson.Field
	PreferredOrder        apijson.Field
	Regions               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DeploymentResponseDeploymentDeploymentSpecification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentResponseDeploymentDeploymentSpecificationJSON) RawJSON() string {
	return r.raw
}

// Function status
type DeploymentResponseDeploymentFunctionStatus string

const (
	DeploymentResponseDeploymentFunctionStatusActive    DeploymentResponseDeploymentFunctionStatus = "ACTIVE"
	DeploymentResponseDeploymentFunctionStatusDeploying DeploymentResponseDeploymentFunctionStatus = "DEPLOYING"
	DeploymentResponseDeploymentFunctionStatusError     DeploymentResponseDeploymentFunctionStatus = "ERROR"
	DeploymentResponseDeploymentFunctionStatusInactive  DeploymentResponseDeploymentFunctionStatus = "INACTIVE"
	DeploymentResponseDeploymentFunctionStatusDeleted   DeploymentResponseDeploymentFunctionStatus = "DELETED"
)

func (r DeploymentResponseDeploymentFunctionStatus) IsKnown() bool {
	switch r {
	case DeploymentResponseDeploymentFunctionStatusActive, DeploymentResponseDeploymentFunctionStatusDeploying, DeploymentResponseDeploymentFunctionStatusError, DeploymentResponseDeploymentFunctionStatusInactive, DeploymentResponseDeploymentFunctionStatusDeleted:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing deployment error
type DeploymentResponseDeploymentHealthInfo struct {
	// Backend/CSP where the GPU powered instance will be launched
	Backend string `json:"backend,required"`
	// Deployment error
	Error string `json:"error,required"`
	// GPU Type as per SDD
	GPU string `json:"gpu,required"`
	// Instance type
	InstanceType string `json:"instanceType"`
	// SIS Request ID
	SisRequestID string                                     `json:"sisRequestId" format:"uuid"`
	JSON         deploymentResponseDeploymentHealthInfoJSON `json:"-"`
}

// deploymentResponseDeploymentHealthInfoJSON contains the JSON metadata for the
// struct [DeploymentResponseDeploymentHealthInfo]
type deploymentResponseDeploymentHealthInfoJSON struct {
	Backend      apijson.Field
	Error        apijson.Field
	GPU          apijson.Field
	InstanceType apijson.Field
	SisRequestID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DeploymentResponseDeploymentHealthInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentResponseDeploymentHealthInfoJSON) RawJSON() string {
	return r.raw
}

type FunctionDeploymentFunctionVersionNewParams struct {
	// Deployment specs with Backend, GPU, instance-type, etc. details
	DeploymentSpecifications param.Field[[]FunctionDeploymentFunctionVersionNewParamsDeploymentSpecification] `json:"deploymentSpecifications,required"`
}

func (r FunctionDeploymentFunctionVersionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing GPU specification.
type FunctionDeploymentFunctionVersionNewParamsDeploymentSpecification struct {
	// GPU name from the cluster
	GPU param.Field[string] `json:"gpu,required"`
	// Instance type, based on GPU, assigned to a Worker
	InstanceType param.Field[string] `json:"instanceType,required"`
	// Maximum number of spot instances for the deployment
	MaxInstances param.Field[int64] `json:"maxInstances,required"`
	// Minimum number of spot instances for the deployment
	MinInstances param.Field[int64] `json:"minInstances,required"`
	// Specific attributes capabilities to deploy functions.
	Attributes param.Field[[]string] `json:"attributes"`
	// List of availability-zones(or clusters) in the cluster group
	AvailabilityZones param.Field[[]string] `json:"availabilityZones"`
	// Backend/CSP where the GPU powered instance will be launched
	Backend param.Field[string] `json:"backend"`
	// Specific clusters within spot instance or worker node powered by the selected
	// instance-type to deploy function.
	Clusters      param.Field[[]string]    `json:"clusters"`
	Configuration param.Field[interface{}] `json:"configuration"`
	// Max request concurrency between 1 (default) and 1024.
	MaxRequestConcurrency param.Field[int64] `json:"maxRequestConcurrency"`
	// Preferred order of deployment if there are several gpu specs.
	PreferredOrder param.Field[int64] `json:"preferredOrder"`
	// List of regions allowed to deploy. The instance or worker node will be in one of
	// the specified geographical regions.
	Regions param.Field[[]string] `json:"regions"`
}

func (r FunctionDeploymentFunctionVersionNewParamsDeploymentSpecification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionDeploymentFunctionVersionUpdateParams struct {
	// Deployment specs with Backend, GPU, instance-type, etc. details
	DeploymentSpecifications param.Field[[]FunctionDeploymentFunctionVersionUpdateParamsDeploymentSpecification] `json:"deploymentSpecifications,required"`
}

func (r FunctionDeploymentFunctionVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing GPU specification.
type FunctionDeploymentFunctionVersionUpdateParamsDeploymentSpecification struct {
	// GPU name from the cluster
	GPU param.Field[string] `json:"gpu,required"`
	// Instance type, based on GPU, assigned to a Worker
	InstanceType param.Field[string] `json:"instanceType,required"`
	// Maximum number of spot instances for the deployment
	MaxInstances param.Field[int64] `json:"maxInstances,required"`
	// Minimum number of spot instances for the deployment
	MinInstances param.Field[int64] `json:"minInstances,required"`
	// Specific attributes capabilities to deploy functions.
	Attributes param.Field[[]string] `json:"attributes"`
	// List of availability-zones(or clusters) in the cluster group
	AvailabilityZones param.Field[[]string] `json:"availabilityZones"`
	// Backend/CSP where the GPU powered instance will be launched
	Backend param.Field[string] `json:"backend"`
	// Specific clusters within spot instance or worker node powered by the selected
	// instance-type to deploy function.
	Clusters      param.Field[[]string]    `json:"clusters"`
	Configuration param.Field[interface{}] `json:"configuration"`
	// Max request concurrency between 1 (default) and 1024.
	MaxRequestConcurrency param.Field[int64] `json:"maxRequestConcurrency"`
	// Preferred order of deployment if there are several gpu specs.
	PreferredOrder param.Field[int64] `json:"preferredOrder"`
	// List of regions allowed to deploy. The instance or worker node will be in one of
	// the specified geographical regions.
	Regions param.Field[[]string] `json:"regions"`
}

func (r FunctionDeploymentFunctionVersionUpdateParamsDeploymentSpecification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionDeploymentFunctionVersionDeleteParams struct {
	// Query param to deactivate function for graceful shutdown
	Graceful param.Field[bool] `query:"graceful"`
}

// URLQuery serializes [FunctionDeploymentFunctionVersionDeleteParams]'s query
// parameters as `url.Values`.
func (r FunctionDeploymentFunctionVersionDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
