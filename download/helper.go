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
	"path/filepath"
	"strings"
	"viewer/main/utils"
)

func validGithubUrl(url string) bool {
	return strings.HasPrefix(url, "https://github.com/") || strings.HasPrefix(url, "https://api.github.com/")
}

// From This function downloads the content from the given url into the specified file-name, and returns a DownloadStatusProvider.
func From(directory string, fileName string, url string) DownloadingStatusProvider {
	if !validGithubUrl(url) {
		return WithInvalidUrl()
	}
	file, err := os.Create(filepath.Join(directory, fileName))
	if err != nil {
		fmt.Println("Error during file creation: ", err)
		return WithDownloadError()
	}
	// [os.File] object closing and error handling.
	defer func(File *os.File) {
		if err := File.Close(); err != nil {
			fmt.Println("Error during File closing: ", err)
		}
	}(file)
	// Make request to the given url and get the [Response] object.
	request := utils.OriginalResponse(url)
	resp := request.Get()
	if resp == nil {
		return WithDownloadError()
	}
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error body's information copying into file: ", err)
		return WithDownloadError()
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Error during body closing: ", err)
		}
	}(resp.Body)
	if size == 0 {
		return WithUnknownAsset()
	}
	return WithAssetDownload(size)
}
