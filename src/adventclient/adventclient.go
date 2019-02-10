package adventclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var baseURL = "https://adventofcode.com/"

func GetString(uri string) string {
	client := &http.Client{}

	url := fmt.Sprintf("%s%s", baseURL, uri)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cookie", "_ga=GA1.2.2010882644.1544457104; _gid=GA1.2.897020045.1549735477; session=53616c7465645f5fe479758764365fa3ca3ef3a270861ea95a77ef7096086abac5d6d0d253729972966e39d219f93f50; _gat=1")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	var bodyString string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString = string(bodyBytes)
	}

	return bodyString
}

func GetStringList(uri string) []string {
	rawString := GetString(uri)
	splitInput := strings.Split(rawString, "\n")
	splitInput = trimLastElement(splitInput)

	return splitInput
}

func trimLastElement(stringArray []string) []string {
	if len(stringArray) > 0 {
		stringArray = stringArray[:len(stringArray)-1]
	}
	return stringArray
}
