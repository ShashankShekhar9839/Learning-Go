// A goroutine is a lightweight thread managed by Go.
// It lets you run functions concurrently (multiple tasks at the same time).

package main
import (
	"fmt"
	"time"
)


/*   normally: 

doTask1()
doTask2()  runs one after another (swquential) 

BUT WITH GOROUTINES 

go doTask1() 
go doTask2()   both run concurrently 

---> Real use cases

handling API requests
background jobs
parallel processing
file downloads
microservices

*/

func sayHello() {
	fmt.Println("Hello");
}

func main() {
   go sayHello()   // run as a goroutine 
   time.Sleep(1 * time.Second)
 }