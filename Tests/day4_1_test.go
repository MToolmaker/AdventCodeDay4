package test

import (
	"AdventCodeDay4/Task"
	"testing"
)

// Here are tests for first part of day 4 task

// Let's test simple example input from task's description.
func TestSimpleExampleInputStrategy1(t *testing.T) {
	/* Let's process file and get useful guard activity information */
	guardSleepingMinutes, guardMidnightMinutes, e := task.ProcessFile("SimpleExampleInput")
	if e != nil {
		t.Errorf(e.Error())
	} else {
		guardIdStrategy1, bestMinuteStrategy1 := task.GetBestCandidateStrategy1(guardSleepingMinutes, guardMidnightMinutes)
		// Example description contains answer: 240
		expectedAnswer := 240
		actualAnswer := guardIdStrategy1 * bestMinuteStrategy1
		if actualAnswer != expectedAnswer {
			t.Errorf("Actual value %d is not equal to expected %d\n", actualAnswer, expectedAnswer)
		}
	}
}

// Let's test input from task's description.
func TestInputStrategy1(t *testing.T) {
	/* Let's process file and get useful guard activity information */
	guardSleepingMinutes, guardMidnightMinutes, e := task.ProcessFile("Input")
	if e != nil {
		t.Errorf(e.Error())
	} else {
		guardIdStrategy1, bestMinuteStrategy1 := task.GetBestCandidateStrategy1(guardSleepingMinutes, guardMidnightMinutes)
		// AdventCode site accepted my answer and says it is correct: 99911
		expectedAnswer := 99911
		actualAnswer := guardIdStrategy1 * bestMinuteStrategy1
		if actualAnswer != expectedAnswer {
			t.Errorf("Actual value %d is not equal to expected %d\n", actualAnswer, expectedAnswer)
		}
	}
}
