package twitchgql

import (
	"time"
)

type Video struct {
	request  VideoRequest
	response VideoResponse
}

type VideoRequest struct {
	params              VideoRequestParams
	animatedPreviewUrl  bool
	broadcastType       bool
	createdAt           bool
	creator             bool
	deletedAt           bool
	description         bool
	download            bool
	duration            bool
	game                bool
	id                  bool
	softDeleted         bool
	language            bool
	offsetSeconds       bool
	playBackAccessToken PlaybackAccessTokenRequest
	publishedAt         bool
	recordedAt          bool
	scope               bool
	previewsUrl         bool
	status              bool
	tags                bool
	thumbnailUrls       bool
	title               bool
	updatedAt           bool
	viewCount           bool
	viewableAt          bool
}

type VideoRequestParams struct {
	id             int
	includePrivate bool // Include private videos. Will return an error if unauthenticated.
}

type VideoResponse struct {
	animatedPreviewUrl  string
	broadcastType       string
	createdAt           time.Time
	deletedAt           time.Time
	description         string
	download            VideoDownload
	duration            time.Duration
	id                  int
	softDeleted         bool
	language            string
	offsetSeconds       int
	PlaybackAccessToken PlaybackAccessTokenResponse
	publishedAt         time.Time
	recordedAt          time.Time
	scope               string
	previewsUrl         string
	status              string
	tags                []string
	thumbnailUrls       []string
	title               string
	updatedAt           time.Time
	viewCount           int
	viewableAt          time.Time
}

type VideoDownload struct {
	status string
	url    string
}
