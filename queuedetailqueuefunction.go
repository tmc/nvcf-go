// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/tmc/nvcf-go/internal/apijson"
	"github.com/tmc/nvcf-go/internal/requestconfig"
	"github.com/tmc/nvcf-go/option"
)

// QueueDetailQueueFunctionService contains methods and other services that help
// with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueDetailQueueFunctionService] method instead.
type QueueDetailQueueFunctionService struct {
	Options  []option.RequestOption
	Versions *QueueDetailQueueFunctionVersionService
}

// NewQueueDetailQueueFunctionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewQueueDetailQueueFunctionService(opts ...option.RequestOption) (r *QueueDetailQueueFunctionService) {
	r = &QueueDetailQueueFunctionService{}
	r.Options = opts
	r.Versions = NewQueueDetailQueueFunctionVersionService(opts...)
	return
}

// Provides details of all the queues associated with the specified function. If a
// function has multiple versions and they are all deployed, then the response
// includes details of all the queues. If the specified function is public, then
// Account Admin cannot perform this operation. Requires a bearer token or an
// api-key with 'queue_details' scope in the HTTP Authorization header.
func (r *QueueDetailQueueFunctionService) List(ctx context.Context, functionID string, opts ...option.RequestOption) (res *QueueDetailQueueFunctionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/queues/functions/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Request queue details of all the functions with same id in an account
type QueueDetailQueueFunctionListResponse struct {
	// Function id
	FunctionID string `json:"functionId,required" format:"uuid"`
	// Details of all the queues associated with same named functions
	Queues []QueueDetailQueueFunctionListResponseQueue `json:"queues,required"`
	JSON   queueDetailQueueFunctionListResponseJSON    `json:"-"`
}

// queueDetailQueueFunctionListResponseJSON contains the JSON metadata for the
// struct [QueueDetailQueueFunctionListResponse]
type queueDetailQueueFunctionListResponseJSON struct {
	FunctionID  apijson.Field
	Queues      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueDetailQueueFunctionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueDetailQueueFunctionListResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a request queue for function version
type QueueDetailQueueFunctionListResponseQueue struct {
	// Function name
	FunctionName string `json:"functionName,required"`
	// Function status
	FunctionStatus QueueDetailQueueFunctionListResponseQueuesFunctionStatus `json:"functionStatus,required"`
	// Function version id
	FunctionVersionID string `json:"functionVersionId,required" format:"uuid"`
	// Approximate number of messages in the request queue
	QueueDepth int64                                         `json:"queueDepth"`
	JSON       queueDetailQueueFunctionListResponseQueueJSON `json:"-"`
}

// queueDetailQueueFunctionListResponseQueueJSON contains the JSON metadata for the
// struct [QueueDetailQueueFunctionListResponseQueue]
type queueDetailQueueFunctionListResponseQueueJSON struct {
	FunctionName      apijson.Field
	FunctionStatus    apijson.Field
	FunctionVersionID apijson.Field
	QueueDepth        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *QueueDetailQueueFunctionListResponseQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueDetailQueueFunctionListResponseQueueJSON) RawJSON() string {
	return r.raw
}

// Function status
type QueueDetailQueueFunctionListResponseQueuesFunctionStatus string

const (
	QueueDetailQueueFunctionListResponseQueuesFunctionStatusActive    QueueDetailQueueFunctionListResponseQueuesFunctionStatus = "ACTIVE"
	QueueDetailQueueFunctionListResponseQueuesFunctionStatusDeploying QueueDetailQueueFunctionListResponseQueuesFunctionStatus = "DEPLOYING"
	QueueDetailQueueFunctionListResponseQueuesFunctionStatusError     QueueDetailQueueFunctionListResponseQueuesFunctionStatus = "ERROR"
	QueueDetailQueueFunctionListResponseQueuesFunctionStatusInactive  QueueDetailQueueFunctionListResponseQueuesFunctionStatus = "INACTIVE"
	QueueDetailQueueFunctionListResponseQueuesFunctionStatusDeleted   QueueDetailQueueFunctionListResponseQueuesFunctionStatus = "DELETED"
)

func (r QueueDetailQueueFunctionListResponseQueuesFunctionStatus) IsKnown() bool {
	switch r {
	case QueueDetailQueueFunctionListResponseQueuesFunctionStatusActive, QueueDetailQueueFunctionListResponseQueuesFunctionStatusDeploying, QueueDetailQueueFunctionListResponseQueuesFunctionStatusError, QueueDetailQueueFunctionListResponseQueuesFunctionStatusInactive, QueueDetailQueueFunctionListResponseQueuesFunctionStatusDeleted:
		return true
	}
	return false
}
