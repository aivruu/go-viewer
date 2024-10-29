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
	"fmt"
	"viewer/main/release"
)

func main() {
	// Request test for repository's latest release.
	releaseRequest := release.NewReleaseRequest(ForRelease("aivruu", "repo-viewer", "v3.4.7"))
	releaseFuture := releaseRequest.Request()
	releaseModel := releaseFuture.Get()
	if releaseModel == nil {
		fmt.Println("Failed to request the request for this repository.")
		return
	}
	fmt.Printf("%s - %s - %s", releaseModel.Author(), releaseModel.Name(), releaseModel.TagName())

	repositoryModel := NewRepositoryRequest(ForRepository("aivruu", "repo-viewer"))
	repositoryFuture := repositoryModel.RequestAndThen(
		func(Model *GithubRepositoryModel) {
			value := *Model
			fmt.Printf("%s - %s - %s", value.Owner(), value.Name(), *value.License())
		},
	)
	if repositoryFuture.Get() == nil {
		fmt.Println("Failed to request the repository.")
		return
	}
	fmt.Println("Repository request is successful.")
}
