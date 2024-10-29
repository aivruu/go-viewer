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

const (
	AssetDownloadedStatus    = byte(0)
	UnknownAssetStatus       = byte(1)
	AssetDownloadErrorStatus = byte(2)
	UnknownAssetDefaultSize  = int64(0)  // Used for non-downloaded (zero read bytes) assets.
	InvalidAssetDefaultSize  = int64(-1) // Used for failed-downloaded assets.
)

// DownloadingStatusProvider This struct is used as status-provider for the repositories' assets' downloads.
type DownloadingStatusProvider struct {
	status byte  // The response's code.
	result int64 // The amount of bytes read from the downloaded file.
}

// WithAssetDownload This method creates a new DownloadingStatusProvider using the given amount of read-bytes, and the
// AssetDownloadedStatus status.
func WithAssetDownload(result int64) *DownloadingStatusProvider {
	return &DownloadingStatusProvider{status: AssetDownloadedStatus, result: result}
}

// WithUnknownAsset This method creates a new DownloadingStatusProvider using the UnknownAssetDefaultSize for result-value,
// and providing the UnknownAssetStatus status.
func WithUnknownAsset() *DownloadingStatusProvider {
	return &DownloadingStatusProvider{status: UnknownAssetStatus, result: UnknownAssetDefaultSize}
}

// WithDownloadError This method creates a new DownloadingStatusProvider using the InvalidAssetDefaultSize for result-value,
// and using the AssetDownloadErrorStatus status.
func WithDownloadError() *DownloadingStatusProvider {
	return &DownloadingStatusProvider{status: AssetDownloadErrorStatus, result: InvalidAssetDefaultSize}
}
