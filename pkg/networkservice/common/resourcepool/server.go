// Copyright (c) 2020-2022 Doc.ai and/or its affiliates.
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

//go:build linux
// +build linux

// Package resourcepool provides chain elements for to select and free VF
package resourcepool

import (
	"context"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"

	"github.com/NikitaSkrynnik/api/pkg/api/networkservice"
	"github.com/NikitaSkrynnik/api/pkg/api/networkservice/mechanisms/common"
	"github.com/NikitaSkrynnik/sdk-kernel/pkg/kernel/networkservice/vfconfig"
	"github.com/NikitaSkrynnik/sdk/pkg/networkservice/core/next"
	"github.com/NikitaSkrynnik/sdk/pkg/networkservice/utils/metadata"
	"github.com/NikitaSkrynnik/sdk/pkg/tools/log"

	"github.com/NikitaSkrynnik/sdk-sriov/pkg/sriov"
	"github.com/NikitaSkrynnik/sdk-sriov/pkg/sriov/config"
	"github.com/NikitaSkrynnik/sdk-sriov/pkg/tools/tokens"
)

type resourcePoolServer struct {
	resourcePool *resourcePoolConfig
}

// NewServer returns a new resource pool server chain element
func NewServer(
	driverType sriov.DriverType,
	resourceLock sync.Locker,
	pciPool PCIPool,
	resourcePool ResourcePool,
	cfg *config.Config,
) networkservice.NetworkServiceServer {
	return &resourcePoolServer{resourcePool: &resourcePoolConfig{
		driverType:   driverType,
		resourceLock: resourceLock,
		pciPool:      pciPool,
		resourcePool: resourcePool,
		config:       cfg,
		selectedVFs:  map[string]string{},
	}}
}

func (s *resourcePoolServer) Request(ctx context.Context, request *networkservice.NetworkServiceRequest) (*networkservice.Connection, error) {
	logger := log.FromContext(ctx).WithField("resourcePoolServer", "Request")
	conn := request.GetConnection()
	tokenID, ok := conn.GetMechanism().GetParameters()[common.DeviceTokenIDKey]
	if !ok {
		return nil, errors.New("no token ID provided")
	}
	if !tokens.IsTokenID(tokenID) {
		return nil, errors.Errorf("no SR-IOV token ID provided, got: %s", tokenID)
	}

	_, vfExists := vfconfig.Load(ctx, metadata.IsClient(s))

	if !vfExists {
		err := assignVF(ctx, logger, conn, tokenID, s.resourcePool, metadata.IsClient(s))
		if err != nil {
			_ = s.resourcePool.close(conn)
			return nil, err
		}
	}

	conn, err := next.Server(ctx).Request(ctx, request)
	if err != nil && !vfExists {
		if closeErr := s.resourcePool.close(conn); closeErr != nil {
			err = errors.Wrapf(err, "connection closed with error: %s", closeErr.Error())
		}
		return nil, err
	}

	return conn, err
}

func (s *resourcePoolServer) Close(ctx context.Context, conn *networkservice.Connection) (*empty.Empty, error) {
	_, err := next.Server(ctx).Close(ctx, conn)

	closeErr := s.resourcePool.close(conn)

	if err != nil && closeErr != nil {
		return nil, errors.Wrapf(err, "failed to free VF: %v", closeErr)
	}
	if closeErr != nil {
		return nil, errors.Wrap(closeErr, "failed to free VF")
	}
	return &empty.Empty{}, err
}
