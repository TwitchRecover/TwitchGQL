package twitchgql

import (
	"errors"
	"strconv"
	"time"

	"github.com/json-iterator/go"
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
	playbackAccessToken PlaybackAccessToken
	publishedAt         bool
	recordedAt          bool
	scope               bool
	previewsUrl         bool
	status              bool
	tags                bool
	previewThumbnailUrl ThumbnailParams
	thumbnailUrls       ThumbnailParams
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
	playbackAccessToken PlaybackAccessToken
	publishedAt         time.Time
	recordedAt          time.Time
	scope               string
	previewsUrl         string
	status              string
	tags                []string
	previewThumbnailURL string
	thumbnailUrls       []string
	title               string
	updatedAt           time.Time
	viewCount           int
	viewableAt          time.Time
}

type ThumbnailParams struct {
	height int
	width  int
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
	if req.playbackAccessToken.request != (PlaybackAccessTokenRequest{}) {
		playbackQuery, err := (req.playbackAccessToken).RequestParser()
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
		query += `seekPreviewsURL,`
	}
	if req.status {
		query += `status,`
	}
	if req.tags {
		query += `tags,`
	}
	if req.previewThumbnailUrl != (ThumbnailParams{}) {
		if req.previewThumbnailUrl.height == 0 || req.previewThumbnailUrl.width == 0 {
			return "", errors.New("Thumbnail height and/or width wasn't specified and both are required")
		}
		query += `previewThumbnailURL(height:` + string(req.previewThumbnailUrl.height) + `,width:` + string(req.previewThumbnailUrl.width) + `),`
	}
	if req.thumbnailUrls != (ThumbnailParams{}) {
		if req.thumbnailUrls.height == 0 || req.thumbnailUrls.width == 0 {
			return "", errors.New("Thumbnail height and/or width wasn't specified and both are required")
		}
		query += `thumbnailURLs(height:` + string(req.thumbnailUrls.height) + `,width:` + string(req.thumbnailUrls.width) + `),`
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

func (v *Video) ResponseParser(res []byte) {
	req := v.request
	if req.animatedPreviewUrl {
		v.response.animatedPreviewUrl = jsoniter.Get(res, "animatedPreviewURL").ToString()
	}
	if req.broadcastType {
		v.response.broadcastType = jsoniter.Get(res, "broadcastType").ToString()
	}
	if req.createdAt {
		v.response.createdAt, _ = time.Parse(time.RFC3339, jsoniter.Get(res, "createdAt").ToString())
	}
	if req.deletedAt {
		deletedAt := jsoniter.Get(res, "deletedAt").ToString()
		if deletedAt != "null" {
			v.response.deletedAt, _ = time.Parse(time.RFC3339, deletedAt)
		}
	}
	if req.description {
		v.response.description = jsoniter.Get(res, "description").ToString()
	}
	if req.download.request != (VideoDownloadRequest{}) {
		download, _ := jsoniter.Marshal(jsoniter.Get(res, "download"))
		(v.request.download).ResponseParser(download)
		v.response.download = v.request.download
	}
	if req.duration {
		durationSeconds := jsoniter.Get(res, "lengthSeconds").ToString() + "s"
		v.response.duration, _ = time.ParseDuration(durationSeconds)
	}
	if req.id {
		v.response.id = jsoniter.Get(res, "id").ToInt()
	}
	if req.softDeleted {
		v.response.softDeleted = jsoniter.Get(res, "isDeleted").ToBool()
	}
	if req.language {
		v.response.language = jsoniter.Get(res, "language").ToString()
	}
	if req.offsetSeconds {
		v.response.offsetSeconds = jsoniter.Get(res, "offsetSeconds").ToInt()
	}
	if req.playbackAccessToken.request != (PlaybackAccessTokenRequest{}) {
		pat, _ := jsoniter.Marshal(jsoniter.Get(res, "playBackAccessToken").ToString())
		(v.request.playbackAccessToken).ResponseParser(pat)
		v.response.playbackAccessToken = v.request.playbackAccessToken
	}
	if req.publishedAt {
		publishedAt := jsoniter.Get(res, "publishedAt").ToString()
		v.response.publishedAt, _ = time.Parse(time.RFC3339, publishedAt)
	}
	if req.recordedAt {
		recordedAt := jsoniter.Get(res, "recordedAt").ToString()
		v.response.recordedAt, _ = time.Parse(time.RFC3339, recordedAt)
	}
	if req.scope {
		v.response.scope = jsoniter.Get(res, "scope").ToString()
	}
	if req.previewsUrl {
		v.response.previewsUrl = jsoniter.Get(res, "seekPreviewsURL").ToString()
	}
	if req.status {
		v.response.status = jsoniter.Get(res, "status").ToString()
	}
	if req.tags {
		jsonTags := jsoniter.Get(res, "tags")
		tags := make([]string, 0)
		for i := 0; i < jsonTags.Size(); i++ {
			tags = append(tags, jsonTags.Get(i).ToString())
		}
		v.response.tags = tags
	}
	if req.previewThumbnailUrl != (ThumbnailParams{}) {
		v.response.previewThumbnailURL = jsoniter.Get(res, "previewThumbnailURL").ToString()
	}
	if req.thumbnailUrls != (ThumbnailParams{}) {
		jsonThumbnailUrls := jsoniter.Get(res, "thumbnailURLs")
		thumbnailUrls := make([]string, 0)
		for i := 0; i < jsonThumbnailUrls.Size(); i++ {
			thumbnailUrls = append(thumbnailUrls, jsonThumbnailUrls.Get(i).ToString())
		}
		v.response.thumbnailUrls = thumbnailUrls
	}
	if req.title {
		v.response.title = jsoniter.Get(res, "title").ToString()
	}
	if req.updatedAt {
		updatedAt := jsoniter.Get(res, "updatedAt").ToString()
		if updatedAt != "null" {
			v.response.updatedAt, _ = time.Parse(time.RFC3339, updatedAt)
		}
	}
	if req.viewCount {
		v.response.viewCount = jsoniter.Get(res, "viewCount").ToInt()
	}
	if req.viewableAt {
		viewableAt := jsoniter.Get(res, "viewableAt").ToString()
		if viewableAt != "null" {
			v.response.viewableAt, _ = time.Parse(time.RFC3339, viewableAt)
		}
	}
}
