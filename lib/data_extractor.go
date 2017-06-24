package lib

/*
  Purpose of this package is to take input given from command line app and
  make an API request to Github.com

  Once that request has been made and JSON returned, make sure it's ready
  for other function consumption
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	_ "github.com/caarlos0/spin"
	"github.com/urfave/cli"
)

type RepoInfo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"html_url"`
	CloneUrl    string `json:"clone_url"`
	Homepage    string `json:"homepage"`
	SshUrl      string `json:"ssh_url"`
	Description string `json:"description"`
	Size        int    `json:"size"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	Issues      int    `json:"open_issues"`
}

func GatherInfo(c *cli.Context) RepoInfo {
	req_url := buildRequest(fixInput(c.Args().Get(0)))
	// possible place to put spinner library
	data := clientRequest(req_url)
	checkData(data)
	return data
}

func fixInput(input string) string {
	if strings.ContainsAny(input, "/") {
		return input
	} else {
		safeInput := url.QueryEscape(input)
		return safeInput + "/" + safeInput
	}
}

func buildRequest(input string) (req *http.Request) {
	url := fmt.Sprintf("https://api.github.com/repos/%s", input)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	return
}

func clientRequest(req *http.Request) RepoInfo {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	var repo RepoInfo
	decodeJson(resp, &repo)
	return repo
}

func decodeJson(resp *http.Response, repo *RepoInfo) {
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&repo); err != nil {
		log.Println(err)
	}
}

func checkData(repo RepoInfo) error {
	if len(repo.Name) == 0 && repo.Id == 0 {
		log.Fatal("Repo ID and Name are empty, Gitinfo can't find this Repo :(")
	}
	return nil
}
