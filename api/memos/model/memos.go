package model

import "time"

type MemosUser struct {
	Id          int       `json:"id" form:"id"`
	UserName    string    `json:"user_name" form:"user_name"`
	MemosLink   string    `json:"memos_link" form:"memos_link"`
	Avatar      string    `json:"avatar" form:"avatar"`
	SiteLink    string    `json:"site_link" form:"site_link"`
	SiteTitle   string    `json:"site_title" form:"site_title"`
	Description string    `json:"description" form:"description"`
	CreateTime  time.Time `json:"create_time" form:"create_time"`
	UpdateTime  time.Time `json:"update_time" form:"update_time"`
}
