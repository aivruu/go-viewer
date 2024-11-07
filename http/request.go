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

package http

import (
	"net/http"
	"time"
	"viewer/main/common"
)

// DefaultClient A default http.Client instance with a defined timeout for any request, this instance is used during with
// no-specified http.Client, the timeout set is to avoid modify original timeout for the built-in default http.Client
// instance.
var DefaultClient = &http.Client{Timeout: 5 * time.Second}

// ResponseOkStatus Correspond to status-code provided if the request was accepted and a response was provided.
const ResponseOkStatus = 200

// RequestModel This interface is used to proportionate request-method to get information from GitHub API to perform multiple
// functions with the received information, such as, check or download content.
type RequestModel[M common.RequestableModel] interface {
	// RequestWith This method request to the URL using the given http.Client and the specified timeout to return the model
	// with the requested information if it is available.
	RequestWith(client *http.Client, timeout time.Duration) *M

	// RequestWithAndThen This method request to the URL using the given http.Client and timeout to provide the model
	// (if it is available) with the requested information. Also, if the model is available, it will be used to execute the
	// specified consumer's logic.
	RequestWithAndThen(client *http.Client, consumer func(*M), timeout time.Duration) *M
}

// Request This function realizes the same execution that RequestAndThen with the difference that this uses a default
// http.Client to make the request.
func Request[M common.RequestableModel](requestModel RequestModel[M], timeout time.Duration) *M {
	return requestModel.RequestWith(nil, timeout)
}

// RequestAndThen This function realizes the same execution that RequestWithAndThen with the difference that this uses a
// default http.Client to make the request.
func RequestAndThen[M common.RequestableModel](requestModel RequestModel[M], consumer func(*M), timeout time.Duration) *M {
	return requestModel.RequestWithAndThen(nil, consumer, timeout)
}
