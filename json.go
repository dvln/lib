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

func init() {
	// Section: BasicGlobal variables to store data (default value only, no overrides)
	// - please add them alphabetically and don't reuse existing opts/vars
	globs.SetDefault("jsonraw", false)
	globs.SetDesc("jsonraw", "Prefer JSON in non-pretty raw format", globs.ExpertUser, globs.BasicGlobal)
	globs.SetDefault("jsonindent", "  ")
	globs.SetDesc("jsonindent", "Override the default two space JSON indent level", globs.ExpertUser, globs.BasicGlobal)
	globs.SetDefault("jsonprefix", "")
	globs.SetDesc("jsonprefix", "Override the default empty JSON prefix string", globs.ExpertUser, globs.BasicGlobal)
}

// PrettyJSON pretty prints JSON data, provide the data and that can be followed
// by two optional arguments, a prefix string and an indent level (both of which
// are strings).  If neither is provided then no prefix used and indent of two
// spaces is the default (see cfgfile:jsonprefix, cfgfile:jsonindent and the
// related DVLN_JSONPREFIX, DVLN_JSONINDENT to adjust indentation and prefix
// as well as cfgfile:jsonraw and DVLN_JSONRAW for skipping pretty printing)
func PrettyJSON(b []byte, fmt ...string) (string, error) {
	jsonraw := globs.GetBool("jsonraw")
	if jsonraw {
		// if there's an override to say pretty JSON is not desired, honor it,
		// Feature: this could be changed to specifically remove carriage
		//          returns and shorten output around {} and :'s and such (?)
		return cast.ToString(b), nil
	}
	prefix := globs.GetString("jsonprefix")
	indent := globs.GetString("jsonindent")
	if len(fmt) == 1 {
		prefix = fmt[0]
	} else if len(fmt) == 2 {
		prefix = fmt[0]
		indent = fmt[1]
	}
	var out bytes.Buffer
	err := json.Indent(&out, b, prefix, indent)
	return cast.ToString(out.Bytes()), err
}
