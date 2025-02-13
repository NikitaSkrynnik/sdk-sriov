// Copyright (c) 2020-2021 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package noop provides a NOOP client chain element
package noop

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	"github.com/NikitaSkrynnik/api/pkg/api/networkservice"
	"github.com/NikitaSkrynnik/api/pkg/api/networkservice/mechanisms/cls"
	"github.com/NikitaSkrynnik/api/pkg/api/networkservice/mechanisms/noop"
	"github.com/NikitaSkrynnik/sdk/pkg/networkservice/core/next"
)

type noopClient struct{}

// NewClient returns a NOOP client chain element
func NewClient() networkservice.NetworkServiceClient {
	return new(noopClient)
}

func (c *noopClient) Request(ctx context.Context, request *networkservice.NetworkServiceRequest, opts ...grpc.CallOption) (*networkservice.Connection, error) {
	request.MechanismPreferences = append(request.MechanismPreferences,
		&networkservice.Mechanism{
			Cls:  cls.LOCAL,
			Type: noop.MECHANISM,
		},
		&networkservice.Mechanism{
			Cls:  cls.REMOTE,
			Type: noop.MECHANISM,
		},
	)
	return next.Client(ctx).Request(ctx, request, opts...)
}

func (c *noopClient) Close(ctx context.Context, conn *networkservice.Connection, opts ...grpc.CallOption) (*empty.Empty, error) {
	return next.Client(ctx).Close(ctx, conn, opts...)
}
