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
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/pseohy/bri/conf"
	"github.com/spf13/cobra"
)

type Page struct {
	Content template.HTML
}

var (
	t          *template.Template
	content    template.HTML
	deviceChan = make(chan int, 1)
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run server that collects data from IOT devices",
	Long: `Run server that collects data from authenticated IOT devices.
Store usage data with encryption.`,
	Run: func(cmd *cobra.Command, args []string) {
		http.Handle("/css/", http.StripPrefix("/css/",
			http.FileServer(http.Dir("static"))))
		var b bytes.Buffer
		t.ExecuteTemplate(&b, "device.html", &conf.DeviceData)
		content = template.HTML(b.String())

		http.HandleFunc("/", displayDevices)
		http.HandleFunc("/device", deviceHandler)
		http.ListenAndServe(":8080", nil)
	},
}

func init() {
	t = template.Must(template.ParseFiles("templates/index.html",
		"templates/device.html"))

	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func displayDevices(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Content: content,
	}
	t.ExecuteTemplate(w, "index.html", p)

	go func() {
		for {
			select {
			case <-deviceChan:
				var b bytes.Buffer
				t.ExecuteTemplate(&b, "device.html", &conf.DeviceData)
				content = template.HTML(b.String())
			default:
				continue
			}
		}
	}()
}

func deviceHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	log.Println("Received usage message")

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

	d := conf.Device{
		Address: h,
		Dtype:   dmsg.Dtype,
		Did:     didInt,
		Status:  true,
		Usage: map[string]int{
			dmsg.Uid: 1000,
		},
	}

	conf.DeviceData.Update(h, d)
	conf.DeviceData.Dump()

	/* send a signal if new device message arrives. */
	deviceChan <- 1
}
