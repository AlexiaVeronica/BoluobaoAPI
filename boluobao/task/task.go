package task

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"time"
)

type Task struct {
}

func (task *Task) GET_TASK() {
	request.Get("user/signInfo").NewRequests().WriteResultString()

}
func (task *Task) GET_TASK_LIST() {

}

func (task *Task) PUT_READING_TIME() {
	readingDate := time.Now().Format("2006-01-02")
	params := map[string]string{
		"seconds":     "3605",
		"readingDate": readingDate,
		"entityType":  "3",
	}
	request.Put("user/readingtime").AddAll(params).NewRequests().WriteResultString()
}
