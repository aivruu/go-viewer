package repository

import (
	json2 "encoding/json"
	"viewer/main/repository/codec"
)

// RepositoryCodecProvider This struct is an implementation used for repository.GithubRepositoryModel deserialization.
type RepositoryCodecProvider struct {
	codec.Provider[GithubRepositoryModel]
}

// From This function's override is used to handle and deserialize correctly the json's information to create a new
// repository.GithubRepositoryModel object.
func (r *RepositoryCodecProvider) From(json string) (*GithubRepositoryModel, error) {
	var model GithubRepositoryModel
	if err := json2.Unmarshal([]byte(json), &model); err != nil {
		return nil, err
	}
	return &model, nil
}
