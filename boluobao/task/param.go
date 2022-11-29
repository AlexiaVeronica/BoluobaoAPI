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

func (task *Task) ReadDate() string {
	return fmt.Sprintf(`{'seconds': 3605, 'readingDate': '%v', 'entityType': 2}`, task.GetDay())
}
func (task *Task) ListenData() string {
	return fmt.Sprintf(`{'seconds': 3605, 'readingDate': '%v', 'entityType': 3}`, task.GetDay())
}

func (task *Task) GetDay() string {
	return time.Now().Format("2006-01-02")
}
