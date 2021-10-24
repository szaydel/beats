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

type pod struct{}

// NewPodEventer creates an eventer that can discover and process pod objects
func NewPodEventer(uuid uuid.UUID, cfg *common.Config, client k8s.Interface, publish func(event []bus.Event)) (Eventer, error) {
	return nil, nil
}

// OnAdd ensures processing of pod objects that are newly added.
func (p *pod) OnAdd(obj interface{}) {
}

// OnUpdate handles events for pods that have been updated.
func (p *pod) OnUpdate(obj interface{}) {}

// OnDelete stops pod objects that are deleted.
func (p *pod) OnDelete(obj interface{}) {}

// GenerateHints creates hints needed for hints builder.
func (p *pod) GenerateHints(event bus.Event) bus.Event {
	return event
}

// Start starts the eventer
func (p *pod) Start() error {
	return nil
}

// Stop stops the eventer
func (p *pod) Stop() {}

// podUpdaterStore is the interface that an object needs to implement to be
// used as a pod updater store.
type podUpdaterStore interface {
	List() []interface{}
}

// namespacePodUpdater notifies updates on pods when their namespaces are updated.
type namespacePodUpdater struct{}

// OnUpdate handles update events on namespaces.
func (n *namespacePodUpdater) OnUpdate(obj interface{}) {}

// OnAdd handles add events on namespaces. Nothing to do, if pods are added to this
// namespace they will generate their own add events.
func (*namespacePodUpdater) OnAdd(interface{}) {}

// OnDelete handles delete events on namespaces. Nothing to do, if pods are deleted from this
// namespace they will generate their own delete events.
func (*namespacePodUpdater) OnDelete(interface{}) {}
