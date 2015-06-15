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

// Package lib is for general duty all purpose dvln specific library
// routines.  The json.go file contains anything needed to handle JSON
// output that isn't covered in more generic packages (like 'lib/api/...')
// such as things that need to rely upon the 'globs' (viper) package settings
// which aren't accessible via the more generic 'api' package.
package lib

import (
	globs "github.com/spf13/viper"
)

func init() {
	// Section: BasicGlobal variables to store data (default value with config file overrides)
	// - please add them alphabetically and don't reuse existing opts/vars
	globs.SetDefault("jsonraw", false)
	globs.SetDesc("jsonraw", "prefer JSON in non-pretty raw format", globs.ExpertUser, globs.BasicGlobal)

	globs.SetDefault("jsonindentlevel", 2)
	globs.SetDesc("jsonindentlevel", "override the default two space JSON indent level", globs.ExpertUser, globs.BasicGlobal)

	globs.SetDefault("jsonprefix", "")
	globs.SetDesc("jsonprefix", "override the default empty JSON prefix string", globs.ExpertUser, globs.BasicGlobal)
}
