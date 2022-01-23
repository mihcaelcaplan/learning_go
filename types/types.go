package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type fucker struct {
	name     string
	birthday time.Time
	secret   uint8
}

func (f fucker) fuck() {
	println("fuck you!!!")
}

func (f *fucker) birth(name string, time time.Time) {
	f.name = name
	f.birthday = time
}

func (f fucker) tellmeasecret() uint8 {
	fmt.Printf("my secret is 0x%02X\n", f.secret)
	return f.secret
}

func (f *fucker) traumatize(secret uint8) {
	f.secret = secret
}
func main() {

	var you fucker

	you.fuck()

	birthday := time.Now()

	you.birth("ass hole", birthday)
	fmt.Printf("you'r biirthday is %s!\n", you.birthday.String())

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
