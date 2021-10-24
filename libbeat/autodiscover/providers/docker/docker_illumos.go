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

package docker

import (
	"github.com/gofrs/uuid"

	"github.com/elastic/beats/v7/libbeat/autodiscover"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/bus"
	"github.com/elastic/beats/v7/libbeat/keystore"
)

func init() {
}

// Provider implements autodiscover provider for docker containers
type Provider struct{}

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

// Start the autodiscover process
func (d *Provider) Start() {}

// Stop the autodiscover process
func (d *Provider) Stop() {}

func (d *Provider) String() string {
	return "docker"
}
