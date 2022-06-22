package ginex

import (
	"github.com/Merry74751/ginex/middleware"
	"github.com/Merry74751/ginex/router"
	"github.com/gin-gonic/gin"
	"testing"
)

type Test struct {
	getById any `router:"path:/getById,params:id,method:GET"`
}

func (t Test) GetById(c *gin.Context, id string) {
	c.JSON(200, gin.H{
		"message": "请求成功",
		"data":    id,
	})
}

func TestRegisterRouter(t *testing.T) {
	router.AddRouter(Test{})
	app := gin.Default()

	router.RegisterRouter(app)

	err := app.Run()
	if err != nil {
		t.Log(err)
	}
}

func TestName(t *testing.T) {
	middleware.Logger.Info("INFO")

}
