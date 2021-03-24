/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert new entries",
	Long: `This command inserts new data to the 
	Phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("insert called")
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//
	insertCmd.Flags().StringP("name", "n", "", "Name value")
	insertCmd.Flags().StringP("surname", "s", "", "Surname value")
	insertCmd.Flags().StringP("tel", "t", "", "Telephone value")
}
