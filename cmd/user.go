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
	"log"

	"github.com/pseohy/bri/conf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long:  `Add or delete users in an encrypted way.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if uDelete {
			/* delete a user from the database */
			h, err := conf.EncryptUser(uInfo[0], uInfo[1])
			if err != nil {
				log.Fatal(err)
			}

			err = conf.UserData.Delete(h)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			/* add a user to the database */
			err = conf.UserData.EncryptAndAdd(uInfo[0], uInfo[1])
			if err != nil {
				log.Fatal(err)
			}

		}
		err = conf.UserData.Dump()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	configCmd.AddCommand(userCmd)

	userCmd.Flags().StringSliceVarP(&uInfo, "info", "i",
		[]string{"", ""}, "User name and phone number")
	userCmd.Flags().BoolVarP(&uDelete, "delete", "d",
		false, "Delete if specified")

	viper.BindPFlag("user", userCmd.Flags().Lookup("user"))
	viper.BindPFlag("delete", userCmd.Flags().Lookup("delete"))
}
