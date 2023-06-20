package db

import (
	"fmt"

	"github.com/yulecd/pp-common/config"
	mysql "github.com/yulecd/pp-common/db"

	"gorm.io/gorm"
)

var BaseMysql *gorm.DB // 旧数据库 pay_server

func InitMysql() {
	var err error
	var mysqlConf map[string]mysql.MysqlConf

	if err = config.Load("mysql", &mysqlConf); err != nil {
		panic(fmt.Sprintf("load mysql config failed, err: %v", err))
	}

	for dbName, dbConf := range mysqlConf {
		switch dbName {
		case "base":
			BaseMysql, err = mysql.InitMysqlClient(dbConf)
		}

		if err != nil {
			panic("mysql connect error: %v" + err.Error())
		}
	}
}
