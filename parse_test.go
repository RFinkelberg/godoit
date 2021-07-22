package godoit

import (
	"fmt"
	"testing"
	"time"
)

func TestParseEmpty(t *testing.T) {
	_, err := ParseTaskString("")
	if err == nil {
		t.Errorf("Parsing empty task string expects an error, but none was given")
	}
}

func TestParseStringTabular(t *testing.T) {
	tests := []struct {
		s        string
		expected Task
	}{
		{
			`x 2021-05-01 2021-04-21 apply to jobs +career @computer`,
			Task{
				Done:          true,
				CompletedDate: TimeFromDay(2021, time.May, 1),
				Project:       "+career",
				Context:       "@computer",
				Body:          `apply to jobs +career @computer`,
				Created:       TimeFromDay(2021, time.April, 21),
			},
		},
		{
			`(A) Thank Mom for the meatballs @phone`,
			Task{
				Done:     false,
				Priority: A,
				Context:  "@phone",
				Body:     `Thank Mom for the meatballs @phone`,
			},
		},
		{
			`(B) Schedule Goodwill pickup +GarageSale @phone`,
			Task{
				Done:     false,
				Priority: B,
				Context:  "@phone",
				Project:  "+GarageSale",
				Body:     `Schedule Goodwill pickup +GarageSale @phone`,
			},
		},
		{
			`Post signs around the neighborhood +GarageSale`,
			Task{
				Done:    false,
				Project: "+GarageSale",
				Body:    `Post signs around the neighborhood +GarageSale`,
			},
		},
		{
			`@GroceryStore Eskimo pies`,
			Task{
				Done:    false,
				Context: "@GroceryStore",
				Body:    `@GroceryStore Eskimo pies`,
			},
		},
		{
			`x 2011-03-03 Call Mom`,
			Task{
				Done:          true,
				CompletedDate: TimeFromDay(2011, time.March, 3),
				Body:          `Call Mom`,
			},
		},
		{
			`x 2011-03-02 2011-03-01 Review Tim's pull request +TodoTxtTouch @github`,
			Task{
				Done:          true,
				CompletedDate: TimeFromDay(2011, time.March, 2),
				Created:       TimeFromDay(2011, time.March, 1),
				Project:       "+TodoTxtTouch",
				Context:       "@github",
				Body:          `Review Tim's pull request +TodoTxtTouch @github`,
			},
		},
		{
			`2011-03-02 Document +TodoTxt task format due:2006-01-02`,
			Task{
				Done:    false,
				Created: TimeFromDay(2011, time.March, 2),
				Project: "+TodoTxt",
				Due:     TimeFromDay(2006, time.January, 2),
				Body:    `Document +TodoTxt task format`,
			},
		},
		{
			`(A) 2011-03-02 Call Mom`,
			Task{
				Done:     false,
				Priority: A,
				Created:  TimeFromDay(2011, time.March, 2),
				Body:     `Call Mom`,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			task, err := ParseTaskString(test.s)
			if err != nil {
				t.Error("ERROR:", err.Error())
			}
			if task != test.expected {
				t.Errorf("\nExpected: %#v,\nActual: %#v", test.expected, task)
			}
		})
	}
}

// Convenience wrapper around time.Date which provides default values for
// time and location consistent with these tests (00:00:00 UTC)
func TimeFromDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
