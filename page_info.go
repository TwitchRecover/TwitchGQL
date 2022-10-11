package twitchgql

import jsoniter "github.com/json-iterator/go"

type PageInfo struct {
	Request  PageInfoRequest
	Response PageInfoResponse
}

type PageInfoRequest struct {
	hasNextPage     bool
	hasPreviousPage bool
}

type PageInfoResponse struct {
	hasNextPage     bool
	hasPreviousPage bool
}

func (pi *PageInfo) RequestParser() (string, error) {
	if pi.Request == (PageInfoRequest{}) {
		return "", nil
	}
	query := `pageInfo{`
	if pi.Request.hasNextPage {
		query += `hasNextPage,`
	}
	if pi.Request.hasPreviousPage {
		query += `hasPreviousPage,`
	}
	return query + `},`, nil
}

func (pi *PageInfo) ResponseParser(res []byte) {
	pi.Response.hasNextPage = jsoniter.Get(res, "hasNextPage").ToBool()
	pi.Response.hasPreviousPage = jsoniter.Get(res, "hasPreviousPage").ToBool()
}
