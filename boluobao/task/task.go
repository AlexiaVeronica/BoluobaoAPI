package task

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type Task struct {
	TaskList Template.Task
}

func (task *Task) GET_SIGN_INFO() {
	var SignIn Template.SignIn
	// "user/signInfo" is old api, now it's "user/newSignInfo"
	Data := fmt.Sprintf(`{'signDate': '%v'}`, task.GetDay())
	request.Put("user/newSignInfo").AddString(Data).NewRequests().Unmarshal(&SignIn)
	if SignIn.Status.HTTPCode == 200 {
		fmt.Println("sign in success, date: ", task.GetDay())
	} else {
		fmt.Println("签到失败,MESSAGE:", SignIn.Status.Msg)
	}

}

func (task *Task) GET_TASKS_LIST() {
	var TaskStruct = Template.Task{}
	request.Get("user/tasks").AddAll(map[string]string{
		"taskCategory": "1",
		"package":      "com.sfacg",
		"deviceToken":  uuid.New().String(),
		"page":         "0",
		"size":         "20",
	}).NewRequests().Unmarshal(&TaskStruct).WriteResultString()
	if TaskStruct.Status.HTTPCode == 200 {
		task.TaskList = TaskStruct
	} else {
		fmt.Println("获取任务列表失败,MESSAGE:", TaskStruct.Status.Msg)
	}
}

func (task *Task) RECEIVE_TASK() {
	task.GET_TASKS_LIST() // init task list
	for _, data := range task.TaskList.Data {
		var TaskInfo = struct {
			Status Template.Status `json:"status"`
		}{}
		request.Post("user/tasks/" + strconv.Itoa(data.TaskId)).AddAll(map[string]string{
			"seconds": "3605", "readingDate": task.GetDay(), "entityType": "3",
		}).NewRequests().Unmarshal(&TaskInfo)
		if TaskInfo.Status.ErrorCode == 200 {
			fmt.Println(data.Name, "领取成功")
		} else if TaskInfo.Status.ErrorCode == 798 {
			fmt.Println(data.Name, "已经领取过了")
		} else {
			fmt.Println(data.Name, "message:", TaskInfo.Status.Msg)
		}
	}
}
func (task *Task) COMPLETE_TASK_LIST() {
	for i, url := range []string{"user/tasks/4", "user/tasks/5", "user/tasks/17"} {
		fmt.Println("正在完成第", i+1, "个任务")
		request.Put(url).AddString(task.ListenData()).NewRequests()
	}

}

func (task *Task) PUT_LISTEN_TIME() {
	request.Put("user/readingtime").AddString(task.ListenData()).NewRequests()
}

func (task *Task) PUT_READING_TIME() {
	request.Put("user/readingtime").AddString(task.ReadingDate()).NewRequests().WriteResultString()
}

func (task *Task) ReadingDate() string {
	return fmt.Sprintf(`{'seconds': 3605, 'readingDate': '%v', 'entityType': 2}`, task.GetDay())
}
func (task *Task) ListenData() string {
	return fmt.Sprintf(`{'seconds': 3605, 'readingDate': '%v', 'entityType': 3}`, task.GetDay())
}

func (task *Task) GetDay() string {
	return time.Now().Format("2006-01-02")
}
