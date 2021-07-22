package godoit

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

const layoutISO = "2006-01-02"

var (
	doneRegex      = regexp.MustCompile(`^x\s+`)
	contextRegex   = regexp.MustCompile(`@\w+`)
	projectRegex   = regexp.MustCompile(`\+\w+`)
	priorityRegex  = regexp.MustCompile(`\([A-F]\)`)
	dueRegex       = regexp.MustCompile(`due:\d\d\d\d-\d\d-\d\d`)
	createdRegex   = regexp.MustCompile(`(^\d\d\d\d-\d\d-\d\d)|(\([A-F]\)\s\d\d\d\d-\d\d-\d\d)`)
	completedRegex = regexp.MustCompile(`^x\s\d\d\d\d-\d\d-\d\d`)
)

// Parses a string in standard todo.txt format into a task struct
func ParseTaskString(s string) (Task, error) {
	if len(s) <= 0 {
		return Task{}, errors.New("Task string may not be empty")
	}
	// TODO
	done := doneRegex.FindStringSubmatch(s) != nil
	context := contextRegex.FindString(s)
	project := projectRegex.FindString(s)
	priority := priorityRegex.FindString(s)
	dueDate := dueRegex.FindString(s)
	completedDate := completedRegex.FindString(s)

	s = strings.TrimLeft(completedRegex.ReplaceAllString(s, ""), " ")
	createdDate := createdRegex.FindString(s)

	// Remove all matched parts from string, leaving only the body of the task
	body := doneRegex.ReplaceAllString(s, "")
	// body = contextRegex.ReplaceAllString(body, "")
	// body = projectRegex.ReplaceAllString(body, "")
	body = priorityRegex.ReplaceAllString(body, "")
	body = dueRegex.ReplaceAllString(body, "")
	body = createdRegex.ReplaceAllString(body, "")

	body = strings.TrimSpace(body)

	task := Task{
		Body:     body,
		Done:     done,
		Context:  context,
		Project:  project,
		Priority: Priority(priority),
	}
	if len(dueDate) > 0 {
		// NOTE we have to remove the "due:" tag key
		// TODO be able to handle generic key:value tag pairs
		due, err := time.Parse(layoutISO, strings.TrimPrefix(dueDate, "due:"))
		if err != nil {
			return task, err
		}
		task.Due = due
	}
	if len(createdDate) > 0 {
		// NOTE: The created date can have multiple different prefixes, so we
		// need a sub regex to extract the date itself. This would be easier
		// if golang regex supported lookbehind (or maybe with capture groups)
		createdDate = regexp.MustCompile(`\d\d\d\d-\d\d-\d\d`).FindString(createdDate)
		created, err := time.Parse(layoutISO, createdDate)
		if err != nil {
			return task, err
		}
		task.Created = created
	}
	if len(completedDate) > 0 {
		// NOTE we need to slice off the "x " before the completed date
		completed, err := time.Parse(layoutISO, completedDate[2:])
		if err != nil {
			return task, err
		}
		task.CompletedDate = completed.In(time.UTC)
	}

	return task, nil
}
