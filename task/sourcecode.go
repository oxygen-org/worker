package task

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/oxygen-org/worker/config"

	"golang.org/x/crypto/ssh"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	go_git_ssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

type CloneCode struct{

}

func (task *CloneCode) Do() *CloneCode{
	keyPath := config.C.General.GitKeyPath
	fmt.Println(keyPath)
	auth := get_ssh_key_auth(keyPath)
	fmt.Println(auth)
	_, err := git.PlainClone("./xxxx", false, &git.CloneOptions{
		URL:           "git@git.coding.net:taceywong163com/go-pusher.git",
		Progress:      os.Stdout,
		Auth:          auth,
	})
	fmt.Println(err)
	return task
}

func get_ssh_key_auth(privateSshKeyFile string) transport.AuthMethod {
	var auth transport.AuthMethod
	sshKey, _ := ioutil.ReadFile(privateSshKeyFile)
	signer, _ := ssh.ParsePrivateKey([]byte(sshKey))
	fmt.Println(sshKey)
	auth = &go_git_ssh.PublicKeys{User: "git", Signer: signer}
	return auth
}
