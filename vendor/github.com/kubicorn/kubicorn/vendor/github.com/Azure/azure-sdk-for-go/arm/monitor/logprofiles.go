package monitor

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// LogProfilesClient is the monitor Management Client
type LogProfilesClient struct {
	ManagementClient
}

// NewLogProfilesClient creates an instance of the LogProfilesClient client.
func NewLogProfilesClient(subscriptionID string) LogProfilesClient {
	return NewLogProfilesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLogProfilesClientWithBaseURI creates an instance of the LogProfilesClient client.
func NewLogProfilesClientWithBaseURI(baseURI string, subscriptionID string) LogProfilesClient {
	return LogProfilesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a log profile in Azure Monitoring REST API.
//
// logProfileName is the name of the log profile. parameters is parameters supplied to the operation.
func (client LogProfilesClient) CreateOrUpdate(logProfileName string, parameters LogProfileResource) (result LogProfileResource, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.LogProfileProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.LogProfileProperties.Locations", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.LogProfileProperties.Categories", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.LogProfileProperties.RetentionPolicy", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.LogProfileProperties.RetentionPolicy.Enabled", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.LogProfileProperties.RetentionPolicy.Days", Name: validation.Null, Rule: true,
								Chain: []validation.Constraint{{Target: "parameters.LogProfileProperties.RetentionPolicy.Days", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}},
						}},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "monitor.LogProfilesClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(logProfileName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client LogProfilesClient) CreateOrUpdatePreparer(logProfileName string, parameters LogProfileResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"logProfileName": autorest.Encode("path", logProfileName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/microsoft.insights/logprofiles/{logProfileName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client LogProfilesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client LogProfilesClient) CreateOrUpdateResponder(resp *http.Response) (result LogProfileResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the log profile.
//
// logProfileName is the name of the log profile.
func (client LogProfilesClient) Delete(logProfileName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(logProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LogProfilesClient) DeletePreparer(logProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"logProfileName": autorest.Encode("path", logProfileName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/microsoft.insights/logprofiles/{logProfileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LogProfilesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LogProfilesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the log profile.
//
// logProfileName is the name of the log profile.
func (client LogProfilesClient) Get(logProfileName string) (result LogProfileResource, err error) {
	req, err := client.GetPreparer(logProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client LogProfilesClient) GetPreparer(logProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"logProfileName": autorest.Encode("path", logProfileName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/microsoft.insights/logprofiles/{logProfileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LogProfilesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LogProfilesClient) GetResponder(resp *http.Response) (result LogProfileResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list the log profiles.
func (client LogProfilesClient) List() (result LogProfileCollection, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client LogProfilesClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/microsoft.insights/logprofiles", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LogProfilesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LogProfilesClient) ListResponder(resp *http.Response) (result LogProfileCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing LogProfilesResource. To update other fields use the CreateOrUpdate method.
//
// logProfileName is the name of the log profile. logProfilesResource is parameters supplied to the operation.
func (client LogProfilesClient) Update(logProfileName string, logProfilesResource LogProfileResourcePatch) (result LogProfileResource, err error) {
	req, err := client.UpdatePreparer(logProfileName, logProfilesResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.LogProfilesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client LogProfilesClient) UpdatePreparer(logProfileName string, logProfilesResource LogProfileResourcePatch) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"logProfileName": autorest.Encode("path", logProfileName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/microsoft.insights/logprofiles/{logProfileName}", pathParameters),
		autorest.WithJSON(logProfilesResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client LogProfilesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client LogProfilesClient) UpdateResponder(resp *http.Response) (result LogProfileResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
