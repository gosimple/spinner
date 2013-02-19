spinner
=======

Package `spinner` implements simple text spinner.

[Documentation online](http://godoc.org/bitbucket.org/gosimple/spinner)

## Features

+ Supported brackets: curly ({}) and square ([])
+ Supported text separators: | and ~
+ Unlimited nesting


	text := "{Hello|Big} {world|people}!"
	fmt.Println(spinner.Spin(text))
	// Will print: Hello world! or Big world! or Hello people! or Big people!

### Requests or bugs? 
<https://bitbucket.org/gosimple/spinner/issues>

## Installation

	go get -u bitbucket.org/gosimple/spinner

## License

The source files are distributed under the 
[Mozilla Public License, version 2.0](http://mozilla.org/MPL/2.0/),
unless otherwise noted.  
Please read the [FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html)
if you have further questions regarding the license.
