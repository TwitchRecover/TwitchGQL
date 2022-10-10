package twitchgql

import (
	"strconv"

	"github.com/json-iterator/go"
)

type PlaybackAccessToken struct {
	request  PlaybackAccessTokenRequest
	response PlaybackAccessTokenResponse
}

type PlaybackAccessTokenRequest struct {
	params    PlaybackAccessTokenRequestParams
	signature bool
	value     bool
}

type PlaybackAccessTokenRequestParams struct {
	platform      string
	playerType    string
	playerBackend string
	hasAdblock    bool
	disableHTTPs  bool
}

type PlaybackAccessTokenResponse struct {
	signature string
	value     string
}

func (pat *PlaybackAccessToken) RequestParser() (string, error) {
	if pat.request == (PlaybackAccessTokenRequest{}) {
		return "", nil
	}
	query := `playbackAccessToken(params:{`
	query += `platform:"` + pat.request.params.platform + `",`
	query += `playerType:"` + pat.request.params.playerType + `",`
	query += `playerBackend:"` + pat.request.params.playerBackend + `",`
	query += `hasAdblock:` + strconv.FormatBool(pat.request.params.hasAdblock) + `,`
	query += `disableHTTPs:` + strconv.FormatBool(pat.request.params.disableHTTPs) + `,`
	query += `}){`
	if pat.request.signature {
		query += `signature,`
	}
	if pat.request.value {
		query += `value,`
	}
	return query + `},`, nil
}

func (pat *PlaybackAccessToken) ResponseParser(res []byte) {
	if pat.request.signature {
		pat.response.signature = jsoniter.Get(res, "signature").ToString()
	}
	if pat.request.value {
		pat.response.value = jsoniter.Get(res, "value").ToString()
	}
}
