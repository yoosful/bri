// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/pseohy/bri/conf"
	"github.com/pseohy/bri/serve"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Page struct {
	Content template.HTML
}

var (
	t          *template.Template
	content    template.HTML
	deviceChan = make(chan int, 1)

	debug bool
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run server that collects data from IOT devices",
	Long: `Run server that collects data from authenticated IOT devices.
Store usage data with encryption.`,
	Run: func(cmd *cobra.Command, args []string) {

		http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
		http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
		http.Handle("/plugins/", http.StripPrefix("/plugins/", http.FileServer(http.Dir("plugins/"))))
		http.Handle("/bootstrap/", http.StripPrefix("/bootstrap/", http.FileServer(http.Dir("bootstrap/"))))

		router := serve.NewRouter()
		router.HandleFunc("/device", CreateDevice).Methods("POST")
		http.Handle("/", router)

		log.Fatal(http.ListenAndServe(":4000", nil))
	},
}

func CreateDevice(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	log.Println("A new message arrived")

	var dmsg DeviceMsg
	err := decoder.Decode(&dmsg)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	h, err := conf.Checksum(dmsg.Dtype, dmsg.Did)
	if err != nil {
		panic(err)
	}

	didInt, err := strconv.ParseInt(dmsg.Did, 10, 64)
	if err != nil {
		panic(err)
	}

	var status bool
	if dmsg.Msg == "on" {
		status = true
	} else if dmsg.Msg == "off" {
		status = false
	} else {
		panic(errors.New("Unexpected message"))
	}

	d := conf.Device{
		Address: h,
		Dtype:   dmsg.Dtype,
		Did:     didInt,
		Status:  status,
		Usage: map[string]int{
			dmsg.Uid: 1000,
		},
	}

	conf.DeviceData.Update(h, d)
	conf.DeviceData.Dump()

	/* send a signal if new device message arrives. */
	deviceChan <- 1

	// params := mux.Vars(r)
	// for _, device := range conf.DeviceData.Data {
	// 	if id, _ := strconv.ParseInt(params["did"], 10, 64); device.Did == id {
	// 		fmt.Println(device.Dtype, "with id ", device.Did, " already exists")
	// 		return
	// 	}
	// 	json.NewEncoder(w).Encode(conf.DeviceData.Data)
	// }
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// var device conf.Device
	// _ = json.NewDecoder(r.Body).Decode(&device)
	// id, _ := strconv.ParseInt(params["did"], 10, 64)
	// device.Did = id
	// device.Dtype = params["dtype"]
	// conf.DeviceData.Data = append(conf.DeviceData.Data, device)
	// json.NewEncoder(w).Encode(conf.DeviceData.Data)
	//
	// fmt.Println("Add ", device.Dtype, "with id ", device.Did)
}

func DeviceDetail(d conf.Device) template.HTML {
	if debug {
		return template.HTML(d.Dtype + " " + strconv.FormatInt(d.Did, 10))
	} else {
		return template.HTML("")
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")

	viper.BindPFlag("debug", serveCmd.Flags().Lookup("debug"))
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
