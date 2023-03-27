package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTextFromURL(url string) (resBody []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)

	resultBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}
	return resultBody, err
}

