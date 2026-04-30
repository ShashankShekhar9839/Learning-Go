package main

import "fmt"

// function definition
func sum(a int, b int) int {
    return a + b
}

func main() {
	var result int
    result = sum(3, 5)
    fmt.Println(result)
}

