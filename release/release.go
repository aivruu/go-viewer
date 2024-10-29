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

package release

import (
	"strconv"
	"strings"
	"viewer/main/async"
	"viewer/main/codec"
	"viewer/main/functional"
	"viewer/main/http"
	"viewer/main/utils"
)

// GithubReleaseModel This struct stores all necessary information for the repository's requested release.
type GithubReleaseModel struct {
	author   string
	tagName  string
	name     string
	uniqueId int
	assets   *[]string
	http.RequestableModel
}

// NewReleaseModel This function creates a new GithubReleaseModel using the given parameters.
func NewReleaseModel(author, tagName, name string, uniqueId int, assets *[]string) *GithubReleaseModel {
	return &GithubReleaseModel{
		author:   author,
		tagName:  tagName,
		name:     name,
		uniqueId: uniqueId,
		assets:   assets,
	}
}

// Type Override to interface's function to return the specific-type for this model.
func Type() string {
	return "request"
}

// Author This method returns this release's author's name.
func (r *GithubReleaseModel) Author() string {
	return r.author
}

// TagName This method returns this release's tag-name.
func (r *GithubReleaseModel) TagName() string {
	return r.tagName
}

// Name This method returns this release's name.
func (r *GithubReleaseModel) Name() string {
	return r.name
}

// UniqueId This method returns this release's unique identifier.
func (r *GithubReleaseModel) UniqueId() int {
	return r.uniqueId
}

// Assets This method returns this release's assets.
func (r *GithubReleaseModel) Assets() *[]string {
	return r.assets
}

// Compare This method compares the given version-number with this release's tag-name (as int) using the specified operator-type
// for the comparison, and return a bool as operation's result.
func (r *GithubReleaseModel) Compare(operatorType Operator, targetVersion int) bool {
	var versionWithoutDots = strings.Replace(r.tagName, "v", "", -1)
	var version, _ = strconv.ParseInt(versionWithoutDots, 8, 16)
	var numVersion = int(version) // Required to could compare both different int-type values.
	switch operatorType {
	case Equal:
		return targetVersion == numVersion
	case Less:
		return targetVersion < numVersion
	case LessOrEqual:
		return targetVersion <= numVersion
	case Greater:
		return targetVersion > numVersion
	case GreaterOrEqual:
		return targetVersion >= numVersion
	default:
		return false
	}
}

// codec.Provider's implementation necessary for this type.
var requestCodec = codec.RequestCodecProvider{}

// RequestModelImpl This codec.Provider implementation is used to handle requests for repositories' releases.
type RequestModelImpl struct {
	http.RequestModel[GithubReleaseModel]
	responseModel *http.ResponseModel
	url           string
}

// NewReleaseRequest This function creates a new RequestModelImpl with the given url.
func NewReleaseRequest(url string) *RequestModelImpl {
	http.Url = url
	return &RequestModelImpl{url: http.Url}
}

func Request() *async.Future[GithubReleaseModel] {
	f := utils.Response(http.Url)
	if f == nil {
		return nil
	}
	responseModel := f.Get()
	return async.NewFuture(func() GithubReleaseModel {
		return *requestCodec.From(responseModel.JSON())
	})
}

func RequestAndThen(consumer functional.RequestConsumer[GithubReleaseModel]) *async.Future[GithubReleaseModel] {
	f := utils.Response(http.Url)
	if f == nil {
		return nil
	}
	releaseModel := requestCodec.From(f.Get().JSON()) // Obtain result from async.Future pass the received JSON (body).
	consumer(releaseModel)                            // Execute consumer logic with deserialized model
	return async.NewFuture(func() GithubReleaseModel {
		return *releaseModel
	})
}
