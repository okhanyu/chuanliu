package pkg

import (
	"encoding/xml"
	"log"
	"rsshub/timer/model"
	"time"
)

func ParseRss(url string) (*model.Rss, *model.Feed, error) {
	response, err := HttpRequestRss(url)
	if err != nil {
		log.Printf("%v请求RSS %s 失败 err is %v", time.Now(), url, err)
		return nil, nil, err
	}
	if rssResult, err := rss(response); err != nil || rssResult.Channel.Title == "" {
		atom, err := atom(response)
		if err != nil {
			return nil, nil, err
		}
		return nil, atom, nil
	} else {
		return rssResult, nil, nil
	}
}

//func ParseRssBak(url string) (*model.Rss, *model.Feed, error) {
//	// 发起HTTP请求获取RSS内容
//	body, err := gohelper_http.HttpPostWithValues(&gohelper_http.RequestInfo{URL: url,
//		Method: http.MethodGet, Timeout: 30 * time.Second},
//		nil,
//		"")
//	//resp, err := http.Get(url)
//	if err != nil {
//		log.Println("请求RSS失败:", err)
//		return nil, nil, err
//	}
//	//defer func(Body io.ReadCloser) {
//	//	err := Body.Close()
//	//	if err != nil {
//	//
//	//	}
//	//}(resp.Body)
//
//	// 读取HTTP响应的内容
//	//body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println("读取响应内容失败:", err)
//		return nil, nil, err
//	}
//
//	if rssResult, err := rss([]byte(body)); err != nil || rssResult.Channel.Title == "" {
//		atom, err := atom([]byte(body))
//		if err != nil {
//			return nil, nil, err
//		}
//		return nil, atom, nil
//	} else {
//		return rssResult, nil, nil
//	}
//}

func atom(body []byte) (*model.Feed, error) {
	// 解析ATOM内容
	var feed model.Feed
	err := xml.Unmarshal(body, &feed)
	if err != nil {
		log.Println("解析 XML 失败:", err)
		return nil, err
	}
	return &feed, nil
}

func rss(body []byte) (*model.Rss, error) {
	// 解析RSS内容
	var rssResult model.Rss
	err := xml.Unmarshal(body, &rssResult)
	if err != nil {
		log.Println("解析RSS失败:", err)
		return nil, err
	}
	return &rssResult, nil
}
