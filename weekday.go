package main

import (
	"math"
	"time"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func GetNextWeekday(start time.Time, weekday Weekday) time.Time {
	daysToAdd := math.Mod(float64(int(weekday)-int(start.Weekday())+7), 7)
	return start.AddDate(0, 0, int(daysToAdd))
}
