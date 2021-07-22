package godoit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type session struct {
	*TaskList
	savePath string
}

func parseCmd(line string) (Command, string, error) {
	tokens := strings.SplitN(strings.TrimSpace(line), " ", 1)
	args := strings.TrimSpace(tokens[1])
	var cmd Command
	switch strings.ToLower(tokens[0]) {
	case "load":
		cmd = cmdLoad
	case "exit":
		cmd = cmdExit
	case "add":
		cmd = cmdAdd
	case "list":
		cmd = cmdList
	case "done":
		cmd = cmdDone
	default:
		return cmd, args, fmt.Errorf("invalid command %s", tokens[0])
	}
	return cmd, args, nil
}

func RunShell() {
	curSession := new(session)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("TODO> ")
		scanner.Scan()
		cmd, args, err := parseCmd(scanner.Text())
		if err != nil {
			fmt.Println(err)
		} else {
			_, err := cmd(curSession, args)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
