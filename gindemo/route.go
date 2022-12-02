/*
 * @Author: tj
 * @Date: 2022-11-03 10:30:39
 * @LastEditors: tj
 * @LastEditTime: 2022-12-02 10:41:15
 * @FilePath: \createApiMarkdown\gindemo\route.go
 */
package gindemo

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	Routes []Route
)

type Route struct {
	Title       string
	Author      string
	Path        string
	HttpMethod  string
	Method      reflect.Value
	Request     interface{}
	SubRequest  interface{}
	Response    interface{}
	SubResponse interface{}
}

func initRouter(e *gin.Engine, isMd2Html bool) error {
	group := e.Group("api")

	for _, v := range Routes {
		switch v.HttpMethod {
		case http.MethodPost:
			group.POST(v.Path, match(v.Path, v))
		case http.MethodGet:
			group.GET(v.Path, match(v.Path, v))
		case http.MethodPut:
			group.PUT(v.Path, match(v.Path, v))
		case http.MethodDelete:
			group.DELETE(v.Path, match(v.Path, v))
		case http.MethodPatch:
			group.PATCH(v.Path, match(v.Path, v))
		default:
			log.Errorln("initRouter unsupported HttpMethod:", v.HttpMethod)
			return errors.New("unsupported HttpMethod")
		}
	}

	err := createApiMd(isMd2Html)
	if err != nil {
		log.Errorln("initRouter createApiMd error:", err)
		return err
	}

	if isMd2Html {
		e.LoadHTMLGlob("./doc.html")
	}

	return nil
}

func match(path string, r Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(Routes) <= 0 {
			return
		}

		// e.g. path:/user/list
		fields := strings.Split(path, "/")
		if len(fields) < 3 {
			return
		}

		arguments := make([]reflect.Value, 0, 1)
		if r.Request != nil {
			arguments = append(arguments, reflect.ValueOf(r.Request))
		}
		arguments = append(arguments, reflect.ValueOf(c))
		r.Method.Call(arguments)
	}
}

// 参考 https://juejin.cn/post/6844904033551908871
func AutoBindWrap(ctrFunc interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctrType := reflect.TypeOf(ctrFunc)
		ctrValue := reflect.ValueOf(ctrFunc)
		// 1. check
		if ctrType.Kind() != reflect.Func {
			panic("not support")
		}

		numIn := ctrType.NumIn()
		if numIn != 2 {
			panic("not support")
		}

		// 2. bind value
		ctrParams := make([]reflect.Value, numIn)
		for i := 0; i < numIn; i++ {
			pt := ctrType.In(i)
			// handle gin.Context
			if pt == reflect.TypeOf(&gin.Context{}) {
				ctrParams[i] = reflect.ValueOf(c)
				continue
			}
			// handle params
			if pt.Kind() == reflect.Ptr && pt.Elem().Kind() == reflect.Struct {
				pv := reflect.New(pt.Elem()).Interface()
				var err error
				switch c.Request.Method {
				case http.MethodGet, http.MethodDelete:
					err = c.ShouldBindQuery(pv)
				case http.MethodPost, http.MethodPut:
					err = c.ShouldBindJSON(pv)
				}
				if err != nil {
					panic(err)
				}

				ctrParams[i] = reflect.ValueOf(pv)
			}
		}
		// 3. call controller
		ctrValue.Call(ctrParams)
	}
}
