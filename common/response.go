package common

import (
	"github.com/Merry74751/yutool/anyutil"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Res struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type ObjResult struct {
	Res
	Data any `json:"data,omitempty"`
}

type ListResult struct {
	Res
	List  any   `json:"list,omitempty"`
	Total int64 `json:"total,omitempty"`
}

func ListRes(list any, total int64) ListResult {
	return ListResult{
		Res{
			Status:  200,
			Message: "请求成功",
		},
		list,
		total,
	}
}

func ObjRes(data any) ObjResult {
	return ObjResult{
		Res{
			Status:  200,
			Message: "请求成功",
		},
		data,
	}
}

func ConvertResult(v any) any {
	kind := anyutil.RealKind(v)
	if kind == reflect.Struct {
		typ := anyutil.Type(v).String()
		if typ == "common.Res" || typ == "common.ListResult" {
			return v
		}
		return ObjRes(v)
	} else if kind == reflect.String {
		return Res{Status: 200, Message: v.(string)}
	}
	return ObjRes(v)
}

func ConvertHandle(f func(c *gin.Context) any) gin.HandlerFunc {
	return func(c *gin.Context) {
		v := f(c)
		result := ConvertResult(v)
		c.JSONP(200, result)
	}
}
