package lib

// Purpose of this package is to take input given from command line app and
// make an API request to Github.com
//
// Once that request has been made and JSON returned, make sure it's ready
// for other function consumption

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	_ "reflect"
	"strings"

	_ "github.com/caarlos0/spin"
	"github.com/urfave/cli"
)

// Repo struct to hold the data I want to collect from the API
type RepoInfo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Size        int    `json:"size"`
	Stargazers  int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
}

func GatherInfo(c *cli.Context) error {
	req_url := buildRequest(fixInput(c.Args().Get(0)))
	// possible place to put spinner library
	clientRequest(req_url)

	return nil
}

// used to correctly format user input before API request
func fixInput(input string) string {
	if strings.ContainsAny(input, "/") {
		return input
	} else {
		safeInput := url.QueryEscape(input)
		return safeInput + "/" + safeInput
	}
}

// method used to build the API request to Github
func buildRequest(input string) (req *http.Request) {
	url := fmt.Sprintf("https://api.github.com/repos/%s", input)
	fmt.Printf("The url is: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	return
}

// used to build the client and make request
func clientRequest(req *http.Request) {

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// return resp use reflection package to see type
	defer resp.Body.Close()

	var repo RepoInfo

	if err := json.NewDecoder(resp.Body).Decode(&repo); err != nil {
		log.Println(err)
	}

	// need to add test to make sure name, url and id are returned otherwise error
	fmt.Println("ID = ", repo.Id)
	fmt.Println("Name   = ", repo.Name)
	fmt.Println("URL  = ", repo.Url)
	fmt.Println("Description   = ", repo.Description)
	fmt.Println("Size   = ", repo.Size)
	fmt.Println("Stargazers   = ", repo.Stargazers)
	fmt.Println("Forks   = ", repo.Forks)
}

func decodeJson() {

}

func checkInfo() {

}
