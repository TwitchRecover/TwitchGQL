package twitchgql

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type Video struct {
	Request  VideoRequest
	Response VideoResponse
}

type VideoRequest struct {
	Params              VideoRequestParams
	AnimatedPreviewUrl  bool
	BroadcastType       bool
	CreatedAt           bool
	Creator             bool
	DeletedAt           bool
	Description         bool
	Download            *VideoDownload
	Duration            bool
	Game                bool
	Id                  bool
	SoftDeleted         bool
	Language            bool
	OffsetSeconds       bool
	PlaybackAccessToken *PlaybackAccessToken
	PublishedAt         bool
	RecordedAt          bool
	Scope               bool
	PreviewsUrl         bool
	Status              bool
	Tags                bool
	PreviewThumbnailUrl ImageParams
	ThumbnailUrls       ImageParams
	Title               bool
	UpdatedAt           bool
	ViewCount           bool
	ViewableAt          bool
}

type VideoRequestParams struct {
	Id             int
	IncludePrivate bool // Include private videos. Will return an error if unauthenticated.
}

type VideoResponse struct {
	AnimatedPreviewUrl  string
	BroadcastType       string
	CreatedAt           time.Time
	DeletedAt           time.Time
	Description         string
	Download            *VideoDownload
	Duration            time.Duration
	Id                  int
	SoftDeleted         bool
	Language            string
	OffsetSeconds       int
	PlaybackAccessToken *PlaybackAccessToken
	PublishedAt         time.Time
	RecordedAt          time.Time
	Scope               string
	PreviewsUrl         string
	Status              string
	Tags                []string
	PreviewThumbnailURL string
	ThumbnailUrls       []string
	Title               string
	UpdatedAt           time.Time
	ViewCount           int
	ViewableAt          time.Time
}

func (v *Video) RequestParser() (string, error) {
	req := v.Request
	query := `video(`
	if req.Params.Id == 0 {
		return "", errors.New("Video ID is required and was not provided")
	}
	query += `id:` + fmt.Sprint(req.Params.Id) + `,`
	query += `options:{includePrivate:` + strconv.FormatBool(req.Params.IncludePrivate) + `}){`
	if req.AnimatedPreviewUrl {
		query += `animatedPreviewURL,`
	}
	if req.BroadcastType {
		query += `broadcastType,`
	}
	if req.CreatedAt {
		query += `createdAt,`
	}
	if req.Creator {
		query += `creator,`
	}
	if req.DeletedAt {
		query += `deletedAt,`
	}
	if req.Description {
		query += `description,`
	}
	if req.Download != nil {
		downloadQuery, err := (req.Download).RequestParser()
		if err != nil {
			return "", err
		}
		query += downloadQuery
	}
	if req.Duration {
		query += `lengthSeconds,`
	}
	if req.Id {
		query += `id,`
	}
	if req.SoftDeleted {
		query += `isDeleted,`
	}
	if req.Language {
		query += `language,`
	}
	if req.OffsetSeconds {
		query += `offsetSeconds,`
	}
	if req.PlaybackAccessToken != nil {
		playbackQuery, err := (req.PlaybackAccessToken).RequestParser()
		if err != nil {
			return "", err
		}
		query += playbackQuery
	}
	if req.PublishedAt {
		query += `publishedAt,`
	}
	if req.RecordedAt {
		query += `recordedAt,`
	}
	if req.Scope {
		query += `scope,`
	}
	if req.PreviewsUrl {
		query += `seekPreviewsURL,`
	}
	if req.Status {
		query += `status,`
	}
	if req.Tags {
		query += `tags,`
	}
	if req.PreviewThumbnailUrl != (ImageParams{}) {
		if req.PreviewThumbnailUrl.Height == 0 || req.PreviewThumbnailUrl.Width == 0 {
			return "", errors.New("Thumbnail height and/or width wasn't specified and both are required")
		}
		query += `previewThumbnailURL(height:` + fmt.Sprint(req.PreviewThumbnailUrl.Height) + `,width:` + fmt.Sprint(req.PreviewThumbnailUrl.Width) + `),`
	}
	if req.ThumbnailUrls != (ImageParams{}) {
		if req.ThumbnailUrls.Height == 0 || req.ThumbnailUrls.Width == 0 {
			return "", errors.New("Thumbnail height and/or width wasn't specified and both are required")
		}
		query += `thumbnailURLs(height:` + fmt.Sprint(req.ThumbnailUrls.Height) + `,width:` + fmt.Sprint(req.ThumbnailUrls.Width) + `),`
	}
	if req.Title {
		query += `title,`
	}
	if req.UpdatedAt {
		query += `updatedAt,`
	}
	if req.ViewCount {
		query += `viewCount,`
	}
	if req.ViewableAt {
		query += `viewableAt,`
	}
	return query + `}`, nil
}

