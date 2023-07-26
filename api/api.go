package api

import (
	"github.com/gin-gonic/gin"
	"github.com/okhanyu/gohelper/gohelper_server"
	"rsshub/api/memos"
	"rsshub/api/rss"
	"rsshub/api/user"
	"rsshub/config"
)

func StartHttpServer() {

	prefix := "api/rss-hub/"

	Gets := make([]*gohelper_server.Router, 0)
	Gets = append(Gets, gohelper_server.NewRouter(prefix, "rss/list", []gin.HandlerFunc{rss.ListRss}))
	Gets = append(Gets, gohelper_server.NewRouter(prefix, "user/list/group", []gin.HandlerFunc{user.ListByGroup}))
	Gets = append(Gets, gohelper_server.NewRouter(prefix, "user/list", []gin.HandlerFunc{user.List}))
	Gets = append(Gets, gohelper_server.NewRouter(prefix, "memos/user/list", []gin.HandlerFunc{memos.List}))
	Gets = append(Gets, gohelper_server.NewRouter(prefix, "rss/tag/list", []gin.HandlerFunc{rss.ListRssTags}))
	Gets = append(Gets, gohelper_server.NewRouter(prefix, "rss/user/list/recent",
		[]gin.HandlerFunc{rss.ListRssUserRecent}))

	Posts := make([]*gohelper_server.Router, 0)
	Posts = append(Posts, gohelper_server.NewRouter(prefix, "rss/watch", []gin.HandlerFunc{rss.WatchArticle}))

	server := gohelper_server.GetServerInstance()
	server.BuildGet(Gets).BuildPost(Posts).BuildUsesFunc(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "*")
	}).Build(config.GlobalConfig.System["port"])
}
