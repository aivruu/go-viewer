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
	"viewer/main/async"
	"viewer/main/codec"
	"viewer/main/functional"
	"viewer/main/http"
	"viewer/main/utils"
)

type GithubReleaseModel struct {
	author   string
	tagName  string
	name     string
	uniqueId int
	assets   *[]string
	http.RequestableModel
}

func NewReleaseModel(author, tagName, name string, uniqueId int, assets *[]string) *GithubReleaseModel {
	return &GithubReleaseModel{
		author:   author,
		tagName:  tagName,
		name:     name,
		uniqueId: uniqueId,
		assets:   assets,
	}
}

func Type() string {
	return "request"
}

func (r *GithubReleaseModel) Author() string {
	return r.author
}

func (r *GithubReleaseModel) TagName() string {
	return r.tagName
}

func (r *GithubReleaseModel) Name() string {
	return r.name
}

func (r *GithubReleaseModel) UniqueId() int {
	return r.uniqueId
}

func (r *GithubReleaseModel) Assets() *[]string {
	return r.assets
}

var requestCodec = codec.RequestCodecProvider{}

type RequestModelImpl struct {
	http.RequestModel[GithubReleaseModel]
	responseModel *http.ResponseModel
	url           string
}

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
	return async.NewFuture(func() *GithubReleaseModel {
		return requestCodec.From(responseModel.JSON())
	})
}

func RequestAndThen(consumer functional.RequestConsumer[GithubReleaseModel]) *async.Future[GithubReleaseModel] {
	f := utils.Response(http.Url)
	if f == nil {
		return nil
	}
	releaseModel := requestCodec.From(f.Get().JSON()) // Obtain result from async.Future pass the received JSON (body).
	consumer(releaseModel)                            // Execute consumer logic with deserialized model
	return async.NewFuture(func() *GithubReleaseModel {
		return releaseModel
	})
}
