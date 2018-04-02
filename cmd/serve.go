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
	"html/template"
	"log"
	"net/http"

	"github.com/pseohy/bri/conf"
	"github.com/pseohy/bri/serve"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Page struct {
	Content template.HTML
}

var (
	t       *template.Template
	content template.HTML

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
		http.Handle("/less/", http.StripPrefix("/less/", http.FileServer(http.Dir("less/"))))

		router := serve.NewRouter()
		http.Handle("/", router)

		log.Fatal(http.ListenAndServe(":4000", nil))
	},
}

func DeviceDetail(d conf.Device) template.HTML {
	if debug {
		return template.HTML(d.Dtype + " " + d.Did)
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
