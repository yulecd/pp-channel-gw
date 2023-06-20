package main

import (
	"fmt"
	"npp-channel-gw/app/pkg/db"
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

	db.InitMysql()
	db.InitRedis()

	// register route and init server
	router := router.InitRouter()
	srv := server.NewServer(router)

	srv.Run()
}
