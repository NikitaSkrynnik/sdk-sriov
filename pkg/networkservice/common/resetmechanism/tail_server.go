// Copyright (c) 2020 Doc.ai and/or its affiliates.
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

package resetmechanism

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/NikitaSkrynnik/api/pkg/api/networkservice"
)

type tailServer struct{}

func (s *tailServer) Request(_ context.Context, request *networkservice.NetworkServiceRequest) (*networkservice.Connection, error) {
	return request.GetConnection(), nil
}

func (s *tailServer) Close(_ context.Context, _ *networkservice.Connection) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
