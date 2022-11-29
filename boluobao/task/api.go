package task

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"github.com/google/uuid"
)

func (task *Task) PUT_LISTEN_TIME() {
	request.Put("user/readingtime").AddString(task.ListenData()).NewRequests()
}

func (task *Task) PUT_READING_TIME() {
	request.Put("user/readingtime").AddString(task.ReadDate()).NewRequests()
}

func (task *Task) POST_RECEIVE_TASK(TaskId string) MessageTask {
	var Message MessageTask
	params := map[string]string{"seconds": "3605", "readingDate": task.GetDay(), "entityType": "3"}
	request.Post("user/tasks/" + TaskId).AddAll(params).NewRequests().Unmarshal(&Message)
	return Message
}

func (task *Task) PUT_SIGN_INFO() Template.SignIn {
	var SignIn Template.SignIn
	// "user/signInfo" is old api, now it's "user/newSignInfo"
	FormData := fmt.Sprintf(`{'signDate': '%v'}`, task.GetDay())
	request.Put("user/newSignInfo").AddString(FormData).NewRequests().Unmarshal(&SignIn)
	return SignIn
}

func (task *Task) GET_TASKS_LIST() Template.Task {
	var TaskStruct = Template.Task{}
	request.Get("user/tasks").AddAll(map[string]string{
		"taskCategory": "1",
		"package":      "com.sfacg",
		"deviceToken":  uuid.New().String(),
		"page":         "0",
		"size":         "20",
	}).NewRequests().Unmarshal(&TaskStruct)
	return TaskStruct
}
