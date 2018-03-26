package serve

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/index.html"))
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func GetDevices(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Devices)

	fmt.Println("Get info of all devices")

}

func GetDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, device := range Devices {
		if device.Id == params["id"] {
			json.NewEncoder(w).Encode(device)
			fmt.Println("Get info of ", device.Type, "with id ", device.Id)
			return
		}
	}
	json.NewEncoder(w).Encode(&Device{})

}

func CreateDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, device := range Devices {
		if device.Id == params["id"] {
			fmt.Println(device.Type, "with id ", device.Id, " already exists")
			return
		}
		json.NewEncoder(w).Encode(Devices)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var device Device
	_ = json.NewDecoder(r.Body).Decode(&device)
	device.Id = params["id"]
	device.Type = params["type"]
	Devices = append(Devices, device)
	json.NewEncoder(w).Encode(Devices)

	fmt.Println("Add ", device.Type, "with id ", device.Id)
}

func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, device := range Devices {
		if device.Id == params["id"] {
			Devices = append(Devices[:index], Devices[index+1:]...)
			fmt.Println("Remove ", device.Type, "with id ", device.Id)
			break
		}
		json.NewEncoder(w).Encode(Devices)
	}

}
