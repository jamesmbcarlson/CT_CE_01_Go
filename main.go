// every Go project needs the main package up top
package main

import (
	// "errors"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
	// "gotutorial/testpackage"
)

// every Go file needs a main to run
func main() {
	stripe.Key = "sk_test_51Jkd2bDg4n49xQRDGH4JezLpg1ErReLCoHlYjfiwCRALuuk4KbVIqcIRZxRRd6Jvk6qjirQg4bCPwxUQEF9Z9gin00F07ahoPT"

	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	log.Println("Listening on localhost:4242...")

	var err error = http.ListenAndServe("localhost:4242", nil)
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

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		// log.Println(err)
		return
	}

	// fmt.Println(req.FirstName)
	// fmt.Println(req.LastName)
	// fmt.Println(req.Address1)
	// fmt.Println(req.Address2)
	// fmt.Println(req.City)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(req.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(paymentIntent.ClientSecret)

	var response struct {
		ClientSecret string `json:"clientSecret"`
	}

	response.ClientSecret = paymentIntent.ClientSecret

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")

	_, err = io.Copy(writer, &buf)
	if err != nil {
		fmt.Println(err)
	}

}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	response := []byte("Server is up and running!")
	_, err := writer.Write(response) // underscore - we're ignoring the first returned value
	if err != nil {
		fmt.Print(err)
	}
}

func calculateOrderAmount(product_id string) int64 {
	switch product_id {
	case "Forever Pants":
		return 26000
	case "Forever Shirt":
		return 15500
	case "Forever Shorts":
		return 30000
	}
	return 0
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
// 	thisIsAString := someString // with := we don't have to write var anymore! - this declares and assigns
// }

// func returnsMultiple() (string, int, bool){
// 	return "string", 1, true
// }
