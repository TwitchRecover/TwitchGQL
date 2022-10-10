package twitchgql

import (
	"strconv"

	"github.com/json-iterator/go"
)

type PlaybackAccessToken struct {
	Request  PlaybackAccessTokenRequest
	Response PlaybackAccessTokenResponse
}

type PlaybackAccessTokenRequest struct {
	Params    PlaybackAccessTokenRequestParams
	Signature bool
	Value     bool
}

type PlaybackAccessTokenRequestParams struct {
	Platform      string
	PlayerType    string
	PlayerBackend string
	HasAdblock    bool
	DisableHTTPs  bool
}

type PlaybackAccessTokenResponse struct {
	Signature string
	Value     string
}

func (pat *PlaybackAccessToken) RequestParser() (string, error) {
	if pat.Request == (PlaybackAccessTokenRequest{}) {
		return "", nil
	}
	query := `playbackAccessToken(params:{`
	query += `platform:\"` + pat.Request.Params.Platform + `\",`
	query += `playerType:\"` + pat.Request.Params.PlayerType + `\",`
	query += `playerBackend:\"` + pat.Request.Params.PlayerBackend + `\",`
	query += `hasAdblock:` + strconv.FormatBool(pat.Request.Params.HasAdblock) + `,`
	query += `disableHTTPS:` + strconv.FormatBool(pat.Request.Params.DisableHTTPs) + `,`
	query += `}){`
	if pat.Request.Signature {
		query += `signature,`
	}
	if pat.Request.Value {
		query += `value,`
	}
	return query + `},`, nil
}

func (pat *PlaybackAccessToken) ResponseParser(res []byte) {
	if pat.Request.Signature {
		pat.Response.Signature = jsoniter.Get(res, "signature").ToString()
	}
	if pat.Request.Value {
		pat.Response.Value = jsoniter.Get(res, "value").ToString()
	}
}
