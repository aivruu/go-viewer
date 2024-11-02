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
	json2 "encoding/json"
	"fmt"
	http2 "net/http"
	"strconv"
	"strings"
	"time"
	"viewer/main/common"
	"viewer/main/download"
	"viewer/main/http"
	"viewer/main/repository/codec"
	"viewer/main/repository/operator"
	"viewer/main/utils"
)

// GithubReleaseModel This struct stores all necessary information for the repository's requested release.
type GithubReleaseModel struct {
	Author struct {
		Login string `json:"login"`
	} `json:"author"`
	TagName  string  `json:"tag_name"`
	Name     string  `json:"name"`
	UniqueId int     `json:"id"`
	Assets   []Asset `json:"assets"`
	common.RequestableModel
}

// Asset This struct stores a release's asset's name and url to be used for downloading later.
type Asset struct {
	Name string `json:"name"`
	Url  string `json:"browser_download_url"`
}

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

// codec.Provider's implementation necessary for this type.
var releaseCodec = CodecRelease{}

// RequestReleaseModelImpl This codec.Provider implementation is used to handle requests for repositories' releases.
type RequestReleaseModelImpl struct {
	http.RequestModel[GithubReleaseModel]
	url string
}

// NewReleaseRequest This function creates a new RequestReleaseModelImpl with the given url.
func NewReleaseRequest(url string) *RequestReleaseModelImpl {
	return &RequestReleaseModelImpl{url: url}
}

func (r *RequestReleaseModelImpl) RequestWith(client *http2.Client, timeout time.Duration) *GithubReleaseModel {
	resp := utils.Response(utils.ValidateAndModifyTimeout(client, timeout), r.url)
	if resp == nil || resp.StatusCode() != http.ResponseOkStatus {
		return nil
	}
	model, err := releaseCodec.From(resp.JSON())
	if err != nil {
		fmt.Println("Error during release-model deserialization: ", err)
	}
	return model
}

func (r *RequestReleaseModelImpl) RequestWithAndThen(client *http2.Client, consumer common.RequestConsumer[GithubReleaseModel], timeout time.Duration) *GithubReleaseModel {
	resp := utils.Response(utils.ValidateAndModifyTimeout(client, timeout), r.url)
	if resp == nil || resp.StatusCode() != http.ResponseOkStatus {
		return nil
	}
	model, err := releaseCodec.From(resp.JSON()) // Obtain result from async.Future pass the received JSON (body).
	if err != nil {
		consumer(model)
	} else {
		fmt.Println("Error during release-model deserialization: ", err)
	}
	return model
}

// CodecRelease This struct is an implementation used for repository.GithubReleaseModel deserialization.
type CodecRelease struct {
	codec.Provider[GithubReleaseModel]
}

// From This function's override is used to handle and deserialize correctly the json's information to create a new
// repository.GithubReleaseModel object.
func (c *CodecRelease) From(json string) (*GithubReleaseModel, error) {
	var release GithubReleaseModel
	err := json2.Unmarshal([]byte(json), &release)
	if err != nil {
		return nil, err
	}
	return &release, nil
}
