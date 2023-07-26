package model

import "time"

type MemosUser struct {
	Id          int
	UserName    string
	MemosLink   string
	Avatar      string
	SiteLink    string
	SiteTitle   string
	Description string
	CreateTime  time.Time
	UpdateTime  time.Time
	Del         int
}
