package tasks

type TaskModel struct {
	// GoRoutinesNum 协程数量
	GoRoutinesNum int
	// Executor 任务执行器
	// index 子任务编号
	Executor func(index int) bool
}
