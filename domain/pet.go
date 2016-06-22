//Package domain defines interfaces and objects
package domain

import "fmt"

// Pet defines the method sets for pets
type Pet interface {
	SayHello()
}

// Cat represent a cat
type Cat struct {
	Name string
}

// Dog represents a dog
type Dog struct {
	Name string
}

// Dogs represents a slice of Dogs
type Dogs []Dog

// PetMap represents a map of pets
type PetMap map[string]Pet

// PetSlice represents a slice of pets
type PetSlice []Pet

// SayHello Cat implements Pet interface
func (c *Cat) SayHello() {
	fmt.Println("Cat says Meow")
}

// SayHello Dog implements Pet interface
func (c *Dog) SayHello() {
	fmt.Println("Dog says Wolf")
}
