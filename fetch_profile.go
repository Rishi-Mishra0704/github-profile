package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GitHubProfile struct {
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	Location    string `json:"location"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	CreatedAt   string `json:"created_at"`
}

func (viewer *GitHubProfileViewer) FetchProfile() (*GitHubProfile, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", viewer.Username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching profile: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching profile: %s", resp.Status)
	}

	var profile GitHubProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, fmt.Errorf("error decoding profile: %v", err)
	}

	return &profile, nil
}
