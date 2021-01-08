package servicebus

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PremiumMessagingRegionsClient is the azure Service Bus client
type PremiumMessagingRegionsClient struct {
	BaseClient
}

// NewPremiumMessagingRegionsClient creates an instance of the PremiumMessagingRegionsClient client.
func NewPremiumMessagingRegionsClient(subscriptionID string) PremiumMessagingRegionsClient {
	return NewPremiumMessagingRegionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPremiumMessagingRegionsClientWithBaseURI creates an instance of the PremiumMessagingRegionsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewPremiumMessagingRegionsClientWithBaseURI(baseURI string, subscriptionID string) PremiumMessagingRegionsClient {
	return PremiumMessagingRegionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets the available premium messaging regions for servicebus
func (client PremiumMessagingRegionsClient) List(ctx context.Context) (result PremiumMessagingRegionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PremiumMessagingRegionsClient.List")
		defer func() {
			sc := -1
			if result.pmrlr.Response.Response != nil {
				sc = result.pmrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pmrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "List", resp, "Failure sending request")
		return
	}

	result.pmrlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.pmrlr.hasNextLink() && result.pmrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client PremiumMessagingRegionsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/premiumMessagingRegions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PremiumMessagingRegionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PremiumMessagingRegionsClient) ListResponder(resp *http.Response) (result PremiumMessagingRegionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PremiumMessagingRegionsClient) listNextResults(ctx context.Context, lastResults PremiumMessagingRegionsListResult) (result PremiumMessagingRegionsListResult, err error) {
	req, err := lastResults.premiumMessagingRegionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PremiumMessagingRegionsClient) ListComplete(ctx context.Context) (result PremiumMessagingRegionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PremiumMessagingRegionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
