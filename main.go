package main

import (
	"github.com/oxygen-org/worker/utils"
	"fmt"
	"os"
	"os/user"
	"github.com/oxygen-org/worker/config"
	"time"

	// "github.com/fsouza/go-dockerclient"
	"github.com/oxygen-org/worker/task"
)

func main() {
	// c:= config.C
	// uPrintlnRedis(c.Redis.Addr, c.Redis.DB, c.Redis.PW)
	// task.CloneCode()
	fmt.Println(user.Current())
	fmt.Println(os.Getpid())
	fmt.Println(config.C.General)
	fmt.Println(config.C.General)
	RegisterMe()
	BeatPing()
	for{
		task.GetJob()
		go func(){
			
		}()
		time.Sleep(1)
	}
	


	time.Sleep(time.Second * 1000000000)

	// fmt.Println(task.TSourceCode{GitURL:"http://"})

}
