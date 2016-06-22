//Packae main shows sample operations of reflect.Slice
package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/intelsdi-x/GoReflectionTalk/domain"
)

func main() {
	// slice of pets
	sPets := domain.PetSlice{}
	pets := map[string]reflect.Value{}

	for i := 0; i < 10; i++ {
		cat := domain.Cat{Name: "cat" + strconv.Itoa(i)}
		sPets = append(sPets, &cat)
	}
	pets["cat"] = reflect.ValueOf(sPets)

	sPets = domain.PetSlice{}
	for i := 0; i < 10; i++ {
		dog := domain.Dog{Name: "dog" + strconv.Itoa(i)}
		sPets = append(sPets, &dog)
	}
	pets["dog"] = reflect.ValueOf(sPets)

	for _, v := range pets {
		switch v.Kind() {
		case reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				fmt.Println(v.Index(i).Elem())
			}
		default:
		}
	}

}
