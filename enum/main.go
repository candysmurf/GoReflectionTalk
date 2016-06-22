//Package main has sample code for reflect.String
package main

import (
	"fmt"
	"reflect"
)

// const constants
const (
	Cat = iota
	Dog
	Raccoon
	Possum
	Gopher
	Pig
	Bird
	Fish
)

var (
	pets = []string{"Cat", "Dog", "Raccoon", "Possum", "Gopher", "Pig", "Bird", "Fish"}
)

// Pets type
type Pets int

func (p Pets) String() string {
	return pets[p]
}

func main() {
	for i := 0; i < 10; i++ {
		v := reflect.ValueOf(i)
		fmt.Println("value=", v, "\t type=", reflect.TypeOf(v), "\t kind=", reflect.TypeOf(v).Kind())
	}
}
