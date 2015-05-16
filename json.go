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
)

// PrettyJSON pretty prints JSON data with two space indent, it will return
// a string result along with an error (if any)
func PrettyJSON(b []byte) (string, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return cast.ToString(out.Bytes()), err
}
