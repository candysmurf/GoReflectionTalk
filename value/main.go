//Package main shows samples operations of reflect.Value
package main

import (
	"fmt"
	"reflect"

	"github.com/intelsdi-x/GoReflectionTalk/domain"
)

func main() {

	v := reflect.ValueOf(10)
	fmt.Println("int 10\t value=", v, "\t dereference=", v, "\t interface=", v.Interface())

	// map of pets
	mPets := domain.PetMap{}

	// slice of pets
	sPets := domain.PetSlice{}

	cat := domain.Cat{Name: "cat"}
	mPets[cat.Name] = &cat
	sPets = append(sPets, &cat)

	dog := domain.Dog{Name: "dog"}
	mPets[dog.Name] = &dog
	sPets = append(sPets, &dog)

	for _, pet := range mPets {
		t := reflect.TypeOf(pet)
		if t.Kind() == reflect.Ptr {
			v := reflect.ValueOf(pet)
			fmt.Println("mPets\t value=", v, "\t dereference=", v.Elem(), "\t interface=", v.Interface())
		}
	}

	for _, pet := range sPets {
		v := reflect.ValueOf(pet)
		fmt.Println("sPets\t value=", v, "\t dereference=", v.Elem(), "\t interface=", v.Interface())
	}
}
