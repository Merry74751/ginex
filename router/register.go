package router

import (
	"github.com/Merry74751/yutool/anyutil"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
	"strings"
)

var router = make([]any, 0)

func AddRouter(v any) {
	if anyutil.IsStruct(v) {
		router = append(router, v)
	}
}

func RegisterRouter(app *gin.Engine) {
	if len(router) == 0 {
		return
	}
	for _, route := range router {
		v := anyutil.Value(route)
		typ := anyutil.Type(route)

		numMethod := v.NumMethod()
		for i := 0; i < numMethod; i++ {
			m := v.Method(i)
			mty := typ.Method(i)

			mName := mty.Name
			mName = firstToLower(mName)

			field, exist := typ.FieldByName(mName)
			if !exist {
				log.Printf("%T are not have field: %s", v, mName)
			}
			tag := field.Tag.Get("router")
			tags := strings.Split(tag, ",")

			var (
				path   string
				params []string
				method string
			)

			for _, t := range tags {
				kv := strings.Split(t, ":")
				k := kv[0]
				v := kv[1]
				if k == "path" {
					path = v
				} else if k == "method" {
					method = v
				} else if k == "params" {
					vs := strings.Split(v, ",")
					params = make([]string, len(vs))
					for i, v := range vs {
						params[i] = v
					}
				}
			}

			vs := make([]reflect.Value, len(params)+1)
			if method == "GET" {
				app.Handle(method, path, func(c *gin.Context) {
					vs[0] = reflect.ValueOf(c)
					for i, param := range params {
						q := c.Query(param)
						vs[i+1] = anyutil.Value(q)
					}
					m.Call(vs)
				})
				return
			}
			// todo post、delete、put请求处理，处理路径变量
			// 新思路 用结构体保存 各个参数：
			// 1、请求路径
			// 2、请求方式
			// 3、路径变量
			// 4、请求体或GET参数或POST表单参数
		}
	}
}

func firstToLower(s string) string {
	bytes := []byte(s)
	bytes[0] = bytes[0] + 32
	return string(bytes)
}
