package _3_动态规划

//题目参考: https://www.bilibili.com/video/av16544031
//抽象时间安排
type TaskEntry struct {
	taskId int //任务ID
	startTime int //开始时间
	endTime int//结束时间
	value int //价值
}
//任务之前可以执行的任务的抽象结构
type prevEntry struct {
	taskId int //自身的任务ID
	prevTaskId int//时间不冲突的任务ID,取最近的
}
//任务列表
var TaskList []*TaskEntry
var PrevList []*prevEntry

//初始数据
func init() {
	TaskList = getTasks()
	PrevList = GetPrev()
}
func getTasks() []*TaskEntry {
	TaskList = make([]*TaskEntry, 8)
	TaskList[0] = &TaskEntry{1,1, 4, 5}
	TaskList[1] = &TaskEntry{2,3, 5, 1}
	TaskList[2] = &TaskEntry{3,0, 6, 8}
	TaskList[3] = &TaskEntry{4, 4, 7, 4}
	TaskList[4] = &TaskEntry{5, 3, 8, 6}
	TaskList[5] = &TaskEntry{6, 5, 9, 3}
	TaskList[6] = &TaskEntry{7, 6, 10, 2}
	TaskList[7] = &TaskEntry{8, 8, 11, 4}
	return TaskList
}
//对任务的选择有两种, 选与不选.如何选则本身+前面prev可以做的任务. 如果不选则下一个任务
func GetPrev() []*prevEntry {
	size := len(TaskList)
	prev := make([]*prevEntry, size)
	for i := 0; i < size; i ++ {
		prev[i] = &prevEntry{taskId:TaskList[i].taskId}
		startTime := TaskList[i].startTime
		for j := i-1; j >= 0; j -- {
			if startTime >= TaskList[j].endTime {
				prev[i] = &prevEntry{TaskList[i].taskId, TaskList[j].taskId}
				break
			}
		}
	}
	return prev
}
//求最优解OPT
func OPT(taskId int) int {
	if taskId == 0 {
		return 0
	}
	return max(OPT(taskId-1), TaskList[taskId-1].value + OPT(PrevList[taskId].taskId))
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}



