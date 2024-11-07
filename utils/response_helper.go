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
	"time"
	"viewer/main/async"
	vhttp "viewer/main/http"
)

// AsyncResponseWith This function makes a request to the given url using the given http.Client and will return an
// async.Future, this object's function may return a http.ResponseModel, or null depending on operation success.
func asyncResponse(client *http.Client, url string) async.Future[vhttp.ResponseModel] {
	return async.NewFuture(func() *vhttp.ResponseModel {
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("Error during request: ", err)
			return nil
		}
		read, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error during reading: ", err)
			return nil
		}
		// Close body after reading.
		defer func(Body io.ReadCloser) {
			if err := Body.Close(); err != nil {
				fmt.Println("Error during Body closing: ", err)
			}
		}(resp.Body)
		return &vhttp.ResponseModel{JSON: string(read), StatusCode: resp.StatusCode, Body: resp.Body}
	})
}

// Response This function calls internally to the asyncResponse and when is available, it will return the http.ResponseModel
// provided by that function.
func Response(client *http.Client, url string) *vhttp.ResponseModel {
	f := asyncResponse(client, url)
	// Return [ResponseModel] object when available.
	return f.Get()
}

// ValidateAndModifyTimeout This function validates the given client and then modifies the client's timeout. If the http.Client
// instance is null or is the http.DefaultClient reference, the function will return the DefaultClient instance instead,
// otherwise, it will modify the client's timeout and return the same client instance (no copies).
func ValidateAndModifyTimeout(client *http.Client, timeout time.Duration) *http.Client {
	if client == nil || client == http.DefaultClient {
		return vhttp.DefaultClient
	}
	client.Timeout = timeout
	return client
}

// OriginalResponse This function makes an async request to the given url, and returns the built-in http.Response object,
// and not a http.ResponseModel.
func OriginalResponse(url string) async.Future[http.Response] {
	return async.NewFuture(func() *http.Response {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error during request: ", err)
			return nil
		}
		return resp
	})
}
