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
	"viewer/main/async"
	"viewer/main/functional"
)

// Url The URL that contains the url to where the request will be sent.
var Url string

// RequestModel This interface is used to proportionate request-method to get information for models-serialization.
type RequestModel[M RequestableModel] interface {
	// Request This method request to the URL and returns an async.Future with the model object.
	Request() *async.Future[M]

	// RequestAndThen This method request to the URL and provides an async.Future with the model. If the model is available,
	// it will be used to execute the specified consumer's logic.
	RequestAndThen(consumer functional.RequestConsumer[M]) *async.Future[M]
}
