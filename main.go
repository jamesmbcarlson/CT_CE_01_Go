// every Go project needs the main package up top
package main

import (
	// "errors"
	"fmt"
	"log"
	"net/http"
	// "gotutorial/testpackage"
)

// every Go file needs a main to run
func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	log.Println("Listening on localhost:4242...")

	var err error =	http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}
	
	
	// calling test functions
	// myFunction();
	// testpackage.TestFunction()
	// var caughtValue string = returnsValue("Hello testing 123")
	// fmt.Println(caughtValue)
	// var err error = returnsError("blehcafjaldfj")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var names []string = []string{"James", "Meggy", "Peaches", "Wendell"}
	// fmt.Println(names)
	// fmt.Println(names[1])
}



func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
	// fmt.Println("Endpoint Called!")
	if request.Method != "POST" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	response := []byte("Server is up and running!")
	_, err := writer.Write(response) // underscore - we're ignoring the first returned value
	if err != nil {
		fmt.Print(err)
	}
}

// local test functionS

// func myFunction() {
// 	var james = "Hello World";
// 	fmt.Println(james);
// }

// how to specify that this method returns a specific type
// func returnsValue(something string) string {
// 	return "something " + something
// }

// func returnsError(password string) error {
// 	var secretPassword string = "supersecretpassword"
// 	if password == secretPassword {
// 		return nil
// 	} else {
// 		return errors.New("invalid password")
// 	}
// }

// func otherFunction() {
// 	someString, someInt, someBool := returnsMultiple()
// 	thisIsAString := someString // with := we don't have to write var anymore!
// }

// func returnsMultiple() (string, int, bool){
// 	return "string", 1, true
// }