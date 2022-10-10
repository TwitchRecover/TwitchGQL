package twitchgql

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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
	RequestParser(interface{}) string
	ResponseParser(string) string
}

func Request(client http.Client, req http.Request, cont []byte) (string, error) {
	req.Body = ioutil.NopCloser(bytes.NewBuffer(cont))
	res, err := client.Do(&req)
	if err != nil || res.StatusCode != 200 {
		if err == nil {
		}
		return "", err
	}
	defer res.Body.Close()
	response, error := io.ReadAll(res.Body)
	if error != nil {
		return "", error
	}
	return string(response), nil
}

func Query(client Client, t Type, reqInterface interface{}) (interface{}, error) {
	req := defReq
	if client.clientId != "" {
		req.Header.Set("Client-ID", client.clientId)
	}
	query := []byte(queryPre + t.RequestParser(reqInterface) + querySuf)
	res, err := Request(client.http_client, req, query)
	if err != nil {
		return nil, err
	}
	return t.ResponseParser(res), nil
}
