package main

import (
	"errors"
	"fmt"
	"time"
)

//structs allows us to manipulate a collection of data more easy

type User struct { //<======== user type is an struct now
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}
func (uStruct User) outputUserData(){
	//this will attach a method (function) into the User struct
	fmt.Printf("\nFirst Name: %v\n", uStruct.firstName)
	fmt.Printf("Last Name :%v\n", uStruct.lastName)
	fmt.Printf("Birth Date:%v\n", uStruct.birthdate)
	fmt.Printf("Created At:%v\n", uStruct.createdAt)

}
func (uStruct *User) clearUserName(){
	//pointer receiver, it can modify the value inside the struct, to make shure that what we creating or editing in a strcut always use the pointer to access to the original struct data and not a copy as we know
	uStruct.firstName =""
	uStruct.lastName = ""
}
func newUser (firstname, lastname, birthdate string) (*User,error){
	if(firstname == "" || lastname == "" || birthdate == "" ){
		return nil, errors.New("All parameters need to be fill")
	}

	return &User{firstName: firstname , lastName: lastname , birthdate: birthdate , createdAt: time.Now()}, nil
}
func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *User
	
	appUser, error := newUser(firstName,lastName,birthdate)
	// ... do something awesome with that gathered data!
	if(error != nil) {
		fmt.Println((error))
		return
	}
	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData()
}
// func outputUserData(uStruct *User){
// 	// go allows this shortcut for struct data, when passing pointer, is not necessary derefference with the * like *uStruct
// 	fmt.Printf("\nFirst Name: %v\n", uStruct.firstName)
// 	fmt.Printf("Last Name :%v\n", uStruct.lastName)
// 	fmt.Printf("Birth Date:%v\n", uStruct.birthdate)
// 	fmt.Printf("Created At:%v\n", uStruct.createdAt)

// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
