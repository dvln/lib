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
	cfg "github.com/spf13/viper"
)

// DvlnVer returns the version of the dvln tool and the build date for
// the binary being used and returns both, will return non-nil error on
// issues (currently only possible error will result in empty build date)
func DvlnVer() (string, string, error) {
	toolVer := cfg.GetString("dvlnToolVer")
	buildDate, err := toolver.BuildDate() // set the build date from executable's mdate
	if err != nil {
		fmt.Errorf("Problem determining build date, error: %s", err)
	}
	return toolVer, buildDate, err
}

// DvlnVerStr returns a string with the version of the dvln tool such
// that it honors verbosity levels as well as look (text/json)
func DvlnVerStr() string {
	toolVer, buildDate, err := DvlnVer()
	// Get current runtime settings around desired verbosity and look (format)
	look := cfg.GetString("look")
	terse := cfg.GetBool("terse")
	verbose := cfg.GetBool("verbose")
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
			case "json": dvlnVerStr = fmt.Sprintf("{\"toolver\": \"%s\"}", toolVer)
			case "text": dvlnVerStr = fmt.Sprint(toolVer)
		}
	} else if verbose {
		switch look {
			case "json": dvlnVerStr = fmt.Sprintf("{\"toolver\": \"%s\", \"builddate\": \"%s\"}", toolVer, buildDate)
			case "text": dvlnVerStr = fmt.Sprintf("Version: %s, Build date: %s", toolVer, buildDate)
		}
	} else {
		switch look {
			case "json": dvlnVerStr = fmt.Sprintf("{\"toolver\": \"%s\"}", toolVer)
			case "text": dvlnVerStr = fmt.Sprintf("Version: %s", toolVer)
		}
	}
	return dvlnVerStr
}

