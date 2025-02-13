// Copyright (c) 2020-2022 Doc.ai and/or its affiliates.
//
// Copyright (c) 2021-2022 Nordix Foundation.
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

// Package token provides chain elements for inserting SRIOV tokens into request and response
package token

import (
	"github.com/NikitaSkrynnik/api/pkg/api/networkservice"

	"github.com/NikitaSkrynnik/sdk-sriov/pkg/networkservice/common/token/multitoken"
)

// NewClient returns a new token client chain element
func NewClient() networkservice.NetworkServiceClient {
	return multitoken.NewClient()
}
