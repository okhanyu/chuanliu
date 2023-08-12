package model

import "time"

type Rss struct {
	Id          int       `json:"id" form:"id"`
	Title       string    `json:"title" form:"title"`
	Summary     string    `json:"summary" form:"summary"`
	Link        string    `json:"link" form:"link"`
	Watch       int       `json:"watch" form:"watch"`
	Like        int       `json:"like" form:"like"`
	UserName    string    `json:"user_name" form:"user_name"`
	Tags        []string  `json:"tags" form:"tags"`
	PubDate     time.Time `json:"pub_date" form:"pub_date"`
	PubDateShow string    `json:"pub_date_show" form:"pub_date_show"`
	Cover       string    `json:"cover" form:"cover"`
	UserId      int       `json:"user_id" form:"user_id"`
}

type GetRss struct {
	Id          int       `json:"id" form:"id"`
	Title       string    `json:"title" form:"title"`
	Summary     string    `json:"summary" form:"summary"`
	Link        string    `json:"link" form:"link"`
	Watch       int       `json:"watch" form:"watch"`
	Like        int       `json:"like" form:"like"`
	UserName    string    `json:"user_name" form:"user_name"`
	Tags        []string  `json:"tags" form:"tags"`
	PubDate     time.Time `json:"pub_date" form:"pub_date"`
	PubDateShow string    `json:"pub_date_show" form:"pub_date_show"`
	Cover       string    `json:"cover" form:"cover"`
	UserId      int       `json:"user_id" form:"user_id"`
	Avatar      string    `json:"avatar" form:"avatar"`
}

type UserRecentArticle struct {
	Id          int       `json:"id" form:"id"`
	UserId      int       `json:"user_id" form:"user_id"`
	UserName    string    `json:"user_name" form:"user_name"`
	Title       string    `json:"title" form:"title"`
	Link        string    `json:"link" form:"link"`
	PubDate     time.Time `json:"pub_date" form:"pub_date"`
	PubDateShow string    `json:"pub_date_show" form:"pub_date_show"`
	//MinPubDateTime time.Time `json:"min_pub_date_time" form:"min_pub_date_time"`
	//MinPubDateShow string    `json:"min_pub_date_show" form:"min_pub_date_show"`
}
