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
	"strconv"
	"strings"
	"viewer/main/download"
	"viewer/main/http"
	"viewer/main/repository"
)

func main() {
	argsAmount := len(os.Args)
	if argsAmount < 2 || argsAmount > 6 {
		fmt.Println("Invalid arguments amount, four for release-info | six for assets download, and two for repository.")
		return
	}
	if (argsAmount == 6) && (strings.Contains(os.Args[3], ".") || strings.Contains(os.Args[3], "latest")) {
		releaseRequest := repository.NewReleaseRequest(ForRelease(os.Args[1], os.Args[2], os.Args[3]))
		model := http.Request(releaseRequest, 5)
		if model == nil {
			fmt.Println("Failed to request the release for asset download.")
			return
		}
		if os.Args[4] == "*" {
			fmt.Println("Downloading all release's assets...")
			downloadAllAssets(os.Args[5], model)
		} else {
			index, err := strconv.Atoi(os.Args[4])
			if err != nil {
				fmt.Println("Not valid index-value for asset download.")
				return
			}
			downloadAsset(os.Args[5], model, index-1)
		}
		return
	}
	if argsAmount == 4 {
		var releaseRequest *repository.RequestReleaseModelImpl
		if os.Args[3] == "latest" {
			releaseRequest = repository.NewReleaseRequest(ForRelease(os.Args[1], os.Args[2], "latest"))
		} else {
			releaseRequest = repository.NewReleaseRequest(ForRelease(os.Args[1], os.Args[2], os.Args[3]))
		}
		if releaseRequest == nil {
			fmt.Println("Failed to create the request.")
			return
		}
		model := http.Request(releaseRequest, 5)
		if model == nil {
			fmt.Println("Failed to request the release for this repository.")
			return
		}
		printReleaseInformation(model)
		return
	}
	repositoryRequest := repository.NewRepositoryRequest(ForRepository(os.Args[1], os.Args[2]))
	model := http.Request(repositoryRequest, 5)
	if model == nil {
		fmt.Println("Failed to request the repository.")
		return
	}
	printRepositoryInformation(model)
}

func downloadAsset(directory string, model *repository.GithubReleaseModel, index int) {
	read := model.Download(directory, index)
	fmt.Println("Downloading asset...")
	if read == download.InvalidAssetDefaultSize || read == download.UnknownAssetDefaultSize {
		fmt.Printf("This asset couldn't be downloaded, may be due to an out of range value, index '%d' assets-amount '%d'", index, len(model.Assets))
		return
	}
	fmt.Printf("Downloaded asset with name '%s' and '%d' read bytes. ", model.Assets[index].Name, read)
}

func downloadAllAssets(directory string, model *repository.GithubReleaseModel) {
	for index := range model.Assets {
		downloadAsset(directory, model, index)
	}
}

func printReleaseInformation(model *repository.GithubReleaseModel) {
	fmt.Println("Showing information for repository's release: ", model.TagName)
	fmt.Println("Title ->", model.Name)
	fmt.Println("Tag ->", model.TagName)
	fmt.Println("Author ->", model.Author)
	fmt.Println("Identifier ->", model.UniqueId)
	fmt.Println("Assets:")
	for index, asset := range model.Assets {
		fmt.Println("  Index ->", index+1)
		fmt.Println("  Name ->", asset.Name)
		fmt.Println("  URL ->", asset.Url)
	}
}

func printRepositoryInformation(model *repository.GithubRepositoryModel) {
	fmt.Println("Showing information for repository:", model.Name)
	fmt.Println()
	fmt.Println("Owner ->", model.Owner.Login)
	fmt.Println("Description ->", model.Description)
	fmt.Println("Topics ->", strings.Join(model.Topics, ", "))
	fmt.Println("Stars ->", model.Stars)
	fmt.Println("Forks ->", model.Forks)
	fmt.Println("Forks Allowed ->", repository.FormatBooleanValue(model.CanFork))
	fmt.Println("Language ->", model.Language)
	fmt.Println("License ->", model.LicenseType.Name)
	fmt.Println("Public ->", repository.FormatBooleanValue(!model.Private))
	fmt.Println("Archived ->", repository.FormatBooleanValue(model.Archived))
	fmt.Println("Disabled ->", repository.FormatBooleanValue(model.Disabled))
}
