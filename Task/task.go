package task

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

/* This function extracts useful information about guard activity from file with given name.
   It returns two maps. First one maps guard ID to total sleeping time. Second one maps guard ID to array of 60 length
   where each element is associated with minute from 00.00 - 00.59 interval. Each element contains number of times when
   a guard was asleep at current minute during all shifts. */
func ProcessFile(fileName string) (map[int]int, map[int]*[60]int, error) {
	bytes, e := ioutil.ReadFile(fileName)
	// Check if file is correctly opened
	if e != nil {
		return nil, nil, e
	} else {
		// We need to split one long string into array of lines.
		lines := strings.Split(string(bytes), "\r\n")
		// We know first tokens are already at format date-time, so we can sort it in lexicographic order
		// and get chronological order
		sort.Strings(lines)
		currentId := -1
		guardSleepingMinutes := make(map[int]int)
		guardMidnightMinutes := make(map[int]*[60]int)
		i := 0
		// This code parses lines
		for i < len(lines) {
			// Check if it's a new shift line
			if strings.Contains(lines[i], "Guard") {
				// Split line with whitespaces into elements
				args := strings.Split(lines[i], " ")
				// Guard ID is third element, but we don't need '#' symbol which is first character of ID part
				// We need to know guard ID who starts new shift
				currentId, _ = strconv.Atoi(args[3][1:])
				// Checking if it's a new guard IT
				if _, b := guardMidnightMinutes[currentId]; !b {
					// Allocating minutes array for him
					guardMidnightMinutes[currentId] = &[60]int{0}
				}
				// Nothing more to do with current line
				i++
				continue
			}
			// At this moment we know that next two lines are "wakes up"/"falls asleep"
			// We can handle both of them in current iteration
			// We know minutes can be only within 0-59 numbers
			// Here we parse minutes from both of lines
			stringTime := strings.Split(lines[i], " ")[1][:5]
			fallMinute := parseMinutes(stringTime)

			stringTime = strings.Split(lines[i+1], " ")[1][:5]
			wakeupMinute := parseMinutes(stringTime)

			// Get difference between wakeup and fall asleep minutes
			timespan := wakeupMinute - fallMinute
			// Adding it to a total time of being asleep for current guard
			guardSleepingMinutes[currentId] += timespan
			// Here we increment counter for every sleeping minute for current guard
			for j := fallMinute; j < wakeupMinute; j++ {
				(*guardMidnightMinutes[currentId])[j]++
			}
			i = i + 2
		}
		return guardSleepingMinutes, guardMidnightMinutes, nil
	}
}

/* This function seeks a guard who sleeps total more than others and returns his ID and minute when he is most likely asleep.
   It gets a map which associate guard ID and his total time being asleep and a map which associate guard ID and
   array of minutes within 0-59, each element of array shows number of times when the guard was asleep at current
   minute */
func GetBestCandidateStrategy1(guardSleepingMinutes map[int]int, guardMidnightMinutes map[int]*[60]int) (int, int) {
	guardId := -1
	maxSleepingMinutes := -1
	// Searching for a guard who sleeps more than others
	for id, sleepingMinutes := range guardSleepingMinutes {
		if maxSleepingMinutes < sleepingMinutes {
			guardId = id
			maxSleepingMinutes = sleepingMinutes
		}
	}
	// Searching for most sleeping minute for current guard
	bestMinute := -1
	counter := -1
	for i := 0; i < 60; i++ {
		tempCounter := (*guardMidnightMinutes[guardId])[i]
		if counter < tempCounter {
			counter = tempCounter
			bestMinute = i
		}
	}
	return guardId, bestMinute
}

/* This function seeks a guard who sleeps more than others at some minute and returns his ID and this minute.
   It gets a map which associate guard ID and array of minutes within 0-59, each element of array shows number
   of times when the guard was asleep at current minute. */
func GetBestCandidateStrategy2(guardMidnightMinutes map[int]*[60]int) (int, int) {
	bestMinute := -1
	counter := -1
	guardId := -1
	// Iterating between all guard ID's
	for id, midnightMinutes := range guardMidnightMinutes {
		// Searching most sleeping minute
		for i := 0; i < 60; i++ {
			tempCounter := (*midnightMinutes)[i]
			if counter < tempCounter {
				counter = tempCounter
				bestMinute = i
				guardId = id
			}
		}
	}
	return guardId, bestMinute
}

// This function parses minutes from string with hours:minutes format
func parseMinutes(stringTime string) int {
	time := strings.Split(stringTime, ":")
	minutes, _ := strconv.Atoi(time[1])
	return minutes
}
