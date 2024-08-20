package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Commit struct {
	Message string `json:"commit.message"`
	Author  string `json:"commit.author.name"`
	Date    string `json:"commit.author.date"`
	URL     string `json:"html_url"`
}

func (viewer *GitHubProfileViewer) FetchCommits(repoName string) ([]Commit, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", viewer.Username, repoName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching commits: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching commits: %s", resp.Status)
	}

	var commits []Commit
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		return nil, fmt.Errorf("error decoding commits: %v", err)
	}

	return commits, nil
}
