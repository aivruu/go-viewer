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

package repository

import "viewer/main/common"

// GithubRepositoryModel This struct represents a requested repository with all its information.
type (
	GithubRepositoryModel struct {
		Owner       Owner    `json:"owner"`
		LicenseType License  `json:"license"`
		Parent      Parent   `json:"parent"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Forked      bool     `json:"fork"`
		CanFork     bool     `json:"allow_forking"`
		Stars       int      `json:"stargazers_count"`
		Forks       int      `json:"forks_count"`
		Private     bool     `json:"private"`
		Archived    bool     `json:"archived"`
		Disabled    bool     `json:"disabled"`
		Language    string   `json:"language"`
		Topics      []string `json:"topics"`
		common.RequestableModel
	}

	Owner struct {
		Login string `json:"login"`
	}

	License struct {
		Name string `json:"name"`
	}

	Parent struct {
		Owner string `json:"owner"`
	}
)

// FormatBooleanValue This function returns a readable string that correspond to the value for the given boolean.
func FormatBooleanValue(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}
