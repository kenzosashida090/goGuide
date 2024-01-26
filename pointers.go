package main

import "fmt"
func main (){
	age :=32 
	var	agePointer *int
	agePointer = &age
	//Dereferencing the pointer to access the value of age through the pointer.
	fmt.Println("Age Pinter Address is :",*agePointer)
	fmt.Println("Age is variable value is :",age)
	getAdultYears(agePointer)
	fmt.Println(age)
	// *agePointer  variable value
	// &agePointer adrress value
}

func getAdultYears(age *int) {
	fmt.Println(&age)
	*age = *age -18;
}