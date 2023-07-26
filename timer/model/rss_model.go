package model

import "encoding/xml"

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	XMLName       xml.Name `xml:"channel"`
	Title         string   `xml:"title"`
	Description   string   `xml:"description"`
	Link          string   `xml:"link"`
	Items         []Item   `xml:"item"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Image         Image    `xml:"image"`
}

type Image struct {
	Url    string `xml:"url"`
	Title  string `xml:"title"`
	Link   string `xml:"link"`
	Width  string `xml:"width"`
	Height string `xml:"height"`
}

type Item struct {
	Title       string     `xml:"title"`
	Description string     `xml:"description"`
	Content     string     `xml:"encoded"`
	Link        string     `xml:"link"`
	Categories  []Category `xml:"category"`
	Author      string     `xml:"creator"`
	PubDate     string     `xml:"pubDate"`
}

type Categories struct {
	XMLName xml.Name `xml:"category"`
	Value   string   `xml:",chardata"`
}
