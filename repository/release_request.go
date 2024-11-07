package repository

import (
	"fmt"
	http2 "net/http"
	"time"
	"viewer/main/http"
	"viewer/main/utils"
)

// releaseCodec codec.Provider's implementation necessary for this type.
var releaseCodec = ReleaseCodecProvider{}

// RequestReleaseModelImpl This http.RequestModel implementation is used to handle http-requests for repositories' releases.
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
	if resp == nil || resp.StatusCode != http.ResponseOkStatus {
		return nil
	}
	model, err := releaseCodec.From(resp.JSON)
	if err != nil {
		fmt.Println("Error during release-model deserialization: ", err)
	}
	return model
}

func (r *RequestReleaseModelImpl) RequestWithAndThen(client *http2.Client, consumer func(*GithubReleaseModel), timeout time.Duration) *GithubReleaseModel {
	resp := utils.Response(utils.ValidateAndModifyTimeout(client, timeout), r.url)
	if resp == nil || resp.StatusCode != http.ResponseOkStatus {
		return nil
	}
	model, err := releaseCodec.From(resp.JSON) // Obtain result from async.Future pass the received JSON (body).
	if err != nil {
		consumer(model)
	} else {
		fmt.Println("Error during release-model deserialization: ", err)
	}
	return model
}
