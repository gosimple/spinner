package main

import (
	"github.com/gosimple/spinner"
	"fmt"
)

func main() {
	str := "I like {{some|a few{ kinds| types|} of} tacos|fish|chicken|pizza}."
	fmt.Println(spinner.Spin(str))
}
