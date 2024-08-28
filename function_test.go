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

func TestFunctionNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nvcf.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Functions.New(context.TODO(), nvcf.FunctionNewParams{
		InferenceURL:  nvcf.F("https://example.com"),
		Name:          nvcf.F("x"),
		APIBodyFormat: nvcf.F(nvcf.FunctionNewParamsAPIBodyFormatPredictV2),
		ContainerArgs: nvcf.F("containerArgs"),
		ContainerEnvironment: nvcf.F([]nvcf.FunctionNewParamsContainerEnvironment{{
			Key:   nvcf.F("key"),
			Value: nvcf.F("value"),
		}, {
			Key:   nvcf.F("key"),
			Value: nvcf.F("value"),
		}, {
			Key:   nvcf.F("key"),
			Value: nvcf.F("value"),
		}}),
		ContainerImage: nvcf.F("https://example.com"),
		Description:    nvcf.F("description"),
		FunctionType:   nvcf.F(nvcf.FunctionNewParamsFunctionTypeDefault),
		Health: nvcf.F(nvcf.FunctionNewParamsHealth{
			ExpectedStatusCode: nvcf.F(int64(0)),
			Port:               nvcf.F(int64(0)),
			Protocol:           nvcf.F(nvcf.FunctionNewParamsHealthProtocolHTTP),
			Timeout:            nvcf.F("PT10S"),
			Uri:                nvcf.F("https://example.com"),
		}),
		HealthUri:            nvcf.F("https://example.com"),
		HelmChart:            nvcf.F("https://example.com"),
		HelmChartServiceName: nvcf.F("helmChartServiceName"),
		InferencePort:        nvcf.F(int64(0)),
		Models: nvcf.F([]nvcf.FunctionNewParamsModel{{
			Name:    nvcf.F("name"),
			Uri:     nvcf.F("https://example.com"),
			Version: nvcf.F("version"),
		}}),
		Resources: nvcf.F([]nvcf.FunctionNewParamsResource{{
			Name:    nvcf.F("name"),
			Uri:     nvcf.F("https://example.com"),
			Version: nvcf.F("version"),
		}}),
		Secrets: nvcf.F([]nvcf.FunctionNewParamsSecret{{
			Name:  nvcf.F("x"),
			Value: nvcf.F("x"),
		}}),
		Tags: nvcf.F([]string{"string"}),
	})
	if err != nil {
		var apierr *nvcf.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunctionListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nvcf.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Functions.List(context.TODO(), nvcf.FunctionListParams{
		Visibility: nvcf.F([]nvcf.FunctionListParamsVisibility{nvcf.FunctionListParamsVisibilityAuthorized}),
	})
	if err != nil {
		var apierr *nvcf.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
