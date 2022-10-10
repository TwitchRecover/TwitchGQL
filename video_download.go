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

func (vd *VideoDownload) RequestParser() (string, error) {
	if vd.request == (VideoDownloadRequest{}) {
		return "", nil
	}
	query := `download{`
	if vd.request.status {
		query += `status,`
	}
	if vd.request.url {
		query += `url,`
	}
	query += `},`
	return query, nil
}

func (vd *VideoDownload) ResponseParser(res []byte) {
	if jsoniter.Get(res, "url").ToBool() {
		vd.response.url = jsoniter.Get(res, "url").ToString()
	}
	if jsoniter.Get(res, "status").ToBool() {
		vd.response.status = jsoniter.Get(res, "status").ToString()
	}
}
