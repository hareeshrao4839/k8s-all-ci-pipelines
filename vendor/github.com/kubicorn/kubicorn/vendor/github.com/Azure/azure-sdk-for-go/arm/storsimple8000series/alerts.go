package storsimple8000series

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

// AlertsClient is the client for the Alerts methods of the Storsimple8000series service.
type AlertsClient struct {
	ManagementClient
}

// NewAlertsClient creates an instance of the AlertsClient client.
func NewAlertsClient(subscriptionID string) AlertsClient {
	return NewAlertsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAlertsClientWithBaseURI creates an instance of the AlertsClient client.
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return AlertsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Clear clear the alerts.
//
// parameters is the clear alert request. resourceGroupName is the resource group name managerName is the manager name
func (client AlertsClient) Clear(parameters ClearAlertRequest, resourceGroupName string, managerName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Alerts", Name: validation.Null, Rule: true, Chain: nil}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple8000series.AlertsClient", "Clear")
	}

	req, err := client.ClearPreparer(parameters, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "Clear", nil, "Failure preparing request")
		return
	}

	resp, err := client.ClearSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "Clear", resp, "Failure sending request")
		return
	}

	result, err = client.ClearResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "Clear", resp, "Failure responding to request")
	}

	return
}

// ClearPreparer prepares the Clear request.
func (client AlertsClient) ClearPreparer(parameters ClearAlertRequest, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/clearAlerts", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ClearSender sends the Clear request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) ClearSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ClearResponder handles the response to the Clear request. The method always
// closes the http.Response Body.
func (client AlertsClient) ClearResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByManager retrieves all the alerts in a manager.
//
// resourceGroupName is the resource group name managerName is the manager name filter is oData Filter options
func (client AlertsClient) ListByManager(resourceGroupName string, managerName string, filter string) (result AlertList, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple8000series.AlertsClient", "ListByManager")
	}

	req, err := client.ListByManagerPreparer(resourceGroupName, managerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "ListByManager", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "ListByManager", resp, "Failure sending request")
		return
	}

	result, err = client.ListByManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "ListByManager", resp, "Failure responding to request")
	}

	return
}

// ListByManagerPreparer prepares the ListByManager request.
func (client AlertsClient) ListByManagerPreparer(resourceGroupName string, managerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/alerts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByManagerSender sends the ListByManager request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) ListByManagerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByManagerResponder handles the response to the ListByManager request. The method always
// closes the http.Response Body.
func (client AlertsClient) ListByManagerResponder(resp *http.Response) (result AlertList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByManagerNextResults retrieves the next set of results, if any.
func (client AlertsClient) ListByManagerNextResults(lastResults AlertList) (result AlertList, err error) {
	req, err := lastResults.AlertListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "ListByManager", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByManagerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "ListByManager", resp, "Failure sending next results request")
	}

	result, err = client.ListByManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "ListByManager", resp, "Failure responding to next results request")
	}

	return
}

// ListByManagerComplete gets all elements from the list without paging.
func (client AlertsClient) ListByManagerComplete(resourceGroupName string, managerName string, filter string, cancel <-chan struct{}) (<-chan Alert, <-chan error) {
	resultChan := make(chan Alert)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByManager(resourceGroupName, managerName, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByManagerNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// SendTestEmail sends a test alert email.
//
// deviceName is the device name parameters is the send test alert email request. resourceGroupName is the resource
// group name managerName is the manager name
func (client AlertsClient) SendTestEmail(deviceName string, parameters SendTestAlertEmailRequest, resourceGroupName string, managerName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.EmailList", Name: validation.Null, Rule: true, Chain: nil}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple8000series.AlertsClient", "SendTestEmail")
	}

	req, err := client.SendTestEmailPreparer(deviceName, parameters, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "SendTestEmail", nil, "Failure preparing request")
		return
	}

	resp, err := client.SendTestEmailSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "SendTestEmail", resp, "Failure sending request")
		return
	}

	result, err = client.SendTestEmailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.AlertsClient", "SendTestEmail", resp, "Failure responding to request")
	}

	return
}

// SendTestEmailPreparer prepares the SendTestEmail request.
func (client AlertsClient) SendTestEmailPreparer(deviceName string, parameters SendTestAlertEmailRequest, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/sendTestAlertEmail", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// SendTestEmailSender sends the SendTestEmail request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) SendTestEmailSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// SendTestEmailResponder handles the response to the SendTestEmail request. The method always
// closes the http.Response Body.
func (client AlertsClient) SendTestEmailResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
