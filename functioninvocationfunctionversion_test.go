// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/tmc/nvcf-go"
	"github.com/tmc/nvcf-go/internal/testutil"
	"github.com/tmc/nvcf-go/option"
)

func TestFunctionInvocationFunctionVersionInvokeWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nvcf.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.FunctionInvocation.Functions.Versions.Invoke(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.FunctionInvocationFunctionVersionInvokeParams{
			Body:                     map[string]interface{}{},
			NvcfInputAssetReferences: nvcf.F([]string{"string", "string", "string"}),
			NvcfPollSeconds:          nvcf.F(int64(0)),
		},
	)
	if err != nil {
		var apierr *nvcf.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
