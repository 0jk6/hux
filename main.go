package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "home")
}

func greetingsHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "greetings")
}

func testHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "post req test")
}


func main(){
	DEBUG = true //to pring logs
	hx := &Hux{}

	hx.Route("GET", "/", homeHandler)
	hx.Route("GET", "/greetings", greetingsHandler)
	hx.Route("POST", "/test", testHandler)

	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", hx)
}