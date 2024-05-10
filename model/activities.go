package model

import "time"

var Acti Activities

type Activities struct {
	AdminId   int
	ActiNum   string
	Stock     int
	StartTime time.Time
	EndTime   time.Time
}
