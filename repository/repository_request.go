package repository

import (
	"fmt"
	http2 "net/http"
	"time"
	"viewer/main/common"
	"viewer/main/http"
	"viewer/main/utils"
)

// codec.Provider's implementation necessary for this type.
var repositoryCodec = RepositoryCodecProvider{}

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
