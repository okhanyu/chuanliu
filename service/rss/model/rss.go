package model

import "time"

type Rss struct {
	Id         int
	Title      string
	Summary    string
	Link       string
	Watch      int
	UserName   string
	Tags       []string
	Del        int
	CreateTime time.Time
	UserId     int
	PubDate    time.Time
	Cover      string
	//Domain         string
	//Protocal       string
}

type GetRss struct {
	Id         int
	Title      string
	Summary    string
	Link       string
	Watch      int
	UserName   string
	Tags       []string
	Del        int
	CreateTime time.Time
	UserId     int
	PubDate    time.Time
	Cover      string
	Avatar     string
	//Domain         string
	//Protocal       string
}

type UserRecentArticle struct {
	Id             int
	UserId         int
	UserName       string
	Title          string
	Link           string
	PubDate        time.Time
	PubDateShow    string
	MinPubDate     time.Time
	MinPubDateShow string
}
