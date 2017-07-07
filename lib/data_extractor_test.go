package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestGatherInfo(t *testing.T) {
	var repo RepoInfo
	json_repo := buildJsonRepo(&repo)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	data := registerMockResponder(json_repo)
	checkData(data)

	if data.Id != repo.Id {
		t.Error("Id of data returned was suppose to be %d, instead got %d", repo.Id, data.Id)
	}

	if data.Name != repo.Name {
		t.Error("Name of data returned was suppose to be %d, instead got %d", repo.Name, data.Name)
	}

	if data.Url != repo.Url {
		t.Error("Url of data returned was suppose to be %d, instead got %d", repo.Url, data.Url)
	}
}

func TestFixInput(t *testing.T) {
	var stringTest = []struct {
		in       string
		expected string
	}{
		{"hey", "hey/hey"},
		{"slash/hey", "slash/hey"},
	}

	for _, tt := range stringTest {
		actual := fixInput(tt.in)
		if actual != tt.expected {
			t.Errorf("fixInput(%d): expected %d, actual %d", tt.in, tt.expected, actual)

		}
	}
}

func TestBuildRequest(t *testing.T) {
	var urlTest = []struct {
		in         string
		method     string
		get_params string
	}{
		{"ruby/ruby", "GET", "/repos/ruby/ruby"},
		{"slash/hey", "GET", "/repos/slash/hey"},
		{"golang/go", "GET", "/repos/golang/go"},
	}

	for _, tt := range urlTest {
		actual := buildRequest(tt.in)

		if actual.Method != tt.method {
			t.Errorf("fixInput(%d): expected %d, actual %d", tt.in, tt.method, actual.Method)
		}
		if actual.URL.Path != tt.get_params {
			t.Errorf("buildRequest(%d): expected %d, actual %d", tt.in, tt.get_params, actual.URL.Path)
		}
	}
}

func TestClientRequest(t *testing.T) {
	var repo RepoInfo
	json_repo := buildJsonRepo(&repo)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	data := registerMockResponder(json_repo)

	if data.Forks != repo.Forks {
		t.Error("Forks of data returned was suppose to be %d, instead got %d", repo.Forks, data.Forks)
	}

	if data.Stars != repo.Stars {
		t.Error("Stars of data returned was suppose to be %d, instead got %d", repo.Stars, data.Stars)
	}
}

func TestCheckData(t *testing.T) {
	var dataTest RepoInfo
	if os.Getenv("CH_DATA") == "1" {
		checkData(dataTest)
		return
	}

	// Refactor into it's own helper function
	cmd := exec.Command(os.Args[0], "-test.run=TestCheckData")
	cmd.Env = append(os.Environ(), "CH_DATA=1")
	stdout, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	gotBytes, _ := ioutil.ReadAll(stdout)
	got := string(gotBytes)
	expected := "Repo ID and Name are empty, Gitinfo can't find this Repo :("
	if !strings.HasSuffix(got[:len(got)-1], expected) {
		t.Fatalf("Unexpected log message. Got %s but should be %s", got[:len(got)-1], expected)
	}
}

func TestDecodeJson(t *testing.T) {
	// REFACTOR: make this function smaller
	httpmock.Activate()
	var test_repo RepoInfo
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.github.com/repos/ruby/ruby",
		httpmock.NewStringResponder(200, string(buildJsonRepo(&test_repo))))

	var decode_repo RepoInfo

	url := fmt.Sprintf("https://api.github.com/repos/ruby/ruby")
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	decodeJson(resp, &decode_repo)

	if decode_repo.Id != test_repo.Id {
		t.Error("Id of data returned was suppose to be %d, instead got %d", test_repo.Id, decode_repo.Id)
	}
}

// Helper functionality for long test functions
func registerMockResponder(json_repo []uint8) RepoInfo {
	httpmock.RegisterResponder("GET", "https://api.github.com/repos/ruby/ruby",
		httpmock.NewStringResponder(200, string(json_repo)))

	req_url := buildRequest("ruby/ruby")
	data := clientRequest(req_url)
	return data
}

func buildJsonRepo(repo *RepoInfo) (json_repo []uint8) {
	repo.Id = 1
	repo.Name = "Ruby"
	repo.Url = "https://github.com/ruby/ruby"
	repo.Description = "the ruby programming language"
	repo.Stars = 1200
	repo.Forks = 450

	json_repo, err := json.Marshal(repo)
	if err != nil {
		fmt.Println("error:", err)
	}
	return
}
