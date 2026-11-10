package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

const (
	Version  = "1.0.0"
	Priority = 1000
)

type Config struct {
	PartnerValue string `json:"partner_value"`
}

func New() interface{} {
	return &Config{}
}

func (conf *Config) Access(kong *pdk.PDK) {
	if conf.PartnerValue == "" {
		exit(kong, 500, "Invalid plugin configuration")
		return
	}

	if err := kong.ServiceRequest.SetHeader("Partner", conf.PartnerValue); err != nil {
		exit(kong, 500, "Internal server error")
	}
}

func exit(kong *pdk.PDK, status int, body string) {
	kong.Response.Exit(status, []byte(body), map[string][]string{
		"Content-Type": {"text/plain; charset=utf-8"},
	})
}

func main() {
	server.StartServer(New, Version, Priority)
}
