package godoit

import (
	"strings"
	"time"
)

type Priority string

const (
	NONE Priority = ""
	A             = "(A)"
	B             = "(B)"
	C             = "(C)"
	D             = "(D)"
	E             = "(E)"
	F             = "(F)"
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
	// TODO decide if we need completion andtherefore creation dates
}

// func NewTask(body string, scheduled, due string, project string, done bool, context string) (*Task, error) {
// 	var task Task
// 	if len(body) <= 0 {
// 		return nil, errors.New("Task cannot have an empty body")
// 	} else {
// 		task.Body = body
// 	}

// 	if len(scheduled) > 0 {
// 		scheduledTime, err := time.Parse(layoutISO, scheduled)
// 		if err != nil {
// 			return nil, err
// 		}
// 		task.Scheduled = scheduledTime
// 	}

// 	if len(due) > 0 {
// 		dueTime, err := time.Parse(layoutISO, due)
// 		if err != nil {
// 			return nil, err
// 		}
// 		task.Scheduled = dueTime
// 	}

// 	if len(project) > 0 {
// 		if strings.HasPrefix(project, "+") {
// 			task.Project = project
// 		} else {
// 			task.Project = "+" + project
// 		}
// 	}

// 	if len(context) > 0 {
// 		if strings.HasPrefix(context, "@") {
// 			task.Context = context
// 		} else {
// 			task.Context = "@" + context
// 		}
// 	}

// 	task.Done = done
// 	return &task, nil
// }

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
	// NOTE the body should contain the project and context
	// if len(t.Project) > 0 {
	// 	builder.WriteString(t.Project + " ")
	// }
	// if len(t.Context) > 0 {
	// 	builder.WriteString(t.Context + " ")
	// }
	builder.WriteString(t.Body)
	if !t.Due.IsZero() {
		builder.WriteString("due:" + t.Due.Format(layoutISO))
	}
	return builder.String()
}

type TaskList map[int]*Task

func (tl TaskList) AddTask(task *Task) {
	tl[len(tl)] = task
}

func (tl TaskList) CompleteTask(id int) bool {
	task, prs := tl[id]
	if prs {
		task.Done = true
		task.CompletedDate = time.Now()
		return true
	}
	return false
}

func (tl TaskList) RemoveTask(id int) bool {
	_, prs := tl[id]
	if prs {
		delete(tl, id)
		return true
	}
	return false
}
