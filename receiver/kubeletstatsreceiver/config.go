// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubeletstatsreceiver

import (
	"time"

	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/config/confignet"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/kubelet"
)

var _ configmodels.Receiver = (*Config)(nil)

type Config struct {
	configmodels.ReceiverSettings `mapstructure:",squash"`
	kubelet.ClientConfig          `mapstructure:",squash"`
	confignet.TCPAddr             `mapstructure:",squash"`
	CollectionInterval            time.Duration `mapstructure:"collection_interval"`

	// ExtraMetadataLabels contains list of extra metadata that should be taken from /pods endpoint
	// and put as extra labels on metrics resource.
	// No additional metadata is fetched by default, so there are no extra calls to /pods endpoint.
	// Only container.id label is supported at the moment.
	ExtraMetadataLabels []kubelet.MetadataLabel `mapstructure:"extra_metadata_labels"`
}
