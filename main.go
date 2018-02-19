// Fun with golang. Courtesy rain forest !!
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	response, err := http.Get("http://letsrevolutionizetesting.com/challenge.json")
	if err != nil {
		fmt.Println("Error while get request")
		return
	}

	var f interface{}

	for true {
		json_resp, err := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(json_resp, &f)
		if err != nil {
			fmt.Println("Error while json parsing breaking!!")
			break
		}

		f2, ok := f.(map[string]interface{})
		next_url, ok := f2["follow"].(string)
		if !ok {
			fmt.Printf("Done with cycle....")
			fmt.Println(f2["message"])
			break
		}
		next_url = strings.Replace(next_url, "challenge", "challenge.json", 1)
		fmt.Printf("%v\n", next_url)
		response, err = http.Get(next_url)
		if err != nil {
			fmt.Println("Error while get request")
			break
		}
	}

	fmt.Printf("Final %v\n", f)
}
