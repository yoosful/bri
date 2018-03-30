package serve

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetDevices",
		"GET",
		"/device",
		GetDevices,
	},
	Route{
		"GetUser",
		"GET",
		"/user",
		GetUser,
	},
	Route{
		"GetDevice",
		"GET",
		"/device/{did}",
		GetDevice,
	},
	// Route{
	// 	"CreateDevice",
	// 	"POST",
	// 	"/device/{did}",
	// 	CreateDevice,
	// },
	Route{
		"DeleteDevice",
		"DELETE",
		"/device/{did}",
		GetDevice,
	},
}
