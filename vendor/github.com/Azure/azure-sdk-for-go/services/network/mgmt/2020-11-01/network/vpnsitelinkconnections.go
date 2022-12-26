package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// VpnSiteLinkConnectionsClient is the network Client
type VpnSiteLinkConnectionsClient struct {
	BaseClient
}

// NewVpnSiteLinkConnectionsClient creates an instance of the VpnSiteLinkConnectionsClient client.
func NewVpnSiteLinkConnectionsClient(subscriptionID string) VpnSiteLinkConnectionsClient {
	return NewVpnSiteLinkConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVpnSiteLinkConnectionsClientWithBaseURI creates an instance of the VpnSiteLinkConnectionsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewVpnSiteLinkConnectionsClientWithBaseURI(baseURI string, subscriptionID string) VpnSiteLinkConnectionsClient {
	return VpnSiteLinkConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieves the details of a vpn site link connection.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
// connectionName - the name of the vpn connection.
// linkConnectionName - the name of the vpn connection.
func (client VpnSiteLinkConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string) (result VpnSiteLinkConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSiteLinkConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, gatewayName, connectionName, linkConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSiteLinkConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnSiteLinkConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSiteLinkConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VpnSiteLinkConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":     autorest.Encode("path", connectionName),
		"gatewayName":        autorest.Encode("path", gatewayName),
		"linkConnectionName": autorest.Encode("path", linkConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections/{linkConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VpnSiteLinkConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VpnSiteLinkConnectionsClient) GetResponder(resp *http.Response) (result VpnSiteLinkConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
