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

package download

import (
	"fmt"
	"io"
	"os"
	"viewer/main/utils"
)

// From This function downloads the content from the given url into the specified file-name, and returns a DownloadStatusProvider.
func From(fileName string, url string) *DownloadingStatusProvider {
	file, err := os.Create(fileName)
	if err != nil {
		return WithDownloadError()
	}
	defer func(File *os.File) {
		err := File.Close()
		if err != nil {
			fmt.Println("Error during File closing: ", err)
		}
	}(file)
	resp := utils.Response(nil, url)
	if resp == nil {
		return WithDownloadError()
	}
	body := resp.Body()
	if body == nil {
		return WithDownloadError()
	}
	defer func(Body *io.ReadCloser) {
		err := (*Body).Close()
		if err != nil {
			fmt.Println("Error during Body closing: ", err)
		}
	}(body)
	size, err := io.Copy(file, *body)
	if err != nil {
		return WithDownloadError()
	}
	if size == 0 {
		return WithUnknownAsset()
	}
	return WithAssetDownload(size)
}
