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

var jokeType string

// converted from JSON responses via https://mholt.github.io/json-to-go/
type ChuckNorrisJoke struct {
	Categories []any  `json:"categories"`
	CreatedAt  string `json:"created_at"`
	IconURL    string `json:"icon_url"`
	ID         string `json:"id"`
	UpdatedAt  string `json:"updated_at"`
	URL        string `json:"url"`
	Value      string `json:"value"`
}

type DadJoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// jokeCmd represents the joke command
var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "Retrieve a random joke from the Internet",
	Long: `Supported joke types are:
* Chuck Norris jokes (-t chucknorris)
* Dad jokes (-t dad)`,
	Run: func(cmd *cobra.Command, args []string) {
		supportedJokeTypes := map[string]bool{
			"chucknorris": true,
			"dad": true,
		}
		_, ok := supportedJokeTypes[jokeType]
		if !ok {
			fmt.Printf("%s is not a supported joke type\n", jokeType)
			os.Exit(1)
		}
		fmt.Printf("Here is your %s joke:\n", jokeType)
		//fmt.Printf("%s is a supported joke type\n", jokeType)
		if jokeType == "chucknorris" {
			resBody, err := httpclient.GetHTTPResponseBody("https://api.chucknorris.io/jokes/random")
			if err != nil {
				os.Exit(1)
			}
			//fmt.Printf("response body: %s\n", string(resBody))
			var result ChuckNorrisJoke
			if err := json.Unmarshal(resBody, &result); err != nil {   // Parse []byte to go struct pointer
				fmt.Println("Error unmarshaling JSON")
				os.Exit(1)
			}
			fmt.Println(result.Value)
			os.Exit(0)
		}
		if jokeType == "dad" {
			resBody, err := httpclient.GetHTTPResponseBody("https://icanhazdadjoke.com/")
			if err != nil {
				os.Exit(1)
			}
			//fmt.Printf("response body: %s\n", string(resBody))
			var result DadJoke
			if err := json.Unmarshal(resBody, &result); err != nil {   // Parse []byte to go struct pointer
				fmt.Println("Error unmarshaling JSON")
				os.Exit(1)
			}
			fmt.Println(result.Joke)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(jokeCmd)
	jokeCmd.PersistentFlags().StringVarP(&jokeType, "type", "t", "", "Joke type: chucknorris (default) or dad")
}
