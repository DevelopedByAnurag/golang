package main

import "fmt"

func main()  {
	dataSlice := []string{"var1", "var2"}
	fmt.Println(dataSlice)
	// Append a new value
	dataSlice = append(dataSlice, "new value")
	fmt.Println(dataSlice)
	// Check loops Dir to check how to iterate over Slice
}