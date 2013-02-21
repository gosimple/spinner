// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package spinner

import (
	"testing"
)

var SpinTests = []struct {
	in  string
	out []string
}{
	{"{Hello|Big} {world|people}!",
		[]string{
			"Big world!",
			"Big people!",
			"Hello world!",
			"Hello people!",
		},
	},
	{"I like {{some|a few{ kinds| types|} of} tacos|pizza}.",
		[]string{
			"I like pizza.",
			"I like some tacos.",
			"I like a few of tacos.",
			"I like a few kinds of tacos.",
			"I like a few types of tacos.",
		},
	},
}

func TestSpin(t *testing.T) {
	for index, st := range SpinTests {
		spinned := Spin(st.in)
		if !contain(st.out, spinned) {
			t.Errorf(
				"%d. Spin(%v) => out = %v, want one of %v",
				index, st.in, spinned, st.out)
		}
	}
}

// Check if list contain value
func contain(list []string, value interface{}) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}
