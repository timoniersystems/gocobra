/*
Copyright © 2023 Timonier Systems

*/

package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHTTPResponseBody(url string) (body []byte, err error) {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept": {"application/json"},
	}
	
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}
	return responseBody, err
}

