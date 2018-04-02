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
	"log"
	"net/http"

	"github.com/pseohy/bri/conf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add new user",
	Long:  `Add new user and grant default privilege`,

	// Run new will send HTTP request to the server to add a user
	// to database
	Run: func(cmd *cobra.Command, args []string) {
		msg := conf.UserMsg{
			Name:  uname,
			Phone: uphone,
		}

		jsonUMsg, err := json.Marshal(&msg)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("POST", userURL, bytes.NewBuffer(jsonUMsg))
		req.Header.Set("Content-type", "application/json")

		client := http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	},
}

func init() {
	userCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&uname, "name", "n", "", "User name to add")
	newCmd.Flags().StringVarP(&uphone, "phone", "p", "", "User phone number to add")
	newCmd.Flags().StringVarP(&userURL, "url", "u", "http://localhost:4000/user/new", "server URL")

	viper.BindPFlag("name", newCmd.Flags().Lookup("name"))
	viper.BindPFlag("phone", newCmd.Flags().Lookup("phone"))
	viper.BindPFlag("url", newCmd.Flags().Lookup("url"))
}
