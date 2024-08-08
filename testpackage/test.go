package testpackage

import "fmt"

// first letter capitalized means this is public
func TestFunction() {
	fmt.Println("Step 1")
	fmt.Println("Step 2")
	fmt.Println("And so on")
	var password string = "Coconut"
	myPrivateFunction(password)
}

// first letter uncapitalized means this is private
func myPrivateFunction(secret string) {
	fmt.Println("This is private method!!")
	fmt.Println("Secrete word: " + secret )
	if secret == "Banana" {
		fmt.Println("Mmm, yummy password!")
	} else if secret == "Apple" {
		fmt.Println("YUCK! I don't like that password") 
	} else {
		fmt.Println("What are you even talking about?")
	}
}