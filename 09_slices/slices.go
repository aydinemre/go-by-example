// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
package main

import "fmt"

func main()  {

	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	// To create an empty slice with non-zero length, use the builtin make.
	// Here we make a slice of strings of length 3 (initially zero-valued).
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices support several more that make them richer than arrays.
	// One is the builtin append, which returns a slice containing one or more new values.
	// Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.



}
