package worker

import (
	"fmt"
	"jawad-ch/orchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("I will collect stats")
}

func (w *Worker) RunTask() {
	fmt.Println("I will run tasks")
}

func (w *Worker) StartTask() {
	fmt.Println("I will start task")
}

func (w *Worker) StopTask() {
	fmt.Println("I will stop task")
}
