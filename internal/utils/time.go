package utils

import "time"

func TimeNow() time.Time {
	// Time in Colombia
	loc, _ := time.LoadLocation("America/Bogota")
	return time.Now().In(loc)
}

func TimeAddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}
