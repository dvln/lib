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
// routines.  Third party libs (under dvln/lib/3rd) should *not* rely
// upon anything in the 'dvln/lib' dir (lib package).  Additionally
// generic standalone packages (eg: dvln/lib/<pkg> like 'dvln/lib/out')
// should also NOT depend upon anything in the dvln specific 'dvln/lib'
// location.  The 'dvln/lib' functionality can, of course, depend upon
// these third party and standalone packages.
package lib

import (
	"fmt"

	"github.com/dvln/out"
	"github.com/dvln/toolver"
	globs "github.com/spf13/viper"
)

func init() {
	// Section: ConstGlobal variables to store data (default value only, no overrides)
	// - please add them alphabetically and don't reuse existing opts/vars
	globs.SetDefault("dvlnToolVer", "0.0.1") // current version of the dvln tool
	globs.SetDesc("dvlnToolVer", "current version of the dvln tool", globs.InternalUse, globs.ConstGlobal)
}

// DvlnToolInfo returns the version of the dvln tool and the build date for
// the binary being used and returns both, will return non-nil error on
// issues (currently only possible error will result in empty build date)
func DvlnToolInfo() (string, string, string, error) {
	toolVer := globs.GetString("dvlnToolVer")
	// get the build date of the current executable
	execName, buildDate, err := toolver.ExecutableInfo()
	if err != nil {
		fmt.Errorf("Problem determining build date, error: %s", err)
	}
	return execName, toolVer, buildDate, err
}

// DvlnVerStr returns a string with the version of the dvln tool such
// that it honors verbosity levels as well as look (text/json)
func DvlnVerStr() string {
	execName, toolVer, buildDate, err := DvlnToolInfo()
	// Get current runtime settings around desired verbosity and look (format)
	look := globs.GetString("look")
	terse := globs.GetBool("terse")
	verbose := globs.GetBool("verbose")
	if err != nil {
		// err in this case is not a big deal, means no build date
		// at the debug output level it won't show up normally unless
		// debugging is on (at which point parsing isn't gonna fly
		// anyhow, so just show the error directly when debugging)
		out.Debugln(err)
	}
	var dvlnVerStr string
	if terse {
		switch look {
		case "json":
			dvlnVerStr, err = PrettyJSON([]byte(fmt.Sprintf("{\"toolver\": \"%s\"}", toolVer)))
		case "text":
			dvlnVerStr = fmt.Sprint(toolVer)
		}
	} else if verbose {
		switch look {
		case "json":
			dvlnVerStr, err = PrettyJSON([]byte(fmt.Sprintf("{\"execname\": \"%s\", \"toolver\": \"%s\", \"builddate\": \"%s\"}", execName, toolVer, buildDate)))
		case "text":
			dvlnVerStr = fmt.Sprintf("Exec Name: %s\nVersion: %s\nBuild Date: %s", execName, toolVer, buildDate)
		}
	} else {
		switch look {
		case "json":
			dvlnVerStr, err = PrettyJSON([]byte(fmt.Sprintf("{\"toolver\": \"%s\", \"builddate\": \"%s\"}", toolVer, buildDate)))
		case "text":
			dvlnVerStr = fmt.Sprintf("Version: %s\nBuild Date: %s", toolVer, buildDate)
		}
	}
	if err != nil {
		out.Fatalln("Failed to parse devline version JSON string:", err)
	}
	return dvlnVerStr
}
