// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package spinner

import (
	"math/rand"
	"strings"
	"time"
)

//=============================================================================

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Spin returns a spinned text.
func Spin(text string) (spinned string) {
	var openBr int
	var braces [][2]int
	repl := strings.NewReplacer("[", "{", "]", "}", "~", "|")
	text = repl.Replace(text)
	endString := text
	if strings.ContainsAny(text, "{}") {
		for ind, sym := range text {
			if string(sym) == "{" {
				temp := [2]int{0, 0}
				braces = append(braces, temp)
				braces[openBr][0] = ind
				openBr += 1
			} else if string(sym) == "}" {
				openBr -= 1
				braces[openBr][1] = ind
				if openBr == 0 && strings.Contains(text[:ind], "{") {
					endString = strings.Replace(
						endString,
						text[braces[openBr][0]:braces[openBr][1]+1],
						Spin(text[braces[openBr][0]+1:braces[openBr][1]]),
						1)
					return Spin(endString)
				} else if openBr == 0 {
					return Spin(text[braces[openBr][0] : braces[openBr][1]+1])
				}
			}
		}
	} else {
		endString = chooseText(text)
	}
	return endString
}

// chooseText will choose random text from a string (text separator: |)
func chooseText(st string) string {
	if strings.Contains(st, "|") {
		choices := strings.Split(st, "|")
		chosen := rand.Intn(len(choices))
		return choices[chosen]
	}
	return st
}
