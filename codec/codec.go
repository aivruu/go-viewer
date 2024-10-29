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

package codec

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	"viewer/main/http"
	"viewer/main/release"
)

// Provider This interface is used to proportionate a scalable way to deserialize json-data into http.RequestableModel structs,
// or implementations.
type Provider[M http.RequestableModel] interface {
	// From Uses the given information to return a deserialized model based on the given json-body, the model returned can
	// be of any type that implements the http.RequestableModel interface.
	From(json string) *M
}

// RepositoryCodecProvider This struct is an implementation used for repository.GithubRepositoryModel deserialization.
type RepositoryCodecProvider struct {
	Provider[GithubRepositoryModel]
}

// From This function's override is used to handle and deserialize correctly the json's information to create a new
// repository.GithubRepositoryModel object.
func (r *RepositoryCodecProvider) From(json string) *GithubRepositoryModel {
	var modelDataMap map[string]interface{} // Map where model's from json data will be stored.
	err := json2.Unmarshal([]byte(json), &modelDataMap)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// Check for possible not found repository.
	if modelDataMap["message"] != nil {
		_ = errors.New("repository requested was not found")
		return nil
	}
	return NewRepositoryModel(
		modelDataMap["owner"].(string),
		modelDataMap["name"].(string),
		modelDataMap["description"].(string),
		modelDataMap["license"].(*string),
		createAttributesContainer(&modelDataMap), // Avoid map copy.
	)
}

// createAttributesContainer This function is used to create a new repository.AttributesContainer object with the given
// parameters.
func createAttributesContainer(data *map[string]interface{}) AttributesContainer {
	mapData := *data
	return NewAttributesContainer(
		mapData["fork"].(bool),
		mapData["parent"].(*string),
		mapData["allow_forking"].(bool),
		mapData["stargazers_count"].(int32),
		mapData["forks_count"].(int32),
		mapData["private"].(bool),
		mapData["archived"].(bool),
		mapData["disabled"].(bool),
		mapData["language"].(string),
		mapData["topics"].(*[]string),
	)
}

// RequestCodecProvider This struct is an implementation used for repository.GithubReleaseModel deserialization.
type RequestCodecProvider struct {
	Provider[release.GithubReleaseModel]
}

// From This function's override is used to handle and deserialize correctly the json's information to create a new
// repository.GithubReleaseModel object.
func (r *RequestCodecProvider) From(json string) *release.GithubReleaseModel {
	var modelDataMap map[string]interface{}
	err := json2.Unmarshal([]byte(json), &modelDataMap)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if modelDataMap["message"] != nil {
		_ = errors.New("repository's release requested was not found")
		return nil
	}
	return release.NewReleaseModel(
		modelDataMap["author"].(string),
		modelDataMap["tag_name"].(string),
		modelDataMap["name"].(string),
		modelDataMap["id"].(int),
		modelDataMap["assets"].(*[]string),
	)
}
