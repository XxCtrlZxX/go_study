package main

import (
	"fmt"

	"github.com/XxCtrlZxX/lecture2/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	/*definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}*/

	/*err := dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search("hello")
	fmt.Println(hello)

	err2 := dictionary.Add("hello", "hi")
	if err2 != nil {
		fmt.Println(err2)
	}*/

	baseword := "hello"
	dictionary.Add(baseword, "First")

	dictionary.Delete(baseword)

	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
