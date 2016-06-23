## Reflection - Standard Library

Reflection is the ability to inspect a value to derive type or other meta-data. Reflection can give our program incredible flexibility to work with data of different types or create values on the fly. Reflection is critical for the encoding and decoding of data.

## Notes

* The reflection package allows us to inspect our types.
* We can add "tags" to our struct fields to store and use meta-data.
* Encoding package leverages reflection and we can as well.

## Links

http://blog.golang.org/laws-of-reflection

## Code Review

### Basics

[Struct Types](basics/struct/struct.go) ([Go Playground](https://play.golang.org/p/V9Lf75Ajaa))  
[Slices](basics/slice/slice.go) ([Go Playground](https://play.golang.org/p/ISWgNKWIxM))  
[Maps](basics/map/map.go) ([Go Playground](https://play.golang.org/p/BNWVXx4QTo))  
[Pointers](basics/pointer/pointer.go) ([Go Playground](https://play.golang.org/p/0bSbKUJk19))  

### Interfaces

[Struct Types](interface/struct/struct.go) ([Go Playground](https://play.golang.org/p/kHC6nuHYty))  
[Slices](interface/slice/slice.go) ([Go Playground](https://play.golang.org/p/UyRIlkjVjW))  
[Maps](interface/map/map.go) ([Go Playground](https://play.golang.org/p/-_niEdmavG))  
[Pointers](interface/pointer/pointer.go) ([Go Playground](https://play.golang.org/p/itFSg3BL0o))  

### Inspection / Decoding

[Struct Types](inspect/struct/struct.go) ([Go Playground](https://play.golang.org/p/ahHLMtun9y))  
[Integers](interface/integer/integer.go) ([Go Playground](https://play.golang.org/p/LmVkzpm57a))  
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
