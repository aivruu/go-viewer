// Copyright 2024 aivruu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
// Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"testing"
	"time"
	"viewer/main/http"
	"viewer/main/repository"
	"viewer/main/repository/operator"
)

func TestReleaseRequest(t *testing.T) {
	// Request test to repository's latest release.
	releaseRequest := repository.NewReleaseRequest(ForRelease("aivruu", "repo-viewer", "v3.4.7"))
	release := http.Request(releaseRequest, 5*time.Second)
	if release == nil {
		t.Error("Failed to request the release for this repository.")
	} else {
		t.Logf("%s - %s - %s", release.Author.Login, release.Name, release.TagName)
		t.Log()
		t.Logf("%t", release.Compare(operator.Less, 134))
		for _, asset := range release.Assets {
			t.Log()
			t.Logf("Asset: %s - %s", asset.Name, asset.Url)
			t.Log()
		}
	}
}

func TestRepositoryRequest(t *testing.T) {
	// Test using "consumer" function to print repository's information if it is available.
	repositoryRequest := repository.NewRepositoryRequest(ForRepository("aivruu", "repo-viewer"))
	model := http.RequestAndThen(repositoryRequest, func(Model *repository.GithubRepositoryModel) {
		t.Logf("%s - %s - %s", Model.Owner, Model.Name, Model.LicenseType.Name)
		t.Log()
		t.Log(Model.Archived)
		t.Log(Model.Forks)
		t.Log(Model.CanFork)
		t.Log(Model.Stars)
		t.Log(Model.Language)
	}, 5*time.Second)
	if model == nil {
		t.Log("Failed to request the repository.")
	}
}
