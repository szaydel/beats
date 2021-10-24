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
	k8s "k8s.io/client-go/kubernetes"

	"github.com/gofrs/uuid"

	"github.com/elastic/beats/v7/libbeat/autodiscover"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/bus"
	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/keystore"
	"github.com/elastic/beats/v7/libbeat/logp"
)

// Eventer allows defining ways in which kubernetes resource events are observed and processed
type Eventer interface {
	kubernetes.ResourceEventHandler
	GenerateHints(event bus.Event) bus.Event
	Start() error
	Stop()
}

// EventManager allows defining ways in which kubernetes resource events are observed and processed
type EventManager interface {
	GenerateHints(event bus.Event) bus.Event
	Start()
	Stop()
}

// Provider implements autodiscover provider for docker containers
type Provider struct{}

// eventerManager implements start/stop methods for autodiscover provider with resource eventer
type eventerManager struct{}

// leaderElectionManager implements start/stop methods for autodiscover provider with leaderElection
type leaderElectionManager struct{}

// AutodiscoverBuilder builds and returns an autodiscover provider
func AutodiscoverBuilder(
	beatName string,
	bus bus.Bus,
	uuid uuid.UUID,
	c *common.Config,
	keystore keystore.Keystore,
) (autodiscover.Provider, error) {
	return nil, nil
}

// Start for Runner interface.
func (p *Provider) Start() {}

// Stop signals the stop channel to force the watch loop routine to stop.
func (p *Provider) Stop() {}

// String returns a description of kubernetes autodiscover provider.
func (p *Provider) String() string {
	return "kubernetes"
}

func NewEventerManager(
	uuid uuid.UUID,
	c *common.Config,
	cfg *Config,
	client k8s.Interface,
	publish func(event []bus.Event),
) (EventManager, error) {
	return nil, nil
}

func NewLeaderElectionManager(
	uuid uuid.UUID,
	cfg *Config,
	client k8s.Interface,
	startLeading func(uuid string, eventID string),
	stopLeading func(uuid string, eventID string),
	logger *logp.Logger,
) (EventManager, error) {
	return nil, nil
}

// Start for EventManager interface.
func (p *eventerManager) Start() {}

// Stop signals the stop channel to force the watch loop routine to stop.
func (p *eventerManager) Stop() {}

// GenerateHints for EventManager interface.
func (p *eventerManager) GenerateHints(event bus.Event) bus.Event {
	return event
}

// Start for EventManager interface.
func (p *leaderElectionManager) Start() {}

// Stop signals the stop channel to force the leader election loop routine to stop.
func (p *leaderElectionManager) Stop() {}

// GenerateHints for EventManager interface.
func (p *leaderElectionManager) GenerateHints(event bus.Event) bus.Event {
	return event
}
