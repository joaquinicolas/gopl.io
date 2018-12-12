package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	printArgs(os.Args)
	printArgsRange(os.Args)
	fmt.Printf("\n%.2f elapsed", time.Since(start).Seconds())
}

// print program arguments calling string join method
func printArgs(args []string)  {
	fmt.Printf("%s: %s", args[0], strings.Join(args[1:], " "))
}

// print program arguments using range
func printArgsRange(args []string)  {
	var sep string
	for i, arg := range args[1:] {
		fmt.Printf("\n%s%d: %s", sep, i, arg)
		sep = " "
	}
}