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

func TestFunctionVersionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Functions.Versions.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.FunctionVersionNewParams{
			InferenceURL:  nvcf.F("https://example.com"),
			Name:          nvcf.F("x"),
			APIBodyFormat: nvcf.F(nvcf.FunctionVersionNewParamsAPIBodyFormatPredictV2),
			ContainerArgs: nvcf.F("containerArgs"),
			ContainerEnvironment: nvcf.F([]nvcf.FunctionVersionNewParamsContainerEnvironment{{
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
			FunctionType:   nvcf.F(nvcf.FunctionVersionNewParamsFunctionTypeDefault),
			Health: nvcf.F(nvcf.FunctionVersionNewParamsHealth{
				ExpectedStatusCode: nvcf.F(int64(0)),
				Port:               nvcf.F(int64(0)),
				Protocol:           nvcf.F(nvcf.FunctionVersionNewParamsHealthProtocolHTTP),
				Timeout:            nvcf.F("PT10S"),
				Uri:                nvcf.F("https://example.com"),
			}),
			HealthUri:            nvcf.F("https://example.com"),
			HelmChart:            nvcf.F("https://example.com"),
			HelmChartServiceName: nvcf.F("helmChartServiceName"),
			InferencePort:        nvcf.F(int64(0)),
			Models: nvcf.F([]nvcf.FunctionVersionNewParamsModel{{
				Name:    nvcf.F("name"),
				Uri:     nvcf.F("https://example.com"),
				Version: nvcf.F("version"),
			}}),
			Resources: nvcf.F([]nvcf.FunctionVersionNewParamsResource{{
				Name:    nvcf.F("name"),
				Uri:     nvcf.F("https://example.com"),
				Version: nvcf.F("version"),
			}}),
			Secrets: nvcf.F([]nvcf.FunctionVersionNewParamsSecret{{
				Name:  nvcf.F("x"),
				Value: nvcf.F("x"),
			}}),
			Tags: nvcf.F([]string{"string"}),
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

func TestFunctionVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Functions.Versions.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.FunctionVersionGetParams{
			IncludeSecrets: nvcf.F(true),
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

func TestFunctionVersionList(t *testing.T) {
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
	_, err := client.Functions.Versions.List(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *nvcf.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunctionVersionDelete(t *testing.T) {
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
	err := client.Functions.Versions.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *nvcf.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
