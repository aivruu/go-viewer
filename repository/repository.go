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
	"time"
	"viewer/main/common"
	"viewer/main/http"
	"viewer/main/repository/codec"
	"viewer/main/utils"
)

// FormatBooleanValue This function returns a readable string that correspond to the value for the given boolean.
func FormatBooleanValue(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}

// GithubRepositoryModel This struct represents a requested repository with all its information.
type GithubRepositoryModel struct {
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
	License struct {
		Name string `json:"name"`
	} `json:"license"`
	Parent struct {
		Owner string `json:"owner"`
	} `json:"parent"`
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

// codec.Provider's implementation necessary for this type.
var repositoryCodec = CodecRepository{}

// RequestRepositoryModelImpl This codec.Provider implementation is used to handle requests for repositories.
type RequestRepositoryModelImpl struct {
	http.RequestModel[GithubRepositoryModel]
	url string
}

// NewRepositoryRequest This function creates a new RequestReleaseModelImpl with the given url.
func NewRepositoryRequest(url string) *RequestRepositoryModelImpl {
	return &RequestRepositoryModelImpl{url: url}
}

func (r *RequestRepositoryModelImpl) RequestWith(client *http2.Client, timeout time.Duration) *GithubRepositoryModel {
	resp := utils.Response(utils.ValidateAndModifyTimeout(client, timeout), r.url)
	if resp == nil || resp.StatusCode() != http.ResponseOkStatus {
		return nil
	}
	model, err := repositoryCodec.From(resp.JSON())
	if err != nil {
		fmt.Println("Error during repository-model deserialization: ", err)
	}
	return model
}

func (r *RequestRepositoryModelImpl) RequestWithAndThen(client *http2.Client, consumer common.RequestConsumer[GithubRepositoryModel], timeout time.Duration) *GithubRepositoryModel {
	resp := utils.Response(utils.ValidateAndModifyTimeout(client, timeout), r.url)
	if resp == nil || resp.StatusCode() != http.ResponseOkStatus {
		return nil
	}
	model, err := repositoryCodec.From(resp.JSON())
	if err == nil {
		consumer(model)
	} else {
		fmt.Println("Error during repository-model deserialization: ", err)
	}
	return model
}

// CodecRepository This struct is an implementation used for repository.GithubRepositoryModel deserialization.
type CodecRepository struct {
	codec.Provider[GithubRepositoryModel]
}

// From This function's override is used to handle and deserialize correctly the json's information to create a new
// repository.GithubRepositoryModel object.
func (r *CodecRepository) From(json string) (*GithubRepositoryModel, error) {
	var repository GithubRepositoryModel
	err := json2.Unmarshal([]byte(json), &repository)
	if err != nil {
		return nil, err
	}
	return &repository, nil
}
