package router

import (
	"github.com/Merry74751/ginex/common"
	"github.com/Merry74751/yutool/anyutil"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

type Handle func(c *gin.Context) (any, error)

type Group interface {
	GroupName() string
}

var router = make([]any, 0)

func AddRouter(r any) {
	router = append(router, r)
}

func RegisterRouter(app *gin.Engine) {
	for _, route := range router {
		var routerGroup *gin.RouterGroup
		group, ok := route.(Group)
		if ok {
			routerGroup = app.Group(group.GroupName())
		}

		methods := anyutil.Methods(route)
		for _, method := range methods {
			if isHandle(method) {
				results := method.Call(nil)
				httpMethod := results[0].Interface().(string)
				path := results[1].Interface().(string)
				h := results[2].Interface().(Handle)
				if ok {
					routerGroup.Handle(httpMethod, path, common.ConvertHandle(h))
					continue
				}
				app.Handle(httpMethod, path, common.ConvertHandle(h))
			}
		}
	}
}

func isHandle(method reflect.Value) bool {
	mt := method.Type().String()
	s := strings.SplitN(mt, " ", 2)[1]
	if s == "(string, string, router.Handle)" {
		return true
	}
	return false
}
