// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package groupbyattrsprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessorMetrics(t *testing.T) {
	expectedViewNames := []string{
		"processor/groupbyattrs/num_grouped_traces",
		"processor/groupbyattrs/num_non_grouped_traces",
		"processor/groupbyattrs/span_groups",
		"processor/groupbyattrs/num_grouped_logs",
		"processor/groupbyattrs/num_non_grouped_logs",
		"processor/groupbyattrs/log_groups",
	}

	views := MetricViews()
	for i, viewName := range expectedViewNames {
		assert.Equal(t, viewName, views[i].Name)
	}
}
