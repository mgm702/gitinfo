package lib

import (
	_ "encoding/json"
	"io/ioutil"
	_ "log"
	_ "net/http"
	_ "net/url"
	"os"
	"os/exec"
	"strings"
	"testing"

	_ "github.com/urfave/cli"
)

func TestGatherInfo(t *testing.T) {
	// Feed cli app with all params into this function
	// =========
	// test to make sure the RepoInfo struct that's returned is equal
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

func TestclientRequest(t *testing.T) {
	// takes a *http.Request client and call this function
	// =========
	// test that it returns a decoded RepoInfo struct filled with info
}

func TestdecodeJson(t *testing.T) {
	// Feed request with certain body
	// =========
	// test to make sure response has stuff placed where it needs to be
}

func TestCheckData(t *testing.T) {
	var dataTest RepoInfo
	if os.Getenv("CH_DATA") == "1" {
		checkData(dataTest)
		return
	}

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
