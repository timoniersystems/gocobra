/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// factCmd represents the fact command
var factCmd = &cobra.Command{
	Use:   "fact",
	Short: "Retrieve a random fact from the Internet",
	Long: `Supported fact types are
	       * Facts about cats`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fact called")
	},
}

func init() {
	rootCmd.AddCommand(factCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// factCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// factCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
