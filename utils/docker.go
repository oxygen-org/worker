package utils

import (
	"github.com/oxygen-org/worker/config"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
)

var DockerC *docker.Client

func init() {
	c := config.C.Docker
	var (
		client *docker.Client
		err    error
	)
	switch {
	case strings.HasPrefix(c.Endpoint, "tcp://"):
		client, err = docker.NewTLSClient(c.Endpoint, c.Cert, c.Key, c.CA)
	case c.Endpoint == "":
		client, err = docker.NewClientFromEnv()
	case strings.HasPrefix(c.Endpoint, "unix://"):
		client, err = docker.NewClient(c.Endpoint)
	default:
		panic("没有有效的docker remote链接")
	}
	if err != nil {
		panic(err)
	}
	DockerC = client

}
