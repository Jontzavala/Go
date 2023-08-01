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