package main

import (
	"fmt"
)

type GitHubProfileViewer struct {
	Username string
}

func NewGitHubProfileViewer(username string) *GitHubProfileViewer {
	return &GitHubProfileViewer{Username: username}
}

func main() {
	fmt.Println("Please Enter Your GitHub Username:")
	var username string
	fmt.Scanln(&username)

	if username == "" {
		fmt.Println("Username cannot be empty. Exiting.")
		return
	}

	viewer := NewGitHubProfileViewer(username)

	profile, err := viewer.FetchProfile()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("GitHub Profile for %s\n", username)
	fmt.Printf("Name: %s\n", profile.Name)
	fmt.Printf("Bio: %s\n", profile.Bio)
	fmt.Printf("Location: %s\n", profile.Location)
	fmt.Printf("Public Repositories: %d\n", profile.PublicRepos)
	fmt.Printf("Followers: %d\n", profile.Followers)
	fmt.Printf("Following: %d\n", profile.Following)
	fmt.Printf("Created at: %s\n", formatDate(profile.CreatedAt))

	repos, err := viewer.FetchRepositories()
	if err != nil {
		fmt.Printf("Error fetching repositories: %v\n", err)
		return
	}

	for _, repo := range repos {
		fmt.Printf("\nRepository: %s\n", repo.Name)
		fmt.Printf("Description: %s\n", repo.Description)
		fmt.Printf("Language: %s\n", repo.Language)
		fmt.Printf("Created at: %s\n", formatDate(repo.CreatedAt))

		prs, err := viewer.FetchPullRequests(repo.Name)
		if err != nil {
			fmt.Printf("Error fetching pull requests: %v\n", err)
			continue
		}

		if len(prs) > 0 {
			fmt.Printf("\nPull Requests for %s:\n", repo.Name)
			for _, pr := range prs {
				fmt.Printf("Title: %s\n", pr.Title)
				fmt.Printf("State: %s\n", pr.State)
				fmt.Printf("Created at: %s\n", formatDate(pr.CreatedAt))
				fmt.Printf("URL: %s\n\n", pr.URL)
			}
		}

		commits, err := viewer.FetchCommits(repo.Name)
		if err != nil {
			fmt.Printf("Error fetching commits: %v\n", err)
			continue
		}

		if len(commits) > 0 {
			fmt.Printf("Commits for %s:\n", repo.Name)
			for _, commit := range commits {
				fmt.Printf("Message: %s\n", commit.Message)
				fmt.Printf("Author: %s\n", commit.Author)
				fmt.Printf("Date: %s\n", formatDate(commit.Date))
				fmt.Printf("URL: %s\n\n", commit.URL)
			}
		}
	}
}
