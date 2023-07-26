package model

import "encoding/xml"

type Feed struct {
	XMLName  xml.Name `xml:"feed"`
	ID       string   `xml:"id"`
	Author   Author   `xml:"author"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
	Link     Link     `xml:"link"`
	Entries  []Entry  `xml:"entry"`
	Updated  string   `xml:"updated"`
}

type Link struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type Entry struct {
	ID        string     `xml:"id"`
	Author    Author     `xml:"author"`
	Title     string     `xml:"title"`
	Summary   string     `xml:"summary"`
	Content   Content    `xml:"content"`
	Link      Link       `xml:"link"`
	Category  []Category `xml:"category"`
	Published string     `xml:"published"`
	Updated   string     `xml:"updated"`
}

type Category struct {
	//Term string `xml:"term,attr"`
	Value string `xml:",chardata"`
}

type Content struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}
