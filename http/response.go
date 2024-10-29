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

import "io"

// ResponseModel This struct represents a provided response's main information, such as Body, json-body, and status-code.
type ResponseModel struct {
	json       string
	statusCode int
	body       io.ReadCloser
}

// NewResponseModel This function creates a new ResponseModel object with the given parameters.
func NewResponseModel(json string, statusCode int, body io.ReadCloser) *ResponseModel {
	return &ResponseModel{json: json, statusCode: statusCode, body: body}
}

// JSON This method returns the json-body for this response.
func (r *ResponseModel) JSON() string {
	return r.json
}

// StatusCode This method returns the status-code for this response.
func (r *ResponseModel) StatusCode() int {
	return r.statusCode
}

// Body This method returns the Body for this response.
func (r *ResponseModel) Body() io.ReadCloser {
	return r.body
}
