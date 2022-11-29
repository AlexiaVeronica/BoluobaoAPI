package task

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"time"
)

type Task struct{}

func (task *Task) GET_SIGN_INFO() {
	var Status Template.Status
	request.Get("user/signInfo").NewRequests().Unmarshal(&Status)
	if Status.HTTPCode == 200 {
		fmt.Println("签到成功")
	} else {
		fmt.Println("签到失败", Status.Msg)
	}

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

func (task *Task) COMPLETE_TASK_LIST() {
	ListenData := map[string]string{
		"seconds":     "3605",
		"readingDate": time.Now().Format("2006-01-02"),
		"entityType":  "3",
	}
	request.Put("user/tasks/4").AddAll(ListenData).NewRequests().WriteResultString()
	request.Put("user/tasks/5").AddAll(ListenData).NewRequests().WriteResultString()
	request.Put("user/tasks/17").AddAll(ListenData).NewRequests().WriteResultString()
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
