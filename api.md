# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#FunctionResponse">FunctionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>

# FunctionManagement

## Functions

### Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListFunctionsResponse">ListFunctionsResponse</a>

Methods:

- <code title="get /v2/nvcf/functions/{functionId}/versions/{functionVersionId}">client.FunctionManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionManagementFunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v2/nvcf/metadata/functions/{functionId}/versions/{functionVersionId}">client.FunctionManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionManagementFunctionVersionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionManagementFunctionVersionUpdateParams">FunctionManagementFunctionVersionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/functions/{functionId}/versions">client.FunctionManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionManagementFunctionVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListFunctionsResponse">ListFunctionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FunctionDeployment

## Functions

### Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#DeploymentResponse">DeploymentResponse</a>

Methods:

- <code title="post /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionNewParams">FunctionDeploymentFunctionVersionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionUpdateParams">FunctionDeploymentFunctionVersionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionDeleteParams">FunctionDeploymentFunctionVersionDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FunctionInvocation

## Functions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationFunctionInvokeResponse">FunctionInvocationFunctionInvokeResponse</a>

Methods:

- <code title="post /v2/nvcf/pexec/functions/{functionId}">client.FunctionInvocation.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationFunctionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationFunctionInvokeParams">FunctionInvocationFunctionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationFunctionInvokeResponse">FunctionInvocationFunctionInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationFunctionVersionInvokeResponse">FunctionInvocationFunctionVersionInvokeResponse</a>

Methods:

- <code title="post /v2/nvcf/pexec/functions/{functionId}/versions/{versionId}">client.FunctionInvocation.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationFunctionVersionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationFunctionVersionInvokeParams">FunctionInvocationFunctionVersionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationFunctionVersionInvokeResponse">FunctionInvocationFunctionVersionInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Status

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationStatusGetResponse">FunctionInvocationStatusGetResponse</a>

Methods:

- <code title="get /v2/nvcf/pexec/status/{requestId}">client.FunctionInvocation.Status.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationStatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationStatusGetParams">FunctionInvocationStatusGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionInvocationStatusGetResponse">FunctionInvocationStatusGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# EnvelopeFunctionInvocation

## Functions

Methods:

- <code title="post /v2/nvcf/exec/functions/{functionId}">client.EnvelopeFunctionInvocation.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#EnvelopeFunctionInvocationFunctionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#EnvelopeFunctionInvocationFunctionInvokeParams">EnvelopeFunctionInvocationFunctionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="post /v2/nvcf/exec/functions/{functionId}/versions/{versionId}">client.EnvelopeFunctionInvocation.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#EnvelopeFunctionInvocationFunctionVersionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#EnvelopeFunctionInvocationFunctionVersionInvokeParams">EnvelopeFunctionInvocationFunctionVersionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Functions

Methods:

- <code title="post /v2/nvcf/functions">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionNewParams">FunctionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#CreateFunctionResponse">CreateFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/functions">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionListParams">FunctionListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListFunctionsResponse">ListFunctionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#CreateFunctionResponse">CreateFunctionResponse</a>

Methods:

- <code title="post /v2/nvcf/functions/{functionId}/versions">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionVersionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionVersionNewParams">FunctionVersionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#CreateFunctionResponse">CreateFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/functions/{functionId}/versions/{functionVersionId}">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## IDs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListFunctionIDsResponse">ListFunctionIDsResponse</a>

Methods:

- <code title="get /v2/nvcf/functions/ids">client.Functions.IDs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionIDService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionIDListParams">FunctionIDListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListFunctionIDsResponse">ListFunctionIDsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Authorizations

## Functions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListAuthorizedPartiesResponse">ListAuthorizedPartiesResponse</a>

Methods:

- <code title="get /v2/nvcf/authorizations/functions/{functionId}">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListAuthorizedPartiesResponse">ListAuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/authorizations/functions/{functionId}">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/nvcf/authorizations/functions/{functionId}">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionAuthorizeParams">AuthorizationFunctionAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="get /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionVersionService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionVersionAuthorizeParams">AuthorizationFunctionVersionAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetResponse">AssetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListAssetsResponse">ListAssetsResponse</a>

Methods:

- <code title="get /v2/nvcf/assets/{assetId}">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetResponse">AssetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/assets">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListAssetsResponse">ListAssetsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/assets/{assetId}">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# AssetManagement

## Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#CreateAssetResponse">CreateAssetResponse</a>

Methods:

- <code title="post /v2/nvcf/assets">client.AssetManagement.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetManagementAssetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetManagementAssetNewParams">AssetManagementAssetNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#CreateAssetResponse">CreateAssetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AuthorizedAccounts

## Functions

Methods:

- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/add">client.AuthorizedAccounts.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionAddParams">AuthorizedAccountFunctionAddParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/remove">client.AuthorizedAccounts.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionRemoveParams">AuthorizedAccountFunctionRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}/add">client.AuthorizedAccounts.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionVersionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionVersionAddParams">AuthorizedAccountFunctionVersionAddParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}/remove">client.AuthorizedAccounts.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionVersionService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionVersionRemoveParams">AuthorizedAccountFunctionVersionRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# QueueDetails

## Queues

### Functions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueueDetailQueueFunctionListResponse">QueueDetailQueueFunctionListResponse</a>

Methods:

- <code title="get /v2/nvcf/queues/functions/{functionId}">client.QueueDetails.Queues.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueueDetailQueueFunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueueDetailQueueFunctionListResponse">QueueDetailQueueFunctionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueueDetailQueueFunctionVersionGetResponse">QueueDetailQueueFunctionVersionGetResponse</a>

Methods:

- <code title="get /v2/nvcf/queues/functions/{functionId}/versions/{versionId}">client.QueueDetails.Queues.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueueDetailQueueFunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueueDetailQueueFunctionVersionGetResponse">QueueDetailQueueFunctionVersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Position

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueueDetailQueuePositionGetResponse">QueueDetailQueuePositionGetResponse</a>

Methods:

- <code title="get /v2/nvcf/queues/{requestId}/position">client.QueueDetails.Queues.Position.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueueDetailQueuePositionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueueDetailQueuePositionGetResponse">QueueDetailQueuePositionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Exec

## Status

Methods:

- <code title="get /v2/nvcf/exec/status/{requestId}">client.Exec.Status.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ExecStatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ClusterGroups

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ClusterGroupsResponse">ClusterGroupsResponse</a>

Methods:

- <code title="get /v2/nvcf/clusterGroups">client.ClusterGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ClusterGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ClusterGroupsResponse">ClusterGroupsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Clients

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ClientGetResponse">ClientGetResponse</a>

Methods:

- <code title="get /v2/nvcf/clients/{clientId}">client.Clients.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ClientService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ClientGetResponse">ClientGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
