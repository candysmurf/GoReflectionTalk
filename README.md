## Reflection - Standard Library

Developers are often told reflection is bad and should be avoided. However, when reflection can help reduce the amount of code needed to solve a problem, it should be considered. Reducing complexity while preserving maintainability must always be a priority. Reflection can be a tool that helps makes this possible. In this talk I am going to show you the power and elegance of the reflection package.

## Links

https://candysmurf.github.io/GoReflectionTalk

http://blog.golang.org/laws-of-reflection

## Code Review

### Basics

Example shows how to reflect over a struct type value.  
[Struct Types](code/basics/struct/struct.go) ([Go Playground](https://play.golang.org/p/V9Lf75Ajaa))    

Example shows how to reflect over a slice of struct type values.  
[Slices](code/basics/slice/slice.go) ([Go Playground](https://play.golang.org/p/ISWgNKWIxM))  

Example shows how to reflect over a map of struct type values.  
[Maps](code/basics/map/map.go) ([Go Playground](https://play.golang.org/p/BNWVXx4QTo))  

Example shows how to reflect over a struct type pointer.  
[Pointers](code/basics/pointer/pointer.go) ([Go Playground](https://play.golang.org/p/0bSbKUJk19))  

### Interfaces

Example shows how to reflect over a struct type value that is stored inside an interface value.  
[Struct Types](code/interface/struct/struct.go) ([Go Playground](https://play.golang.org/p/kHC6nuHYty))  

Example shows how to reflect over a slice of struct type values that are stored inside an interface value.  
[Slices](code/interface/slice/slice.go) ([Go Playground](https://play.golang.org/p/UyRIlkjVjW))  

Example shows how to reflect over a map of struct type values that are stored inside an interface value.  
[Maps](code/interface/map/map.go) ([Go Playground](https://play.golang.org/p/-_niEdmavG))  

Example shows how to reflect over a struct type pointer that is stored inside an interface value.  
[Pointers](code/interface/pointer/pointer.go) ([Go Playground](https://play.golang.org/p/itFSg3BL0o))  

### Inspection / Decoding

Example shows how to inspect a structs fields and display the field name, type and value.  
[Struct Types](code/inspect/struct/struct.go) ([Go Playground](https://play.golang.org/p/ahHLMtun9y))  

Example shows how to use reflection to decode an integer.  
[Integers](code/interface/integer/integer.go) ([Go Playground](https://play.golang.org/p/LmVkzpm57a))  
___
All material is licensed under the [Apache License Version 2.0, June 2016](http://www.apache.org/licenses/LICENSE-2.0).
