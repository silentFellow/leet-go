package leetcode

import "container/heap"

type Task struct {
	userId   int
	taskId   int
	priority int
	index    int
}

type TaskHeap []*Task

func (this TaskHeap) Len() int { return len(this) }

func (this TaskHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this TaskHeap) Less(i, j int) bool {
	if this[i].priority == this[j].priority {
		return this[i].taskId > this[j].taskId
	}
	return this[i].priority > this[j].priority
}

func (this *TaskHeap) Push(v any) {
	task := v.(*Task)
	task.index = len(*this)
	*this = append(*this, task)
}

func (this *TaskHeap) Pop() any {
	old := *this
	n := len(old)
	task := old[n-1]
	task.index = -1
	*this = old[:n-1]
	return task
}

type TaskManager struct {
	taskMap  map[int]*Task
	taskHeap *TaskHeap
}

func Constructor(tasks [][]int) TaskManager {
	h := &TaskHeap{}
	heap.Init(h)
	taskMap := make(map[int]*Task)

	for _, task := range tasks {
		userId, taskId, priority := task[0], task[1], task[2]
		task := &Task{userId: userId, taskId: taskId, priority: priority}
		heap.Push(h, task)
		taskMap[taskId] = task
	}

	return TaskManager{
		taskHeap: h,
		taskMap:  taskMap,
	}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	task := &Task{
		userId:   userId,
		taskId:   taskId,
		priority: priority,
	}
	heap.Push(this.taskHeap, task)
	this.taskMap[taskId] = task
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	task := this.taskMap[taskId]
	task.priority = newPriority
	heap.Fix(this.taskHeap, task.index)
}

func (this *TaskManager) Rmv(taskId int) {
	task := this.taskMap[taskId]
	heap.Remove(this.taskHeap, task.index)
	delete(this.taskMap, task.taskId)
}

func (this *TaskManager) ExecTop() int {
	if this.taskHeap.Len() == 0 {
		return -1
	}

	task := heap.Pop(this.taskHeap).(*Task)
	delete(this.taskMap, task.taskId)
	return task.userId
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
