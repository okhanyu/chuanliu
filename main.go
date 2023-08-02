package main

import (
	"log"
	"rsshub/api"
	"rsshub/config"
	"rsshub/dao/db"
	"rsshub/timer"
	"strconv"
)

var dbDsn = "./db/rss.db"

func main() {
	config.InitConfig()
	dbDsn = config.GlobalConfig.System["db"]
	log.Printf("db: %s", dbDsn)
	db.InitDB(dbDsn)
	timer.GetNotionTimer()
	timer.GetRssTimer()
	timer.GetNotionMemosTimer()
	timer.LenContent, _ = strconv.Atoi(config.GlobalConfig.System["len"])
	if timer.LenContent == 0 {
		timer.LenContent = 500
	}
	api.StartHttpServer()
}
