package main

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
	"strings"
)

//go:generate go install github.com/elliotchance/pie@latest
//go:generate pie MyStringList.Contains
////go:generate pie MyStringList.* // or All functions
//go:generate go mod tidy

type MyStringList []string

func main() {
	// example usage of built-in functions
	names := pie.FilterNot([]string{"iTem1", "Item2", "Item3", "iTem4"},
		func(name string) bool {
			return strings.HasPrefix(name, "i")
		})
	fmt.Println(names) // "[Item2 Item3]"

	// example usage of custom-type generated functions
	var myList = MyStringList(names)
	if myList.Contains("Item3") {
		fmt.Println("yes, found Item3")
	}
}
