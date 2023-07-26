package timer

import (
	"fmt"
	"rsshub/config"
	rssDao "rsshub/dao/rss"
	rssModel "rsshub/dao/rss/model"
	userDao "rsshub/dao/user"
	userModel "rsshub/dao/user/model"
	"rsshub/pkg"
	"rsshub/timer/model"
	"strconv"
	"strings"
	"time"
)

var LenContent int

const (
	RssGetTimer = "rss_get_timer"
	HttpPrefix  = "http"
	DataPrefix  = "data:image"
)

func GetRssTimer() {
	timeSetting, _ := strconv.Atoi(config.GlobalConfig.System[RssGetTimer])
	ticker := time.NewTicker(time.Duration(timeSetting) * time.Second)
	go func() {
		for range ticker.C {
			fmt.Printf("[获取Rss定时任务执行开始 %v]\n", time.Now())
			RssTask()
			fmt.Printf("[获取Rss定时任务执行完毕 %v]\n", time.Now())
		}
	}()
	// time.Sleep(5 * time.Second)
	// ticker.Stop()
	// fmt.Println("定时任务已停止")
}

func RssTask() {
	userList, err := userDao.GetUserList()
	if err != nil {
		return
	}

	for _, userObj := range userList {
		rssLink := userObj.RssLink
		rssContent, atomResult, err := pkg.ParseRss(rssLink)
		if err != nil {
			continue
		}

		//var items []rssModel.Rss
		if rssContent != nil {
			_ = eachRss(rssContent, userObj)
		}

		if atomResult != nil {
			_ = eachAtom(atomResult, userObj)
		}

		//for _, item := range items {
		//	rssDao.AddRss(item)
		//}
	}
}

func eachRss(rssContent *model.Rss, userObj userModel.User) []rssModel.Rss {
	var items []rssModel.Rss
	for _, item := range rssContent.Channel.Items {
		cate := ""
		for i, s := range item.Categories {
			if i > 0 {
				cate = fmt.Sprintf("%s,%s", cate, s.Value)
			} else {
				cate = s.Value
			}
		}
		pubDate, _ := pkg.PubDateTimeConvert(item.PubDate)
		image := getRssImg(item)

		var protocolDomain string
		urls := strings.SplitAfter(item.Link, "://")
		if len(urls) > 1 {
			uris := strings.SplitAfterN(urls[1], "/", 2)
			protocolDomain = urls[0] + uris[0]
		} else {
			protocolDomain = urls[0]
		}

		if image != "" && !strings.HasPrefix(image, HttpPrefix) {
			image = fmt.Sprintf("%s%s", protocolDomain, image)
		}

		if len(item.Description) > LenContent {
			item.Description = item.Description[:LenContent]
		}

		userName := userObj.UserName
		if item.Author != "" {
			userName = item.Author
		}
		//items = append(items, rssModel.Rss{
		//	Title:       item.Title,
		//	Summary:     item.Description,
		//	Content:     item.Content,
		//	Link:        item.Link,
		//	PubDateTime: pubDate,
		//	UserName:    userName,
		//	Tags:        cate,
		//	Cover:       image,
		//	UserId:      userObj.Id,
		//	CreateTime:  time.Now(),
		//})
		rssDao.AddRss(rssModel.Rss{
			Title:       item.Title,
			Summary:     item.Description,
			Content:     item.Content,
			Link:        item.Link,
			PubDateTime: pubDate,
			UserName:    userName,
			Tags:        cate,
			Cover:       image,
			UserId:      userObj.Id,
			CreateTime:  time.Now(),
		})
	}
	return items
}

