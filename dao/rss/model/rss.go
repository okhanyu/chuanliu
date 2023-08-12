package model

import "time"

type Rss struct {
	Id          int
	UserId      int
	Title       string
	Summary     string
	Cover       string
	Content     string
	Link        string
	Watch       int
	Like        int
	UserName    string
	Tags        string
	Del         int
	CreateTime  time.Time
	PubDateTime time.Time
}

type GetRss struct {
	Id          int
	UserId      int
	Title       string
	Summary     string
	Cover       string
	Content     string
	Link        string
	Watch       int
	Like        int
	UserName    string
	Tags        string
	Del         int
	CreateTime  time.Time
	PubDateTime time.Time
	Avatar      string
}

type UserRecentArticle struct {
	Id          int
	UserId      int
	UserName    string
	Title       string
	Link        string
	PubDateTime time.Time
	//MinPubDateTime time.Time
}
