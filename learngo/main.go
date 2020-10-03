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

// 1.5 for, range, ...args
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

// 1.6 if with a Twist
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
/*func canIDrink(age int) bool {
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
}*/

// 1.8 pointers
/*func main() {
	a := 2
	b := &a
	*b = 20
	fmt.Println(a, &a, b, *b)
}*/

// 1.9 arrays and slices
/*func main() {
	names := []string{"a", "b", "c"}
	names = append(names, "d")
	fmt.Println(names)
}*/

// 1.10 Maps
/*func main() {
	nico := map[string]string{"name": "nico", "age": "12"}
	for key, value := range nico {
		fmt.Println(key, value)
	}
}*/

// 1.11 Structs
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
}
