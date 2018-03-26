package serve

type Device struct {
	Type string
	Id   string
	On   bool
}

var Devices []Device
