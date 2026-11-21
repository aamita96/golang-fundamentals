# What is Go?

- Go is a cross-platform, open source programming language
- Go can be used to create high-performance applications
- Go is a fast, statically typed, compiled language known for its simplicity and efficiency
- Go was developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson in 2007
- Go's syntax is similar to C++
- Go is a pass-by-value language


Some Important points to know about go before starting - 
- Go (Golang) does not support classes in the traditional OOP sense like Java, C++, or Python.
 - However, Go supports most object-oriented programming concepts using other features.
 - Struct is a blueprint which describes data similar to classes in other programming languages.
- OOP Concepts in Traditional OOP Languages vs Go:

	1. Class              -> Struct
	2. Object             -> Struct Instance
	3. Methods             -> Methods with Receivers
	4. Inheritance         -> Composition (Struct Embedding)
	5. Interface           -> Interface (Implicit Implementation)
	6. Constructor         -> Constructor Function
	7. Encapsulation       -> Exported and Unexported Fields/Methods
	8. Polymorphism        -> Interfaces
	9. Abstraction         -> Interfaces
	10. Method Overloading -> Not Supported Directly
	11. Method Overriding  -> Not Supported Directly (Method Shadowing via Embedding)
	12. this / self        -> Receiver (e.g., func (u User) ...)
	13. Static Members     -> Package-level Variables and Functions
	14. Multiple Inheritance -> Not Supported (Composition Instead)

# What is Go used for?
- Web development (server-side)
- Developing network-based programs
- Developing cross-platform enterprise applications
- Cloud-native development

# Why use Go?
- Go is fun and easy to learn
- Go has fast run time and compilation time
- Go supports concurrency
- Go has memory management
- Go works on different platforms (Windows, Mac, Linux, Raspberry Pi, etc.)

| Go | Python | C++ |
|---|---|---|
| Statically typed | Dynamically typed | Statically typed |
| Fast run time | Slow run time | Fast run time |
| Compiled | Interpreted | Compiled |
| Fast compile time | Interpreted | Slow compile time |
| Supports concurrency through goroutines and channel | No built-in concurrency mechanism | Supports concurrency through threads |
| Has automatic garbage collection | Has automatic garbage collection | Does not have automatic garbage collection |
| Does not support classes and objects | Has classes and objects | Has classes and objects |
| Does not support inheritance | Supports inheritance | Supports inheritance |


# Go Syntax

A Go file consists of the following parts:

- Package declaration
- Import packages
- Functions
- Statements and expressions

## Look at the following code, to understand it better:

```
package main
import "fmt"

func main() {
  fmt.Println("Hello World!")
}
```

# Go Comments 

- A comment is a text that is ignored upon execution.
- Comments can be used to explain the code, and to make it more readable.
- Go supports single-line or multi-line comments.

### Go Single-line Comments
Single-line comments start with two forward slashes (`//`).
### Go Multi-line Comments
Multi-line comments start with `/*` and ends with `*/`.

# Go variables
Variables are containers for storing data values.


# Declaring (Creating) Variables
### In Go, there are two ways to declare a variable:

1. With the var keyword:

    Use the var keyword, followed by variable name and type:

    `var name = "Mario"`

    `Note:` You always have to specify either type or value (or both).

2. With the := sign:

    Use the := sign, followed by the variable value:

    ```
    name := "Mario" //type is inferred
    age := 2        //type is inferred
    ```

    `Note:` In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

    `Note:` It is not possible to declare a variable using :=, without assigning a value to it.

# Difference Between var and :=
There are some small differences between the var var `:=` -

| var                                        | :=                                   |
|--------------------------------------------|--------------------------------------|
|Can be used inside and outside of functions | Can only be used inside functions
| Variable declaration and value assignment can be done separately|Variable declaration and value assignment cannot be done separately (must be done in the same line)


# Go Variable Declaration in a Block
Multiple variable declarations can also be grouped together into a block for greater readability:

Example -
```
package main
import "fmt"

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```

# Go Constants
If a variable should have a fixed value that cannot be changed, you can use the `const` keyword.

The `const` keyword declares the variable as "constant", which means that it is unchangeable and read-only.

Example - `const PI = 3.14`

`Note:` The value of a constant must be assigned when you declare it.

# Constants: Unchangeable and Read-only
When a constant is declared, it is not possible to change the value later:

Example - 
```
package main
import "fmt"

func main() {
  const A = 1
  A = 2
  fmt.Println(A)
}
```

Result:

`./prog.go:8:7: cannot assign to A`

# Multiple Constants Declaration
Multiple constants can be grouped together into a block for readability:

Example -

```
package main
import "fmt"

const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}
```
# Go Data Types 
### [Link for more details](https://go.dev/ref/spec#Types)

Data type is an important concept in programming. Data type specifies the size and type of variable values.

Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.

Go has three basic data types:

- `bool`: represents a boolean value and is either true or false
- `Numeric`: represents integer types, floating point values, and complex types
- `string`: represents a string value

Go Floating Data Types
- The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.

- The float data type has two keywords:

