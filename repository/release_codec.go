package repository

import (
	json2 "encoding/json"
	"viewer/main/repository/codec"
)

// ReleaseCodecProvider This struct is an implementation used for repository.GithubReleaseModel deserialization.
type ReleaseCodecProvider struct {
	codec.Provider[GithubReleaseModel]
}

// From This function's override is used to handle and deserialize correctly the json's information to create a new
// repository.GithubReleaseModel object.
func (c *ReleaseCodecProvider) From(json string) (*GithubReleaseModel, error) {
	var model GithubReleaseModel
	if err := json2.Unmarshal([]byte(json), &model); err != nil {
		return nil, err
	}
	return &model, nil
}
