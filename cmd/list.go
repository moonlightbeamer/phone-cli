/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all entries.",
	Long:  `List the contents of Phone book.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create request
		URL := SERVER + PORT + "/list"

		// Send request to server
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Read data
		responseData, err := ioutil.ReadAll(data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Check HTTP Status Code

		fmt.Print(string(responseData))

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
