package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range list {
		if tmp := v / 2; v%2 == 0 {
			fmt.Println(i, "Even", tmp, v)
		} else {
			fmt.Println(i, "Odd", tmp, v)
		}

	}
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

}
