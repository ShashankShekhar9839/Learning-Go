package main

import "fmt"

// embeded struct 

type Address struct {
	City string
}

type User struct {
    id       int
    userName string
    isActive bool
	Address
}

func (u User) greet() string {
	return  "Hello " + u.userName;
}

// production way of using struct. --> Using constructor function 

func NewUser(id int, name string, isActive bool, city string) User {
	return  User {
		id: id, 
		userName: name,
		isActive: isActive,
		Address: Address{
			City: city,
		},
	}
}

func main() {

    user := User{
        id:       1,
        userName: "shashank",
        isActive: true,
    }

	user2 := User{
		id: 2, 
		userName: "shekhar", 
		isActive: false,
	}

    // anonymous struct
    tempConfig := struct {
        port int
        env  string
    }{
        port: 8080,
        env:  "production",
    }

	userUsingConstructor := NewUser(3, "yadav", true, "Bangalore")

	user.City = "Varanasi"; 

    fmt.Println("Shashank", user, tempConfig);
	fmt.Println("Shekhar", user2);
	fmt.Println(user.greet());
	fmt.Println(user2.greet());
	fmt.Println(userUsingConstructor);
}