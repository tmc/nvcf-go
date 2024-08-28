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

func TestFunctionDeploymentFunctionVersionDeleteDeploymentWithOptionalParams(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.DeleteDeployment(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.FunctionDeploymentFunctionVersionDeleteDeploymentParams{
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

func TestFunctionDeploymentFunctionVersionInitiateDeployment(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.InitiateDeployment(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.FunctionDeploymentFunctionVersionInitiateDeploymentParams{
			DeploymentSpecifications: nvcf.F([]nvcf.FunctionDeploymentFunctionVersionInitiateDeploymentParamsDeploymentSpecification{{
				GPU:                   nvcf.F("gpu"),
				InstanceType:          nvcf.F("instanceType"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				Attributes:            nvcf.F([]string{"string"}),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				Backend:               nvcf.F("backend"),
				Clusters:              nvcf.F([]string{"string"}),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				PreferredOrder:        nvcf.F(int64(1)),
				Regions:               nvcf.F([]string{"string"}),
			}, {
				GPU:                   nvcf.F("gpu"),
				InstanceType:          nvcf.F("instanceType"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				Attributes:            nvcf.F([]string{"string"}),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				Backend:               nvcf.F("backend"),
				Clusters:              nvcf.F([]string{"string"}),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				PreferredOrder:        nvcf.F(int64(1)),
				Regions:               nvcf.F([]string{"string"}),
			}, {
				GPU:                   nvcf.F("gpu"),
				InstanceType:          nvcf.F("instanceType"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				Attributes:            nvcf.F([]string{"string"}),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				Backend:               nvcf.F("backend"),
				Clusters:              nvcf.F([]string{"string"}),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				PreferredOrder:        nvcf.F(int64(1)),
				Regions:               nvcf.F([]string{"string"}),
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

func TestFunctionDeploymentFunctionVersionGetDeployment(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.GetDeployment(
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

func TestFunctionDeploymentFunctionVersionUpdateDeployment(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.UpdateDeployment(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.FunctionDeploymentFunctionVersionUpdateDeploymentParams{
			DeploymentSpecifications: nvcf.F([]nvcf.FunctionDeploymentFunctionVersionUpdateDeploymentParamsDeploymentSpecification{{
				GPU:                   nvcf.F("gpu"),
				InstanceType:          nvcf.F("instanceType"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				Attributes:            nvcf.F([]string{"string"}),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				Backend:               nvcf.F("backend"),
				Clusters:              nvcf.F([]string{"string"}),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				PreferredOrder:        nvcf.F(int64(1)),
				Regions:               nvcf.F([]string{"string"}),
			}, {
				GPU:                   nvcf.F("gpu"),
				InstanceType:          nvcf.F("instanceType"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				Attributes:            nvcf.F([]string{"string"}),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				Backend:               nvcf.F("backend"),
				Clusters:              nvcf.F([]string{"string"}),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				PreferredOrder:        nvcf.F(int64(1)),
				Regions:               nvcf.F([]string{"string"}),
			}, {
				GPU:                   nvcf.F("gpu"),
				InstanceType:          nvcf.F("instanceType"),
				MaxInstances:          nvcf.F(int64(0)),
				MinInstances:          nvcf.F(int64(0)),
				Attributes:            nvcf.F([]string{"string"}),
				AvailabilityZones:     nvcf.F([]string{"string", "string", "string"}),
				Backend:               nvcf.F("backend"),
				Clusters:              nvcf.F([]string{"string"}),
				Configuration:         nvcf.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvcf.F(int64(1)),
				PreferredOrder:        nvcf.F(int64(1)),
				Regions:               nvcf.F([]string{"string"}),
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
