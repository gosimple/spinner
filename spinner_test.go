// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package spinner

import (
	"testing"
)

// Check if list contain value
func contain(list []string, value interface{}) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

func TestSpin(t *testing.T) {
	testCases := []struct {
		in   string
		want []string
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
		{"{I like some}}} pizza.",
			[]string{
				"I like some}} pizza.",
			},
		},
		{"I like some}{} pizza.",
			[]string{
				"I like some} pizza.",
			},
		},
		{"{{{I like some} pizza.",
			[]string{
				"{{I like some pizza.",
			},
		},
	}

	for _, tc := range testCases {
		got := Spin(tc.in)
		if !contain(tc.want, got) {
			t.Errorf(
				"Spin(%#v) = %#v; want one of %#v", tc.in, got, tc.want)
		}
	}
}
