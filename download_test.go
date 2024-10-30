package main

import (
	"testing"
	"viewer/main/download"
	"viewer/main/http"
	"viewer/main/repository"
)

func TestAssetDownload(t *testing.T) {
	releaseRequest := repository.NewReleaseRequest(ForRelease("aivruu", "repo-viewer", "v3.4.7"))
	release := http.Request(releaseRequest)
	if release == nil {
		t.Error("Failed to request the release for this repository.")
		return
	}
	for _, asset := range release.Assets {
		status := download.From(asset.Name, asset.Url)
		switch status.Status() {
		case download.AssetDownloadedStatus:
			t.Logf("Asset downloaded: %d bytes read", status.Result())
		case download.UnknownAssetStatus:
			t.Logf("Asset not downloaded")
		case download.AssetDownloadErrorStatus:
			t.Error("Asset failed to download")
		}
	}
}
