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

func TestAuthorizedAccountFunctionAddWithOptionalParams(t *testing.T) {
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
	_, err := client.AuthorizedAccounts.Functions.Add(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.AuthorizedAccountFunctionAddParams{
			AuthorizedParty: nvcf.F(nvcf.AuthorizedAccountFunctionAddParamsAuthorizedParty{
				NcaID:    nvcf.F("ncaId"),
				ClientID: nvcf.F("clientId"),
			}),
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

func TestAuthorizedAccountFunctionRemoveWithOptionalParams(t *testing.T) {
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
	_, err := client.AuthorizedAccounts.Functions.Remove(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.AuthorizedAccountFunctionRemoveParams{
			AuthorizedParty: nvcf.F(nvcf.AuthorizedAccountFunctionRemoveParamsAuthorizedParty{
				NcaID:    nvcf.F("ncaId"),
				ClientID: nvcf.F("clientId"),
			}),
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
