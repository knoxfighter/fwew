//	This file is part of Fwew.
//	Fwew is free software: you can redistribute it and/or modify
// 	it under the terms of the GNU General Public License as published by
// 	the Free Software Foundation, either version 3 of the License, or
// 	(at your option) any later version.
//
//	Fwew is distributed in the hope that it will be useful,
//	but WITHOUT ANY WARRANTY; without even implied warranty of
//	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//	GNU General Public License for more details.
//
//	You should have received a copy of the GNU General Public License
//	along with Fwew.  If not, see http://gnu.org/licenses/

// Package util handles general program stuff. txt.go handles program strings.
package util

import (
	"fmt"
	"os/user"
	"path/filepath"
)

var usr, err = user.Current()
var texts = map[string]string{}

func init() {
	// main program strings
	texts["name"] = "Fwew"
	texts["author"] = "Tirea Aean"
	texts["header"] = fmt.Sprintf("%s - Na'vi Dictionary Search - by %s\n`fwew -h` for usage, `fwew -v` for version\n", texts["name"], texts["author"])
	texts["languages"] = "de, eng, est, hu, nl, pl, ru, sv"
	texts["prompt"] = "Fwew:> "

	// flag strings
	texts["usageDebug"] = "Show extremely verbose probing"
	texts["usageV"] = "Show program & dictionary version numbers"
	texts["usageL"] = "Use specified language \n\tValid values: " + texts["languages"]
	texts["usageI"] = "Display infix location data"
	texts["usageIPA"] = "Display IPA data"
	texts["usageP"] = "Search for word(s) with specified part of speech"
	texts["usageR"] = "Reverse the lookup direction from Na'vi->Local to Local->Na'vi"
	texts["usageA"] = "Find all matches by using affixes to match the input word"
	texts["usageN"] = "Convert Numbers Octal<->Decimal"
	texts["defaultFilter"] = "all"

	// file strings
	texts["homeDir"], _ = filepath.Abs(usr.HomeDir)
	texts["dataDir"] = filepath.Join(texts["homeDir"], ".fwew")
	texts["config"] = filepath.Join(texts["dataDir"], "config.json")
	texts["database"] = filepath.Join(texts["dataDir"], "database.sqlite")
	texts["dictionary"] = filepath.Join(texts["dataDir"], "dictionary.txt")
	texts["prefixes"] = filepath.Join(texts["dataDir"], "prefixes.txt")
	texts["infixes"] = filepath.Join(texts["dataDir"], "infixes.txt")
	texts["suffixes"] = filepath.Join(texts["dataDir"], "suffixes.txt")

	// general message strings
	texts["cset"] = "Currently set"
	texts["set"] = "set"
	texts["unset"] = "unset"

	// error message strings
	texts["none"] = "No Results\n"
	texts["noDataError"] = "Error 1: Failed to open and/or read dictionary file (" + texts["dictionary"] + ")"
	texts["fileError"] = "Error 2: Failed to open and/or read configuration file (" + texts["config"] + ")"
	texts["noOptionError"] = "Error 3: Invalid option"
	texts["invalidIntError"] = "Error 4: Input must be a decimal integer in range 0 <= n <= 32767 or octal integer in range 0 <= n <= 77777"
	texts["invalidOctalError"] = "Error 5: Invalid octal integer"
	texts["invalidDecimalError"] = "Error 6: Invalid decimal integer"
}

// Text function is the accessor for []string texts
func Text(s string) string {
	return texts[s]
}

// SetText is the setter for []string texts
func SetText(i, s string) {
	texts[i] = s
}
