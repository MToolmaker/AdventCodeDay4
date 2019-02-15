package test

import (
	"AdventCodeDay4/Task"
	"testing"
)

// Here are tests for first part of day 4 task

// Let's test simple example input from task's description.
func TestSimpleExampleInputStrategy2(t *testing.T) {
	/* Let's process file and get useful guard activity information */
	_, guardMidnightMinutes, e := task.ProcessFile("SimpleExampleInput")
	if e != nil {
		t.Errorf(e.Error())
	} else {
		guardIdStrategy1, bestMinuteStrategy2 := task.GetBestCandidateStrategy2(guardMidnightMinutes)
		// Example description contains answer: 240
		expectedAnswer := 4455
		actualAnswer := guardIdStrategy1 * bestMinuteStrategy2
		if actualAnswer != expectedAnswer {
			t.Errorf("Actual value %d is not equal to expected %d\n", actualAnswer, expectedAnswer)
		}
	}
}

// Let's test input from task's description.
func TestInputStrategy2(t *testing.T) {
	/* Let's process file and get useful guard activity information */
	_, guardMidnightMinutes, e := task.ProcessFile("Input")
	if e != nil {
		t.Errorf(e.Error())
	} else {
		guardIdStrategy2, bestMinuteStrategy2 := task.GetBestCandidateStrategy2(guardMidnightMinutes)
		// AdventCode site accepted my answer and says it is correct: 99911
		expectedAnswer := 65854
		actualAnswer := guardIdStrategy2 * bestMinuteStrategy2
		if actualAnswer != expectedAnswer {
			t.Errorf("Actual value %d is not equal to expected %d\n", actualAnswer, expectedAnswer)
		}
	}
}
