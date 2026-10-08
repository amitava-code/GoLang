package main

import "fmt"

// for -> only construct in go for looping

func main() {

	// while loop

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//clasic for loop

	for i := 0; i <= 10; i++ {

		// break
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}

	// range

	for i := range 4 {

		fmt.Println(i)
	}

}
