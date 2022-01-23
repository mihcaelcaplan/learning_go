package main

import "fmt"

func create_slice1d() []int {
	return make([]int, 10)

}

func create_slice2d() [][]int {
	return make([][]int, 10)
}

func init_and_print_slice_1x1() {
	slice1 := create_slice1d()

	for i := range slice1 {
		slice1[i] = len(slice1) - i
	}

	fmt.Println("Slice1")
	for i, el := range slice1 {
		fmt.Printf("#: %d, val: %d\n", i, el)
	}

}

func init_and_print_slice_2x1() [][]int {
	slice2 := create_slice2d()

	for i := range slice2 {
		a := []int{len(slice2) - i, i}
		slice2[i] = a
	}

	fmt.Println("Slice2")
	for i, el := range slice2 {
		fmt.Printf("#: %d, [ %d , %d ]\n", i, el[0], el[1])
	}

	return slice2
}

func flatten_slice_2x1(full_slice [][]int) []int {

	flattened := []int{}

	for _, el := range full_slice {
		flattened = append(flattened, el[0], el[1])
	}

	return flattened
}

func main() {
	init_and_print_slice_1x1()
	fmt.Println()
	slice_2d := init_and_print_slice_2x1()
	flattened := flatten_slice_2x1(slice_2d)

	fmt.Println("flattened")
	for _, el := range flattened {
		fmt.Printf("%d ", el)
	}
}
