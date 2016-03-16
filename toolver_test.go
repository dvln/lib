// Copyright © 2015 Erik Brady <brady@dvln.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lib

import (
	"encoding/json"
	"testing"

	"github.com/dvln/testify/assert"
	globs "github.com/dvln/viper"
)

func TestTextDvlnverOutput(t *testing.T) {
	// Try "standard" verbosity with text output, see if right items set
	output := DvlnVerStr()
	assert.Contains(t, output, "Version:")
	assert.Contains(t, output, "API Rev:")
	assert.Contains(t, output, "Build Date:")
	assert.NotContains(t, output, "Exec Name:")

	// Try "verbose" output
	globs.Set("verbose", true)
	output = DvlnVerStr()
	assert.Contains(t, output, "Version:")
	assert.Contains(t, output, "API Rev:")
	assert.Contains(t, output, "Build Date:")
	assert.Contains(t, output, "Exec Name:")

	// Try "terse" output
	globs.Set("verbose", false)
	globs.Set("terse", true)
	output = DvlnVerStr()
	assert.Contains(t, output, "Version:")
	assert.NotContains(t, output, "API Rev:")
	assert.NotContains(t, output, "Build Date:")
	assert.NotContains(t, output, "Exec Name:")
	globs.Set("terse", false)
}

func TestJSONDvlnverOutput(t *testing.T) {
	// Try "standard" verbosity with JSON output, see if right items set
	globs.Set("look", "json")
	output := DvlnVerStr()
	assert.Contains(t, output, "  \"context\": \"dvlnVersion\"")
	assert.Contains(t, output, "    \"items\":")
	assert.Contains(t, output, "        \"toolVersion\":")
	assert.Contains(t, output, "        \"apiVersion\":")
	assert.Contains(t, output, "        \"buildDate\":")
	assert.NotContains(t, output, "        \"execName\":")
	// Verify the JSON is parsable
	var result interface{}
	err := json.Unmarshal([]byte(output), &result)
	if err != nil {
		t.Fatalf("Unable to marshal regular JSON output: %s", err)
	}
	// FEATURE: could get more results out of returned structure... below also

	// Try "verbose" output
	globs.Set("verbose", true)
	output = DvlnVerStr()
	assert.Contains(t, output, "  \"context\": \"dvlnVersion\"")
	assert.Contains(t, output, "    \"items\":")
	assert.Contains(t, output, "        \"toolVersion\":")
	assert.Contains(t, output, "        \"apiVersion\":")
	assert.Contains(t, output, "        \"buildDate\":")
	assert.Contains(t, output, "        \"execName\":")
	err = json.Unmarshal([]byte(output), &result)
	if err != nil {
		t.Fatalf("Unable to unmarshal verbose JSON output: %s", err)
	}

	// Try "terse" output
	globs.Set("verbose", false)
	globs.Set("terse", true)
	output = DvlnVerStr()
	assert.Contains(t, output, "  \"context\": \"dvlnVersion\"")
	assert.Contains(t, output, "    \"items\":")
	assert.Contains(t, output, "        \"toolVersion\":")
	assert.NotContains(t, output, "        \"apiVersion\":")
	assert.NotContains(t, output, "        \"buildDate\":")
	assert.NotContains(t, output, "        \"execName\":")
	err = json.Unmarshal([]byte(output), &result)
	if err != nil {
		t.Fatalf("Unable to marshal terse JSON output: %s", err)
	}

	// Reset settings in case later tests
	globs.Set("terse", false)
	globs.Set("look", "text")
}
