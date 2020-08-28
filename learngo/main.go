package main

import (
	"fmt"
	"strings"
)

// naked return
func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm dome") // defer : 함수가 끝났을 때 실행
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	length, upper := lenAndUpper2("abcde")
	fmt.Println(length, upper)
}
