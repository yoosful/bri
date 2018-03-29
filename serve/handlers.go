package serve

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pseohy/bri/conf"
)

func Index(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/templates/index.html"))
	tmpl.ExecuteTemplate(w, "index.html", conf.DeviceData.Data)
}

func GetDevices(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(conf.DeviceData.Data)

	fmt.Println("Get info of all devices")

	tmpl := template.Must(template.ParseFiles("static/templates/device.html"))

	tmpl.ExecuteTemplate(w, "device.html", conf.DeviceData.Data)
}

func GetDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, device := range conf.DeviceData.Data {
		if id, _ := strconv.ParseInt(params["did"], 10, 64); device.Did == id {
			json.NewEncoder(w).Encode(device)
			fmt.Println("Get info of ", device.Dtype, "with id ", device.Did)
			return
		}
	}
	json.NewEncoder(w).Encode(&Device{})
}

func CreateDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, device := range conf.DeviceData.Data {
		if id, _ := strconv.ParseInt(params["did"], 10, 64); device.Did == id {
			fmt.Println(device.Dtype, "with id ", device.Did, " already exists")
			return
		}
		json.NewEncoder(w).Encode(conf.DeviceData.Data)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var device conf.Device
	_ = json.NewDecoder(r.Body).Decode(&device)
	id, _ := strconv.ParseInt(params["did"], 10, 64)
	device.Did = id
	device.Dtype = params["dtype"]
	conf.DeviceData.Data = append(conf.DeviceData.Data, device)
	json.NewEncoder(w).Encode(conf.DeviceData.Data)

	fmt.Println("Add ", device.Dtype, "with id ", device.Did)
}

func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, device := range conf.DeviceData.Data {
		if id, _ := strconv.ParseInt(params["did"], 10, 64); device.Did == id {
			conf.DeviceData.Data = append(conf.DeviceData.Data[:index], conf.DeviceData.Data[index+1:]...)
			fmt.Println("Remove ", device.Dtype, "with id ", device.Did)
			break
		}
		json.NewEncoder(w).Encode(conf.DeviceData.Data)
	}
}
