package timer

import (
	"encoding/json"
	"fmt"
	"github.com/okhanyu/gohelper/gohelper_http"
	"gorm.io/gorm"
	"net/http"
	"rsshub/config"
	memosUserDao "rsshub/dao/memos"
	memosUserModel "rsshub/dao/memos/model"
	timerModel "rsshub/timer/model"
	"strconv"
	"strings"
	"time"
)

const (
	NotionMemosGetTimer = "notion_memos_get_timer"
	MemosDatabasesUrl   = "memos_databases_url"
)

func GetNotionMemosTimer() {
	timeConfig := config.GlobalConfig.System[NotionMemosGetTimer]
	timeConfigInt, err := strconv.Atoi(timeConfig)
	if err != nil {
		timeConfigInt = 1200
		fmt.Printf("无法读取和转换配置中获取NOTION数据的时间，使用降级值：%d", timeConfigInt)
	}
	ticker := time.NewTicker(time.Duration(timeConfigInt) * time.Second)
	go func() {
		for range ticker.C {
			fmt.Printf("[获取Notion Memos定时任务执行开始 %v]\n", time.Now())
			getMemosUserTask()
			fmt.Printf("[获取Notion Memos定时任务执行完毕 %v]\n", time.Now())
		}
	}()
	// time.Sleep(5 * time.Second)
	// ticker.Stop()
	// fmt.Println("定时任务已停止")
}

func getMemosUserTask() {
	notionDatabasesUrl := config.GlobalConfig.Notion[MemosDatabasesUrl]
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

		userObj := memosUserModel.MemosUser{}
		// 获取姓名
		for _, title := range result.Properties.Name.Title {
			if title.Text.Content != "" {
				userObj.UserName = title.Text.Content
			}
		}
		// 获取MEMOS_LINK
		for _, s := range result.Properties.Link.RichText {
			if s.Text.Content != "" || s.Text.Link.Url != "" {
				userObj.MemosLink = compareMemosLink(s.Text.Content, s.Text.Link.Url)
			}
		}
		// 获取Avatar
		for _, s := range result.Properties.Avatar.RichText {
			if s.Text.Content != "" || s.Text.Link.Url != "" {
				userObj.Avatar = compareMemosLink(s.Text.Content, s.Text.Link.Url)
			}
		}

		if userObj.UserName == "" || userObj.MemosLink == "" {
			continue
		}

		getUser, err := memosUserDao.GetMemosUserByMemosLink(memosUserModel.GetMemosReq{MemosLink: userObj.MemosLink})
		if err != nil && err != gorm.ErrRecordNotFound {
			continue
		}

		if getUser.Id != 0 { //存在
			if result.Properties.UpdateTime.LastEditedTime.After(getUser.UpdateTime) {
				if active == "off" || active == "" {
					userObj.Del = 1
				}

				if active == "on" {
					userObj.Del = 0
				}
				err := memosUserDao.UpdateUser(userObj)
				if err != nil {
					continue
				}
			}
		} else { //不存在
			if err == nil && userObj.MemosLink != "" && active == "on" {
				timeNow := time.Now()
				userObj.CreateTime = timeNow
				userObj.UpdateTime = timeNow
				err := memosUserDao.AddMemosUser(userObj)
				if err != nil {
					continue
				}
			}
		}

	}
}

func compareMemosLink(a string, b string) string {
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
