package model

import "time"

type AwsCost struct {
	StartDate time.Time
	EndDate time.Time
	Amount string
}
