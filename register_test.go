package ginex

import (
	"github.com/Merry74751/ginex/common"
	"github.com/Merry74751/ginex/middleware"
	"github.com/Merry74751/ginex/router"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	v := common.Res{}
	ty := reflect.TypeOf(v)
	t.Log(ty.PkgPath())
}

func TestRegister(t *testing.T) {
	app := gin.Default()
	router.AddRouter(Test{})
	router.RegisterRouter(app)

	err := app.Run()
	if err != nil {
		t.Log(err)
	}
}

func Register2(app *gin.Engine) {
	test := Test{}
	v := reflect.ValueOf(test)
	method := v.MethodByName("Insert")
	typ := method.Type().In(0)
	value := reflect.New(typ)
	bv := value.Interface()

	app.POST("/test", func(c *gin.Context) {
		err := c.BindJSON(bv)
		if err != nil {
			log.Println("bind error:", err)
		}
		bvv := reflect.ValueOf(bv).Elem()
		values := method.Call([]reflect.Value{bvv})
		values[0].Interface().(gin.HandlerFunc)(c)
	})
}

func TestRegister2(t *testing.T) {
	app := gin.Default()
	Register2(app)

	err := app.Run()
	if err != nil {
		t.Log(err)
	}
}

func TestLog(t *testing.T) {
	middleware.Logger.Info("error")
}

type Test struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func (receiver Test) GroupName() string {
	return "test"
}

func (receiver Test) Test() (string, string, router.Handle) {
	return "GET", "getById", func(c *gin.Context) any {
		return "hello"
	}
}

func (t Test) Insert(test Test) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSONP(200, test)
	}
}
