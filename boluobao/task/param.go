package task

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"time"
)

var ReadTaskList = []string{"user/tasks/4", "user/tasks/5", "user/tasks/17"}

type Task struct {
	AccountId string
}

type MessageTask struct {
	Status Template.Status `json:"status"`
}

func (task *Task) Data(entityType int) string {
	return fmt.Sprintf(`{'seconds': 3605, 'readingDate': '%v', 'entityType': %v}`, task.GetDay(), entityType)
}
func (task *Task) ReadDate() string {
	return task.Data(2)
}
func (task *Task) ListenData() string {
	return task.Data(3)
}

func (task *Task) GetDay() string {
	return time.Now().Format("2006-01-02")
}
