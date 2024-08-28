# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#FunctionResponse">FunctionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#GetQueuesResponse">GetQueuesResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>

# UserSecretManagement

## Functions

### Versions

Methods:

- <code title="put /v2/nvcf/secrets/functions/{functionId}/versions/{versionId}">client.UserSecretManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#UserSecretManagementFunctionVersionService.UpdateSecrets">UpdateSecrets</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#UserSecretManagementFunctionVersionUpdateSecretsParams">UserSecretManagementFunctionVersionUpdateSecretsParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# FunctionManagement

## Functions

### Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#CreateFunctionResponse">CreateFunctionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListFunctionsResponse">ListFunctionsResponse</a>

Methods:

- <code title="put /v2/nvcf/metadata/functions/{functionId}/versions/{functionVersionId}">client.FunctionManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionManagementFunctionVersionService.UpdateMetadata">UpdateMetadata</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionManagementFunctionVersionUpdateMetadataParams">FunctionManagementFunctionVersionUpdateMetadataParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## IDs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListFunctionIDsResponse">ListFunctionIDsResponse</a>

# FunctionDeployment

## Functions

### Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#DeploymentResponse">DeploymentResponse</a>

Methods:

- <code title="delete /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionService.DeleteDeployment">DeleteDeployment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionDeleteDeploymentParams">FunctionDeploymentFunctionVersionDeleteDeploymentParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionService.InitiateDeployment">InitiateDeployment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionInitiateDeploymentParams">FunctionDeploymentFunctionVersionInitiateDeploymentParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionService.GetDeployment">GetDeployment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionService.UpdateDeployment">UpdateDeployment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionDeploymentFunctionVersionUpdateDeploymentParams">FunctionDeploymentFunctionVersionUpdateDeploymentParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

# EnvelopeFunctionInvocation

## Functions

Methods:

- <code title="post /v2/nvcf/exec/functions/{functionId}">client.EnvelopeFunctionInvocation.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#EnvelopeFunctionInvocationFunctionService.InvokeEnvelope">InvokeEnvelope</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#EnvelopeFunctionInvocationFunctionInvokeEnvelopeParams">EnvelopeFunctionInvocationFunctionInvokeEnvelopeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="post /v2/nvcf/exec/functions/{functionId}/versions/{versionId}">client.EnvelopeFunctionInvocation.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#EnvelopeFunctionInvocationFunctionVersionService.InvokeEnvelope">InvokeEnvelope</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#EnvelopeFunctionInvocationFunctionVersionInvokeEnvelopeParams">EnvelopeFunctionInvocationFunctionVersionInvokeEnvelopeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Functions

Methods:

- <code title="post /v2/nvcf/functions">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionNewParams">FunctionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#CreateFunctionResponse">CreateFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/functions">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionListParams">FunctionListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListFunctionsResponse">ListFunctionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Versions

Methods:

- <code title="post /v2/nvcf/functions/{functionId}/versions">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionVersionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionVersionNewParams">FunctionVersionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#CreateFunctionResponse">CreateFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/functions/{functionId}/versions/{functionVersionId}">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/functions/{functionId}/versions">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListFunctionsResponse">ListFunctionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/functions/{functionId}/versions/{functionVersionId}">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## IDs

Methods:

- <code title="get /v2/nvcf/functions/ids">client.Functions.IDs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionIDService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#FunctionIDListParams">FunctionIDListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListFunctionIDsResponse">ListFunctionIDsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AuthorizedAccounts

## Functions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListAuthorizedPartiesResponse">ListAuthorizedPartiesResponse</a>

Methods:

- <code title="get /v2/nvcf/authorizations/functions/{functionId}">client.AuthorizedAccounts.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListAuthorizedPartiesResponse">ListAuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/authorizations/functions/{functionId}">client.AuthorizedAccounts.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/nvcf/authorizations/functions/{functionId}">client.AuthorizedAccounts.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionAuthorizeParams">AuthorizedAccountFunctionAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="get /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.AuthorizedAccounts.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.AuthorizedAccounts.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.AuthorizedAccounts.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionVersionService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizedAccountFunctionVersionAuthorizeParams">AuthorizedAccountFunctionVersionAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetResponse">AssetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#CreateAssetResponse">CreateAssetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListAssetsResponse">ListAssetsResponse</a>

Methods:

- <code title="post /v2/nvcf/assets">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetNewParams">AssetNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#CreateAssetResponse">CreateAssetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/assets/{assetId}">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetResponse">AssetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/assets">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#ListAssetsResponse">ListAssetsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/assets/{assetId}">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AssetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Authorizations

## Functions

Methods:

- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/add">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionAddParams">AuthorizationFunctionAddParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/remove">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionRemoveParams">AuthorizationFunctionRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}/add">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionVersionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionVersionAddParams">AuthorizationFunctionVersionAddParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}/remove">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionVersionService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#AuthorizationFunctionVersionRemoveParams">AuthorizationFunctionVersionRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Queues

## Functions

Methods:

- <code title="get /v2/nvcf/queues/functions/{functionId}">client.Queues.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueueFunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#GetQueuesResponse">GetQueuesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="get /v2/nvcf/queues/functions/{functionId}/versions/{versionId}">client.Queues.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueueFunctionVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go/shared#GetQueuesResponse">GetQueuesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Position

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#GetPositionInQueueResponse">GetPositionInQueueResponse</a>

Methods:

- <code title="get /v2/nvcf/queues/{requestId}/position">client.Queues.Position.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#QueuePositionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#GetPositionInQueueResponse">GetPositionInQueueResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Pexec

## Status

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#PexecStatusGetResponse">PexecStatusGetResponse</a>

Methods:

- <code title="get /v2/nvcf/pexec/status/{requestId}">client.Pexec.Status.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#PexecStatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#PexecStatusGetParams">PexecStatusGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvcf-go#PexecStatusGetResponse">PexecStatusGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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
