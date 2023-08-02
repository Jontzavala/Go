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
## Creating and Debugging Programs
- CLI Tools
    - OS package allows you to interact with the operating system
    - stdin (keyboard)
    - stdout (monitor)
    - stderr (allows you to write error information out)
    - fmt package
    - Scan* (read information in from various sources)
    - Print*
    - bufio package (buffered input output, allows you to gather group of text together for more useful units)
- Demo: Building a CLI Application
    - see building_cli_application.go
- Web Services
    - Two computers that are talking to eachother
    - The client (laptop) send a Request to the server and the server sends a response back to the client
    - Inside the server there is a front controller pattern, and then there are back controllers pattern.
    - The front controller recieves all the request and then determines which back controller to send it to.
    - Then after the back controller processes the request it sends it back to the front controller where the front controller send the response back to the client.
    - all of the functionality above is provided in one Go package. https://pkg.go.dev/net/http
- Demo: Building a Web Service
    - see building_web_service.go
    - Go provides front controllers for us
- Demo: Debugging a Program
    - see debugging_a_program.go
    - In VScode click on the play button with the little bug
    - click the create a launch.json file
    - click to the left of the number on the file you want to test, this will place your debugger on that line
    - Hit the play button next to Run and Debug to run the debugger
    - Call Stack allows you to walk back through your code to see how you go to the debugger
    - Watch allows you to watch certain variables if you'd like
    - You can hover over any variable to see what its been through until it hit the debugger.
## Aggregate Data Types
- Concept: Array Types
    - An array is a fixed size collection of data element of all the same type
- Creating and Using Arrays
    - ```var arr [3]int  //array of 3 ints```
    - ```fmt.Println(arr)    // [0 0 0]```
    - ```arr = [3]int{1, 2, 3}``` array literal
    - ```fmt.Println(arr[1])     // 2```
    - ```arr[1] = 99``` updated the value
    - ```fmt.Println(arr)    // [1 99, 3]``` changed the value at index 1 with the above line.
    - ```fmt.Println(len(arr))``` returns the length of the array.
    - ```arr := [3]string{"foo", "bar", "baz"}```
    - ```arr2 := arr```
    - ```fmt.Println(arr2)     // {"foo" "bar" "baz"}```
    - ```arr[0] = quux``` changing the first value of arr
    - ```fmt.Println(arr)      // {"quux" "bar" "baz"}```
    - ```fmt.Println(arr2)     // {"foo" "bar" "baz"}``` arr2 is still the original arr even after the first item in arr was changed
    - arr2 is a copy of arr
    - ```arr == arr2   false - arrays are comparable```
- Demo: Arrays
    - see array.go
- Concept: Slice Types
    - has the ability to grow and shrink
    - slice points to data in some array some where
    - if the array is changed the slice will automatically update
    - if a value is editied in the slice the underlying array will be updated
- Creating and Using Slices
    - ```var s []int    // slices of ints```
    - ```fmt.Println(s)   // [] (nil)```
    - A declared but uninitialized slice will return nil
    - ```s = []int{1, 2, 3}```    // slice literal
    - ```fmt.Println(s[1])   // 2```
    - ```s[1] = 99   // update value```
    - ```fmt.Println(s)   // [1 99 3]```
    - ```s = append(s, 5, 10, 15)   // add elements to the slice```
    - ```fmt.Println(s)   // [1 99 3 5 10 15]```
    - ```s = slices.Delete(s, 1, 3)   // removes indices 1, 2 from slice (golang.ord/x/exp/slices)```
    - ```fmt.Println(s)   // [1 5 10 15] ```
- Demo: Slices
    - see slices.go
- Concept: Map Types
    - value {string}
    - key {string}
    - Dynamically sized like a slice
