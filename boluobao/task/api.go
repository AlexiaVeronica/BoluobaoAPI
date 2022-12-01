package task

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/config"
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

func (task *Task) PUT_SHARE(account_id string) {
	// This interaction logic is so bad, I will refactor it, but now I will leave it like this.
	// I will refactor it when I have time.
	var change bool
	if config.AppConfig.App {
		change = true
		config.AppConfig.App = false
	} else {
		change = false
	}
	request.Put("user/tasks?taskId=4&userId=" + account_id).AddString(`{"env": 0}`).NewRequests()

	if change {
		config.AppConfig.App = true // change back to app
	}
}
