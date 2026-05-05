// Enums = a set of named constant values
// no dedicated enum keyword.

// Why enums are used ---> To restrict a variable to a fixed set of valid values

// Instead of letting anything go in you define:
// 👉 “only these values are allowed

package main

import "fmt"


type Status int 

// we can define strings also instead of int -- iota is a concept in go. So the subsequent values will increase

const (
	Pending = iota
	Approved
	Rejected 
)

func changeStatus(status Status) {
	fmt.Println("current status is ", status)
}

func main() {
     changeStatus(Pending)
}