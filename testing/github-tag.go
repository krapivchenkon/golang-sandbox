package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
	"os"
)

// taken from http://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/

type ReleasesInfo struct {
	Id      uint   `json:"id"`
	TagName string `json:"tag_name"`
}

type ReleaseInfoer interface {
	GetLatestReleaseTag(string) (string, error)
}

type GithubReleaseInfoer struct{}

// Function to actually query the GitHub API for the release information.

func (gh GithubReleaseInfoer) GetLatestReleaseTag(repo string) (string, error) {
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	response, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	releases := []ReleasesInfo{}

	if err := json.Unmarshal(body, &releases); err != nil {
		return "", err
	}

	tag := releases[0].TagName

	return tag, nil
}

// Function to get the message to display to the end user.
func getReleaseTagMessage(ri ReleaseInfoer, repo string) (string, error) {
	tag, err := ri.GetLatestReleaseTag(repo)
	if err != nil {
		return "", fmt.Errorf("Error query GitHub API: %s", err)
	}

	return fmt.Sprintf("The latest release is %s", tag), nil
}

func main() {
	gh := GithubReleaseInfoer{}
	msg, err := getReleaseTagMessage(gh, "docker/machine")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(msg)
}
