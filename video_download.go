package twitchgql

import (
	"github.com/json-iterator/go"
)

type VideoDownload struct {
	request  VideoDownloadRequest
	response VideoDownloadResponse
}

type VideoDownloadRequest struct {
	status bool
	url    bool
}

type VideoDownloadResponse struct {
	status string
	url    string
}

func (vd *VideoDownload) RequestParser(vdr VideoDownloadRequest) (string, error) {
	if vdr == (VideoDownloadRequest{}) {
		return "", nil
	}
	query := `download{`
	if vdr.status {
		query += `status,`
	}
	if vdr.url {
		query += `url,`
	}
	query += `},`
	return query, nil
}

func (vd *VideoDownload) ResponseParser(res []byte) (VideoDownloadResponse, error) {
	response := VideoDownloadResponse{}
	if jsoniter.Get(res, "url").ToBool() {
		response.url = jsoniter.Get(res, "url").ToString()
	}
	if jsoniter.Get(res, "status").ToBool() {
		response.status = jsoniter.Get(res, "status").ToString()
	}
	return response, nil
}
