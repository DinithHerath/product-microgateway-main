/*
 *  Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package routercallbacks

import (
	"context"
	"fmt"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	logger "github.com/wso2/product-microgateway/adapter/internal/loggers"
	"github.com/wso2/product-microgateway/adapter/pkg/logging"
)

const instanceIdentifierKey string = "instanceIdentifier"

// Callbacks is used to debug the xds server related communication.
type Callbacks struct {
}

// Report logs the fetches and requests.
func (cb *Callbacks) Report() {}

// OnStreamOpen prints debug logs
func (cb *Callbacks) OnStreamOpen(_ context.Context, id int64, typ string) error {
	logger.LoggerRouterXdsCallbacks.Debugf("stream %d open for %s\n", id, typ)
	return nil
}

// OnStreamClosed prints debug logs
func (cb *Callbacks) OnStreamClosed(id int64) {
	logger.LoggerRouterXdsCallbacks.Debugf("stream %d closed\n", id)
}

// OnStreamRequest prints debug logs
func (cb *Callbacks) OnStreamRequest(id int64, request *discovery.DiscoveryRequest) error {
	nodeIdentifier := getNodeIdentifier(request)
	logger.LoggerRouterXdsCallbacks.Debugf("stream request on stream id: %d, from node: %s, version: %s, for type: %s",
		id, nodeIdentifier, request.VersionInfo, request.TypeUrl)
	if request.ErrorDetail != nil {
		logger.LoggerRouterXdsCallbacks.ErrorC(logging.ErrorDetails{
			Message: fmt.Sprintf("Stream request for type %s on stream id: %d, from node: %s, Error: %s", request.GetTypeUrl(),
				id, nodeIdentifier, request.ErrorDetail.Message),
			Severity:  logging.CRITICAL,
			ErrorCode: 1401,
		})
	}
	return nil
}

// OnStreamResponse prints debug logs
func (cb *Callbacks) OnStreamResponse(context context.Context, id int64, request *discovery.DiscoveryRequest, response *discovery.DiscoveryResponse) {
	nodeIdentifier := getNodeIdentifier(request)
	logger.LoggerRouterXdsCallbacks.Debugf("stream response on stream id: %d, to node: %s, version: %s, for type: %v", id,
		nodeIdentifier, response.VersionInfo, response.TypeUrl)
}

// OnFetchRequest prints debug logs
func (cb *Callbacks) OnFetchRequest(_ context.Context, req *discovery.DiscoveryRequest) error {
	logger.LoggerRouterXdsCallbacks.Debugf("fetch request from node %s, version: %s, for type %s", req.Node.Id, req.VersionInfo, req.TypeUrl)
	return nil
}

// OnFetchResponse prints debug logs
func (cb *Callbacks) OnFetchResponse(req *discovery.DiscoveryRequest, res *discovery.DiscoveryResponse) {
	logger.LoggerRouterXdsCallbacks.Debugf("fetch response to node: %s, version: %s, for type %s", req.Node.Id, req.VersionInfo, res.TypeUrl)
}

// OnDeltaStreamOpen is unused.
func (cb *Callbacks) OnDeltaStreamOpen(_ context.Context, id int64, typ string) error {
	return nil
}

// OnDeltaStreamClosed is unused.
func (cb *Callbacks) OnDeltaStreamClosed(id int64) {
}

// OnStreamDeltaResponse is unused.
func (cb *Callbacks) OnStreamDeltaResponse(id int64, req *discovery.DeltaDiscoveryRequest, res *discovery.DeltaDiscoveryResponse) {
}

// OnStreamDeltaRequest is unused.
func (cb *Callbacks) OnStreamDeltaRequest(id int64, req *discovery.DeltaDiscoveryRequest) error {
	return nil
}

func getNodeIdentifier(request *discovery.DiscoveryRequest) string {
	metadataMap := request.Node.Metadata.AsMap()
	nodeIdentifier := request.Node.Id
	if identifierVal, ok := metadataMap[instanceIdentifierKey]; ok {
		nodeIdentifier = request.Node.Id + ":" + identifierVal.(string)
	}
	return nodeIdentifier
}