| Type                  | Size            | Range       |
|-----------------------|-----------------|-------------|
float32                 | 32 bits         |-3.4e+38 to 3.4e+38.
float64                 | 64 bits         |-1.7e+308 to +1.7e+308.

`Tip`: The default type for float is float64. If you do not specify a type, the type will be float64.

### Numeric types

- `uint8`    -    the set of all unsigned  8-bit integers (0 to 255)
- `uint16`   -    the set of all unsigned 16-bit integers (0 to 65535)
- `uint32`   -    the set of all unsigned 32-bit integers (0 to 4294967295)
- `uint64`   -    the set of all unsigned 64-bit integers (0 to 18446744073709551615)

- `int8`     -   the set of all signed  8-bit integers (-128 to 127)
- `int16`    -   the set of all signed 16-bit integers (-32768 to 32767)
- `int32`    -   the set of all signed 32-bit integers (-2147483648 to 2147483647)
- `int64`    -   the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

- `float32`     the set of all IEEE 754 32-bit floating-point numbers
- `float64`     the set of all IEEE 754 64-bit floating-point numbers

- `complex64`   the set of all complex numbers with float32 real and imaginary parts
- `complex128`  the set of all complex numbers with float64 real and imaginary parts

- `byte`        alias for uint8
- `rune`        alias for int32


# Important Format Specifiers to remember 
### [Link for more details](https://pkg.go.dev/fmt@go1.27.4)

- `%d` → Integer
- `%f` → Float
- `%s` → String
- `%t` → Boolean
- `%T` → Data type
- `%v` → Any value in its default format
- `%p` → Pointer address
- `%q` → Quoted string


## Segregating types with `pass-by-value` & `pass-by-reference`
Go makes `copies` of values when passed into functions.

Non-Pointer Values | Pointer Wrapper Values
--                 | --
Strings            | Slices
Ints               | Maps
Floats             | Functions
Booleans           |
Arrays             |
Structs            |


## Structs
- A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

- While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.
- A struct can be useful for grouping data together to create records.

### Declare a Struct

```
  type struct_name struct {
    member1 datatype;
    member2 datatype;
    member3 datatype;
    ...
  }
```

### Example - 
```
  type Person struct {
    name string
    age int
    job string
    salary int
  }
```

#### Best practices

- Use UserProfile, not userProfile, for exported struct types.
- Use userProfile, not UserProfile, for unexported struct types.
- Use meaningful, concise names.
- Avoid underscores in struct names.
- Go uses capitalization to determine visibility:
  - Starts with uppercase → exported.
  - Starts with lowercase → unexported.

`Recommendation: Follow Go's convention: PascalCase for exported structs and lowercase for unexported structs.`

## Interfaces

An interface is a type that defines a set of method signatures. 

Any type that implements all the methods declared by an interface automatically satisfies that interface.

The most important thing to understand is that Go interfaces are implemented implicitly. Unlike Java or C++, you don't need to explicitly declare that a struct implements an interface.

### Why do we need interfaces?

Imagine you have two types: Dog and Cat.

Both can make sounds, but they make different sounds.

Instead of writing separate functions for each animal, you can define a common interface called Animal.

### Basic interface example

```
package main

import "fmt"

// Define an interface
type Animal interface {
    Sound()
}

// Dog struct
type Dog struct{}

func (d Dog) Sound() {
    fmt.Println("Dog barks")
}

// Cat struct
type Cat struct{}

func (c Cat) Sound() {
    fmt.Println("Cat meows")
}

// Function accepting an interface
func MakeSound(a Animal) {
    a.Sound()
}

func main() {
    dog := Dog{}
    cat := Cat{}

    MakeSound(dog)
    MakeSound(cat)
}
```

Output - 

```
Dog barks
Cat meows
```

### What's happening here?
- We created an Animal interface with one method: Sound().
- Dog implements the Sound() method.
- Cat also implements the Sound() method.
- Both automatically satisfy the Animal interface.
- The MakeSound() function accepts any value that implements Animal.

Notice that we never wrote something like:

```
  // Go doesn't require this
  implements Animal
```
That's implicit interface implementation.

### How interfaces work internally

Consider this:

```
  type Shape interface {
      Area() float64
  }

  type Circle struct {
      Radius float64
  }

  func (c Circle) Area() float64 {
      return 3.14 * c.Radius * c.Radius
  }
```
Since `Circle` has the `Area() float64` method, it automatically implements `Shape`.

You can use it like this:
```
  func PrintArea(s Shape) {
      fmt.Println(s.Area())
  }

  func main() {
      c := Circle{Radius: 5}
      PrintArea(c)
  }
```
Output - 
```
78.5
```

### Empty interface (`interface{}`)
An empty interface has no methods.

```
interface{}
```

Because it requires no methods, any value can be assigned to it.

Example:
```
  func PrintValue(value interface{}) {
      fmt.Println(value)
  }

  func main() {
      PrintValue(10)
      PrintValue("Hello")
      PrintValue(true)
  }
```
In modern Go, you can use `any` instead of `interface{}`.

```
  func PrintValue(value any) {
      fmt.Println(value)
  }
```

`Note` - A type can satisfy multiple interfaces simultaneously. 