package twitchgql

import (
	"net/http"
)

type Client struct {
	token       string
	http_client http.Client
}

type Type interface {
	Query(Client) interface{}
	Mutation(interface{}, Client) interface{}
}
