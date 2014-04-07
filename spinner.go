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
	var brOpenPos, brClosePos int
	repl := strings.NewReplacer("[", "{", "]", "}", "~", "|")
	text = repl.Replace(text)
	endString := text
	if strings.ContainsAny(text, "{}") {
		for ind, sym := range text {
			if string(sym) == "{" {
				if openBr == 0 {
					brOpenPos = ind
				}
				openBr += 1
			} else if string(sym) == "}" {
				if openBr > 0 {
					openBr -= 1
					brClosePos = ind
					if openBr == 0 || !strings.Contains(text[ind+1:], "}") {
						endString = strings.Replace(
							endString,
							text[brOpenPos:brClosePos+1],
							Spin(text[brOpenPos+1:brClosePos]),
							1)
						return Spin(endString)
					}
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
