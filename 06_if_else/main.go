package main

import "fmt"

func main() {

	// age := 10

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else {
	// 	fmt.Println("person is not adult")
	// }

	age := 8

	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is a kid")
	}

	if age := 20; age >= 18 {

		fmt.Println("person is and asult", age)

	} else if age >= 12 {

		fmt.Println("person is a teenager", age)

	}
}
