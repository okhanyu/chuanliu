package rss

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"math"
	"rsshub/config"
	"rsshub/dao/cons"
	"rsshub/dao/db"
	"rsshub/dao/rss/model"
	"rsshub/pkg"
	"strconv"
	"strings"
	"time"
)

func GetRssDB() *gorm.DB {
	return db.DB.Table(cons.TableRss)
}

func GetRssDBWithTx() *gorm.DB {
	return db.DB.Table(cons.TableRss).Begin()
}

func GetTags(_ *gin.Context, param model.GetTagsReq) ([]string, error) {
	var tags []string
	where := fmt.Sprintf(" %s.del = 0 and %s.user_id in ( select id as user_id from %s where del = 0 )  ",
		cons.TableRss, cons.TableRss, cons.TableUser)
	err := GetRssDB().Select("tags").Offset(param.PageNum * param.PageSize).Limit(param.PageSize).
		Where(where).Find(&tags).Error
	return tags, err
}

// GetRssList 获取RSS内容列表
func GetRssList(_ *gin.Context, param model.GetRssListReq) ([]model.GetRss, error) {
	var rssList []model.GetRss

	if param.PageSize > 50 {
		param.PageSize = 50
	}

	order := " pub_date_time desc "
	where := fmt.Sprintf(" %s.del = 0 and %s.user_id in ( select id as user_id from %s where del = 0 )  ",
		cons.TableRss, cons.TableRss, cons.TableUser)
	// where := fmt.Sprintf(" %s.del = 0 and %s.del = 0", TableRss, user.TableUser)
	timeCondition := config.GlobalConfig.SqlCondition[cons.RandsTime]
	randsCount := config.GlobalConfig.SqlCondition[cons.RandsCount]
	randsCountInt, err := strconv.Atoi(randsCount)
	if err != nil {
		log.Println(err)
		randsCountInt = 3
		log.Println("randsCountInt降级为3")
	}
	// 按照观看量、发布时间倒序
	if param.Order == 1 {
		order = fmt.Sprintf("  watch DESC, %s ", order)
	}

	// 按照喜欢、发布时间倒序
	if param.Order == 2 {
		order = fmt.Sprintf("  `like` DESC, %s ", order)
	}

	// 按照观看量、喜欢、发布时间倒序
	if param.Order == 3 {
		order = fmt.Sprintf("   watch DESC, `like` DESC, %s ", order)
	}

	// 按照喜欢、观看量、发布时间倒序
	if param.Order == 4 {
		order = fmt.Sprintf("  `like` DESC,  watch DESC, %s ", order)
	}

	// 查找除热点外的其他文章
	if param.Where == 1 {
		where = fmt.Sprintf(" %s and %s.id not in (SELECT * FROM (SELECT id FROM `%s` WHERE   del = 0  and "+
			"pub_date_time >= (NOW() - INTERVAL %s) ORDER BY  watch DESC,  "+
			"pub_date_time desc   LIMIT %d ) AS subquery )",
			where, cons.TableRss, cons.TableRss, timeCondition, randsCountInt)
	}
	// 查找热点文章
	if param.Where == 2 {
		where = fmt.Sprintf(" %s and pub_date_time >= (NOW() - INTERVAL %s) ", where, timeCondition)
	}

	// 查找有头图的文章
	if param.Where == 3 {
		where = fmt.Sprintf(" %s and cover != '' ", where)
	}

	// 查找全部文章
	if param.Where == 4 {

	}

	if param.Tag != "" {
		where = fmt.Sprintf(" %s and tags LIKE '%%%s%%' ", where, param.Tag)
	}
	// join := fmt.Sprintf("JOIN %s ON %s.user_id = %s.id", user.TableUser, TableRss, user.TableUser)
	// selectField := fmt.Sprintf("%s.*,%s.id as userId,%s.del as userDel", TableRss, user.TableUser, user.TableUser)

	// join联动查头像信息
	selectField := fmt.Sprintf("%s.*,%s.avatar", cons.TableRss, cons.TableUser)
	join := fmt.Sprintf("LEFT JOIN %s ON %s.user_id = %s.id", cons.TableUser, cons.TableRss, cons.TableUser)

	err = GetRssDB().Model(&rssList).Select(selectField).Joins(join).Order(order).Offset(param.PageNum * param.
		PageSize).Limit(param.PageSize).Where(where).Find(&rssList).Error
	return rssList, err
}

