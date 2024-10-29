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

package utils

import (
	"io"
	"net/http"
	"viewer/main/async"
	"viewer/main/codec"
	"viewer/main/functional"
	vhttp "viewer/main/http"
	status "viewer/main/http/response"
)

// Response This function makes a request to the given url, and returns a pointer of async.Future, which contains the http.ResponseModel
// with this request's response's main-information.
func Response(url string) *async.Future[vhttp.ResponseModel] {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	read, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	// Close body after reading.
	defer func(Body io.ReadCloser) {
		panic(Body.Close())
	}(resp.Body)
	return async.NewFuture(func() vhttp.ResponseModel {
		return *vhttp.NewResponseModel(string(read), resp.StatusCode, resp.Body)
	})
}

// VerifyAndProvideResponse This function verifies the given response (if it is available), and uses their status-code to verify
// and provide a http.ResponseStatusProvider depending on given response-code.
//
// Additionally, if the response is valid, the model will be used to execute the given consumer.
func VerifyAndProvideResponse[M vhttp.RequestableModel](resp *vhttp.ResponseModel, consumer *functional.RequestConsumer[M], codecProvider codec.Provider[M]) *vhttp.ResponseStatusProvider[M] {
	if resp == nil {
		return vhttp.WithUnauthorizedResponse[M]()
	}
	switch resp.StatusCode() {
	case status.NotFound:
		return vhttp.WithInvalidResponse[M]()
	case status.Unauthorized:
		return vhttp.WithUnauthorizedResponse[M]()
	case status.MovedPermanently:
		return vhttp.WithMovedPermanentlyResponse[M]()
	case status.Forbidden:
		return vhttp.WithForbiddenResponse[M]()
	case status.OK:
		model := codecProvider.From(resp.JSON())
		if model == nil {
			return vhttp.WithInvalidResponse[M]()
		}
		if consumer != nil {
			(*consumer)(model)
		}
		return vhttp.WithValidResponse(model)
	default:
		return vhttp.WithInvalidResponse[M]()
	}
}
