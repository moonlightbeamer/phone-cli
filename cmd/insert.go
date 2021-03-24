/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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

	insertCmd.Flags().StringP("name", "n", "", "Name value")
	insertCmd.Flags().StringP("surname", "s", "", "Surname value")
	insertCmd.Flags().StringP("tel", "t", "", "Telephone value")
}
