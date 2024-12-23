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
	"strings"
)

const (
	GithubApiUrl        = "https://api.github.com/repos/%s/%s"
	GithubApiReleaseUrl = GithubApiUrl + "/releases/tags/%s"
)

// ForRepository This function formats the GithubApiUrl to include the author and repository specified to create a valid
// url for a request.
func ForRepository(author, repository string) string {
	return fmt.Sprintf(GithubApiUrl, author, repository)
}

// ForRelease This function formats the GithubApiReleaseUrl to include the author, repository and tag specified to create a valid
// url for a request.
func ForRelease(author, repository, tag string) string {
	var version string
	if strings.Index(tag, "latest") < 0 {
		version = tag
	} else {
		version = strings.Replace(GithubApiReleaseUrl, "tags/", "", 1)
	}
	return fmt.Sprintf(GithubApiReleaseUrl, author, repository, version)
}
