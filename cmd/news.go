/*
Copyright © 2023 Timonier Systems

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"math/rand"
	"time"
	"github.com/spf13/cobra"
	httpclient "github.com/timoniersystems/gocobra/internal/httpclient"
)

var newsCategory string

// converted from JSON responses via https://mholt.github.io/json-to-go/
type News struct {
	Category string `json:"category"`
	Data     []struct {
		Author      string `json:"author"`
		Content     string `json:"content"`
		Date        string `json:"date"`
		ID          string `json:"id"`
		ImageURL    string `json:"imageUrl"`
		ReadMoreURL string `json:"readMoreUrl"`
		Time        string `json:"time"`
		Title       string `json:"title"`
		URL         string `json:"url"`
	} `json:"data"`
	Success bool `json:"success"`
}

// newsCmd represents the news command
var newsCmd = &cobra.Command{
	Use:   "news",
	Short: "Retrieve news from the Internet",
	Long: `Supported news categories are:
* business
* world
* politics
* technology`,
	Run: func(cmd *cobra.Command, args []string) {
		supportedNewsCategories := map[string]bool{
			"business": true,
			"world": true,
			"politics": true,
			"technology": true,
		}
		_, ok := supportedNewsCategories[newsCategory]
		if !ok {
			fmt.Printf("%s is not a supported news category.\n", newsCategory)
			newsCategory = "world"
			fmt.Printf("However, here is your %s news:\n", newsCategory)
		} else {
			fmt.Printf("Here is your %s news:\n", newsCategory)
		}
		url := fmt.Sprintf("https://inshorts.deta.dev/news?category=%s", newsCategory)

		resBody, err := httpclient.GetHTTPResponseBody(url)
		if err != nil {
			os.Exit(1)
		}
		//fmt.Printf("response body: %s\n", string(resBody))
		var result News
		if err := json.Unmarshal(resBody, &result); err != nil {   // Parse []byte to go struct pointer
			fmt.Println("Error unmarshaling JSON")
			os.Exit(1)
		}
		// pick a random element from Data array
		rand.Seed(time.Now().Unix())
		length := len(result.Data)
		n := rand.Int() % length
		fmt.Println(result.Data[n].Content)
	},
}

func init() {
	rootCmd.AddCommand(newsCmd)
	newsCmd.PersistentFlags().StringVarP(&newsCategory, "category", "c", "world", "News category")
}
