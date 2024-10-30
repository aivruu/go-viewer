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

import "viewer/main/common"

const (
	ValidResponseStatus            = byte(0) // The response was provided.
	UnauthorizedResponseStatus     = byte(1) // The response was not provided because the user is not authorized.
	MovedPermanentlyResponseStatus = byte(2) // The response was not provided because the resource has moved permanently.
	ForbiddenResponseStatus        = byte(3) // The response was not provided because the user is forbidden from accessing the resource.
	InvalidResponseStatus          = byte(4) // The response was not provided because the request was invalid.
)

// ResponseStatusProvider This struct represents a status-provider used for requests' responses.
type ResponseStatusProvider[M common.RequestableModel] struct {
	status byte // The response's code.
	result *M   // The model provided for the response.
}

// WithValidResponse This method returns a new ResponseStatusProvider using the given model, and the ValidResponseStatus value.
func WithValidResponse[M common.RequestableModel](model *M) *ResponseStatusProvider[M] {
	return &ResponseStatusProvider[M]{status: ValidResponseStatus, result: model}
}

// WithUnauthorizedResponse This method returns a new ResponseStatusProvider using the given model, and the UnauthorizedResponseStatus value.
func WithUnauthorizedResponse[M common.RequestableModel]() *ResponseStatusProvider[M] {
	return &ResponseStatusProvider[M]{status: UnauthorizedResponseStatus, result: nil}
}

// WithMovedPermanentlyResponse This method returns a new ResponseStatusProvider using the given model, and the MovedPermanentlyResponseStatus value.
func WithMovedPermanentlyResponse[M common.RequestableModel]() *ResponseStatusProvider[M] {
	return &ResponseStatusProvider[M]{status: MovedPermanentlyResponseStatus, result: nil}
}

// WithForbiddenResponse This method returns a new ResponseStatusProvider using the given model, and the ForbiddenResponseStatus value.
func WithForbiddenResponse[M common.RequestableModel]() *ResponseStatusProvider[M] {
	return &ResponseStatusProvider[M]{status: ForbiddenResponseStatus, result: nil}
}

// WithInvalidResponse This method returns a new ResponseStatusProvider using the given model, and the InvalidResponseStatus value.
func WithInvalidResponse[M common.RequestableModel]() *ResponseStatusProvider[M] {
	return &ResponseStatusProvider[M]{status: InvalidResponseStatus, result: nil}
}

// Valid This method returns whether the response's status is ValidResponseStatus.
func (r *ResponseStatusProvider[M]) Valid() bool {
	return r.status == ValidResponseStatus
}

// Unauthorized This method returns whether the response's status is UnauthorizedResponseStatus.
func (r *ResponseStatusProvider[M]) Unauthorized() bool {
	return r.status == UnauthorizedResponseStatus
}

// MovedPermanently This method returns whether the response's status is MovedPermanentlyResponseStatus.
func (r *ResponseStatusProvider[M]) MovedPermanently() bool {
	return r.status == MovedPermanentlyResponseStatus
}

// Forbidden This method returns whether the response's status is ForbiddenResponseStatus.
func (r *ResponseStatusProvider[M]) Forbidden() bool {
	return r.status == ForbiddenResponseStatus
}

// Invalid This method returns whether the response's status is InvalidResponseStatus.
func (r *ResponseStatusProvider[M]) Invalid() bool {
	return r.status == InvalidResponseStatus
}

// Status This method returns the response's status.
func (r *ResponseStatusProvider[M]) Status() byte {
	return r.status
}

// Result This method returns the response's model.
func (r *ResponseStatusProvider[M]) Result() *M {
	return r.result
}
