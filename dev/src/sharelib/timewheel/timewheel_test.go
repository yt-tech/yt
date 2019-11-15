package timewheel

import (
	"fmt"
	"time"
)

func main() {

	//初始化时间轮盘
	//参数：interval 时间间隔
	//参数：slotNum  轮盘大小
	tw := New(time.Second, 160)

	tw.Start()

	key := "task1"
	//添加定时任务
	//参数：interval 时间间隔
	//参数：times 执行次数 -1 表示周期任务 >0 执行指定次数
	//参数：key 任务唯一标识符 用户更新任务和删除任务
	//参数：taskData 回调函数参数
	//参数：job 回调函数
	tw.AddTask(time.Second, -1, key, TaskData{"name": "john"},
		func(params TaskData) {
			fmt.Println(time.Now().Unix(), params["name"])
		})

	//更新任务参数
	time.Sleep(time.Second * 10)
	tw.UpdateTask(key, time.Second*3, TaskData{"name": "terry"})

	//删除定时任务
	time.Sleep(time.Second * 10)
	tw.RemoveTask(key)

	//轮盘停止
	tw.Stop()
}
