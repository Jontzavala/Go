package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	// Used to register the back controller. http package is to work with web services.
	http.HandleFunc("/", Handler)

	// Starts the server, nil allows us to use a front controller created by Go, instead of making one.
	http.ListenAndServe("localhost:3000", nil)
}

/* This function is creating a back controller, http.ResponseWrites allows the back controller to send its response to the front controller.
The *http.Request allows the back controller to get the response from the front controller. */
func Handler(w http.ResponseWriter, r *http.Request){

	// This uses the Open function from the os package to read the menu from menu.txt and it's stored in the variable f
	f, _ := os.Open("./menu.txt")
	
	// io.Copy allows you to copy from a read source to write source. Takes the contents of the menu file and copies it to the response
	io.Copy(w, f)
}