func (v *Video) ResponseParser(res []byte) {
	req := v.Request
	res, _ = jsoniter.Marshal(jsoniter.Get(res, "video"))
	if req.AnimatedPreviewUrl {
		v.Response.AnimatedPreviewUrl = jsoniter.Get(res, "animatedPreviewURL").ToString()
	}
	if req.BroadcastType {
		v.Response.BroadcastType = jsoniter.Get(res, "broadcastType").ToString()
	}
	if req.CreatedAt {
		v.Response.CreatedAt, _ = time.Parse(time.RFC3339, jsoniter.Get(res, "createdAt").ToString())
	}
	if req.DeletedAt {
		deletedAt := jsoniter.Get(res, "deletedAt").ToString()
		if deletedAt != "null" {
			v.Response.DeletedAt, _ = time.Parse(time.RFC3339, deletedAt)
		}
	}
	if req.Description {
		v.Response.Description = jsoniter.Get(res, "description").ToString()
	}
	if req.Download != nil {
		download, _ := jsoniter.Marshal(jsoniter.Get(res, "download"))
		(v.Request.Download).ResponseParser(download)
		v.Response.Download = v.Request.Download
	}
	if req.Duration {
		durationSeconds := jsoniter.Get(res, "lengthSeconds").ToString() + "s"
		v.Response.Duration, _ = time.ParseDuration(durationSeconds)
	}
	if req.Id {
		v.Response.Id = jsoniter.Get(res, "id").ToInt()
	}
	if req.SoftDeleted {
		v.Response.SoftDeleted = jsoniter.Get(res, "isDeleted").ToBool()
	}
	if req.Language {
		v.Response.Language = jsoniter.Get(res, "language").ToString()
	}
	if req.OffsetSeconds {
		v.Response.OffsetSeconds = jsoniter.Get(res, "offsetSeconds").ToInt()
	}
	if req.PlaybackAccessToken != nil {
		pat, _ := jsoniter.Marshal(jsoniter.Get(res, "playbackAccessToken"))
		(v.Request.PlaybackAccessToken).ResponseParser(pat)
		v.Response.PlaybackAccessToken = v.Request.PlaybackAccessToken
	}
	if req.PublishedAt {
		publishedAt := jsoniter.Get(res, "publishedAt").ToString()
		v.Response.PublishedAt, _ = time.Parse(time.RFC3339, publishedAt)
	}
	if req.RecordedAt {
		recordedAt := jsoniter.Get(res, "recordedAt").ToString()
		v.Response.RecordedAt, _ = time.Parse(time.RFC3339, recordedAt)
	}
	if req.Scope {
		v.Response.Scope = jsoniter.Get(res, "scope").ToString()
	}
	if req.PreviewsUrl {
		v.Response.PreviewsUrl = jsoniter.Get(res, "seekPreviewsURL").ToString()
	}
	if req.Status {
		v.Response.Status = jsoniter.Get(res, "status").ToString()
	}
	if req.Tags {
		jsonTags := jsoniter.Get(res, "tags")
		tags := make([]string, 0)
		for i := 0; i < jsonTags.Size(); i++ {
			tags = append(tags, jsonTags.Get(i).ToString())
		}
		v.Response.Tags = tags
	}
	if req.PreviewThumbnailUrl != (ImageParams{}) {
		v.Response.PreviewThumbnailURL = jsoniter.Get(res, "previewThumbnailURL").ToString()
	}
	if req.ThumbnailUrls != (ImageParams{}) {
		jsonThumbnailUrls := jsoniter.Get(res, "thumbnailURLs")
		thumbnailUrls := make([]string, 0)
		for i := 0; i < jsonThumbnailUrls.Size(); i++ {
			thumbnailUrls = append(thumbnailUrls, jsonThumbnailUrls.Get(i).ToString())
		}
		v.Response.ThumbnailUrls = thumbnailUrls
	}
	if req.Title {
		v.Response.Title = jsoniter.Get(res, "title").ToString()
	}
	if req.UpdatedAt {
		updatedAt := jsoniter.Get(res, "updatedAt").ToString()
		if updatedAt != "null" {
			v.Response.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
		}
	}
	if req.ViewCount {
		v.Response.ViewCount = jsoniter.Get(res, "viewCount").ToInt()
	}
	if req.ViewableAt {
		viewableAt := jsoniter.Get(res, "viewableAt").ToString()
		if viewableAt != "null" {
			v.Response.ViewableAt, _ = time.Parse(time.RFC3339, viewableAt)
		}
	}
}
