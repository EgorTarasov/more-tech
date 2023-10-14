package service

import (
	"fmt"
	"more-tech/internal/model"
	"strconv"
	"strings"
)

func GetClosestTimeSlot(travelTime float64, workload []model.Workload, customTimeSlot string) string {
	// arrivalTime := time.Now().Local().Add(time.Second * time.Duration(travelTime))
	// weekdayKey := arrivalTime.Weekday().String()
	// weekdays := map[string]int{
	// 	"Monday":    0,
	// 	"Tuesday":   1,
	// 	"Wednesday": 2,
	// 	"Thursday":  3,
	// 	"Friday":    4,
	// 	"Saturday":  5,
	// 	"Sunday":    6,
	// }
	// weekday := weekdays[weekdayKey]
	weekday := 0

	if customTimeSlot == "" {
		for idx, timeSlot := range workload[weekday].LoadHours {
			splitted := strings.Split(timeSlot.Hour, "-")
			if strings.Split(splitted[0], ":")[0] == fmt.Sprintf("%d", 12) { // arrivalTime.Hour()
				minutes, _ := strconv.Atoi(strings.Split(splitted[0], ":")[1])
				if minutes > 0 { // arrivalTime.Minute()
					return workload[weekday].LoadHours[idx-1].Hour
				} else {
					return workload[weekday].LoadHours[idx].Hour
				}
			}
		}
		return "no timeslot found"
	} else {
		splitted := strings.Split(customTimeSlot, "-")
		if strings.Split(splitted[0], ":")[0] == fmt.Sprintf("%d", 12) { // arrivalTime.Hour()
			minutes, _ := strconv.Atoi(strings.Split(splitted[0], ":")[1])
			if minutes > 0 { // arrivalTime.Minute()
				return "no timeslot found"
			} else {
				return customTimeSlot
			}
		}
		return "no timeslot found"
	}
}
