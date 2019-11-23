package config

import (
	"bytes"
	"errors"
	"net"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"

	"github.com/koding/multiconfig"
)

var C = new(github.com/oxygen-orgWorker)

func init() {
	multiconfig.MustLoadWithPath("config.json", C)
	C.General.PID = os.Getpid()

	if C.General.HostName == "" {
		C.General.HostName, _ = os.Hostname()
	}
	if C.General.HostIP == "" {
		C.General.HostIP, _ = externalIP()
	}

	if strings.HasPrefix(C.General.GitKeyPath, "~/") {
		home, _ := getHome()
		C.General.GitKeyPath = path.Join(home, C.General.GitKeyPath[1:])
	}

}

type GeneralConfig struct {
	Debug      bool   `defasult:"true"`
	HostName   string `required:"false"`
	HostIP     string `required:"false"`
	PID        int    `required:"false"`
	GitKeyPath string `required:"true"`
}

type RedisConfig struct {
	Addr string `default:"localhost:6379"`
	DB   int    `default:"0"`
	PW   string `default:""`
}

type DockerConfig struct{
	Endpoint string `required:"false"`
	CA string `required:"false"`
	Cert string `required:"false"`
	Key string `required:"false"`
}

type github.com/oxygen-orgWorker struct {
	General GeneralConfig
	Redis   RedisConfig
	Docker DockerConfig
}

func getHome() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("Blank output when reading home directory")
	}
	return result, nil

}

// Get external IP of local host
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // 不是IPV4地址
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("x")
}
