package main

import (
	"fmt"
	"strconv"
)

func hello(name string) {
	fmt.Println("Hello", name, "!")
}

// starts with upper case = public"
func Sum(a, b int) int {
	return a + b
}

// starts with lower case = "private"
func convertAndSum(a int, b string) (total int, err error) {
	// err from return declaration "(total int, err error)"
	i, err := strconv.Atoi(b)

	if err != nil {
		return
	}

	// total from return declaration "(total int, err error)"
	total = a + i

	return
}

func main() {
	hello("Matheus")
	fmt.Println("sum 1 + 1 = ", Sum(1, 1))

	total, err := convertAndSum(1, "123")

	if err != nil {
		fmt.Println("convertAndSum Error:", err)
		return
	}

	fmt.Println("convertAndSum 1 + '123' = ", total)
}
