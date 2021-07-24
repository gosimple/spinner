spinner
=======

[![Go Reference](https://pkg.go.dev/badge/github.com/gosimple/spinner.svg)](https://pkg.go.dev/github.com/gosimple/spinner)
[![Tests](https://github.com/gosimple/spinner/actions/workflows/tests.yml/badge.svg)](https://github.com/gosimple/spinner/actions/workflows/tests.yml)

Package `spinner` implements text spinner. Useful for generating semirandom text.

## Features

+ Supported brackets: curly `{}` and square `[]`
+ Supported text separators: `|` and `~`
+ Unlimited nesting

## Example

	text := "{Hello|Big} {world|people}!"
	fmt.Println(spinner.Spin(text))
	// Possible outputs: 
	//   Hello world!
	//   Big world!
	//   Hello people!
	//   Big people!

### Requests or bugs? 
<https://github.com/gosimple/spinner/issues>

## Installation

	go get -u github.com/gosimple/spinner

## License

The source files are distributed under the 
[Mozilla Public License, version 2.0](http://mozilla.org/MPL/2.0/),
unless otherwise noted.  
Please read the [FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html)
if you have further questions regarding the license.
