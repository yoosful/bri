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
		"GetDevice",
		"GET",
		"/device/{id}",
		GetDevice,
	},
	Route{
		"CreateDevice",
		"POST",
		"/device/{id}",
		CreateDevice,
	},
	Route{
		"DeleteDevice",
		"DELETE",
		"/device/{id}",
		GetDevice,
	},
}
