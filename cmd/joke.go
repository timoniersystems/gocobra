/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// jokeCmd represents the joke command
var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "Retrieve a random joke from the Internet",
	Long: `Supported joke types are
	       * Chuck Norris jokes
		   * Dad jokes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("joke called")
	},
}

func init() {
	rootCmd.AddCommand(jokeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jokeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jokeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
