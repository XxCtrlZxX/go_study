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
/*func superAdd(numbers ...int) int {
	var total int = 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}*/

// 1.6 if
/*func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge > 18 {
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
}*/

// 1.7 switch
func canIDrink(age int) bool {
	switch {
	case age > 18:
		return true
	case age <= 18:
		return false
	}
	return false
}

func main() {
	fmt.Println(canIDrink(19))
}
