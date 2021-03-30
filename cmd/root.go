/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

var SERVER = "http://localhost"
var PORT = ":1234"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "phone-cli",
	Short: "Client for Phonebook application",
	Long:  `A command line utility of the Phone Book application.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

}
