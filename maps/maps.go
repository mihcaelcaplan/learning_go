package main

import (
	"fmt"
	"math"
)

func divider() {
	print("\n")
}

func hash(input string) (uint32, string) {
	// sum the product of string and p^n
	// mod it by m
	// where p is a prime near the number of expected input values
	//  m is a large number since collisions scale with 1/m
	// stole the current values from https://cp-algorithms.com/string/string-hashing.html
	p := 31
	m := math.Pow(10, 9) + 9
	var summer float64
	var err string

	for i, s := range input {
		summer += float64(s) * math.Pow(float64(p), float64(i))
	}
	partial := int(summer) % int(m)
	if partial >= 4294967296 {
		err = fmt.Sprintf("error: hash of %s too big!", input)
	}
	result := uint32(partial)
	return result, err
}

func fuckamap(input map[string]string) map[uint32]uint32 {
	output := map[uint32]uint32{}

	for k, v := range input {
		hk, err := hash(k)
		if len(err) > 0 {
			print(err)
		}
		hv, err := hash(v)
		if len(err) > 0 {
			print(err)
		}
		output[hk] = hv
	}

	return output
}

func main() {
	//  declare and define a map
	map1 := map[string]string{"fuck": "you", "who": "me?"}

	fmt.Printf("%s\n", map1)
	divider()

	for k, v := range map1 {
		fmt.Printf("%s:%s\n", k, v)
	}
	divider()

	hmap1 := fuckamap(map1)
	fmt.Printf("%d\n", hmap1)
	divider()

	for k, v := range hmap1 {
		fmt.Printf("%d:%d\n", k, v)
	}

}
