// Generics let you write one function or type that works with multiple data types 

package main 

import "fmt"

// without generics we have create two functions one for int and one for float; 
// But with generics we can handle this with a single function 

func sumInt(a, b int) int {
	 return  a + b;
}

func sumFloat(a, b float32) float32 {
	return a+ b;
}

func sum[T int | float64](a, b T) T {
	 return  a + b;
} 

func main() {
     fmt.Println(sumInt(4,6));
	 fmt.Println(sumFloat(4.2, 6.0)); 
	 fmt.Println(sum(4, 6));
	 fmt.Println(sum(4, 0.9));
}