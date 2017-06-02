package models

import (
	"time"
	"os"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	. "github.com/yikeso/goblog/common"
	_ "github.com/mattn/go-sqlite3"
	"path"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              Long
	Title           string
	Created         time.Time        `orm:"index"`
	Views           Long             `orm:"inxex"`
	TopicTime       time.Time        `orm:"inxex"`
	TopicCount      Long
	TopicLastUserId Long
}

type Topic struct {
	Id               Long
	Uid              Long
	Title            string
	Content          string `orm:"size(5000)"`
	Attachment       string
	Created          time.Time        `orm:"inxex"`
	Updated          time.Time        `orm:"inxex"`
	Views            Long        `orm:"inxex"`
	Author           string
	ReplyTime        time.Time        `orm:"inxex"`
	ReplyCount       Long
	RepleyLastUserId Long
}

/**
 * 注册数据库
 */
func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)//创建文件夹
		os.Create(_DB_NAME)//创建文件
	}
	orm.RegisterModel(new(Category),new(Topic))//注册表结构实体
	orm.RegisterDriver(_SQLITE3_DRIVER,orm.DRSqlite)//注册数据库类型，驱动
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)//设置默认数据库
}
