package main

import (
	"AdventCodeDay4/Task"
	"fmt"
)

func main() {
	/* Let's process file and get useful guard activity information */
	guardSleepingMinutes, guardMidnightMinutes, e := task.ProcessFile("Tests/Input")
	if e != nil {
		fmt.Println(e.Error())
	} else {
		guardIdStrategy1, bestMinuteStrategy1 := task.GetBestCandidateStrategy1(guardSleepingMinutes, guardMidnightMinutes)

		fmt.Printf("Most sleeping guard has id number: %d\n", guardIdStrategy1)
		fmt.Printf("Best minute for sneaking is %d according strategy 1\n", bestMinuteStrategy1)
		fmt.Printf("His best minute multiplied by his ID is %d\n", guardIdStrategy1*bestMinuteStrategy1)

		guardIdStrategy2, bestMinuteStrategy2 := task.GetBestCandidateStrategy2(guardMidnightMinutes)

		fmt.Printf("Most sleeping guard has id number: %d\n", guardIdStrategy2)
		fmt.Printf("Best minute for sneaking is %d according strategy 2\n", bestMinuteStrategy2)
		fmt.Printf("His best minute multiplied by his ID is %d\n", guardIdStrategy2*bestMinuteStrategy2)
	}
}
