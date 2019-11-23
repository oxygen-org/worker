package main

import (
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "Omit new line")
var s = flag.String("s", " ", "sep")

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(*n,*s)
}
