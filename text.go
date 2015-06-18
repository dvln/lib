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
// routines.  The text.go file contains anything needed to handle text
// output that isn't covered in more generic packages (like 'lib/3rd/pretty/..')
// such as things that need to rely upon the 'globs' (viper) package settings
// which aren't accessible via the generic 'pretty' package (but in the general
// 'dvln' library here we can prep globs settings no problem and they will be
// at an init() level that is done for any testing of lib modules to work).
package lib

import (
	globs "github.com/dvln/viper"
)

func init() {
	// Section: BasicGlobal variables to store data (default value with config file overrides)
	// - please add them alphabetically and don't reuse existing opts/vars
	globs.SetDefault("texthumanize", true)
	globs.SetDesc("texthumanize", "prefer human readable text, turn off for Go-like data struct", globs.ExpertUser, globs.BasicGlobal)

	globs.SetDefault("textindentlevel", 2)
	globs.SetDesc("textindentlevel", "override the default two space text indent level", globs.ExpertUser, globs.BasicGlobal)

	globs.SetDefault("textprefix", "")
	globs.SetDesc("textprefix", "override the default empty text prefix string", globs.ExpertUser, globs.BasicGlobal)
}
