package godoit

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Priority string

const (
	NONE Priority = ""
	A    Priority = "(A)"
	B    Priority = "(B)"
	C    Priority = "(C)"
	D    Priority = "(D)"
	E    Priority = "(E)"
	F    Priority = "(F)"
)

type Task struct {
	Body     string
	Priority Priority
	Project  string
	// Scheduled     time.Time
	Due           time.Time
	Created       time.Time
	CompletedDate time.Time
	Done          bool
	Context       string
}

func (t Task) String() string {
	builder := strings.Builder{}
	if t.Done {
		builder.WriteString("x ")
		builder.WriteString(t.CompletedDate.Format(layoutISO) + " ")
	} else {
		builder.WriteString(string(t.Priority) + " ")
	}
	if !t.Created.IsZero() {
		builder.WriteString(t.Created.Format(layoutISO) + " ")
	}
	builder.WriteString(t.Body)
	if !t.Due.IsZero() {
		builder.WriteString("due:" + t.Due.Format(layoutISO))
	}
	return builder.String()
}

type TaskList struct {
	tasks map[int]*Task
}

func NewTaskList(tasks []*Task) TaskList {
	tl := TaskList{}
	for _, task := range tasks {
		tl.Add(task)
	}
	return tl
}

func (tl *TaskList) Add(task *Task) {
	if tl.tasks == nil {
		tl.tasks = make(map[int]*Task)
	}
	tl.tasks[len(tl.tasks)] = task
}

func (tl *TaskList) Complete(id int) bool {
	task, prs := tl.tasks[id]
	if prs {
		task.Done = true
		task.CompletedDate = time.Now()
		return true
	}
	return false
}

func (tl *TaskList) RemoveTask(id int) bool {
	_, prs := tl.tasks[id]
	if prs {
		delete(tl.tasks, id)
		return true
	}
	return false
}

func (tl *TaskList) SaveFile(path string) error {
	// f, err := os.Open(path)
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	for _, task := range tl.tasks {
		_, err := f.WriteString(task.String() + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file %s: %v", path, err)
		}
	}
	return nil
}
