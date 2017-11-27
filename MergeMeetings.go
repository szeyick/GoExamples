/**
To do this, you’ll need to know when any team is having a meeting. In HiCal, a meeting is stored as an
object of a Meeting class with integer variables startTime and endTime. These integers represent the number of
30-minute blocks past 9:00am.

Given: [Meeting(0, 1), Meeting(3, 5), Meeting(4, 8), Meeting(10, 12), Meeting(9, 10)]
Then:  [Meeting(0, 1), Meeting(3, 8), Meeting(9, 12)]

*/
package main

import "fmt"

/**
 The struct to hold the values.
 */
type meeting struct {
	startTime int
	endTime int
}

/**
	The application main.
 */
func main() {
	fmt.Println("Hello there")
	meetings := initMeetings()
	printMeeting(meetings)
	sort(meetings)

	mergedMeetings := mergeMeetings(meetings)

	printMeeting(mergedMeetings)
	fmt.Println("Capacity: ", cap(mergedMeetings), " Length: ", len(mergedMeetings))
}

/**
	Generate some meetings.
 */
func initMeetings() ([]meeting){
	meetings := make([]meeting, 5, 5)

	meetings[0] = meeting {0,1}
	meetings[1] = meeting {3, 5}
	meetings[2] = meeting {4, 8}
	meetings[3] = meeting {10, 12}
	meetings[4] = meeting {9, 10}

	return meetings
}

/**
	Merge the meetings together.
 */
func mergeMeetings(meetings []meeting) ([]meeting) {
	mergeMeetings := make([]meeting, len(meetings), len(meetings))
	previousMeeting := meetings[0]
	var mergeMeetIndex int = 0
	for i := 1; i < len(meetings); i++ {
		currentMeeting := meetings[i]
		// If the current meeting cannot be merged into the previous meeting, then add the current meeting to
		// the merged list and set the current meeting as the previous.
		if currentMeeting.startTime > previousMeeting.startTime && currentMeeting.startTime <= previousMeeting.endTime {
			if currentMeeting.endTime >= previousMeeting.endTime {
				previousMeeting.endTime = currentMeeting.endTime
			}
		} else {
			mergeMeetings[mergeMeetIndex] = previousMeeting
			mergeMeetIndex++
			previousMeeting = currentMeeting
		}
		// Special Case: We need to add the last meeting into the merge.
		if i + 1 == len(meetings) {
			mergeMeetings[mergeMeetIndex] = previousMeeting
			mergeMeetIndex++
		}
	}
	// Drop the remaining capacity of the merged meetings.
	mergeMeetings = mergeMeetings[:mergeMeetIndex]
	return mergeMeetings
}

/**
	Implementetation of a basic sort to sort the meeting times by start.
 */
func sort(meetings []meeting) {
	for i := 0; i < len(meetings); i++ {
		earliestMeeting := meetings[i]
		earliestIndex := i
		for j := i + 1; j < len(meetings); j++ {
			// Search through the rest of the array to find something smaller.
			nextMeeting := meetings[j]
			if nextMeeting.startTime < earliestMeeting.startTime {
				earliestIndex = j
				earliestMeeting = meetings[j]
			}
		}
		if earliestIndex != i {
			// Swap current with earliest index.
			swap(i, earliestIndex, meetings)
		}
	}
}

/**
	Swap the entries.
 */
func swap(fromIndex int, toIndex int, meetings[] meeting) {
	meetingTmp := meetings[fromIndex]
	meetings[fromIndex] = meetings[toIndex]
	meetings[toIndex] = meetingTmp
}

/**
	Print the contents of the meeting.
 */
func printMeeting(meetings []meeting) {
	for i := 0; i < len(meetings); i++ {
		fmt.Println("Meeting Start: ", meetings[i].startTime, " - Meeting End: ", meetings[i].endTime)
	}
	fmt.Println()
}