- Creating and Using Maps
    - ```var m map[string]int    // declare a map```
    - ```fmt.Println(m)    // map[] (nil)``` just like a slice a map starts with nil, which means its a referece type like a slice meaning it doesn't have it's own data.
    - ```m = map[string]int{"foo": 1, "bar": 2}``` map literal
    - ```fmt.Println(m["foo"])    // lookup value in map```
    - ```m["bar"] = 99    // update value in map```
    - ```delete(m, "foo")    // remove entry from map```
    - ```m["baz"] = 418    // add value to map```
    - ```fmt.Println(m["foo"])    // 0 - queries always return results```
    - ```v, ok := m["foo"]    // comma okay syntax verifies presents```
    - ```fmt.Println(v, ok)    // 0, false```
    - Updating maps
    - ```m := map[string]int {"foo":1, "bar":2, "baz":3}```
    - ```m2 := m```
    - ```m["foo"], m2["bar"] = 99, 42    // update values in map```
    - ```fmt.Println(m)    //map[foo:99 bar:42 baz:3]```
    - ```fmt.Println(m2)    // map[foo:99 bar:42 baz:3, data is shared```
    - maps are copied by reference
    - use maps.Clone to clone
    - ```m == m2    //compile time error - maps are not comparable```
- Demo: Maps
    - see maps.go
- Concept: Struct Types
    - fixed size aggregate type but it's values can be different types
    - heterogenious
- Creating and Using Structs
    - ```var s struct {name string id int}   // declare an anonymous struct```
    - ```fmt.Println(s)    // {"" 0}```
    - ```s.name = "Arthur"    //assign value to field```
    - ```
        type myStruct struct {
        name string
        id int
        }    //create custom type based on struct
        ```
    - ```var s myStuct    // declare variable with custom type```
    - ```fmt.Println(s)    // {"" 0}```
    - ```
        s = myStruct{
            name: "Arthur",
            id:42}    // Struct literal
        ```
    - ```fmt.Println(s)    // {"Arthur" 42}```
    - ```s2 := s```
    - ```s.name = "Tricia"```
    - ```fmt.Println(s, s2)    // {"Tricia" 42} {"Arthur" 42}``` changing s doesn't change s2, structs are copied by value, value type
    - ```s == s2    // false - structs are comparable```
- Demo: Structs
    - see structs.go
## Looping
- Concept: Looping
    - Every loop is a for loop in Go
    - ```
        func main() {
            statement 1
            loop {
                statement 2
                statement 3
            }
            statement n
        }    // looping execution
        ```
- Basic Loops
    - ```for {...}    //infinite loop```
    - ```for condition {...}    // loop till condition```
    - ```for initializer; test; post clause {...}    // counter-based loop```
    - initializer allow us to set up the loop for whatever we need to do
    - test which is the same as the condition in the conditional loop
    - Infinite Loops
    - ```
        i := 1
        for {
            fmt.Println(i)
            i += 1
        }
        ```
    - Loop till Condition
    - ```
        i := 1
        for i < 3 {
            fmt.Println(i)
            i += 1
        }
        fmt.Println("Done!:)
        ```
    - Counter-base Loop
    - ```
        for i:=1; i < 3; i++ {
            fmt.Println(i)
        }
        fmt.Println("Done!")
        ```
    - The first statement go is fgoing to look at is the initalizer (i:= 1) it only going to run one time
    - Then go is going to check the condition (i < 3) if condition is true Go enters the loop and prints out ```i```
    - After the loop body executes the Go will move to the post clause (i++), it's typically used to increment a counter in the loop.
    - i++ is the increment operator in this loop
    - After incrementing the Go goes back to the condition and checks it and continues until condition is met.
- Demo: Looping
    - see looping.go
- Looping through Collections
    - ```for key, value := range collection {...}    // a collection can be an array, slice or map```
    - ```for key := range collection {...}    // you can pull back a single value```
    - ```for _, value := range collection {...}    // if you just want the values, "_" is the blank identifier```
    - How does it work?
    - ```
        arr := [3]int{101, 102, 103}
        for i, v := range arr {
            fmt.Println(i, v)
        }
        fmt.Println("Done!")    // ["0, 101" "1, 102" "2, 103"]
        ```
- Demo: Looping through Collections
    - see looping_through_collections.go
    - To learn more about formatting verbs go to pkg.go.dev/fmt