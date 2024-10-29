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
	"viewer/main/async"
	"viewer/main/codec"
	"viewer/main/functional"
	"viewer/main/http"
	"viewer/main/utils"
)

// GithubRepositoryModel This struct represents a requested repository with all its information.
type GithubRepositoryModel struct {
	owner       string
	name        string
	description string
	license     *string // The repository's current license-type, nil if no license is set.
	attributes  AttributesContainer
	http.RequestableModel
}

// NewRepositoryModel This method creates a new GithubRepositoryModel instance using the given parameters.
func NewRepositoryModel(owner, name, description string, license *string, attributes AttributesContainer) *GithubRepositoryModel {
	return &GithubRepositoryModel{
		owner:       owner,
		name:        name,
		description: description,
		license:     license,
		attributes:  attributes,
	}
}

// Type Override to interface's function to return the specific-type for this model.
func Type() string {
	return "repository"
}

// Owner This method returns this repository's owner's name.
func (r *GithubRepositoryModel) Owner() string {
	return r.owner
}

// Name This method returns this repository's name.
func (r *GithubRepositoryModel) Name() string {
	return r.name
}

// Description This method returns this repository's description.
func (r *GithubRepositoryModel) Description() string {
	return r.description
}

// License This method returns this repository's license-type (if it is available).
func (r *GithubRepositoryModel) License() *string {
	return r.license
}

// Attributes This method returns this repository's attributes.
func (r *GithubRepositoryModel) Attributes() *AttributesContainer {
	return &r.attributes
}

// codec.Provider's implementation necessary for this type.
var repositoryCodec = codec.RepositoryCodecProvider{}

// RequestModelImpl This codec.Provider implementation is used to handle requests for repositories.
type RequestModelImpl struct {
	http.RequestModel[GithubRepositoryModel]
	responseModel *http.ResponseModel
	url           string
}

// NewRepositoryRequest This function creates a new RequestModelImpl with the given url.
func NewRepositoryRequest(url string) *RequestModelImpl {
	http.Url = url
	return &RequestModelImpl{url: http.Url}
}

func Request() *async.Future[GithubRepositoryModel] {
	f := utils.Response(http.Url)
	if f == nil {
		return nil
	}
	responseModel := f.Get()
	return async.NewFuture(func() GithubRepositoryModel {
		return *repositoryCodec.From(responseModel.JSON())
	})
}

func RequestAndThen(consumer functional.RequestConsumer[GithubRepositoryModel]) *async.Future[GithubRepositoryModel] {
	f := utils.Response(http.Url)
	if f == nil {
		return nil
	}
	repositoryModel := repositoryCodec.From(f.Get().JSON()) // Obtain result from async.Future pass the received JSON (body).
	consumer(repositoryModel)                               // Execute consumer logic with deserialized model
	return async.NewFuture(func() GithubRepositoryModel {
		return *repositoryModel
	})
}
