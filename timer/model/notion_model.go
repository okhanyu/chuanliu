package model

import "time"

type Result struct {
	Object  string `json:"object,omitempty"`
	Results []Page `json:"results,omitempty"`
}

type Page struct {
	Object         string `json:"object,omitempty"`
	ID             string `json:"id,omitempty"`
	CreatedTime    string `json:"created_time,omitempty"`
	LastEditedTime string `json:"last_edited_time,omitempty"`
	Properties     struct {
		Name struct {
			ID        string      `json:"id,omitempty"`
			Type      string      `json:"type,omitempty"`
			Title     []TitleText `json:"title,omitempty"`
			PlainText string      `json:"plain_text,omitempty"`
		} `json:"Name,omitempty"`
		Link struct {
			ID       string `json:"id,omitempty"`
			Type     string `json:"type,omitempty"`
			RichText []struct {
				Type      string   `json:"type,omitempty"`
				Text      RichText `json:"text,omitempty"`
				PlainText string   `json:"plain_text,omitempty"`
				Href      string   `json:"href,omitempty"`
			} `json:"rich_text,omitempty"`
		} `json:"link,omitempty"`
		Avatar struct {
			ID       string `json:"id,omitempty"`
			Type     string `json:"type,omitempty"`
			RichText []struct {
				Type      string   `json:"type,omitempty"`
				Text      RichText `json:"text,omitempty"`
				PlainText string   `json:"plain_text,omitempty"`
				Href      string   `json:"href,omitempty"`
			} `json:"rich_text,omitempty"`
		} `json:"avatar,omitempty"`
		Active struct {
			ID     string `json:"id,omitempty"`
			Type   string `json:"type,omitempty"`
			Select struct {
				ID    string `json:"id,omitempty"`
				Name  string `json:"name,omitempty"`
				Color string `json:"color,omitempty"`
			} `json:"select,omitempty"`
		} `json:"active,omitempty"`
		UpdateTime struct {
			Id             string    `json:"id"`
			Type           string    `json:"type"`
			LastEditedTime time.Time `json:"last_edited_time"`
		} `json:"update_time,omitempty"`
	} `json:"properties,omitempty"`
}

type TitleText struct {
	Type      string `json:"type,omitempty"`
	Text      Text   `json:"text,omitempty"`
	PlainText string `json:"plain_text,omitempty"`
	Href      string `json:"href,omitempty"`
}

type RichText struct {
	Content string `json:"content,omitempty"`
	Link    struct {
		Url string `json:"url,omitempty"`
	} `json:"link,omitempty"`
}
type Text struct {
	Content string `json:"content,omitempty"`
	Link    string `json:"link,omitempty"`
}

/****/

//type Row struct {
//	Name   string
//	RssLink   string
//	Avatar string
//}
