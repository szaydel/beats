// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package kubernetes

import (
	"github.com/gofrs/uuid"
	k8s "k8s.io/client-go/kubernetes"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/bus"
)

type service struct{}

// NewServiceEventer creates an eventer that can discover and process service objects
func NewServiceEventer(uuid uuid.UUID, cfg *common.Config, client k8s.Interface, publish func(event []bus.Event)) (Eventer, error) {
	return nil, nil
}

// OnAdd ensures processing of service objects that are newly created
func (s *service) OnAdd(obj interface{}) {}

// OnUpdate ensures processing of service objects that are updated
func (s *service) OnUpdate(obj interface{}) {}

// OnDelete ensures processing of service objects that are deleted
func (s *service) OnDelete(obj interface{}) {}

// GenerateHints creates hints needed for hints builder
func (s *service) GenerateHints(event bus.Event) bus.Event {
	return event
}

// Start starts the eventer
func (s *service) Start() error {
	return nil
}

// Stop stops the eventer
func (s *service) Stop() {}
