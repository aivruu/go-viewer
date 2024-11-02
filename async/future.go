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

package async

import "sync"

// Future This struct is used to represents and manages the result of an asynchronous computation.
type Future[T any] struct {
	result chan T
}

// NewFuture This method creates a new Future object using the specified function, this function may return a value, or nil.
func NewFuture[T any](fn func() *T) Future[T] {
	rwMut := sync.RWMutex{}
	f := Future[T]{make(chan T)}
	rwMut.RLock()
	// Use goroutines for channel-to-channel communication.
	go func() {
		result := fn()
		// Avoid dereferencing for a null pointer
		if result == nil {
			result = new(T)
		}
		// Give to channel the pointer's value.
		f.result <- *result
	}()
	rwMut.RUnlock()
	return f
}

// Get This method returns this Future's channel's value (result) when it becomes available.
func (f *Future[T]) Get() *T {
	value := <-f.result
	return &value
}
