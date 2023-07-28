package model

import "time"

type User struct {
	Id          int
	UserName    string
	RssLink     string
	Avatar      string
	SiteLink    string
	SiteTitle   string
	Description string
	CreateTime  time.Time
	UpdateTime  time.Time
	Del         int
}

type UserStatistics struct {
	Id          int
	UserName    string
	RssLink     string
	Avatar      string
	SiteLink    string
	SiteTitle   string
	Description string
	CreateTime  time.Time
	UpdateTime  time.Time
	Del         int
	Watch       int
	Like        int
	Total       int
}
