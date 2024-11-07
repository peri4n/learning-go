package main

import "fmt"

func main() {
	list := []int{1, 2, 3}
	fmt.Println("Initial slice", list)

	list = append(list, []int{7, 8, 9}...)
	fmt.Println("Slice after appending", list)

	list = make([]int, 0, 10)
	fmt.Println("Create a new slice", list)

	list = append(list, 4, 5, 6)
	fmt.Println("Slice after appending", list)

	fmt.Println("Slicing the slice ':2'", list[:2])
	fmt.Println("Slicing the slice '2:'", list[2:])

	testMap := map[int][]string{}

	testMap[1] = []string{"foo", "bar"}
	fmt.Println("Print a map ", testMap)

	type person struct {
		name string
		age  int
		pet  string
	}

	fmt.Println("A person", person{name: "Fred", age: 42, pet: "Simba"})
}
