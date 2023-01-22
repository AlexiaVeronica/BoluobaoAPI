package task

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/api"
	"github.com/google/uuid"
)

func (task *Task) PUT_LISTEN_TIME() {
	api.Put("user/readingtime").AddString(task.ListenData()).NewRequests()
}

func (task *Task) PUT_READING_TIME() {
	api.Put("user/readingtime").AddString(task.ReadDate()).NewRequests()
}

func (task *Task) POST_RECEIVE_TASK(TaskId string) MessageTask {
	var Message MessageTask
	params := map[string]string{"seconds": "3605", "readingDate": task.GetDay(), "entityType": "3"}
	api.Post("user/tasks/" + TaskId).AddAll(params).NewRequests().Unmarshal(&Message)
	return Message
}

func (task *Task) PUT_SIGN_INFO() Template.SignIn {
	var SignIn Template.SignIn
	// "user/signInfo" is old api, now it's "user/newSignInfo"
	FormData := fmt.Sprintf(`{'signDate': '%v'}`, task.GetDay())
	api.Put("user/newSignInfo").AddString(FormData).NewRequests().Unmarshal(&SignIn)
	return SignIn
}

func (task *Task) GET_TASKS_LIST() Template.Task {
	var TaskStruct = Template.Task{}
	api.Get("user/tasks").AddAll(map[string]string{
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
	if api.Request.App {
		change = true
		api.Request.App = false
	} else {
		change = false
	}
	api.Put("user/tasks?taskId=4&userId=" + account_id).AddString(`{"env": 0}`).NewRequests()

	if change {
		api.Request.App = true // change back to app
	}
}
