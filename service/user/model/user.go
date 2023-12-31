package model

import "time"

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
	Total       int
	Like        int
}

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
