package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%+v\n", Flatten([]interface{}{[]interface{}{1, 2, []int{3}}, 4}))
}

func Flatten(A []interface{}) (O []int) {
	for _, i := range A {
		switch v := i.(type) {
		case int:
			O = append(O, v)
		case []int:
			O = append(O, v...)
		case []interface{}:
			O = append(O, Flatten(v)...)
		}
	}
	return
}
