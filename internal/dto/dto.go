package dto

import "time"

type RequestCreategRPCInServic struct {
	Name        string
	Description string
	Timestamp   time.Time
}