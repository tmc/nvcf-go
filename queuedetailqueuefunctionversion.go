// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/nvcf-go/internal/apijson"
	"github.com/stainless-sdks/nvcf-go/internal/requestconfig"
	"github.com/stainless-sdks/nvcf-go/option"
)

// QueueDetailQueueFunctionVersionService contains methods and other services that
// help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueDetailQueueFunctionVersionService] method instead.
type QueueDetailQueueFunctionVersionService struct {
	Options []option.RequestOption
}

// NewQueueDetailQueueFunctionVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewQueueDetailQueueFunctionVersionService(opts ...option.RequestOption) (r *QueueDetailQueueFunctionVersionService) {
	r = &QueueDetailQueueFunctionVersionService{}
	r.Options = opts
	return
}

// Provides details of all the queues associated with the specified function. If a
// function has multiple versions and they are all deployed, then the response
// includes details of all the queues. If the specified function is public, then
// Account Admin cannot perform this operation. Requires a bearer token or an
// api-key with 'queue_details' scope in the HTTP Authorization header.
func (r *QueueDetailQueueFunctionVersionService) Get(ctx context.Context, functionID string, versionID string, opts ...option.RequestOption) (res *QueueDetailQueueFunctionVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required versionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/queues/functions/%s/versions/%s", functionID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Request queue details of all the functions with same id in an account
type QueueDetailQueueFunctionVersionGetResponse struct {
	// Function id
	FunctionID string `json:"functionId,required" format:"uuid"`
	// Details of all the queues associated with same named functions
	Queues []QueueDetailQueueFunctionVersionGetResponseQueue `json:"queues,required"`
	JSON   queueDetailQueueFunctionVersionGetResponseJSON    `json:"-"`
}

// queueDetailQueueFunctionVersionGetResponseJSON contains the JSON metadata for
// the struct [QueueDetailQueueFunctionVersionGetResponse]
type queueDetailQueueFunctionVersionGetResponseJSON struct {
	FunctionID  apijson.Field
	Queues      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueDetailQueueFunctionVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueDetailQueueFunctionVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a request queue for function version
type QueueDetailQueueFunctionVersionGetResponseQueue struct {
	// Function name
	FunctionName string `json:"functionName,required"`
	// Function status
	FunctionStatus QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatus `json:"functionStatus,required"`
	// Function version id
	FunctionVersionID string `json:"functionVersionId,required" format:"uuid"`
	// Approximate number of messages in the request queue
	QueueDepth int64                                               `json:"queueDepth"`
	JSON       queueDetailQueueFunctionVersionGetResponseQueueJSON `json:"-"`
}

// queueDetailQueueFunctionVersionGetResponseQueueJSON contains the JSON metadata
// for the struct [QueueDetailQueueFunctionVersionGetResponseQueue]
type queueDetailQueueFunctionVersionGetResponseQueueJSON struct {
	FunctionName      apijson.Field
	FunctionStatus    apijson.Field
	FunctionVersionID apijson.Field
	QueueDepth        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *QueueDetailQueueFunctionVersionGetResponseQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueDetailQueueFunctionVersionGetResponseQueueJSON) RawJSON() string {
	return r.raw
}

// Function status
type QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatus string

const (
	QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatusActive    QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatus = "ACTIVE"
	QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatusDeploying QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatus = "DEPLOYING"
	QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatusError     QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatus = "ERROR"
	QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatusInactive  QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatus = "INACTIVE"
	QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatusDeleted   QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatus = "DELETED"
)

func (r QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatus) IsKnown() bool {
	switch r {
	case QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatusActive, QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatusDeploying, QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatusError, QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatusInactive, QueueDetailQueueFunctionVersionGetResponseQueuesFunctionStatusDeleted:
		return true
	}
	return false
}
