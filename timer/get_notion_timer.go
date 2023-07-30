package timer

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/okhanyu/gohelper/gohelper_http"
	"github.com/okhanyu/gohelper/gohelper_server"
	"gorm.io/gorm"
	"net/http"
	"rsshub/config"
	userDao "rsshub/dao/user"
	userModel "rsshub/dao/user/model"
	timerModel "rsshub/timer/model"
	"strconv"
	"strings"
	"time"
)

const (
	NotionGetTimer = "notion_get_timer"
	DatabasesUrl   = "databases_url"
	Token          = "token"
	Authorization  = "Authorization"
	NotionVersion  = "Notion-Version"
	Version        = "version"
	Bearer         = "Bearer"
)

func GetNotionTimer() {
	timeConfig := config.GlobalConfig.System[NotionGetTimer]
	timeConfigInt, err := strconv.Atoi(timeConfig)
	if err != nil {
		timeConfigInt = 1200
		fmt.Printf("无法读取和转换配置中获取NOTION数据的时间，使用降级值：%d", timeConfigInt)
	}
	ticker := time.NewTicker(time.Duration(timeConfigInt) * time.Second)
	go func() {
		for range ticker.C {
			fmt.Printf("[获取Notion定时任务执行开始 %v]\n", time.Now())
			getUserTask()
			fmt.Printf("[获取Notion定时任务执行完毕 %v]\n", time.Now())
		}
	}()
	// time.Sleep(5 * time.Second)
	// ticker.Stop()
	// fmt.Println("定时任务已停止")
}

var userExeFlag = true

func GetUserTask(c *gin.Context) {
	if userExeFlag {
		go func() {
			userExeFlag = false
			fmt.Printf("[获取Notion User主动任务执行开始 %v]\n", time.Now())
			getUserTask()
			fmt.Printf("[获取Notion User主动任务执行完毕 %v]\n", time.Now())
			userExeFlag = true
		}()
		gohelper_server.Success(c, "执行成功")
	} else {
		gohelper_server.Success(c, "任务正在执行中，本次不执行")
	}
}

func getUserTask() {
	notionDatabasesUrl := config.GlobalConfig.Notion[DatabasesUrl]
	header := make(map[string]string)
	header[Authorization] = fmt.Sprintf("%s %s", Bearer, config.GlobalConfig.Notion[Token])
	header[NotionVersion] = config.GlobalConfig.Notion[Version]
	getNotionResult, err := gohelper_http.HttpPostWithValues(&gohelper_http.RequestInfo{
		Header:  header,
		URL:     notionDatabasesUrl,
		Method:  http.MethodPost,
		Timeout: 30 * time.Second},
		nil,
		"")
	if err != nil {
		fmt.Printf("请求NOTION失败: %v", err)
		return
	}

	fmt.Printf("请求NOTION成功: %s", notionDatabasesUrl)
	response := timerModel.Result{}
	err = json.Unmarshal([]byte(getNotionResult), &response)
	if err != nil {
		fmt.Printf("解析NOTION结果失败: %v", err)
		return
	}

	// 遍历Notion结果
	for _, result := range response.Results {

		// 获取Active
		active := result.Properties.Active.Select.Name

		userObj := userModel.User{}
		// 获取姓名
		for _, title := range result.Properties.Name.Title {
			if title.Text.Content != "" {
				userObj.UserName = title.Text.Content
			}
		}
		// 获取RSS_LINK
		for _, s := range result.Properties.Link.RichText {
			if s.Text.Content != "" || s.Text.Link.Url != "" {
				userObj.RssLink = compareLink(s.Text.Content, s.Text.Link.Url)
			}
		}
		// 获取Avatar
		for _, s := range result.Properties.Avatar.RichText {
			if s.Text.Content != "" || s.Text.Link.Url != "" {
				userObj.Avatar = compareLink(s.Text.Content, s.Text.Link.Url)
			}
		}

		if userObj.UserName == "" || userObj.RssLink == "" {
			continue
		}

		getUser, err := userDao.GetUserByRssLink(userModel.GetUserReq{RssLink: userObj.RssLink})
		if err != nil && err != gorm.ErrRecordNotFound {
			continue
		}

		//// 获取并解析RSS内容(用于验证RSS)
		//rssGetResult, atomResult, err := pkg.ParseRss(userObj.RssLink)
		//if err != nil {
		//	continue
		//}
		//if rssGetResult != nil {
		//	userObj.Description = rssGetResult.Channel.Description
		//	userObj.SiteTitle = rssGetResult.Channel.Title
		//	userObj.SiteLink = rssGetResult.Channel.Link
		//	userObj.Avatar = rssGetResult.Channel.Image.Url
		//	siteUpdated, _ := pkg.PubDateTimeConvert(rssGetResult.Channel.LastBuildDate)
		//	userObj.UpdateTime = siteUpdated
		//}
		//if atomResult != nil {
		//	userObj.Description = atomResult.Subtitle
		//	userObj.SiteTitle = atomResult.Title
		//	userObj.SiteLink = atomResult.ID
		//	userObj.UserName = atomResult.Author.Name
		//	siteUpdated, _ := pkg.PubDateTimeConvert(atomResult.Updated)
		//	userObj.UpdateTime = siteUpdated
		//}

		if getUser.Id != 0 { //存在
			if result.Properties.UpdateTime.LastEditedTime.After(getUser.UpdateTime) {
				if active == "off" || active == "" {
					userObj.Del = 1
				}

				if active == "on" {
					userObj.Del = 0
				}
				err := userDao.UpdateUser(userObj)
				if err != nil {
					continue
				}
			}
		} else { //不存在
			if err == nil && userObj.RssLink != "" && active == "on" {
				timeNow := time.Now()
				userObj.CreateTime = timeNow
				userObj.UpdateTime = timeNow
				err := userDao.AddUser(userObj)
				if err != nil {
					continue
				}
			}
		}

	}
}

func compareLink(a string, b string) string {
	cStr := "http"
	if strings.Contains(a, cStr) && strings.Contains(b, cStr) {
		if len(a) > len(b) {
			return a
		} else {
			return b
		}
	}
	if strings.Contains(a, cStr) && !strings.Contains(b, cStr) {
		return a
	}
	if !strings.Contains(a, cStr) && strings.Contains(b, cStr) {
		return b
	}
	return b
}
