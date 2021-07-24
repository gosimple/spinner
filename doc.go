// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

/*
Package spinner implements text spinner. Useful for generating semirandom text.

## Features

+ Supported brackets: curly ({}) and square ([])
+ Supported text separators: | and ~
+ Unlimited nesting

## Example

	text := "{Hello|Big} {world|people}!"
	fmt.Println(spinner.Spin(text))
	// Posible outputs: 
	//   Hello world!
	//   Big world!
	//   Hello people!
	//   Big people!

Requests or bugs?

https://github.com/gosimple/spinner/issues
*/
package spinner
