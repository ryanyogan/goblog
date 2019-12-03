package service

import "net/http"

// Route defines a single route and the shit
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the the type Routes which is a slice of Route
type Routes []Route

var routes = Routes{
	Route{
		"GetAccount",
		"Get",
		"/accounts/{accountID}",
		GetAccount,
	},
}
