package user

import (
	"fmt"
	"gorm.io/gorm"
	"rsshub/config"
	"rsshub/dao/cons"
	"rsshub/dao/db"
	"rsshub/dao/user/model"
	"time"
)

func GetUserDB() *gorm.DB {
	return db.DB.Table(cons.TableUser)
}

func GetUserDBWithTx() *gorm.DB {
	return db.DB.Table(cons.TableUser).Begin()
}

func GetUserList() ([]model.User, error) {
	var userList []model.User
	err := GetUserDB().Model(&userList).Where(" del = 0 ").Order(" create_time desc ").Find(&userList).Error
	return userList, err
}

func GetUserListByGroup(param model.GetListReq) ([]model.UserStatistics, error) {
	var userList []model.UserStatistics
	selectField := fmt.Sprintf("%s.*, sum(%s.watch) as watch, sum(%s.`like`) as `like`, count(%s.user_id) as total ",
		cons.TableUser,
		cons.TableRss, cons.TableRss, cons.TableRss)
	join := fmt.Sprintf("LEFT JOIN %s ON %s.id = %s.user_id ", cons.TableRss, cons.TableUser, cons.TableRss)
	group := fmt.Sprintf("%s.id", cons.TableUser)
	where := fmt.Sprintf(" %s.del = 0 ", cons.TableUser)
	order := " create_time desc "
	if param.Where == 1 {
		rankTime := config.GlobalConfig.SqlCondition[cons.RankTime]
		where = fmt.Sprintf("%s and pub_date_time >= (NOW() - INTERVAL %s )  ", where, rankTime)
	}
	if param.Order == 1 {
		order = " total desc "
	}
	if param.Order == 2 {
		order = " watch desc "
	}
	if param.Order == 3 {
		order = " `like` desc "
	}
	err := GetUserDB().Model(&userList).Select(selectField).Joins(join).Where(where).
		Order(order).Group(group).Offset(param.PageNum * param.
		PageSize).Limit(param.PageSize).Find(&userList).Error
	return userList, err
}

func GetUserByRssLink(param model.GetUserReq) (model.User, error) {
	var userObj model.User
	err := GetUserDB().Model(&userObj).Order(" create_time desc ").Where(" rss_link = ? ",
		param.RssLink).Find(&userObj).Error
	return userObj, err
}

func AddUser(param model.User) error {
	//userObj, err := GetUserByRssLink(model.GetUserReq{
	//	RssLink: param.RssLink,
	//})
	//if userObj.RssLink == "" || err == gorm.ErrRecordNotFound {
	//	tx := GetUserDBWithTx()
	//	err = tx.Model(&param).Create(&param).Error
	//	if err != nil {
	//		tx.Rollback()
	//	} else {
	//		tx.Commit()
	//	}
	//}
	tx := GetUserDBWithTx()
	err := tx.Model(&param).Create(&param).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}

func UpdateUser(param model.User) error {
	//userObj, err := GetUserByRssLink(model.GetUserReq{
	//	RssLink: param.RssLink,
	//})
	//if userObj.RssLink == "" {
	//	fmt.Println("user信息不存在，无法更新")
	//	return errors.New("user信息不存在，无法更新")
	//}
	tx := GetUserDBWithTx()
	err := tx.Model(&param).Where("rss_link = ? ", param.RssLink).Updates(map[string]interface{}{
		"user_name": param.UserName,
		//"description": param.Description,
		"update_time": time.Now(),
		"avatar":      param.Avatar,
		//"site_title":  param.SiteTitle,
		"del": param.Del,
	}).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}
