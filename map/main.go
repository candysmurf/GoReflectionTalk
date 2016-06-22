//Package main shows sample operations for reflect.Map
package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/intelsdi-x/GoReflectionTalk/domain"
)

func main() {
	s := []domain.PetMap{}

	m := domain.PetMap{}
	for i := 0; i < 10; i++ {
		cat := domain.Cat{Name: "cat" + strconv.Itoa(i)}
		m[cat.Name] = &cat
	}
	s = append(s, m)

	m = domain.PetMap{}
	for i := 0; i < 10; i++ {
		dog := domain.Dog{Name: "dog" + strconv.Itoa(i)}
		m[dog.Name] = &dog
	}
	s = append(s, m)

	for _, a := range s {
		v := reflect.ValueOf(a)

		switch v.Kind() {
		case reflect.Map:
			for i, key := range v.MapKeys() {
				fmt.Println(i, ":", v.MapIndex(key).Elem())
			}
		default:
		}
	}

}
