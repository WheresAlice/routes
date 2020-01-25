package main

import (
	"fmt"
	"github.com/wheresalice/routes"
)

//PrintHandler a function which does nothing except printing the string it is provided
func PrintHandler(input string) {
	fmt.Println(input)
}

func main() {
	// Create a router
	router := routes.NewRouter()

	// Add a route which responds to the full string "alice" and returns the function PrintHandler when called
	router.AddRoute("^alice$", PrintHandler)

	// Route the input of "alice", returning the handler associated with the matching regexp
	a := router.Exec("alice")
	// Run the handler that was returned, with the original input
	a("alice")

	// Route the input of "bob"
	b := router.Exec("bob")
	// Since bob doesn't match any pattern, we get a nil back.  Which means we need to handle it.
	if b != nil {
		b("bob")
	}
}
