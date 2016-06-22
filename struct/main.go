// package main shows sample operations of reflect.Struct
package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

// Pet defines pet's properties
type Pet struct {
	Name  string    `json:"name"`
	Age   int       `json:"age"`
	Label []string  `json:"label"`
	Draw  io.Writer `json:"draw"`
}

func (p *Pet) reflect() {
	val := reflect.ValueOf(p).Elem()

	for i := 0; i < val.NumField(); i++ {
		// reflect.Value
		vField := val.Field(i)
		// reflect.StructField
		tField := val.Type().Field(i)
		// reflect.StructTag
		tag := tField.Tag
		fmt.Printf("Field Name=%s,\t  Value=%v,\t  Tag=%s\n", tField.Name, vField.Interface(), tag.Get("json"))
	}
}

func (p *Pet) update() {
	val := reflect.ValueOf(p).Elem()
	for i := 0; i < val.NumField(); i++ {
		vField := val.Field(i)
		if vField.CanSet() {
			switch vField.Kind() {
			case reflect.String:
				vField.SetString("Animal")
			case reflect.Int:
				vField.SetInt(10)
			case reflect.Slice:
				for i := 0; i < vField.Len(); i++ {
					sField := vField.Index(i)
					switch sField.Kind() {
					case reflect.String:
						sField.SetString("cutie")
					default:
					}
				}
			default:
			}
		}
	}
}

func main() {
	pet := &Pet{Name: "Gopher", Age: 1, Label: []string{"raccoon", "possum"}, Draw: os.Stdout}
	pet.reflect()
	pet.update()
	pet.reflect()
}
