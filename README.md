spinner
=======

Package `spinner` implements simple text spinner.

[![GoDoc](https://godoc.org/github.com/gosimple/spinner?status.png)](https://godoc.org/github.com/gosimple/spinner)
[![Build Status](https://drone.io/github.com/gosimple/spinner/status.png)](https://drone.io/github.com/gosimple/spinner/latest)

[Documentation online](http://godoc.org/github.com/gosimple/spinner)

## Features

+ Supported brackets: curly ({}) and square ([])
+ Supported text separators: | and ~
+ Unlimited nesting

## Example

	text := "{Hello|Big} {world|people}!"
	fmt.Println(spinner.Spin(text))
	// Will print: Hello world! or Big world! or Hello people! or Big people!

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
