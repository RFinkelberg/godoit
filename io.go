package godoit

import (
	"bufio"
	"os"
)

func LoadFromFile(path string) (TaskList, error) {
	//TODO
	f, err := os.Open(path)
	if err != nil {
		return TaskList{}, err
	}
	defer f.Close()

	var tasks TaskList
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		text := scanner.Text()
		// match := pattern.FindStringSubmatch(text)
		task, err := ParseTaskString(text)
		if err != nil {
			return tasks, err
		}
		tasks.Add(&task)
	}
	return tasks, nil
}
