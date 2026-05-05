package main
import "fmt"

// Interface -- an interface defines behavior(methods) not data

// It says any type that has these methods can be used here 

type Speaker interface {
	Speak() string 
} 

type Dog struct {}

func (d Dog) Speak() string  {
	return "Woof"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return  "Meow";
}

func makeSounds(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
    makeSounds(Dog{})
}