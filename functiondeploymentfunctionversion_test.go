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

func TestFunctionDeploymentFunctionVersionNew(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.FunctionDeploymentFunctionVersionNewParams{
			DeploymentSpecifications: nvcf.F([]nvcf.FunctionDeploymentFunctionVersionNewParamsDeploymentSpecification{{
				GPU:                   nvcf.F("gpu"),
				Backend:               nvcf.F("backend"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				InstanceType:          nvcf.F("instanceType"),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
			}, {
				GPU:                   nvcf.F("gpu"),
				Backend:               nvcf.F("backend"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				InstanceType:          nvcf.F("instanceType"),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
			}, {
				GPU:                   nvcf.F("gpu"),
				Backend:               nvcf.F("backend"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				InstanceType:          nvcf.F("instanceType"),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
			}}),
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

func TestFunctionDeploymentFunctionVersionGet(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.Get(
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

func TestFunctionDeploymentFunctionVersionUpdate(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.FunctionDeploymentFunctionVersionUpdateParams{
			DeploymentSpecifications: nvcf.F([]nvcf.FunctionDeploymentFunctionVersionUpdateParamsDeploymentSpecification{{
				GPU:                   nvcf.F("gpu"),
				Backend:               nvcf.F("backend"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				InstanceType:          nvcf.F("instanceType"),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
			}, {
				GPU:                   nvcf.F("gpu"),
				Backend:               nvcf.F("backend"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				InstanceType:          nvcf.F("instanceType"),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
			}, {
				GPU:                   nvcf.F("gpu"),
				Backend:               nvcf.F("backend"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				InstanceType:          nvcf.F("instanceType"),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
			}}),
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

func TestFunctionDeploymentFunctionVersionDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.FunctionDeploymentFunctionVersionDeleteParams{
			Graceful: nvcf.F(true),
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
