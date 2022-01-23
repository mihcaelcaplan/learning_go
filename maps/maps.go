package main

import "fmt"

func divider() {
	println("\n")
}

func main() {
	//  declare and define a map
	map1 := map[string]string{"fuck": "you", "who": "me?"}
	fmt.Printf("%s", map1)
	divider()
	for k, v := range map1 {
		fmt.Printf("%s:%s\n", k, v)
	}

}
