package godoit

import (
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
				CompletedDate: time.Date(2021, time.May, 1, 0, 0, 0, 0, time.UTC),
				Project:       "+career",
				Context:       "@computer",
				Body:          `apply to jobs +career @computer`,
				Created:       time.Date(2021, time.April, 21, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			`(A) Thank Mom for the meatballs @phone`,
			Task{
				Done:     false,
				Priority: A,
				Context:  "@phone",
				Body:     `(A) Thank Mom for the meatballs @phone`,
			},
		},
	}

	for _, test := range tests {
		task, err := ParseTaskString(test.s)
		if err != nil {
			t.Error("ERROR:", err.Error())
		}
		if task != test.expected {
			t.Errorf("\nExpected: %#v,\nActual: %#v", test.expected, task)
		}
	}
}

// func TestParseCompleted(t *testing.T) {
// 	s := `x 2021-05-01 2021-04-21 apply to jobs +career @computer`
// 	want := Task{
// 		Done:          true,
// 		CompletedDate: time.Date(2021, time.May, 1, 0, 0, 0, 0, time.UTC),
// 		Project:       "+career",
// 		Context:       "@computer",
// 		Body:          `apply to jobs +career @computer`,
// 		Created:       time.Date(2021, time.April, 21, 0, 0, 0, 0, time.UTC),
// 	}

// 	task, err := ParseTaskString(s)
// 	if err != nil {
// 		t.Error("ERROR:", err.Error())
// 	}
// 	if task != want {
// 		t.Errorf("\nExpected: %#v,\nActual: %#v", want, task)
// 	}
// }
