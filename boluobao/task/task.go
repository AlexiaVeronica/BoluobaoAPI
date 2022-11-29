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
func (task *Task) POST_TASK_LIST() {
	ListenData := map[string]string{
		"seconds":     "3605",
		"readingDate": time.Now().Format("2006-01-02"),
		"entityType":  "3",
	}
	request.Post("user/tasks/4").AddAll(ListenData).NewRequests().WriteResultString()
	request.Post("user/tasks/5").AddAll(ListenData).NewRequests().WriteResultString()
	request.Post("user/tasks/17").AddAll(ListenData).NewRequests().WriteResultString()
}

func (task *Task) PUT_LISTEN_TIME() {
	ListenData := map[string]string{
		"seconds":     "3605",
		"readingDate": time.Now().Format("2006-01-02"),
		"entityType":  "3",
	}
	request.Put("user/readingtime").AddAll(ListenData).NewRequests().WriteResultString()
}

func (task *Task) PUT_READING_TIME() {
	ReadData := map[string]string{
		"seconds":     "3605",
		"readingDate": time.Now().Format("2006-01-02"),
		"entityType":  "2",
	}
	request.Put("user/readingtime").AddAll(ReadData).NewRequests().WriteResultString()
}
