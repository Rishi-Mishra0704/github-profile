package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PullRequest struct {
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	State     string `json:"state"`
	URL       string `json:"html_url"`
}

func (viewer *GitHubProfileViewer) FetchPullRequests(repoName string) ([]PullRequest, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls", viewer.Username, repoName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching pull requests: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching pull requests: %s", resp.Status)
	}

	var prs []PullRequest
	if err := json.NewDecoder(resp.Body).Decode(&prs); err != nil {
		return nil, fmt.Errorf("error decoding pull requests: %v", err)
	}

	return prs, nil
}