func eachAtom(atomResult *model.Feed, userObj userModel.User) []rssModel.Rss {
	var items []rssModel.Rss
	for _, item := range atomResult.Entries {
		cate := ""
		for i, s := range item.Category {
			s.Value = strings.ReplaceAll(s.Value, "\n", "")
			s.Value = strings.TrimSpace(s.Value)
			if i > 0 {
				cate = fmt.Sprintf("%s,%s", cate, s.Value)
			} else {
				cate = s.Value
			}
		}
		if item.Published == "" {
			item.Published = item.Updated
		}
		pubDate, _ := pkg.PubDateTimeConvert(item.Published)
		image := getAtomImg(item)

		if len(item.Content.Value) > LenContent {
			item.Content.Value = item.Content.Value[:LenContent]
		}

		userName := userObj.UserName
		if item.Author.Name != "" {
			userName = item.Author.Name
		}
		urlLink := item.Link.Href
		if urlLink == "" {
			urlLink = item.ID
		}

		var protocolDomain string
		urls := strings.SplitAfter(urlLink, "://")
		if len(urls) > 1 {
			uris := strings.SplitN(urls[1], "/", 2)
			protocolDomain = urls[0] + uris[0]
		} else {
			protocolDomain = urls[0]
		}

		if image != "" && !strings.HasPrefix(image, HttpPrefix) {
			image = fmt.Sprintf("%s%s", protocolDomain, image)
		}
		//items = append(items, rssModel.Rss{
		//	Title:       item.Title,
		//	Summary:     item.Summary,
		//	Content:     item.Content.Value,
		//	Link:        urlLink,
		//	PubDateTime: pubDate,
		//	UserName:    userName,
		//	Tags:        cate,
		//	Cover:       image,
		//	UserId:      userObj.Id,
		//	CreateTime:  time.Now(),
		//})
		rssDao.AddRss(rssModel.Rss{
			Title:       item.Title,
			Summary:     item.Summary,
			Content:     item.Content.Value,
			Link:        urlLink,
			PubDateTime: pubDate,
			UserName:    userName,
			Tags:        cate,
			Cover:       image,
			UserId:      userObj.Id,
			CreateTime:  time.Now(),
		})
	}
	return items
}

func getAtomImg(item model.Entry) string {
	rssLink := item.Link.Href
	if rssLink == "" {
		rssLink = item.ID
	}
	image, exists := pkg.GetOpenGraphImage(rssLink)
	if image != "" && exists {
		return image
	}

	content := item.Content.Value
	if item.Content.Value == "" {
		content = item.Summary
	}
	image, exists = pkg.GetContentImage(content)
	if strings.HasPrefix(image, DataPrefix) {
		image = ""
	}
	if image != "" && exists {
		return image
	}
	image, exists = pkg.GetFirstImage(rssLink)
	if strings.HasPrefix(image, DataPrefix) {
		image = ""
	}
	return image
}

func getRssImg(item model.Item) string {
	image, exists := pkg.GetOpenGraphImage(item.Link)
	if image != "" && exists {
		return image
	}
	image, exists = pkg.GetContentImage(item.Description)
	if strings.HasPrefix(image, DataPrefix) {
		image = ""
	}
	if image != "" && exists {
		return image
	}
	image, exists = pkg.GetFirstImage(item.Link)
	if strings.HasPrefix(image, DataPrefix) {
		image = ""
	}
	return image
}

//
//
//func getAtomImg(item model.Entry) string {
//	content := item.Content.Value
//	if item.Content.Value == "" {
//		content = item.Summary
//	}
//	image, exists := pkg.GetContentImage(content)
//	if image != "" && exists {
//		return image
//	}
//	rssLink := item.Link.Href
//	if rssLink == "" {
//		rssLink = item.ID
//	}
//	image, exists = pkg.GetFirstImage(rssLink)
//	if strings.HasPrefix(image, DataPrefix) {
//		image = ""
//	}
//	return image
//}
//
//func getRssImg(item model.Item) string {
//	image, exists := pkg.GetContentImage(item.Description)
//	if image != "" && exists {
//		return image
//	}
//	image, exists = pkg.GetFirstImage(item.Link)
//	if strings.HasPrefix(image, DataPrefix) {
//		image = ""
//	}
//	return image
//}
