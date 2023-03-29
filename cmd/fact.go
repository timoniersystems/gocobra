/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/spf13/cobra"
	httpclient "github.com/timoniersystems/gocobra/internal/httpclient"
)

var factType string
var numberValue int

// converted from JSON responses via https://mholt.github.io/json-to-go/
type CatsFact struct {
	Data []string `json:"data"`
}

type NumbersFact struct {
	Text   string `json:"text"`
	Number int    `json:"number"`
	Found  bool   `json:"found"`
	Type   string `json:"type"`
}

// factCmd represents the fact command
var factCmd = &cobra.Command{
	Use:   "fact",
	Short: "Retrieve a random fact from the Internet",
	Long: `Supported fact types are:
* Cat facts (-t cats)
* Number facts (-t numbers)`,
	Run: func(cmd *cobra.Command, args []string) {
		supportedFactTypes := map[string]bool{
			"cats": true,
			"numbers": true,
		}
		_, ok := supportedFactTypes[factType]
		if !ok {
			fmt.Printf("%s is not a supported fact type\n", factType)
			os.Exit(1)
		}
		
		if factType == "cats" {
			fmt.Printf("Here is your %s fact:\n", factType)
			resBody, err := httpclient.GetHTTPResponseBody("https://meowfacts.herokuapp.com/")
			if err != nil {
				os.Exit(1)
			}
			//fmt.Printf("response body: %s\n", string(resBody))
			var result CatsFact
			if err := json.Unmarshal(resBody, &result); err != nil {   // Parse []byte to go struct pointer
				fmt.Println("Error unmarshaling JSON")
				os.Exit(1)
			}
			fmt.Println(result.Data[0])
			os.Exit(0)
		}
		if factType == "numbers" {
			fmt.Printf("Here is your %s fact for n=%d:\n", factType, numberValue)
			url := fmt.Sprintf("http://numbersapi.com/%d/trivia?json", numberValue)
			resBody, err := httpclient.GetHTTPResponseBody(url)
			if err != nil {
				os.Exit(1)
			}
			//fmt.Printf("response body: %s\n", string(resBody))
			var result NumbersFact
			if err := json.Unmarshal(resBody, &result); err != nil {   // Parse []byte to go struct pointer
				fmt.Println("Error unmarshaling JSON")
				os.Exit(1)
			}
			fmt.Println(result.Text)
			os.Exit(0)
		}
		
	},
}

func init() {
	rootCmd.AddCommand(factCmd)
	factCmd.PersistentFlags().StringVarP(&factType, "type", "t", "cats", "Fact type: cats (default) or numbers")
	factCmd.PersistentFlags().IntVarP(&numberValue, "number", "n", 42, "Number to get a fact for")
}
