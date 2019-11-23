package task

import (
	"bufio"
	"os"
	"github.com/oxygen-org/worker/utils"

	docker "github.com/fsouza/go-dockerclient"
)

type PullImage struct{

}

func (task *PullImage) Do(imgName string, tag string) *PullImage {

	readP, writeP, _ := os.Pipe()
	if tag == "" {
		tag = "latest"
	}
	opts := docker.PullImageOptions{
		Repository:   imgName,
		OutputStream: writeP,
		Tag:          tag,
	}
	auth := docker.AuthConfiguration{}
	syn := make(chan int) //sync channel
	go func(syn chan<- int) {
		scanner := bufio.NewScanner(readP)
		for scanner.Scan() {
			
		}
		syn <- 0
	}(syn)
	utils.DockerC.PullImage(opts, auth)
	writeP.Close()
	<-syn
	return task
}

func CreateContainer() {

}

func StartContainer() {

}

func BuildImage() {

}
