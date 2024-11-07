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
	AssetDownloadedStatus    = byte(0)   // Asset was downloaded.
	UnknownAssetStatus       = byte(1)   // The asset wasn't downloaded, may be unknown.
	InvalidAssetUrlStatus    = byte(2)   // The asset's URL is not valid.
	AssetDownloadErrorStatus = byte(3)   // The asset couldn't be downloaded.
	UnknownAssetDefaultSize  = int64(0)  // Used for non-downloaded (zero read bytes) assets.
	InvalidAssetDefaultSize  = int64(-1) // Used for failed-downloaded assets.
)

// DownloadingStatusProvider This struct is used as status-provider for the repositories' assets' downloads.
type DownloadingStatusProvider struct {
	Status byte  // The response's code.
	Result int64 // The amount of bytes read from the downloaded file.
}

// WithAssetDownload This method creates a new DownloadingStatusProvider using the given amount of read-bytes, and the
// AssetDownloadedStatus status.
func WithAssetDownload(result int64) DownloadingStatusProvider {
	return DownloadingStatusProvider{Status: AssetDownloadedStatus, Result: result}
}

// WithUnknownAsset This method creates a new DownloadingStatusProvider using the UnknownAssetDefaultSize for result-value,
// and providing the UnknownAssetStatus status.
func WithUnknownAsset() DownloadingStatusProvider {
	return DownloadingStatusProvider{Status: UnknownAssetStatus, Result: UnknownAssetDefaultSize}
}

// WithInvalidUrl This method creates a new DownloadingStatusProvider using the InvalidAssetDefaultSize for result-value,
// and using the InvalidAssetUrlStatus status.
func WithInvalidUrl() DownloadingStatusProvider {
	return DownloadingStatusProvider{Status: InvalidAssetUrlStatus, Result: InvalidAssetDefaultSize}
}

// WithDownloadError This method creates a new DownloadingStatusProvider using the InvalidAssetDefaultSize for result-value,
// and using the AssetDownloadErrorStatus status.
func WithDownloadError() DownloadingStatusProvider {
	return DownloadingStatusProvider{Status: AssetDownloadErrorStatus, Result: InvalidAssetDefaultSize}
}

// Downloaded This method return whether the status-code is AssetDownloadedStatus.
func (d *DownloadingStatusProvider) Downloaded() bool {
	return d.Status == AssetDownloadedStatus
}

// Unknown This method return whether the status-code is UnknownAssetStatus.
func (d *DownloadingStatusProvider) Unknown() bool {
	return d.Status == UnknownAssetStatus
}

// InvalidUrl This method return whether the status-code is InvalidAssetUrlStatus.
func (d *DownloadingStatusProvider) InvalidUrl() bool {
	return d.Status == InvalidAssetUrlStatus
}

// Error This method return whether the status-code is AssetDownloadErrorStatus.
func (d *DownloadingStatusProvider) Error() bool {
	return d.Status == AssetDownloadErrorStatus
}
