//Package main shows samples operations of reflect.Type
package main

import (
	"fmt"
	"reflect"

	"github.com/intelsdi-x/GoReflectionTalk/domain"
)

func main() {
	// map of pets
	mPets := domain.PetMap{}
	t := reflect.TypeOf(mPets)
	fmt.Println("mPets\t dynamicType=", t, "\t concreteType=", t.Kind())

	// slice of pets
	sPets := domain.PetSlice{}
	t = reflect.TypeOf(sPets)
	fmt.Println("sPets\t dynamicType=", t, "\t concreteType=", t.Kind())

	cat := domain.Cat{Name: "cat"}
	fmt.Println("cat\t dynamicType=", reflect.TypeOf(cat), "\t concreteType=", reflect.TypeOf(cat).Kind())

	mPets[cat.Name] = &cat
	sPets = append(sPets, &cat)

	dog := domain.Dog{Name: "dog"}
	fmt.Println("dog\t dynamicType=", reflect.TypeOf(dog), "\t concreteType=", reflect.TypeOf(dog).Kind())

	mPets[dog.Name] = &dog
	sPets = append(sPets, &dog)

	for _, pet := range mPets {
		t := reflect.TypeOf(pet)
		fmt.Println("mPets\t dynamicType=", t, "\t concreteType=", t.Kind())
	}

	for _, pet := range mPets {
		t := reflect.TypeOf(pet)
		if t.Kind() == reflect.Ptr {
			elem := reflect.ValueOf(pet).Elem()
			eType := reflect.TypeOf(elem)
			fmt.Println("mPets\t dynamicType=", eType, "\t concreteType=", eType.Kind())
		}
	}

	for _, pet := range sPets {
		t := reflect.TypeOf(pet)
		fmt.Println("sPets\t dynamicType=", t, "\t concreteType=", t.Kind())
	}
}
