// every Go project needs the main package up top
package main

import (
	"fmt"
	"gotutorial/testpackage"
)

// every Go file needs a main to run
func main() {
	myFunction();
	testpackage.TestFunction()
}

func myFunction() {
	var james = "Hello World";
	fmt.Println(james);
}