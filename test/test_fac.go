package main

import (
	f "Lacs/test/fac"
	"fmt"
)

func Factory(name string) f.Animal {
	switch name {
	case "dog":
		return &f.Dog{MaxAge: 20}
	case "cat":
		return &f.Cat{MaxAge: 10}
	default:
		panic("No such animal")
	}
}

func main()  {
	animal := Factory("dog")
	animal.Sleep()
	fmt.Printf("%s max age is: %d", animal.Type(), animal.Age())
}