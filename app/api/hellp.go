package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/richkeyu/gocommons/util/app"
	"github.com/yulecd/pp-common/plog"
	"github.com/yulecd/pp-common/server"
)

type BaseApi struct{}

func (b BaseApi) Greeter(c *gin.Context) {
	name, found := c.Get("name")
	if !found {
		name = "world"
	}

	ctx := server.NewContext(context.Background(), c)

	plog.Info(ctx, "Greeter")
	data := "hello " + name.(string)

	app.Response(c, http.StatusOK, 1, data)
}
