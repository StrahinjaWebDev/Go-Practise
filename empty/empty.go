package main

import "fmt"

func main() {
	var i any

	i = 7
	fmt.Println(i)

	i = "hi"
	fmt.Println(i)

	s := i.(string) // type assertion
	fmt.Println("S:", s)

	/*
		n := i.(int) // will panic
		fmt.Println("n:", n)

	*/

	n, ok := i.(int) // won't panic
	if ok {
		fmt.Println("n:", n)
	} else {
		fmt.Println("not an int")
	}

	switch i.(type) {
	case int:
		fmt.Println("an int")
	case string:
		fmt.Println("a string")
	default:
		fmt.Println("a unknow type %T\n", i)

	}
	fmt.Println(maxInts([]int{3, 1, 2}))
	fmt.Println(maxFloat64s([]float64{3, 1, 2}))

	fmt.Println(max([]int{3, 1, 2}))
	fmt.Println(max([]float64{3, 1, 2}))
}

type Number interface {
	int | float64
}

func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

func maxInts(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

func maxFloat64s(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}
