package tasks

import (
	"fmt"
	"strconv"
	"sync"
)

var taskList map[string]TaskModel = make(map[string]TaskModel)

// RegisterTasks 注册任务
func RegisterTasks(name string, task TaskModel) {
	var _, ok = taskList[name]
	if ok {
		fmt.Errorf("Task " + name + " Exists!")
		return
	}
	taskList[name] = task
}

// Run 运行任务
func Run() {
	var wg sync.WaitGroup //定义一个同步等待的组
	for name, task := range taskList {
		if task.Concurrence == 0 {
			task.Concurrence = task.GoRoutinesNum
		}
		var groups = task.GoRoutinesNum/task.Concurrence + 1
		var i, j = 0, 0
		for j = 0; j < groups; j++ {
			for i = j * task.Concurrence; i < (j+1)*task.Concurrence; i++ {
				wg.Add(1)
				var temp = i
				go func() {
					var succ = task.Executor(temp)
					if !succ {
						fmt.Println("task " + name + " " + strconv.Itoa(temp) + " failed!")
					}
					wg.Done()
				}()
			}
			wg.Wait()
		}
	}
}
