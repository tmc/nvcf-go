// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/nvcf-go"
	"github.com/stainless-sdks/nvcf-go/internal/testutil"
	"github.com/stainless-sdks/nvcf-go/option"
)

func TestAssetManagementAssetNew(t *testing.T) {
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
	_, err := client.AssetManagement.Assets.New(context.TODO(), nvcf.AssetManagementAssetNewParams{
		ContentType: nvcf.F("contentType"),
		Description: nvcf.F("description"),
	})
	if err != nil {
		var apierr *nvcf.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
