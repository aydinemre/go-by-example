/*
This first program will print the classic “hello world” message.

To run the program use go run.
- go run hello-world.go

Sometimes we’ll want to build our programs into binaries. We can do this using go build.
- go build hello-world.go

We can then execute the built binary directly.
- ./hello-world

 */
package main

import "fmt"

func main()  {
	fmt.Println("hello world")
}

