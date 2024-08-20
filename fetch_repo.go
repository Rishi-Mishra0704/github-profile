package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
	CreatedAt   string `json:"created_at"`
}

func (viewer *GitHubProfileViewer) FetchRepositories() ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", viewer.Username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching repositories: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching repositories: %s", resp.Status)
	}

	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, fmt.Errorf("error decoding repositories: %v", err)
	}

	return repos, nil
}
