package pkg

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func GetOpenGraphImage(url string) (string, bool) {

	response, err := HttpRequestRss(url)
	//resp, err := http.Get(url)
	if err != nil {
		log.Println("请求RSS失败:", err)
		return "", false
	}

	// 将 io.Reader 转换为 io.ReadCloser
	readCloser := ioutil.NopCloser(strings.NewReader(string(response)))
	defer func(readCloser io.ReadCloser) {
		err := readCloser.Close()
		if err != nil {

		}
	}(readCloser)

	doc, err := goquery.NewDocumentFromReader(readCloser)
	if err != nil {
		log.Println("解析HTML失败:", err)
		return "", false
	}
	content, exists := "", false
	doc.Find("meta[property='og:image']").Each(func(i int, s *goquery.Selection) {
		if content == "" || exists == false {
			content, exists = s.Attr("content")
		}
	})
	return content, exists
}

func GetContentImage(content string) (string, bool) {
	// 将字符串转换为 io.Reader
	reader := strings.NewReader(content)
	// 将 io.Reader 转换为 io.ReadCloser
	readCloser := ioutil.NopCloser(reader)
	// 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(readCloser)
	if err != nil {
		log.Println("搜索content image失败:", err)
		return "", false
	}

	// 查找第一个图像元素
	//imageURL := ""
	//doc.Find("body img").Each(func(i int, s *goquery.Selection) {
	//	if i == 0 {
	//		imageURL, _ = s.Attr("src")
	//		return
	//	}
	//})

	// 查找body下的第一个img标签
	img := doc.Find("img").First()
	// 获取img标签的src属性值
	src, exists := img.Attr("src")
	// 关闭 readCloser
	readCloser.Close()
	return src, exists

}

func GetFirstImage(url string) (string, bool) {
	if url == "" {
		return "", false
	}
	// 发起HTTP请求获取网页内容
	//response, err := http.Get(url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer response.Body.Close()

	response, err := HttpRequestRss(url)
	//resp, err := http.Get(url)
	if err != nil {
		log.Println("请求RSS失败:", err)
		return "", false
	}

	// 将 io.Reader 转换为 io.ReadCloser
	readCloser := ioutil.NopCloser(strings.NewReader(string(response)))
	defer func(readCloser io.ReadCloser) {
		err := readCloser.Close()
		if err != nil {

		}
	}(readCloser)

	// 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(readCloser)
	if err != nil {
		log.Println("解析image失败:", err)
		return "", false
	}

	// 查找第一个图像元素
	//imageURL := ""
	//doc.Find("body img").Each(func(i int, s *goquery.Selection) {
	//	if i == 0 {
	//		imageURL, _ = s.Attr("src")
	//		return
	//	}
	//})

	firstImg := doc.Find("body").Find("article").First().Find("img").Not("header img").
		Not("nav img").First()
	//firstImg := doc.Find("body").Find("article").First().Not("header").Find("img").First()
	src, exists := firstImg.Attr("src")
	if !exists {
		firstImg = doc.Find("body").Find("figure").First().Find("img").Not("header img").
			Not("nav img").First()
		src, exists = firstImg.Attr("src")
	}
	//if exists {
	return src, exists
	//}
	//firstImg = doc.Find("body").Find("img").First()
	//src, exists = firstImg.Attr("src")
	//if exists {
	//	return src, exists
	//}
	//return "", false
}
