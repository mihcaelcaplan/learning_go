package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// type struct
// when choosing to use pointer receivers or value receivers for methods
// think, do i want to always make a copy of this type of thing,
// so i can pass them around and in and out of mutating functions?
// or do i want to pass a reference of this thing around,
// and do operations on the underlying?

// the nice thing is that GO will make things make sense under the hood in
// my methods, don't have to keep track of passing as a pointer or val

// for this type, which represents a "fucker", a person essentially,
// i want it to represent a coherent person. essentially a reference

type fucker struct {
	name     string
	birthday time.Time
	secret   uint8
}

// method that doesn't take any info
// even tho it doesn't take any info, because I've chosen to treat a "fucker"
// like a nonprimitive type, i give it a reference
func (f *fucker) fuck() {
	println("fuck you!!!")
}

// method that changes the values
func (f *fucker) birth(name string, time time.Time) {
	f.name = name
	f.birthday = time
}

// method that returns a value
func (f *fucker) tellmeasecret() uint8 {
	fmt.Printf("my secret is 0x%02X\n", f.secret)
	return f.secret
}

// method that writes a secret
func (f *fucker) traumatize(secret uint8) {
	f.secret = secret
}

// main
func main() {

	var you fucker

	you.fuck()

	birthday := time.Now().AddDate(-42.0, 69, -69)

	you.birth("ass hole", birthday)
	fmt.Printf("you'r biirthday is %s!\n", you.birthday.String())
	fmt.Printf("you'r name is %s!\n", you.name)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	secret := uint8(r.Intn(int(math.Pow(2, 8))))
	// fmt.Printf("hehe, secret: %d\n", secret)
	fmt.Printf("hehe, secret: 0x%02X\n", secret)

	you.traumatize(secret)

	your_secret := you.tellmeasecret()

	if your_secret == secret {
		println("looks like noone else has traumatizd you :)")
		fmt.Printf("ur secret is the one i gave you, 0x%02X", secret)
	}

}
