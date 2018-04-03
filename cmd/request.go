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

var userRequestURL string

// requestCmd represents the request command
var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "Request permission to use a device",
	Long:  `Request permission to use a privileged device.`,

	Run: func(cmd *cobra.Command, args []string) {
		// We are assuming that the request contains of device id, to
		// simplify testing condition.

		// We will assume here the user knows the address of user after
		// he made a transaction...
		msg := conf.UserMsg{
			Name:  uname,
			Phone: uphone,
			Type:  requestType,
			Id:    requestId,
		}

		jsonUMsg, err := json.Marshal(&msg)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("POST", userRequestURL, bytes.NewBuffer(jsonUMsg))
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
	userCmd.AddCommand(requestCmd)

	requestCmd.Flags().StringVarP(&uname, "name", "n", "", "Name of requesting user")
	requestCmd.Flags().StringVarP(&uphone, "phone", "p", "", "Phone number of requeting user")
	requestCmd.Flags().StringVarP(&userRequestURL, "url", "u",
		"http://localhost:4000/user/request", "Requested device id")
	requestCmd.Flags().StringVarP(&requestType, "type", "t", "", "Type of requested device")
	requestCmd.Flags().StringVarP(&requestId, "id", "i", "", "Id of requested device")

	viper.BindPFlag("name", requestCmd.Flags().Lookup("name"))
	viper.BindPFlag("phone", requestCmd.Flags().Lookup("phone"))
	viper.BindPFlag("url", requestCmd.Flags().Lookup("url"))
	viper.BindPFlag("request", requestCmd.Flags().Lookup("request"))
}
