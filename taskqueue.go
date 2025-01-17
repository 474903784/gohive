package gohive

import (
	"github.com/pkg/errors"
	"sync"
)

//DefaultQueueSize indicates the default size of the TaskQueue
const DefaultQueueSize = 10

//TaskQueue represents a queue that holds tasks which
//are in waiting for workers from the pool
type TaskQueue struct {
	//queue that holds tasks
	que []Task

	//number of tasks that currently resides in the queue
	totalTasks int

	//Mutex used for atomic operations
	locker sync.Mutex
}

//NewTaskQueue returns new TaskQueue with the default capacity
func NewTaskQueue() *TaskQueue {
	wtQue := TaskQueue{que: make([]Task, 0, DefaultQueueSize), totalTasks: 0}
	return &wtQue
}

//EnqueueTask puts a new task in the TaskQueue
func (wq *TaskQueue) EnqueueTask(task Task) {
	wq.locker.Lock()
	defer wq.locker.Unlock()
	wq.que = append(wq.que, task)
	wq.totalTasks++
}

//DequeueTask returns a task and removes it from the TaskQueue
func (wq *TaskQueue) DequeueTask() (Task, error) {
	wq.locker.Lock()
	defer wq.locker.Unlock()
	if wq.totalTasks > 0 {
		task := wq.que[0]
		wq.que = append(wq.que[:0], wq.que[1:]...)
		wq.totalTasks--
		return task, nil
	}
	return Task{}, errors.New("Queue is Empty")
}

//IsNotEmpty returns whether the TaskQueue is empty or not
func (wq *TaskQueue) IsNotEmpty() bool {
	wq.locker.Lock()
	defer wq.locker.Unlock()
	return wq.totalTasks > 0
}
