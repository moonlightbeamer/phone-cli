/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an entry",
	Long: `This commands deletes an existing entry from
	the phone book application given a phone number.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		number, _ := cmd.Flags().GetString("tel")
		if number == "" {
			fmt.Println("Number is empty!")
			return
		}

		// Create request
		URL := "http://" + SERVER + ":" + PORT + "/delete/" + number

		// Send request to server
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Check HTTP Status Code
		if data.StatusCode != http.StatusOK {
			fmt.Println("Status code:", data.StatusCode)
			return
		}

		// Read data
		responseData, err := ioutil.ReadAll(data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("tel", "t", "", "Telephone number to delete")
}
