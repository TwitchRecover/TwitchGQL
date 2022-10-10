package twitchgql

import (
	"errors"
	"strconv"
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
	download            VideoDownload
	duration            bool
	game                bool
	id                  bool
	softDeleted         bool
	language            bool
	offsetSeconds       bool
	playBackAccessToken PlaybackAccessToken
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
	PlaybackAccessToken PlaybackAccessToken
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

func (v *Video) RequestParser() (string, error) {
	req := v.request
	query := `video(`
	if req.params.id == 0 {
		return "", errors.New("Video ID is required and was not provided")
	}
	query += `id:` + string(req.params.id) + `,`
	query += `options:{includePrivate:` + strconv.FormatBool(req.params.includePrivate) + `}){`
	if req.animatedPreviewUrl {
		query += `animatedPreviewURL,`
	}
	if req.broadcastType {
		query += `broadcastType,`
	}
	if req.createdAt {
		query += `createdAt,`
	}
	if req.creator {
		query += `creator,`
	}
	if req.deletedAt {
		query += `deletedAt,`
	}
	if req.description {
		query += `description,`
	}
	if req.download.request != (VideoDownloadRequest{}) {
		downloadQuery, err := (req.download).RequestParser()
		if err != nil {
			return "", err
		}
		query += downloadQuery
	}
	if req.duration {
		query += `lengthSeconds,`
	}
	if req.id {
		query += `id,`
	}
	if req.softDeleted {
		query += `isDeleted,`
	}
	if req.language {
		query += `language,`
	}
	if req.offsetSeconds {
		query += `offsetSeconds,`
	}
	if req.playBackAccessToken.request != (PlaybackAccessTokenRequest{}) {
		playbackQuery, err := (req.playBackAccessToken).RequestParser()
		if err != nil {
			return "", err
		}
		query += playbackQuery
	}
	if req.publishedAt {
		query += `publishedAt,`
	}
	if req.recordedAt {
		query += `recordedAt,`
	}
	if req.scope {
		query += `scope,`
	}
	if req.previewsUrl {
		query += `previewThumbnailURL,`
	}
	if req.status {
		query += `status,`
	}
	if req.tags {
		query += `tags,`
	}
	if req.thumbnailUrls {
		query += `thumbnailURLs,`
	}
	if req.title {
		query += `title,`
	}
	if req.updatedAt {
		query += `updatedAt,`
	}
	if req.viewCount {
		query += `viewCount,`
	}
	if req.viewableAt {
		query += `viewableAt,`
	}
	return query + `}`, nil
}
