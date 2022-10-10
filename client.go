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
	gqlHost     = "gql.twitch.tv"
	gqlPath     = "gql"
	gqlClientId = "kimne78kx3ncx6brgo4mv6wki5h1ko"
	queryPre    = `{"query":"query{`
	querySuf    = `}"}`
)

var (
	gqlUrl = url.URL{
		Scheme: "https",
		Host:   gqlHost,
		Path:   gqlPath,
	}
	headers = map[string][]string{
		"Client-ID": {gqlClientId},
	}
	defReq = http.Request{
		URL:    &gqlUrl,
		Header: headers,
		Method: "POST",
	}
)

type Client struct {
	clientId    string
	http_client http.Client
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
	if client.clientId != "" {
		req.Header.Set("Client-ID", client.clientId)
	}
	parsedReq, err := t.RequestParser()
	if err != nil {
		return nil, err
	}
	query := []byte(queryPre + parsedReq + querySuf)
	res, error := Request(client.http_client, req, query)
	if error != nil {
		return nil, error
	}
	data, _ := jsoniter.Marshal(jsoniter.Get(res, "data"))
	t.ResponseParser(data)
	return t, nil
}
