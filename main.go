package main

import "fmt"

func greaterThat() func(int) []int {

	var greaterV int = 0
	var greaterS = make([]int, 0)

	return func(num int) []int {
		if num > greaterV {
			fmt.Println(num, "greater that", greaterV)
			greaterS = append(greaterS, num)
			greaterV = num
		} else {
			fmt.Println(greaterV, "greater that", num)
		}

		return greaterS
	}
}

func main() {
	var greater = greaterThat()

	fmt.Println(greater(2))
	fmt.Println(greater(6))
	fmt.Println(greater(3))
	fmt.Println(greater(5))
	fmt.Println(greater(9))
}
