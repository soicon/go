package service

import "net/http"

// Defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routes = Routes{

	Route{
		"GetBrand",                                     // Name
		"GET",                                            // HTTP method
		"/brands/",                          // Route pattern
		GetBrand,
	},{
		"GetAllProduct",                                     // Name
		"GET",                                            // HTTP method
		"/brands/{brand}/products/",                          // Route pattern
		GetAllProduct,
	},
	{
		"GetProduct",                                     // Name
		"GET",                                            // HTTP method
		"/brands/{brandId}/products/{productId}",                          // Route pattern
		GetProduct,
	},
	{
		"AddProduct",                                     // Name
		"POST",                                            // HTTP method
		"/new",                          // Route pattern
		AddProduct,
	},
}