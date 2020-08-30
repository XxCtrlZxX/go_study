package main

import (
	"fmt"
)

// 1.4 naked return
/*func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm dome") // defer : 함수가 끝났을 때 실행
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	length, upper := lenAndUpper2("abcde")
	fmt.Println(length, upper)
}*/

// 1.5 for
func superAdd(numbers ...int) int {
	var total int = 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
