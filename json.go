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

// The dvln/lib/json.go module is for routines that might be useful
// for manipulating json beyond (or wrapping) the Go standard lib

package lib

import (
	"bytes"
	"encoding/json"

	"github.com/spf13/cast"
	globs "github.com/spf13/viper"
)

// insure that the globs (viper) packages env prefix is set *before* any
// init() functions are run so they all honor hte DVLN_ prefix
var envPrefix = globs.SetEnvPrefix("DVLN")

func init() {
	// Section: BasicGlobal variables to store data (default value only, no overrides)
	// - please add them alphabetically and don't reuse existing opts/vars
	globs.SetDefault("rawjson", false)
	globs.SetDesc("rawjson", "Used to print JSON in non-pretty raw format", globs.ExpertUser, globs.BasicGlobal)
}

// PrettyJSON pretty prints JSON data with two space indent, it will return
// a string result along with an error (if any)
func PrettyJSON(b []byte) (string, error) {
	rawjson := globs.GetBool("rawjson")
	if rawjson {
		// if there's an override to say pretty JSON is not desired, honor it
		return cast.ToString(b), nil
	}
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return cast.ToString(out.Bytes()), err
}
