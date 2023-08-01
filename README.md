# Go: The Big Picture

## Why Go?
- Why Create Go
    - Created within Google
    - Go has efficient compilation, efficient execution and ease of programming
- Guiding Principles
    - Simplicity as a core value
    - Commitment to backward compatibility
    - Made for concurrent, network enabled applications
    - Holistic approach to solve common problems
- Language Characteristics
    - Strong static type system
    - C-inspired syntax
    - Garbage collected
    - Fully compiled
    - Rapid compilation
    - Single binary output
    - Cross platform
- Primary Use Cases
    - Used for cloud and network services
    - Used for command-line interfaces
    - Used for Cloud infrastructure (Docker, Kubernetes, Terraform)
## Programming with Go
- Language Characteristics and Evolution
    - Starts with C programming language
    - Modern C (Faster computers, Multiple CPUs, Network aware)
- Demo: Hello, World
    - See hello_world.go file for code.
    - Everything in go is contained within packages
    - Create Entry point (func main() {})
    - Package fmt to print strings
    - "go run (name of file)" to run and see result
- Demo: A Simple Web Service
    - Writers are interfaces that can be written to.
- Demo: CLI Application
    - flag package strongly type command line paramaters as flags
    - "*" is a pointer
## Leveraging the Go Ecosystem
- Development Tools
    - VScode, Goland, Vim, Emacs, Eclipse
    - Gopls unifies the tools
- Common Frameworks
    - Go Kit (Comprehensive microservice framework)
    - Gin (Fast, lightweight web framework)
    - Gorilla Toolkit (Collection of useful tools without framework lock-in)
    - Cobra (Framework for building command-line interface applications)
    - Standard library (Zero dependency APIs for managing simple CLIs)
    - Docker (Containerize applications to simplify deployment)
    - Kubernetes (System that builds, deploys and scales containerized apps)
    - Terraform (Cloud infrastructure management platform)
- Overview of Go's website
    - go.dev
    - Docs
    - Effective Go (Describes how Go code is written)
    - Playground
    - Packages
    - /std (takes you to standard library)
- Learning Resources
    - Pluralsight
    - The Go Blog
    - github.com/golang/go/wiki
    - go.dev/learn
- Online Community
    - Gophers Slack Channel
    - Go discord channel
    - Go reddit page
- Conferences
    - go wiki /conferences

# Go Fundamentals

## Variables and Simple Data Types
- Simple Data Types
    - Strings
    - Numbers
    - Booleans
    - Errors
- The String Type
    - quoted string "Hello" (interpereted string)
    - backtick mark `Hello` (raw string)
    - "\n" in a quoted string makes a new line but with backtick string it's not interpereted.
    - raw strings ingnore new lines
- Numeric Types
    - Integers (Int)
    - Unsigned integers (uint) lowest it can go is zero
    - Floating point numbers (float32, float64)
    - Complex numbers (complex64, complex128)
- Boolean Types
    - true
    - false
    - usually used in a comparision operation
- Error Types
    - The error built-in interface type is the conventional interface for representing an error condition, with the nil value representing no error
    - simple definition: An error occured
    - interface, has the ability to tell you want the error was.
    - Use the error function to tell you were the error is.
- Finding Documentation for Built-in Types
    - pkg.go.dev/std
    - go to bultin package (It's documentation only)
    - Shows the bultin functionality of Go
    - Also shows the different types (just scroll down further)
- Declaring Variables
    - var myName string (declare variable, commonly used)
    - var myName string = "Mike" (declare and initialize, commonly used)
    - var myName = "Mike" (initialize with inferred type, rarley used)
    - myName := "Mike" (short declaration syntax, most used used)
- Type Conversions
    - Go doesn't support implicit conversions (```var i int = 32 var f float32     f = i```  output: error)
    - ```f = float32(i)``` (Type conversions allow explicit conversions)
- Demo: Using Simple Types and Type Conversions
    - see simple_types_and_conversions.go
- Common Arithmetic Operators
    - ```a, b := 10, 5``` (Go allows multiple variable to be initialized at once)
    - ```c := a + b``` (Addition, output is: 15)
    - ```c = a - b``` (Subtraction, output is 5. You don't have to use the walrus opperator because ```c``` was already initialized.)
    - ```*``` for multiplication
    - ```/``` division
    - ```a/3``` integer division
    - ```%``` remainder of division (Modulus operator)
- Common Comparision Operators
    - ```a, b := 10, 5```
    - ```c := a == b``` the ```==``` operator means equals and in this specific case it's false.
    - Errors don't support comparision operators
    - ```!=``` is the not equal to operator
    - ```<, <=, >, >=``` Less than, less than or equal too, greater than, greater than or equal too.
- Demo: Using Arithmetic and Comparision Operators
    - see using_arthmetic_and_comparison_operators.go
    - Go to the go docs and go to language specifications
    - Find arithmetic and comparision operator options
- Constants, Constant Experssions and Iota
    - constants cannot change its value
    - ```const a = 42``` constant implicitly typed (treated as a literal value, can be a string, int, etc)
    - ```const b string = "hello, world"``` explicitly typed constant
    - ```const c = a``` one constant can be assigned to another.
    - ```const ( d = true e = 3.14)``` You can group constants so you don't have to keep declaring constants within your code.
    - ```const (a = "foo" b  //"foo")``` unassigned constants recieve previous value. So because b was not given a value, it took on the value of a.
    - ```const c = 2 * 5``` constant expression
    - ```const d = "hello, "+"world" ``` must be calculable at compile time
    - ```const e = someFunction()``` this won't work it can't be evaluated at compile time.
    - ```const a = iota    //0``` iota is related to position in constant group
    - ```const (b = iota  //0    c   // 1    d = 3 * iota   //6)``` iota starts at zero on first line, constant expression copied iota increments, iota increments again
    - If you go into another constant block iota restarts at 0
- Demo: Constants, Constant Experssions and Iota
    - see constants_constant_expressions_iota.go
    - go.dev/doc/effective_go#constants to see iota examples.
- Pointers and Values
    - ```a := 42``` creating ```a``` variable a with the value of 42
    - ```b := a``` creating variable of ```b``` and assiging it to the current value of ```a``` (42)
    - ```a = 27``` changing the value of ```a``` to 27 but the value of ```b``` stays the same at 42
    - With pointers it's different
    - ```a := 42``` creating a variable ```a``` with the value of 42
    - ```b := &a``` declaring variable ```b``` and having it point to ```a``` the ```&``` takes the address of ```a```. ```b``` is now a pointer
    - ```*b   //42``` dereferencing the pointer to see what value the pointer is pointing to.
    - ```a = 27``` changing the value of ```a```
    - ```*b   //27``` Now ```b``` is 27 just like ```a``` this is what a pointer does, it holds memory address.
    - ```a := "foo"``` Create a string variable
    - ```b := &a``` address operator returns the address of the ```a``` variable
    - all pointers are strongly typed in Go
    - ```*b = "bar"``` dereference a pointer with asterisk, dereferencing ```b``` to change the value of ```a```
    - ```c = new(int)``` built-in "new" function creates pointer to anonymous variable
    - POINTERS ARE PRIMARILY USED TO SHARE MEMORY. USE COPIES WHENEVER POSSIBLE!
    - Using copies avoids data races because it's using shared memory.
- Demo: Creating and Using Pointers
    - see creating_and_using_pointers.go