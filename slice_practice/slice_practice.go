package main

import "fmt"

func divider() {
	print("\n\n")
}

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

	fmt.Println("a 2d slice")
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

func subslicing() {
	// create a slice
	mother := []string{"fuck", "you", "my", "guy"}

	// create a smaller slice
	child1 := mother[0:2]

	// show how the data is linked
	fmt.Printf("mother: %s\n", mother)
	fmt.Printf("child: %s\n", child1)
	println("after changing mother[1] from you to me")
	mother[1] = "me"
	fmt.Printf("mother: %s\n", mother)
	fmt.Printf("child: %s\n", child1)

	// grab a subslice with capacity set to the length so it snips off from larger slice
	child2 := mother[0:2:2] //las [i:j:k] where k = j+i

	// show the data is now unlinked
	println("now change mother back, child changes even tho we thot we detached it")
	mother[1] = "you"
	fmt.Printf("mother: %s\n", mother)
	fmt.Printf("child: %s\n", child2)

	// as you can see, this didn't work :(),
	// that's because the detachment only happens on an append call
	// and really only prevents changes to the child from backpropagating to the mother
	// so
	child2 = append(child2, "you", "fuck")
	println("after appending to the child, it changes, but mother stays the same")
	fmt.Printf("mother: %s\n", mother)
	fmt.Printf("child: %s\n", child2)
}

func main() {

	// use make to init a slice
	new := make([]int, 10)
	// print the 0 inited slice
	fmt.Println("a slice ccreated with make([]int, 10)")
	for _, el := range new {
		fmt.Printf("%d ", el)
	}

	divider() //divider

	// use the functions
	init_and_print_slice_1x1()

	divider() //divider

	slice_2d := init_and_print_slice_2x1()
	flattened := flatten_slice_2x1(slice_2d)

	fmt.Println("flattened")
	for _, el := range flattened {
		fmt.Printf("%d ", el)
	}

	divider() //divider

	subslicing()
}
