// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/nvcf-go"
	"github.com/stainless-sdks/nvcf-go/internal/testutil"
	"github.com/stainless-sdks/nvcf-go/option"
)

func TestUsage(t *testing.T) {
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
	createFunctionResponse, err := client.Functions.New(context.TODO(), nvcf.FunctionNewParams{
		InferenceURL: nvcf.F("https://example.com"),
		Name:         nvcf.F("x"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", createFunctionResponse.Function)
}
