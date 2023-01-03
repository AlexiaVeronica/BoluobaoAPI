package task

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"strconv"
)

func (task *Task) NovelCompleteTas() {
	task.NovelCompleteSignInfo()
	TaskList := task.GET_TASKS_LIST()
	if TaskList.Status.HTTPCode == 200 {
		task.NovelCompleteReceiveTaskAll(TaskList) // receive all task
		task.NovelCompleteTaskList()               // complete all task
	} else {
		fmt.Println("获取任务列表失败,MESSAGE:", TaskList.Status.Msg)
	}
}

func (task *Task) NovelCompleteSignInfo() {
	SignIn := task.PUT_SIGN_INFO()
	if SignIn.Status.HTTPCode == 200 {
		fmt.Println("sign in success, date: ", task.GetDay())
	} else {
		fmt.Println("签到失败:", SignIn.Status.Msg)
	}

}

func (task *Task) NovelCompleteReceiveTaskAll(TaskList Template.Task) {
	for _, data := range TaskList.Data {
		TaskInfo := task.POST_RECEIVE_TASK(strconv.Itoa(data.TaskId))
		if TaskInfo.Status.ErrorCode == 200 {
			fmt.Println(data.Name, "领取成功")
		} else if TaskInfo.Status.ErrorCode == 798 {
			fmt.Println(data.Name, "已经领取过了")
		} else {
			fmt.Println(data.Name, "message:", TaskInfo.Status.Msg)
		}
	}
}

func (task *Task) NovelCompleteTaskList() {
	for i, url := range ReadTaskList {
		fmt.Println("complete task No:", i+1)
		if url == "user/tasks/4" && task.AccountId != "" {
			task.PUT_SHARE(task.AccountId)
		} else {
			task.PUT_LISTEN_TIME()
			task.PUT_READING_TIME()
		}
		request.Put(url).AddString(task.ListenData()).NewRequests()
	}

}
