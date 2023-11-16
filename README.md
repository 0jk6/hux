# hux

http router in golang similar to mux, but with less features

no need to install this, just copy the file `hux.go` and use it in your projects


### Usage
```
func main(){
	DEBUG = true //to print logs
	hx := &Hux{}

	hx.Route("GET", "/", homeHandler)
	hx.Route("GET", "/greetings", greetingsHandler)
	hx.Route("POST", "/test", testHandler)

	fmt.Println("listening on port 8080")

	http.ListenAndServe(":8080", hx)
}
```

checking for existing routes can be optimized with the help of a Trie data structure, I'm planning to add it in the future