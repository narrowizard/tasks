package tasks

type TaskModel struct {
	// GoRoutinesNum 子任务数量
	GoRoutinesNum int
	// Concurrence 允许同时运行的子任务数量
	Concurrence int
	// Executor 任务执行器
	// index 子任务编号
	Executor func(index int) bool
}
