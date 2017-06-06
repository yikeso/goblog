package models

import (
	"time"
	"os"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"path"
	"errors"
	"fmt"
	"strconv"
	"database/sql"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64           `orm:"auto"`
	Title           string
	Created         time.Time        `orm:"index"`
	Views           sql.NullInt64             `orm:"index"`
	TopicTime       time.Time        `orm:"index"`
	TopicCount      sql.NullInt64
	TopicLastUserId sql.NullInt64
}

type Topic struct {
	Id               int64            `orm:"auto"`
	Uid              sql.NullInt64
	Title            string
	Content          string `	orm:"size(5000)"`
	Attachment       string
	Created          time.Time        `orm:"index"`
	Updated          time.Time        `orm:"index"`
	Views            sql.NullInt64        `	orm:"index"`
	Author           string
	ReplyTime        time.Time        `orm:"index"`
	ReplyCount       sql.NullInt64
	RepleyLastUserId sql.NullInt64
}

/**
 * 注册数据库
 */
func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm) //创建文件夹
		os.Create(_DB_NAME)                          //创建文件
	}
	orm.RegisterModel(new(Category), new(Topic))                   //注册表结构实体
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)              //注册数据库类型，驱动
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10) //设置默认数据库
}

/**
 * 插入类型
 */
func AddCategory(name string) error {
	o := orm.NewOrm()                                                          //获取orm结构
	cate := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()} //创建插入结构
	qs := o.QueryTable("category")                                             //获取查询表结构
	err := qs.Filter("title", name).One(cate)                                  //获取一个title为name的cate结构存入cate
	if err == nil { //没有错误，查询成功，
		return errors.New("类型名重复,插入失败")
	}
	_, err = o.Insert(cate)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 查出所有的类别
 */
func FindAllCategory() (categories []*Category, err error) {
	o := orm.NewOrm() //获取orm结构
	categories = make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err = qs.All(&categories)
	return
}

/**
 * 根据id删除类型
 */
func DeleteById(id string) (err error) {
	n, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		o := orm.NewOrm()
		cate := &Category{Id: n}
		_, err = o.Delete(cate)
	}
	return
}
