package memos

import (
	"gorm.io/gorm"
	"rsshub/dao/cons"
	"rsshub/dao/db"
	"rsshub/dao/memos/model"
	"time"
)

func GetMemosUserDB() *gorm.DB {
	return db.DB.Table(cons.TableMemos)
}

func GetMemosUserDBWithTx() *gorm.DB {
	return db.DB.Table(cons.TableMemos).Begin()
}

func GetMemosUserByMemosLink(param model.GetMemosReq) (model.MemosUser, error) {
	var memosUser model.MemosUser
	err := GetMemosUserDB().Model(&memosUser).Order(" create_time desc ").Where(" memos_link = ? ",
		param.MemosLink).Find(&memosUser).Error
	return memosUser, err
}

func GetMemosUserList() ([]model.MemosUser, error) {
	var memosUserList []model.MemosUser
	err := GetMemosUserDB().Model(&memosUserList).Where(" del = 0 ").Order(" create_time desc ").
		Find(&memosUserList).Error
	return memosUserList, err
}

func AddMemosUser(param model.MemosUser) error {
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
	tx := GetMemosUserDBWithTx()
	err := tx.Model(&param).Create(&param).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}

func UpdateUser(param model.MemosUser) error {
	//userObj, err := GetUserByRssLink(model.GetUserReq{
	//	RssLink: param.RssLink,
	//})
	//if userObj.RssLink == "" {
	//	fmt.Println("user信息不存在，无法更新")
	//	return errors.New("user信息不存在，无法更新")
	//}
	tx := GetMemosUserDBWithTx()
	err := tx.Model(&param).Where("memos_link = ? ", param.MemosLink).Updates(map[string]interface{}{
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
