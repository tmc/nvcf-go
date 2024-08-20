// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/nvcf-go/internal/apijson"
	"github.com/stainless-sdks/nvcf-go/internal/requestconfig"
	"github.com/stainless-sdks/nvcf-go/option"
)

// ClusterGroupService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterGroupService] method instead.
type ClusterGroupService struct {
	Options []option.RequestOption
}

// NewClusterGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClusterGroupService(opts ...option.RequestOption) (r *ClusterGroupService) {
	r = &ClusterGroupService{}
	r.Options = opts
	return
}

// Lists Cluster Groups for the current account. The response includes cluster
// groups defined specifically in the current account and publicly available
// cluster groups such as GFN, OCI, etc. Requires a bearer token with
// 'list_cluster_groups' scope in HTTP Authorization header.
func (r *ClusterGroupService) List(ctx context.Context, opts ...option.RequestOption) (res *ClusterGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/clusterGroups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ClusterGroupsResponse struct {
	ClusterGroups []ClusterGroupsResponseClusterGroup `json:"clusterGroups"`
	JSON          clusterGroupsResponseJSON           `json:"-"`
}

// clusterGroupsResponseJSON contains the JSON metadata for the struct
// [ClusterGroupsResponse]
type clusterGroupsResponseJSON struct {
	ClusterGroups apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ClusterGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clusterGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type ClusterGroupsResponseClusterGroup struct {
	ID               string                                      `json:"id" format:"uuid"`
	AuthorizedNcaIDs []string                                    `json:"authorizedNcaIds"`
	Clusters         []ClusterGroupsResponseClusterGroupsCluster `json:"clusters"`
	GPUs             []ClusterGroupsResponseClusterGroupsGPU     `json:"gpus"`
	Name             string                                      `json:"name"`
	NcaID            string                                      `json:"ncaId"`
	JSON             clusterGroupsResponseClusterGroupJSON       `json:"-"`
}

// clusterGroupsResponseClusterGroupJSON contains the JSON metadata for the struct
// [ClusterGroupsResponseClusterGroup]
type clusterGroupsResponseClusterGroupJSON struct {
	ID               apijson.Field
	AuthorizedNcaIDs apijson.Field
	Clusters         apijson.Field
	GPUs             apijson.Field
	Name             apijson.Field
	NcaID            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ClusterGroupsResponseClusterGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clusterGroupsResponseClusterGroupJSON) RawJSON() string {
	return r.raw
}

type ClusterGroupsResponseClusterGroupsCluster struct {
	ID         string                                        `json:"id"`
	K8sVersion string                                        `json:"k8sVersion"`
	Name       string                                        `json:"name"`
	JSON       clusterGroupsResponseClusterGroupsClusterJSON `json:"-"`
}

// clusterGroupsResponseClusterGroupsClusterJSON contains the JSON metadata for the
// struct [ClusterGroupsResponseClusterGroupsCluster]
type clusterGroupsResponseClusterGroupsClusterJSON struct {
	ID          apijson.Field
	K8sVersion  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClusterGroupsResponseClusterGroupsCluster) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clusterGroupsResponseClusterGroupsClusterJSON) RawJSON() string {
	return r.raw
}

type ClusterGroupsResponseClusterGroupsGPU struct {
	InstanceTypes []ClusterGroupsResponseClusterGroupsGPUsInstanceType `json:"instanceTypes"`
	Name          string                                               `json:"name"`
	JSON          clusterGroupsResponseClusterGroupsGPUJSON            `json:"-"`
}

// clusterGroupsResponseClusterGroupsGPUJSON contains the JSON metadata for the
// struct [ClusterGroupsResponseClusterGroupsGPU]
type clusterGroupsResponseClusterGroupsGPUJSON struct {
	InstanceTypes apijson.Field
	Name          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ClusterGroupsResponseClusterGroupsGPU) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clusterGroupsResponseClusterGroupsGPUJSON) RawJSON() string {
	return r.raw
}

type ClusterGroupsResponseClusterGroupsGPUsInstanceType struct {
	Default     bool                                                   `json:"default"`
	Description string                                                 `json:"description"`
	Name        string                                                 `json:"name"`
	JSON        clusterGroupsResponseClusterGroupsGPUsInstanceTypeJSON `json:"-"`
}

// clusterGroupsResponseClusterGroupsGPUsInstanceTypeJSON contains the JSON
// metadata for the struct [ClusterGroupsResponseClusterGroupsGPUsInstanceType]
type clusterGroupsResponseClusterGroupsGPUsInstanceTypeJSON struct {
	Default     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClusterGroupsResponseClusterGroupsGPUsInstanceType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clusterGroupsResponseClusterGroupsGPUsInstanceTypeJSON) RawJSON() string {
	return r.raw
}
