package pkg

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HttpRequestRss(url string) ([]byte, error) {
	client := &http.Client{}
	client.Timeout = 30 * time.Second
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Close = true
	req.Header.Set("Content-Type", "application/rss+xml; charset=UTF-8")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应内容失败:", err)
		return nil, err
	}

	return response, nil
}

func HttpRequest(url string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, body)
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	//req.SetBasicAuth("user", "pass")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应内容失败:", err)
		return nil, err
	}

	return response, nil
}

func HttpRequestWithHeader(url string, body io.Reader, header http.Header) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, body)
	req.Close = true
	req.Header = header
	//req.SetBasicAuth("user", "pass")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应内容失败:", err)
		return nil, err
	}

	return response, nil

}
