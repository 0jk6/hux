package main

import (
	"fmt"
	"net/http"
)


//to pring debug logs
var DEBUG bool = false

func debugLog(msg string){
	if DEBUG == true{
		fmt.Println(msg)
	}
}

/*
the following struct will store the route entries
eg: /getData, /api/v1/authorize, etc.

it will have, URL Path, Request Method and a HandlerFunc to handle incoming requests
*/
type RouteEntryTable struct {
	Path string
	Method string
	handlerFunc http.Handler
}


//following method will check if there's an existing route entry matching with incoming request
func (routeEntry *RouteEntryTable) RouteExists(r *http.Request) bool {
	if r.Method == routeEntry.Method && r.URL.Path == routeEntry.Path {
		return true
	}

	return false
}


//let's create a router struct with a name Hux and initalize route entires
type Hux struct {
	routes []RouteEntryTable
}


//following method will add/append new routes to the existing routes array
func (hx *Hux) Route(method, path string, fn http.HandlerFunc){
	
	debugLog(fmt.Sprintf("Added \"%s %s\"", method, path))

	newRouteEntry := RouteEntryTable{
		Path: path,
		Method: method,
		handlerFunc: fn,
	}

	hx.routes = append(hx.routes, newRouteEntry)
}


//all incoming requests will go through this following request, refer to http.Handler
func (hx *Hux) ServeHTTP(w http.ResponseWriter, r *http.Request){
	//loop through all the stored route entries in the route entry table and check
	//if any matching route exists, if yes, then serve the request
	for _, routeEntry := range hx.routes {
		if routeEntry.RouteExists(r) == true {
			//call the ServeHTTP method attached to the handlerFunc in the route entry
			debugLog(fmt.Sprintf("\"%s, %s\"", r.Method, r.URL.Path))
			routeEntry.handlerFunc.ServeHTTP(w, r)
			return
		}
	}

	//if nothing matches, return 404 not found
	debugLog(fmt.Sprintf("404 Not Found \"%s, %s\"", r.Method, r.URL.Path))
	http.NotFound(w, r)
}