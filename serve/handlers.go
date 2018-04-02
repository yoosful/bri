package serve

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pseohy/bri/conf"
)

func Index(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := map[string]interface{}{
		"Devices": conf.DeviceData.Data,
		"Users":   conf.UserData.Data,
	}

	tmpl.ExecuteTemplate(w, "index.html", data)

}

func GetDevices(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(conf.DeviceData.Data)

	fmt.Println("Get info of all devices")

	tmpl := template.Must(template.ParseFiles("templates/device.html"))

	tmpl.ExecuteTemplate(w, "device.html", conf.DeviceData.Data)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(conf.DeviceData.Data)

	fmt.Println("Get info of all users")

	tmpl := template.Must(template.ParseFiles("templates/user.html"))

	tmpl.ExecuteTemplate(w, "user.html", conf.UserData.Data)
}

// func GetDevice(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for _, device := range conf.DeviceData.Data {
// 		if id, _ := strconv.ParseInt(params["did"], 10, 64); device.Did == id {
// 			json.NewEncoder(w).Encode(device)
// 			fmt.Println("Get info of ", device.Dtype, "with id ", device.Did)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&conf.DeviceData.Data)
// }

func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, device := range conf.DeviceData.Data {
		if id := params["did"]; device.Did == id {
			conf.DeviceData.Data = append(conf.DeviceData.Data[:index], conf.DeviceData.Data[index+1:]...)
			fmt.Println("Remove ", device.Dtype, "with id ", device.Did)
			break
		}
		json.NewEncoder(w).Encode(conf.DeviceData.Data)
	}
}

func RefreshDevices(w http.ResponseWriter, r *http.Request) {
	conf.DeviceData.Init()
	json.NewEncoder(w).Encode(&conf.DeviceData.Data)
}

func RefreshUsers(w http.ResponseWriter, r *http.Request) {
	conf.UserData.Init()
	json.NewEncoder(w).Encode(&conf.UserData.Data)
}

// NewUser handles NewUser Request from /user/new
func NewUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var umsg conf.UserMsg
	err := decoder.Decode(&umsg)
	if err != nil {
		panic(err)
		return
	}
	defer req.Body.Close()

	err = conf.UserData.EncryptAndAdd(umsg.Name, umsg.Phone)
	if err != nil {
		panic(err)
		return
	}

	err = conf.UserData.Dump()
	if err != nil {
		panic(err)
		return
	}
}
