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
// routines.  The globals.go file contains globs (viper) pkg variables
// used to drive the functionality of anything 'lib' related.  Right now
// it bootstraps some text and JSON settings that are needed early.
// Using such settings we can drive packages like lib/3rd/pretty and
// lib/api to set up dvln specific settings for indentation and such.
// It is done here as methods in 'lib' use those packages so this init()
// will be fired by the time those packages are called.
package lib

import (
	"dvln/lib/out"
	"os"

	globs "github.com/dvln/viper"
)

func doPkgLibGlobsInit() {
	// Section: BasicGlobal variables to store data (default value with config file overrides)
	// - please add them alphabetically and don't reuse existing opts/vars
	globs.SetDefault("texthumanize", true)
	globs.SetDesc("texthumanize", "prefer human readable text, turn off for Go-like data struct (bool)", globs.ExpertUser, globs.BasicGlobal)

	globs.SetDefault("textindentlevel", 2)
	globs.SetDesc("textindentlevel", "override the default two space text indent level (int)", globs.ExpertUser, globs.BasicGlobal)

	globs.SetDefault("textprefix", "")
	globs.SetDesc("textprefix", "override the default empty text prefix string (string)", globs.ExpertUser, globs.BasicGlobal)

	globs.SetDefault("jsonraw", false)
	globs.SetDesc("jsonraw", "prefer JSON in non-pretty raw format (bool)", globs.ExpertUser, globs.BasicGlobal)

	globs.SetDefault("jsonindentlevel", 2)
	globs.SetDesc("jsonindentlevel", "override the default two space JSON indent level (int)", globs.ExpertUser, globs.BasicGlobal)

	globs.SetDefault("jsonprefix", "")
	globs.SetDesc("jsonprefix", "override the default empty JSON prefix string (string)", globs.ExpertUser, globs.BasicGlobal)

	// Section: ConstGlobal variables to store data (default value only, no overrides)
	// - please add them alphabetically and don't reuse existing opts/vars
	globs.SetDefault("toolver", toolVer) // current version of the dvln tool
	globs.SetDesc("toolver", "current version of the dvln tool", globs.InternalUse, globs.ConstGlobal)

	globs.SetDefault("apiver", apiVer) // current version of the dvln tool
	globs.SetDesc("apiver", "current version of the dvln api", globs.InternalUse, globs.ConstGlobal)
	if err := os.Setenv("PKG_API_APIVER", apiVer); err != nil {
		out.Fatalln("Failed to set PKG_API_APIVER in the environment:", err)
	}
}

func init() {
	// Section: BasicGlobal variables to store data (default value with config file overrides)
	// - please add them alphabetically and don't reuse existing opts/vars
	doPkgLibGlobsInit()
}
