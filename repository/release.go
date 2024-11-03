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

import (
	"strconv"
	"strings"
	"viewer/main/common"
	"viewer/main/download"
	"viewer/main/repository/operator"
)

// GithubReleaseModel This struct stores all necessary information for the repository's requested release.
type (
	GithubReleaseModel struct {
		Author   Author  `json:"author"`
		TagName  string  `json:"tag_name"`
		Name     string  `json:"name"`
		UniqueId int     `json:"id"`
		Assets   []Asset `json:"assets"`
		common.RequestableModel
	}

	// Author This struct only stores the release's author's github-username from the request.
	Author struct {
		Login string `json:"login"`
	}

	// Asset This struct stores a release's asset's name and url to be used for downloading later.
	Asset struct {
		Name string `json:"name"`
		Url  string `json:"browser_download_url"`
	}
)

// Download This method tries to download the asset-specified for this release from the array of assets into specified directory,
// and will return a boolean value whether the asset-number is valid, and asset was downloaded correctly.
func (r *GithubReleaseModel) Download(directory string, assetNum int) int64 {
	if assetNum < 0 {
		return download.UnknownAssetDefaultSize
	}
	assetsAmount := len(r.Assets)
	if assetsAmount == 0 || assetNum >= assetsAmount {
		return download.InvalidAssetDefaultSize
	}
	asset := r.Assets[assetNum]
	downloadStatus := download.From(directory, asset.Name, asset.Url)
	return downloadStatus.Result()
}

// Compare This method compares the given version-number with this release's tag-name (as int) using the specified operator-type
// for the comparison, and return a bool as operation's result.
func (r *GithubReleaseModel) Compare(operatorType operator.Operator, targetVersion int) bool {
	var versionFormatted string
	if r.TagName[0] == 'v' {
		versionFormatted = r.TagName[1:]
	} else {
		versionFormatted = r.TagName
	}
	versionWithoutDots := strings.Replace(versionFormatted, ".", "", 3)
	version, _ := strconv.Atoi(versionWithoutDots)
	switch operatorType {
	case operator.Equal:
		return targetVersion == version
	case operator.Less:
		return targetVersion < version
	case operator.LessOrEqual:
		return targetVersion <= version
	case operator.Greater:
		return targetVersion > version
	case operator.GreaterOrEqual:
		return targetVersion >= version
	default:
		return false
	}
}
