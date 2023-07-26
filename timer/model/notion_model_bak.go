package model

//type Result struct {
//	Object  string `json:"object,omitempty"`
//	Results []Page `json:"results,omitempty"`
//}
//
//type NotionUser struct {
//	Object string `json:"object,omitempty"`
//	ID     string `json:"id,omitempty"`
//}
//
//type Page struct {
//	Object         string     `json:"object,omitempty"`
//	ID             string     `json:"id,omitempty"`
//	CreatedTime    string     `json:"created_time,omitempty"`
//	LastEditedTime string     `json:"last_edited_time,omitempty"`
//	CreatedBy      NotionUser `json:"created_by,omitempty"`
//	LastEditedBy   NotionUser `json:"last_edited_by,omitempty"`
//	Cover          string     `json:"cover,omitempty"`
//	Icon           string     `json:"icon,omitempty"`
//	Parent         struct {
//		Type       string `json:"type,omitempty"`
//		DatabaseID string `json:"database_id,omitempty"`
//	} `json:"parent,omitempty"`
//	Archived   bool `json:"archived,omitempty"`
//	Properties struct {
//		Tags struct {
//			ID          string `json:"id,omitempty"`
//			Type        string `json:"type,omitempty"`
//			MultiSelect []struct {
//				ID    string `json:"id,omitempty"`
//				Name  string `json:"name,omitempty"`
//				Color string `json:"color,omitempty"`
//			} `json:"multi_select,omitempty"`
//		} `json:"Tags,omitempty"`
//
//		RssLink struct {
//			ID       string `json:"id,omitempty"`
//			Type     string `json:"type,omitempty"`
//			RichText []struct {
//				Type        string   `json:"type,omitempty"`
//				Text        RichText `json:"text,omitempty"`
//				Annotations struct {
//					Bold          bool   `json:"bold,omitempty"`
//					Italic        bool   `json:"italic,omitempty"`
//					Strikethrough bool   `json:"strikethrough,omitempty"`
//					Underline     bool   `json:"underline,omitempty"`
//					Code          bool   `json:"code,omitempty"`
//					Color         string `json:"color,omitempty"`
//				} `json:"annotations,omitempty"`
//				PlainText string `json:"plain_text,omitempty"`
//				Href      string `json:"href,omitempty"`
//			} `json:"rich_text,omitempty"`
//		} `json:"link,omitempty"`
//
//		FilesMedia struct {
//			ID    string `json:"id,omitempty"`
//			Type  string `json:"type,omitempty"`
//			Files []struct {
//				// add fields here
//			} `json:"files,omitempty"`
//		} `json:"Files & media,omitempty"`
//		Name struct {
//			ID        string      `json:"id,omitempty"`
//			Type      string      `json:"type,omitempty"`
//			SiteTitle     []TitleText `json:"title,omitempty"`
//			PlainText string      `json:"plain_text,omitempty"`
//		} `json:"Name,omitempty"`
//	} `json:"properties,omitempty"`
//	URL       string `json:"url,omitempty"`
//	PublicURL string `json:"public_url,omitempty"`
//}
//
//type TitleText struct {
//	Type        string `json:"type,omitempty"`
//	Text        Text   `json:"text,omitempty"`
//	PlainText   string `json:"plain_text,omitempty"`
//	Href        string `json:"href,omitempty"`
//	Annotations struct {
//		Bold          bool   `json:"bold,omitempty"`
//		Italic        bool   `json:"italic,omitempty"`
//		Strikethrough bool   `json:"strikethrough,omitempty"`
//		Underline     bool   `json:"underline,omitempty"`
//		Code          bool   `json:"code,omitempty"`
//		Color         string `json:"color,omitempty"`
//	} `json:"annotations,omitempty"`
//}
//
//type RichText struct {
//	Content string `json:"content,omitempty"`
//	RssLink    struct {
//		Url string `json:"url,omitempty"`
//	} `json:"link,omitempty"`
//}
//type Text struct {
//	Content string `json:"content,omitempty"`
//	RssLink    string `json:"link,omitempty"`
//}
//
//type Data struct {
//	Object         string `json:"object,omitempty"`
//	NextCursor     string `json:"next_cursor,omitempty"`
//	HasMore        bool   `json:"has_more,omitempty"`
//	Type           string `json:"type,omitempty"`
//	PageOrDatabase struct {
//		// add fields here
//	} `json:"page_or_database,omitempty"`
//}
//
///****/
//
//type Row struct {
//	SiteTitle  string
//	Text   string
//	RssLink   string
//	Avatar string
//}
