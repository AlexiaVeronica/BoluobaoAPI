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
	today := time.Now().Format("2006-01-02")
	request.Put("user/newSignInfo").Add("signDate", today).NewRequests().Unmarshal(&SignIn)
	if SignIn.Status.HTTPCode == 200 {
		fmt.Println("sign in success, date: ", today)
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
			"seconds":     "3605",
			"readingDate": time.Now().Format("2006-01-02"),
			"entityType":  "3",
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
	ReadData := fmt.Sprintf(`{'seconds': 3605, 'readingDate': '%v', 'entityType': 2}`, time.Now().Format("2006-01-02"))
	request.Put("user/readingtime").AddString(ReadData).NewRequests().WriteResultString()
}
