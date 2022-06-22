package ginex

import (
	"github.com/Merry74751/ginex/router"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRegister(t *testing.T) {
	app := gin.Default()
	router.AddRouter(Test{})
	router.RegisterRouter(app)

	err := app.Run()
	if err != nil {
		t.Log(err)
	}
}

type Test struct {
}

func (t Test) GetById() (string, string, router.Handle) {
	return "GET", "/getById", func(c *gin.Context) (any, error) {
		query := c.Query("id")
		return query, nil
	}
}

func (t Test) GroupName() string {
	return "user"
}
