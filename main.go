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
	"os"
	"strings"
	"viewer/main/http"
	"viewer/main/repository"
)

func main() {
	argsAmount := len(os.Args)
	if argsAmount < 2 {
		fmt.Println("Missing arguments.")
		os.Exit(1)
		return
	}
	if argsAmount < 3 || argsAmount > 4 {
		fmt.Println("Invalid arguments amount, four for release-info, five for repository.")
		os.Exit(2)
		return
	}
	if argsAmount == 4 {
		releaseRequest := repository.NewReleaseRequest(ForRelease(os.Args[1], os.Args[2], os.Args[3]))
		model := http.Request(releaseRequest)
		if model == nil {
			fmt.Println("Failed to request the release for this repository.")
			os.Exit(3)
			return
		}
		printReleaseInformation(model)
		return
	}
	repositoryRequest := repository.NewRepositoryRequest(ForRepository(os.Args[1], os.Args[2]))
	model := http.Request(repositoryRequest)
	if model == nil {
		fmt.Println("Failed to request the repository.")
		os.Exit(3)
		return
	}
	printRepositoryInformation(model)
}

func printReleaseInformation(model *repository.GithubReleaseModel) {
	fmt.Println("Showing information for repository's release: ", model.TagName)
	fmt.Println("Title -> ", model.Name)
	fmt.Println("Tag -> ", model.TagName)
	fmt.Println("Author -> ", model.Author)
	fmt.Println("Identifier -> ", model.UniqueId)
	fmt.Println("Assets:")
	for _, asset := range model.Assets {
		fmt.Println("  Name -> ", asset.Name)
		fmt.Println("  URL -> ", asset.Url)
	}
}

func printRepositoryInformation(model *repository.GithubRepositoryModel) {
	fmt.Println("Showing information for repository: ", model.Name)
	fmt.Println()
	fmt.Println("Owner -> ", model.Owner.Login)
	fmt.Println("Description -> ", model.Description)
	fmt.Println("Topics -> ", strings.Join(model.Topics, ", "))
	fmt.Println("Stars -> ", model.Stars)
	fmt.Println("Forks -> ", model.Forks)
	fmt.Println("Forks Allowed -> ", repository.FormatBooleanValue(model.CanFork))
	fmt.Println("Language -> ", model.Language)
	fmt.Println("License -> ", model.License.Name)
	fmt.Println("Public -> ", repository.FormatBooleanValue(!model.Private))
	fmt.Println("Archived -> ", repository.FormatBooleanValue(model.Archived))
	fmt.Println("Disabled -> ", repository.FormatBooleanValue(model.Disabled))
}
