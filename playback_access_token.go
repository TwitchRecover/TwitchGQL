package twitchgql

type PlaybackAccessToken struct {
	request  PlaybackAccessTokenRequest
	response PlaybackAccessTokenResponse
}

type PlaybackAccessTokenRequest struct {
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
