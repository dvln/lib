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
	"os"

	"github.com/dvln/api"
	"github.com/dvln/out"
	"github.com/dvln/toolver"
	"github.com/dvln/pretty"
	globs "github.com/dvln/viper"
)

const (
	toolVer = "0.0.1"
	apiVer  = "0.1"
)

func init() {
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

// DvlnToolInfo returns the version of the dvln tool and the build date for
// the binary being used and returns both, will return non-nil error on
// issues (currently only possible error will result in empty build date)
func DvlnToolInfo() (string, string, string, string, error) {
	toolVer := globs.GetString("toolver")
	apiVer := globs.GetString("apiver")
	// get the build date of the current executable
	execName, buildDate, err := toolver.ExecutableInfo()
	if err != nil {
		fmt.Errorf("Problem determining build date, error: %s", err)
	}
	return execName, toolVer, buildDate, apiVer, err
}

// DvlnVerStr returns a string with the version of the dvln tool such
// that it honors verbosity levels as well as look (text/json)
func DvlnVerStr() string {
	execName, toolVer, buildDate, apiVer, err := DvlnToolInfo()
	// Get current runtime settings around desired verbosity and look (format)
	look := globs.GetString("look")
	terse := globs.GetBool("terse")
	verbose := globs.GetBool("verbose")
	if err != nil {
		// err in this case is not a big deal, means no build date
		// at the debug output level it won't show up normally unless
		// debugging is on (at which point parsing isn't gonna fly
		// anyhow, so just display the error directly when debugging)
		out.Debugln(err)
	}
	type itemData struct {
		ToolVersion string `json:"toolVersion" pretty:"Version"`
		APIVersion  string `json:"apiVersion,omitempty" pretty:"API Rev,omitempty"`
		BuildDate   string `json:"buildDate,omitempty" pretty:"Build Date,omitempty"`
		ExecName    string `json:"execName,omitempty" pretty:"Exec Name,omitempty"`
	}
	fields := make([]string, 0, 0)
	items := make([]interface{}, 0, 0)
	verbosity := "regular"
	var newItem itemData
	newItem.ToolVersion = toolVer
	fields = append(fields, "toolVersion")
	if !terse {
		newItem.APIVersion = apiVer
		fields = append(fields, "apiVersion")
		newItem.BuildDate = buildDate
		fields = append(fields, "buildDate")
	} else {
		verbosity = "terse"
	}
	if verbose {
		verbosity = "verbose"
		newItem.ExecName = execName
		fields = append(fields, "execName")
	}
	items = append(items, newItem)
	if look == "json" {
		// see lib/json.go for these json* variables
		jsonLevel := globs.GetInt("jsonindentlevel")
		api.SetJSONIndentLevel(jsonLevel)
		raw := globs.GetBool("jsonraw")
		api.SetJSONRaw(raw)
		jsonPrefix := globs.GetString("jsonprefix")
		api.SetJSONPrefix(jsonPrefix)
		output, fatalProblem := api.GetJSONString("", "dvlnVersion", "version", verbosity, fields, items)
		if fatalProblem {
			out.Print(output)
			out.Exit(-1)
		}
		return output
	}
	// see lib/text.go for these text* variables
	humanize := globs.GetBool("texthumanize")
	pretty.SetHumanize(humanize)
	textLevel := globs.GetInt("textindentlevel")
	pretty.SetOutputIndentLevel(textLevel)
	textPrefixStr := globs.GetString("textprefix")
	pretty.SetOutputPrefixStr(textPrefixStr)
	return pretty.Sprintf("%# v", items)
}
