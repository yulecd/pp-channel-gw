package main

import (
	"fmt"
	"npp-channel-gw/router"
	"os"

	"github.com/yulecd/pp-common/client"
	"github.com/yulecd/pp-common/config"
	"github.com/yulecd/pp-common/plog"
	"github.com/yulecd/pp-common/server"
	"github.com/yulecd/pp-common/util"
	"github.com/yulecd/pp-common/wrapper"
)

func main() {

	util.ChangeTimeToUTC()

	if _, err := config.NewConfig(os.Getenv(config.AppEnvName)); err != nil {
		panic(fmt.Sprintf("init config failed, err: %v", err))
	}

	plog.InitWithPath("log", "prod")

	client.AddDefaultWrappers(wrapper.HttpClientTrace)

	// db.InitMysql()
	// db.InitRedis()
	// db.InitLocalCache()
	// middleware.InitValidatorTranslate()
	// snow.InitSnowFlake()
	// conf.InitConfig()
	// ip.Init()
	// bigquery.InitBigQuery()
	//queue.InitGProducerConf() //发MQ消息必须先加载配置

	// register route and init server
	router := router.InitRouter()
	srv := server.NewServer(router)

	srv.Run()
}
