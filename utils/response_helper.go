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
	"fmt"
	"io"
	"net/http"
	"viewer/main/async"
	vhttp "viewer/main/http"
)

// AsyncResponseWith This function makes a request to the given url using the given http.Client, and returns an async.Future,
// this object's function may return a http.ResponseModel, or nil depending on operation success.
func asyncResponseWith(client *http.Client, url string) async.Future[vhttp.ResponseModel] {
	return async.NewFuture(func() *vhttp.ResponseModel {
		resp, err := client.Get(url)
		if err != nil {
			return nil
		}
		read, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil
		}
		// Close body after reading.
		defer func(Body *io.ReadCloser) {
			err := (*Body).Close()
			if err != nil {
				fmt.Println("Error during Body closing: ", err)
			}
		}(&resp.Body)
		return vhttp.NewResponseModel(string(read), resp.StatusCode, &resp.Body)
	})
}

// Response This function realizes a request and provides an async response of async.Future type, this function checks if the
// provided http.Client is nil, if it is, it will call to the asyncResponse function, otherwise, it will use asyncResponseWith
func Response(client *http.Client, url string) *vhttp.ResponseModel {
	var response async.Future[vhttp.ResponseModel]
	if client == nil {
		response = asyncResponse(url)
	} else {
		response = asyncResponseWith(client, url)
	}
	return response.Get()
}

// AsyncResponse This function makes a request to the given url using a default client and returns an async.Future,
// this object's function may return a http.ResponseModel, or nil depending on operation success.
func asyncResponse(url string) async.Future[vhttp.ResponseModel] {
	return async.NewFuture(func() *vhttp.ResponseModel {
		resp, err := vhttp.DefaultClient.Get(url)
		if err != nil {
			return nil
		}
		read, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil
		}
		// Close body after reading.
		defer func(Body *io.ReadCloser) {
			err := (*Body).Close()
			if err != nil {
				fmt.Println("Error during Body closing: ", err)
			}
		}(&resp.Body)
		return vhttp.NewResponseModel(string(read), resp.StatusCode, &resp.Body)
	})
}
