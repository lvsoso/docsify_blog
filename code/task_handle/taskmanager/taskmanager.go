package taskmanager

import "github.com/hibiken/asynq"

type TaskManager struct {
	inspector asynq.Inspector
}

func NewTaskManager(redisAddr string) *TaskManager {
	inspector := asynq.NewInspector(asynq.RedisClientOpt{
		Addr: redisAddr,
	})
	return &TaskManager{
		inspector: *inspector,
	}
}

func (tm *TaskManager) CancelProcessing(taskId string) error {
	return tm.inspector.CancelProcessing(taskId)
}

func (tm *TaskManager) DeleteTask(queue, taskId string) error {
	return tm.inspector.DeleteTask(queue, taskId)
}

func (tm *TaskManager) GetTaskInfo(queue, taskId string) (*asynq.TaskInfo, error) {
	return tm.inspector.GetTaskInfo(queue, taskId)
}
