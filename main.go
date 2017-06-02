package main

import (
	_ "github.com/yikeso/goblog/routers"//下划线执行该包的初始化方法
	"github.com/astaxie/beego"
	"github.com/yikeso/goblog/models"
	"github.com/astaxie/beego/orm"
)
/**
 * 初始化方法，初始化数据库连接
 */
func init(){
	models.RegisterDB()
}

func main() {
	//持久层框架设置为debug模式
	orm.Debug = true
	//在default库中建表，force是否强制建表，启动时，即使有表也会删除重建
	//verbose是否打印相关信息
	orm.RunSyncdb("default",false,true)
	//日志打印级别为info
	beego.SetLevel(beego.LevelInformational)
	beego.Run()
}

