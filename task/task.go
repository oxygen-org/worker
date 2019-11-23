package task

import (
	"fmt"
)

//Triger 任务触发器接口
type Triger interface {
	String()
}

//TImage 镜像任务触发器
type TImage struct {
	Name string
	Tag  string
}

func (img TImage) String() string {
	return fmt.Sprintf("ImageTriger->\n\tName:%s\n\tTag:%s\n",
		img.Name, img.Tag)
}

//TSourceCode 源码获取相关的触发器
type TSourceCode struct {
	GitURL    string
	Branch    string
	Commit    string
	TargetDir string
}

func (ts TSourceCode) String() string {
	return fmt.Sprintf("SourceCodeTriger->\n\tGitURL:%s\n\tBranch:%s\n\tCommit:%s\n\tTargetDir:%s\n",
		ts.GitURL, ts.Branch, ts.Commit, ts.TargetDir)
}

//TData 数据获取触发器
type TData struct {
	URL  string
	Type string
}

// ITask 定义任务接口
type ITask interface {
	Do(<-chan Triger, chan<- Triger) ITask
	LogEvent(end <-chan int)
	LogOutput(msg string)
}

// LogEvent 记录事件
func LogEvent(t string, msg string) {
	fmt.Println(t, msg)
}

// LogOutput 记录运行输出
func LogOutput(msg string) {
	fmt.Println(msg)
}

// PipeInterface 任务流水线接口
type PipeInterface interface {
	Execute()
	AddTask(ITask, <-chan Triger, chan<- Triger)
}

// Pipe 基础管线
type Pipe struct {
	TaskList      []ITask
	InTrigerList  []<-chan Triger
	OutTrigerList []chan<- Triger
}

//AddTask 向管线添加任务(组装流水线)
func (pipe *Pipe) AddTask(task ITask, inTriger <-chan Triger,
	outTriger chan<- Triger) {
	pipe.TaskList = append(pipe.TaskList, task)
	pipe.InTrigerList = append(pipe.InTrigerList, inTriger)
	pipe.OutTrigerList = append(pipe.OutTrigerList, outTriger)
}

//Execute 执行流水线任务
func (pipe *Pipe) Execute() {
	for index, task := range pipe.TaskList {
		inTriger := pipe.InTrigerList[index]
		outTriger := pipe.OutTrigerList[index]
		task.Do(inTriger, outTriger)
	}
}
