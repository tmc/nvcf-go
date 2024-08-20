// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/tmc/nvcf-go/internal/apierror"
	"github.com/tmc/nvcf-go/shared"
)

type Error = apierror.Error

// Parties authorized to invoke function
//
// This is an alias to an internal type.
type AuthorizedPartiesResponse = shared.AuthorizedPartiesResponse

// Data Transfer Object(DTO) representing a function with authorized accounts
//
// This is an alias to an internal type.
type AuthorizedPartiesResponseFunction = shared.AuthorizedPartiesResponseFunction

// Data Transfer Object(DTO) representing an authorized party.
//
// This is an alias to an internal type.
type AuthorizedPartiesResponseFunctionAuthorizedParty = shared.AuthorizedPartiesResponseFunctionAuthorizedParty

// Response body with function details
//
// This is an alias to an internal type.
type FunctionResponse = shared.FunctionResponse

// Data Transfer Object (DTO) representing a function
//
// This is an alias to an internal type.
type FunctionResponseFunction = shared.FunctionResponseFunction

// Used to indicate a STREAMING function. Defaults to DEFAULT.
//
// This is an alias to an internal type.
type FunctionResponseFunctionFunctionType = shared.FunctionResponseFunctionFunctionType

// This is an alias to an internal value.
const FunctionResponseFunctionFunctionTypeDefault = shared.FunctionResponseFunctionFunctionTypeDefault

// This is an alias to an internal value.
const FunctionResponseFunctionFunctionTypeStreaming = shared.FunctionResponseFunctionFunctionTypeStreaming

// Function status
//
// This is an alias to an internal type.
type FunctionResponseFunctionStatus = shared.FunctionResponseFunctionStatus

// This is an alias to an internal value.
const FunctionResponseFunctionStatusActive = shared.FunctionResponseFunctionStatusActive

// This is an alias to an internal value.
const FunctionResponseFunctionStatusDeploying = shared.FunctionResponseFunctionStatusDeploying

// This is an alias to an internal value.
const FunctionResponseFunctionStatusError = shared.FunctionResponseFunctionStatusError

// This is an alias to an internal value.
const FunctionResponseFunctionStatusInactive = shared.FunctionResponseFunctionStatusInactive

// This is an alias to an internal value.
const FunctionResponseFunctionStatusDeleted = shared.FunctionResponseFunctionStatusDeleted

// Data Transfer Object(DTO) representing a spot instance
//
// This is an alias to an internal type.
type FunctionResponseFunctionActiveInstance = shared.FunctionResponseFunctionActiveInstance

// Instance status
//
// This is an alias to an internal type.
type FunctionResponseFunctionActiveInstancesInstanceStatus = shared.FunctionResponseFunctionActiveInstancesInstanceStatus

// This is an alias to an internal value.
const FunctionResponseFunctionActiveInstancesInstanceStatusActive = shared.FunctionResponseFunctionActiveInstancesInstanceStatusActive

// This is an alias to an internal value.
const FunctionResponseFunctionActiveInstancesInstanceStatusErrored = shared.FunctionResponseFunctionActiveInstancesInstanceStatusErrored

// This is an alias to an internal value.
const FunctionResponseFunctionActiveInstancesInstanceStatusPreempted = shared.FunctionResponseFunctionActiveInstancesInstanceStatusPreempted

// This is an alias to an internal value.
const FunctionResponseFunctionActiveInstancesInstanceStatusDeleted = shared.FunctionResponseFunctionActiveInstancesInstanceStatusDeleted

// Invocation request body format
//
// This is an alias to an internal type.
type FunctionResponseFunctionAPIBodyFormat = shared.FunctionResponseFunctionAPIBodyFormat

// This is an alias to an internal value.
const FunctionResponseFunctionAPIBodyFormatPredictV2 = shared.FunctionResponseFunctionAPIBodyFormatPredictV2

// This is an alias to an internal value.
const FunctionResponseFunctionAPIBodyFormatCustom = shared.FunctionResponseFunctionAPIBodyFormatCustom

// Data Transfer Object(DTO) representing a container environment entry
//
// This is an alias to an internal type.
type FunctionResponseFunctionContainerEnvironment = shared.FunctionResponseFunctionContainerEnvironment

// Data Transfer Object(DTO) representing a function ne
//
// This is an alias to an internal type.
type FunctionResponseFunctionHealth = shared.FunctionResponseFunctionHealth

// HTTP/gPRC protocol type for health endpoint
//
// This is an alias to an internal type.
type FunctionResponseFunctionHealthProtocol = shared.FunctionResponseFunctionHealthProtocol

// This is an alias to an internal value.
const FunctionResponseFunctionHealthProtocolHTTP = shared.FunctionResponseFunctionHealthProtocolHTTP

// This is an alias to an internal value.
const FunctionResponseFunctionHealthProtocolGRpc = shared.FunctionResponseFunctionHealthProtocolGRpc

// Data Transfer Object(DTO) representing an artifact
//
// This is an alias to an internal type.
type FunctionResponseFunctionModel = shared.FunctionResponseFunctionModel

// Data Transfer Object(DTO) representing an artifact
//
// This is an alias to an internal type.
type FunctionResponseFunctionResource = shared.FunctionResponseFunctionResource

// Response body with result from a request for executing a job/task as a cloud
// function using a GPU powered spot/on-demand instance.
//
// This is an alias to an internal type.
type InvokeFunctionResponse = shared.InvokeFunctionResponse

// Status of the task/job executing cloud function.
//
// This is an alias to an internal type.
type InvokeFunctionResponseStatus = shared.InvokeFunctionResponseStatus

// This is an alias to an internal value.
const InvokeFunctionResponseStatusErrored = shared.InvokeFunctionResponseStatusErrored

// This is an alias to an internal value.
const InvokeFunctionResponseStatusInProgress = shared.InvokeFunctionResponseStatusInProgress

// This is an alias to an internal value.
const InvokeFunctionResponseStatusFulfilled = shared.InvokeFunctionResponseStatusFulfilled

// This is an alias to an internal value.
const InvokeFunctionResponseStatusPendingEvaluation = shared.InvokeFunctionResponseStatusPendingEvaluation

// This is an alias to an internal value.
const InvokeFunctionResponseStatusRejected = shared.InvokeFunctionResponseStatusRejected
