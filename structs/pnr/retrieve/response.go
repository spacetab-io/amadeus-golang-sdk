package PNR_Information

import (
	"time"
)

type Response struct {
	TravellesInfo []TravellerInfo
}

type TravellerInfo struct {
	FirstName   string
	LastName    string
	Type        PaxType // ADT, CHD, INF
	DateOfBirth time.Time
	Quaifier    string // PT, PI
	Number      int    // reference number
}

type PaxType string
