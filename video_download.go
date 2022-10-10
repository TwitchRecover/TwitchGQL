package twitchgql

import (
	"github.com/json-iterator/go"
)

type VideoDownload struct {
	Request  VideoDownloadRequest
	Response VideoDownloadResponse
}

type VideoDownloadRequest struct {
	Status bool
	Url    bool
}

type VideoDownloadResponse struct {
	Status string
	Url    string
}

func (vd *VideoDownload) RequestParser() (string, error) {
	if vd.Request == (VideoDownloadRequest{}) {
		return "", nil
	}
	query := `download{`
	if vd.Request.Status {
		query += `status,`
	}
	if vd.Request.Url {
		query += `url,`
	}
	return query + `},`, nil
}

func (vd *VideoDownload) ResponseParser(res []byte) {
	if vd.Request.Url {
		vd.Response.Url = jsoniter.Get(res, "url").ToString()
	}
	if vd.Request.Status {
		vd.Response.Status = jsoniter.Get(res, "status").ToString()
	}
}
