// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package visualstudio

import original "github.com/Azure/azure-sdk-for-go/services/visualstudio/mgmt/2014-04-01-preview/visualstudio"

type AccountsClient = original.AccountsClient

func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type ExtensionsClient = original.ExtensionsClient

func NewExtensionsClient(subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClient(subscriptionID)
}
func NewExtensionsClientWithBaseURI(baseURI string, subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClientWithBaseURI(baseURI, subscriptionID)
}

type AccountResource = original.AccountResource
type AccountResourceListResult = original.AccountResourceListResult
type AccountResourceRequest = original.AccountResourceRequest
type CheckNameAvailabilityParameter = original.CheckNameAvailabilityParameter
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type ExtensionResource = original.ExtensionResource
type ExtensionResourceListResult = original.ExtensionResourceListResult
type ExtensionResourcePlan = original.ExtensionResourcePlan
type ExtensionResourceRequest = original.ExtensionResourceRequest
type Operation = original.Operation
type OperationListResult = original.OperationListResult
type OperationProperties = original.OperationProperties
type ProjectResource = original.ProjectResource
type ProjectResourceListResult = original.ProjectResourceListResult
type ProjectsCreateFuture = original.ProjectsCreateFuture
type Resource = original.Resource
type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}

type ProjectsClient = original.ProjectsClient

func NewProjectsClient(subscriptionID string) ProjectsClient {
	return original.NewProjectsClient(subscriptionID)
}
func NewProjectsClientWithBaseURI(baseURI string, subscriptionID string) ProjectsClient {
	return original.NewProjectsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
