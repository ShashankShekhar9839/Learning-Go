package main 
import "fmt"

func main() {
    
	// uninitialized slice is nil in go 
   var nums [] int
   nums = append(nums, 2)
   nums =  append(nums, 4)

   fmt.Println(nums)
   fmt.Println(len(nums))

}