package twitchgql

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/json-iterator/go"
)

const (
	GqlHost     = "gql.twitch.tv"
	GqlPath     = "gql"
	GqlClientId = "kimne78kx3ncx6brgo4mv6wki5h1ko"
	queryPre    = `{"query":"query{`
	querySuf    = `}"}`
)

var (
	GqlUrl = url.URL{
		Scheme: "https",
		Host:   GqlHost,
		Path:   GqlPath,
	}
	headers = map[string][]string{
		"Client-ID": {GqlClientId},
	}
	defReq = http.Request{
		URL:    &GqlUrl,
		Header: headers,
		Method: "POST",
	}
)

type Client struct {
	ClientId string
	Client   http.Client
}

type Type interface {
	RequestParser() (string, error)
	ResponseParser([]byte)
}

func Request(client http.Client, req http.Request, cont []byte) ([]byte, error) {
	req.Body = ioutil.NopCloser(bytes.NewBuffer(cont))
	res, err := client.Do(&req)
	if err != nil || res.StatusCode != 200 {
		if err == nil {
		}
		return nil, err
	}
	defer res.Body.Close()
	response, error := io.ReadAll(res.Body)
	if error != nil {
		return nil, error
	}
	return response, nil
}

func Query(client Client, t Type) (Type, error) {
	req := defReq
	if client.ClientId != "" {
		req.Header.Set("Client-ID", client.ClientId)
	}
	parsedReq, err := t.RequestParser()
	if err != nil {
		return nil, err
	}
	query := []byte(queryPre + parsedReq + querySuf)
	res, error := Request(client.Client, req, query)
	if error != nil {
		return nil, error
	}
	data, _ := jsoniter.Marshal(jsoniter.Get(res, "data"))
	t.ResponseParser(data)
	return t, nil
}
