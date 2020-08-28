package main

import (
	"fmt"
	"strings"
)

// several returns
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// naked return
func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	len, up := lenAndUpper("abcd")
	fmt.Println(len, up)
	length, upper := lenAndUpper2("efghijk")
	fmt.Println(length, upper)
}