// Watch 增加观看
func Watch(_ *gin.Context, param int) error {
	tx := GetRssDBWithTx()
	err := tx.Model(&param).Where("id = ? and watch < ? ", param, math.MaxInt).
		UpdateColumn("watch", gorm.Expr("watch + ?", 1)).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}

// Like 增加喜欢
func Like(_ *gin.Context, param int) error {
	tx := GetRssDBWithTx()
	err := tx.Model(&param).Where("id = ? and `like` < ? ", param, math.MaxInt).
		UpdateColumn("like", gorm.Expr("`like` + ?", 1)).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}

// AddRss 增加RSS记录
func AddRss(param model.Rss) error {

	/*
		SELECT
		    substr(link, 1, instr(link, '://') - 1) AS part1,
		    substr(link, instr(link, '://') + length('://')) AS part2
		FROM
		    rss_data;

		SELECT
	*/
	param.Summary = pkg.HandleUTF8(param.Summary)
	param.Content = pkg.HandleUTF8(param.Content)
	param.UserName = pkg.HandleUTF8(param.UserName)

	var rssObj model.Rss
	param.Link = strings.TrimSpace(param.Link)
	err := GetRssDB().Model(&rssObj).Where(" substr(link, instr(link, "+
		"'://') + length('://')) = ? ", strings.Split(param.Link, "://")[1]).Find(&rssObj).Error
	if rssObj.Link == "" || err == gorm.ErrRecordNotFound {
		tx := GetRssDBWithTx()
		err = tx.Model(&param).Create(&param).Error
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	} else if param.Cover != "" && param.Cover != rssObj.Cover {
		txUpdate := GetRssDBWithTx()
		errUpdate := txUpdate.Model(&param).Where("link = ? ", param.Link).Updates(map[string]interface{}{
			"update_time": time.Now(),
			"cover":       param.Cover,
		}).Error
		if errUpdate != nil {
			txUpdate.Rollback()
		} else {
			txUpdate.Commit()
		}
	}
	return err
}

// GetUserRecentArticleListByGroup 每个用户最近3条
func GetUserRecentArticleListByGroup(_ *gin.Context, param model.GetRssListReq) ([]model.UserRecentArticle, error) {
	var userList []model.UserRecentArticle

	//sql := fmt.Sprintf("SELECT id,%s.user_id,user_name,title,link, %s.pub_date_time,"+
	//	"min_pub_date_time  FROM %s LEFT JOIN ( SELECT t.user_id, "+
	//	"( SELECT pub_date_time FROM %s WHERE user_id = t."+
	//	"user_id ORDER BY pub_date_time DESC LIMIT 1 OFFSET 2 ) AS min_pub_date_time FROM ("+
	//	"SELECT user_id AS user_id  FROM %s GROUP BY user_id "+
	//	") AS t ) AS jt ON jt.user_id = %s.user_id where %s.pub_date_time >= jt.min_pub_date_time  and %s.del != 1",
	//	cons.TableRss, cons.TableRss, cons.TableRss, cons.TableRss, cons.TableRss,
	//	cons.TableRss, cons.TableRss, cons.TableRss)

	sql := fmt.Sprintf("SELECT subquery.`id`,subquery.`user_id`,subquery.`title`,subquery.`link`,"+
		"subquery.`pub_date_time`, user.user_name as `user_name` FROM ( SELECT *, "+
		"ROW_NUMBER() OVER ( PARTITION BY user_id ORDER BY pub_date_time desc ) AS row_num FROM %s WHERE del = 0 "+
		") AS subquery LEFT JOIN %s AS user ON user.`id` = subquery.user_id and user."+
		"del = 0 WHERE row_num <= 3 and user.del = 0", cons.TableRss, cons.TableUser)
	//  order by  user.`sort` asc , user.`id` asc

	err := GetRssDB().Raw(sql).Scan(&userList).Error

	return userList, err
}
