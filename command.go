package godoit

import (
	"errors"
	"fmt"
	"strconv"
)

type Command func(*session, string) (bool, error)

var (
	cmdAdd = func(s *session, args string) (bool, error) {
		if s.TaskList == nil {
			s.TaskList = &TaskList{}
		}

		task, err := ParseTaskString(args)
		if err != nil {
			return false, err
		}
		s.Add(&task)
		return true, nil
	}
	cmdList = func(s *session, _ string) (bool, error) {
		for i, task := range s.tasks {
			fmt.Printf("(%d):\t%s", i, task.String())
		}
		return true, nil
	}
	cmdLoad = func(s *session, args string) (bool, error) {
		var err error
		*s.TaskList, err = LoadFromFile(args)
		if err != nil {
			return false, err
		}
		s.savePath = args
		return true, nil
	}
	cmdExit = func(s *session, args string) (bool, error) {
		if len(args) == 0 {
			args = s.savePath
		}
		if err := s.SaveFile(args); err != nil {
			return false, err
		}
		return true, nil
	}
	cmdDone = func(s *session, args string) (bool, error) {
		taskIdx, err := strconv.Atoi(args)
		if err != nil {
			return false, fmt.Errorf("error parsing args to done: %v", err)
		}
		if ok := s.Complete(taskIdx); !ok {
			return false, errors.New("failed to complete task")
		}
		return true, nil
	}
)
