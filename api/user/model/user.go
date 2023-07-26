package model

import "time"

type UserStatistics struct {
	Id          int       `json:"id" form:"id"`
	UserName    string    `json:"user_name" form:"user_name"`
	RssLink     string    `json:"rss_link" form:"rss_link"`
	Avatar      string    `json:"avatar" form:"avatar"`
	SiteLink    string    `json:"site_link" form:"site_link"`
	SiteTitle   string    `json:"site_title" form:"site_title"`
	Description string    `json:"description" form:"description"`
	CreateTime  time.Time `json:"create_time" form:"create_time"`
	UpdateTime  time.Time `json:"update_time" form:"update_time"`
	Del         int       `json:"del" form:"del"`
	Watch       int       `json:"watch" form:"watch"`
	Total       int       `json:"total" form:"total"`
}

type User struct {
	Id          int       `json:"id" form:"id"`
	UserName    string    `json:"user_name" form:"user_name"`
	RssLink     string    `json:"rss_link" form:"rss_link"`
	Avatar      string    `json:"avatar" form:"avatar"`
	SiteLink    string    `json:"site_link" form:"site_link"`
	SiteTitle   string    `json:"site_title" form:"site_title"`
	Description string    `json:"description" form:"description"`
	CreateTime  time.Time `json:"create_time" form:"create_time"`
	UpdateTime  time.Time `json:"update_time" form:"update_time"`
}
