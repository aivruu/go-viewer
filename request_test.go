package main

import (
	"testing"
	"viewer/main/http"
	"viewer/main/repository"
)

func TestReleaseRequest(t *testing.T) {
	// Request test to repository's latest release.
	releaseRequest := repository.NewReleaseRequest(ForRelease("aivruu", "repo-viewer", "v3.4.7"))
	release := http.Request(releaseRequest)
	if release == nil {
		t.Error("Failed to request the release for this repository.")
	} else {
		t.Logf("%s - %s - %s", release.Author.Login, release.Name, release.TagName)
		t.Log()
		for _, asset := range release.Assets {
			t.Logf("Asset: %s - %s", asset.Name, asset.Url)
			t.Log()
		}
	}
}

func TestRepositoryRequest(t *testing.T) {
	// Test using "consumer" function to print repository's information if it is available.
	repositoryRequest := repository.NewRepositoryRequest(ForRepository("aivruu", "repo-viewer"))
	model := http.RequestAndThen(repositoryRequest, func(Model *repository.GithubRepositoryModel) {
		t.Logf("%s - %s - %s", Model.Owner, Model.Name, Model.License.Name)
		t.Log()
		t.Log(Model.Archived)
		t.Log(Model.Forks)
		t.Log(Model.CanFork)
		t.Log(Model.Stars)
		t.Log(Model.Language)
	})
	if model == nil {
		t.Log("Failed to request the repository.")
	}
}
