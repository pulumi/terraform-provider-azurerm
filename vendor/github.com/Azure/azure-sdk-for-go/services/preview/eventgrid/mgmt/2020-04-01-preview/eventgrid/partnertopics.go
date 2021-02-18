package eventgrid

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

// PartnerTopicsClient is the azure EventGrid Management Client
type PartnerTopicsClient struct {
	BaseClient
}

// NewPartnerTopicsClient creates an instance of the PartnerTopicsClient client.
func NewPartnerTopicsClient(subscriptionID string) PartnerTopicsClient {
	return NewPartnerTopicsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPartnerTopicsClientWithBaseURI creates an instance of the PartnerTopicsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPartnerTopicsClientWithBaseURI(baseURI string, subscriptionID string) PartnerTopicsClient {
	return PartnerTopicsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Activate activate a newly created partner topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerTopicName - name of the partner topic.
func (client PartnerTopicsClient) Activate(ctx context.Context, resourceGroupName string, partnerTopicName string) (result PartnerTopic, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicsClient.Activate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ActivatePreparer(ctx, resourceGroupName, partnerTopicName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Activate", nil, "Failure preparing request")
		return
	}

	resp, err := client.ActivateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Activate", resp, "Failure sending request")
		return
	}

	result, err = client.ActivateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Activate", resp, "Failure responding to request")
		return
	}

	return
}

// ActivatePreparer prepares the Activate request.
func (client PartnerTopicsClient) ActivatePreparer(ctx context.Context, resourceGroupName string, partnerTopicName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partnerTopicName":  autorest.Encode("path", partnerTopicName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/activate", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ActivateSender sends the Activate request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicsClient) ActivateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ActivateResponder handles the response to the Activate request. The method always
// closes the http.Response Body.
func (client PartnerTopicsClient) ActivateResponder(resp *http.Response) (result PartnerTopic, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deactivate deactivate specific partner topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerTopicName - name of the partner topic.
func (client PartnerTopicsClient) Deactivate(ctx context.Context, resourceGroupName string, partnerTopicName string) (result PartnerTopic, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicsClient.Deactivate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeactivatePreparer(ctx, resourceGroupName, partnerTopicName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Deactivate", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeactivateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Deactivate", resp, "Failure sending request")
		return
	}

	result, err = client.DeactivateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Deactivate", resp, "Failure responding to request")
		return
	}

	return
}

// DeactivatePreparer prepares the Deactivate request.
func (client PartnerTopicsClient) DeactivatePreparer(ctx context.Context, resourceGroupName string, partnerTopicName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partnerTopicName":  autorest.Encode("path", partnerTopicName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/deactivate", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeactivateSender sends the Deactivate request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicsClient) DeactivateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeactivateResponder handles the response to the Deactivate request. The method always
// closes the http.Response Body.
func (client PartnerTopicsClient) DeactivateResponder(resp *http.Response) (result PartnerTopic, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete existing partner topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerTopicName - name of the partner topic.
func (client PartnerTopicsClient) Delete(ctx context.Context, resourceGroupName string, partnerTopicName string) (result PartnerTopicsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, partnerTopicName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PartnerTopicsClient) DeletePreparer(ctx context.Context, resourceGroupName string, partnerTopicName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partnerTopicName":  autorest.Encode("path", partnerTopicName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicsClient) DeleteSender(req *http.Request) (future PartnerTopicsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PartnerTopicsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("eventgrid.PartnerTopicsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PartnerTopicsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get properties of a partner topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerTopicName - name of the partner topic.
func (client PartnerTopicsClient) Get(ctx context.Context, resourceGroupName string, partnerTopicName string) (result PartnerTopic, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, partnerTopicName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PartnerTopicsClient) GetPreparer(ctx context.Context, resourceGroupName string, partnerTopicName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partnerTopicName":  autorest.Encode("path", partnerTopicName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PartnerTopicsClient) GetResponder(resp *http.Response) (result PartnerTopic, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup list all the partner topics under a resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client PartnerTopicsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32) (result PartnerTopicsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.ptlr.Response.Response != nil {
				sc = result.ptlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.ptlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.ptlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.ptlr.hasNextLink() && result.ptlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client PartnerTopicsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client PartnerTopicsClient) ListByResourceGroupResponder(resp *http.Response) (result PartnerTopicsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client PartnerTopicsClient) listByResourceGroupNextResults(ctx context.Context, lastResults PartnerTopicsListResult) (result PartnerTopicsListResult, err error) {
	req, err := lastResults.partnerTopicsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client PartnerTopicsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32) (result PartnerTopicsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter, top)
	return
}

// ListBySubscription list all the partner topics under an Azure subscription.
// Parameters:
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client PartnerTopicsClient) ListBySubscription(ctx context.Context, filter string, top *int32) (result PartnerTopicsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.ptlr.Response.Response != nil {
				sc = result.ptlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.ptlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.ptlr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.ptlr.hasNextLink() && result.ptlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client PartnerTopicsClient) ListBySubscriptionPreparer(ctx context.Context, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/partnerTopics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client PartnerTopicsClient) ListBySubscriptionResponder(resp *http.Response) (result PartnerTopicsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client PartnerTopicsClient) listBySubscriptionNextResults(ctx context.Context, lastResults PartnerTopicsListResult) (result PartnerTopicsListResult, err error) {
	req, err := lastResults.partnerTopicsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client PartnerTopicsClient) ListBySubscriptionComplete(ctx context.Context, filter string, top *int32) (result PartnerTopicsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, filter, top)
	return
}

// Update asynchronously updates a partner topic with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerTopicName - name of the partner topic.
// partnerTopicUpdateParameters - partnerTopic update information.
func (client PartnerTopicsClient) Update(ctx context.Context, resourceGroupName string, partnerTopicName string, partnerTopicUpdateParameters PartnerTopicUpdateParameters) (result PartnerTopic, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, partnerTopicName, partnerTopicUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PartnerTopicsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, partnerTopicName string, partnerTopicUpdateParameters PartnerTopicUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partnerTopicName":  autorest.Encode("path", partnerTopicName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}", pathParameters),
		autorest.WithJSON(partnerTopicUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PartnerTopicsClient) UpdateResponder(resp *http.Response) (result PartnerTopic, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
