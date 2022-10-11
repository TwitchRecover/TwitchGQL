package twitchgql

import (
	"strings"

	"github.com/json-iterator/go"
)

type VideoEdge struct {
	Request  VideoEdgeRequest
	Response VideoEdgeResponse
}

type VideoEdgeRequest struct {
	cursor bool
	node   *Video
}

type VideoEdgeResponse struct {
	cursor string
	node   *Video
}

func (ve *VideoEdge) RequestParser() (string, error) {
	if ve.Request == (VideoEdgeRequest{}) {
		return "", nil
	}
	query := ``
	if ve.Request.cursor {
		query += `cursor,`
	}
	if ve.Request.node != nil {
		videoQuery, _ := (ve.Request.node).RequestParser()
		query += `node{` + videoQuery[strings.Index(videoQuery, "{")+1:] + `,`
	}
	return query, nil
}

func (ve *VideoEdge) ResponseParser(res []byte) {
	if jsoniter.Get(res, "cursor").Size() > 0 {
		ve.Response.cursor = jsoniter.Get(res, "cursor").ToString()
	}
	if jsoniter.Get(res, "node").Size() > 0 {
		videoNode := []byte(`{"video":` + jsoniter.Get(res, "node").ToString() + `}`)
		(ve.Response.node).ResponseParser(videoNode)
	}
	return
}
