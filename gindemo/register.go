/*
 * @Author: tj
 * @Date: 2022-11-02 17:40:48
 * @LastEditors: tj
 * @LastEditTime: 2022-12-15 09:50:22
 * @FilePath: \createApiMarkdown\gindemo\register.go
 */
package gindemo

import (
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	pathPrefix string = "/"
)

var (
	HasLogin bool = false
)

// 注册路由
func register(controller interface{}) error {
	ctrlName := reflect.TypeOf(controller).String()

	module := ctrlName
	if strings.Contains(ctrlName, ".") {
		module = strings.ToLower(ctrlName[strings.Index(ctrlName, ".")+1:])
	}

	v := reflect.ValueOf(controller)
	for i := 0; i < v.NumMethod(); i++ {
		action := strings.ToLower(v.Type().Method(i).Name)
		path, err := url.JoinPath(pathPrefix, module, action)
		if err != nil {
			log.Errorln("register JoinPath error:", err)
			return err
		}

		httpMethod := http.MethodPost
		methodStr := checkMethod(action)
		if methodStr != "" {
			httpMethod = methodStr
		}

		method := v.Method(i)
		request, subRequestMap := getRequest(method)
		response, subResponse := getResponse(method)

		route := Route{
			Title:       module + "-" + action,
			Author:      "",
			Path:        path,
			HttpMethod:  httpMethod,
			Method:      method,
			Request:     request,
			SubRequest:  subRequestMap,
			Response:    response,
			SubResponse: subResponse,
		}

		Routes = append(Routes, route)
	}

	return nil
}

func checkMethod(action string) string {
	if strings.Contains(strings.ToUpper(action), http.MethodGet) {
		return http.MethodGet
	}

	if strings.Contains(strings.ToUpper(action), http.MethodPut) {
		return http.MethodPut
	}

	if strings.Contains(strings.ToUpper(action), http.MethodDelete) {
		return http.MethodDelete
	}

	if strings.Contains(strings.ToUpper(action), http.MethodPatch) {
		return http.MethodPatch
	}

	return ""
}

func getRequest(method reflect.Value) (interface{}, interface{}) {
	var request interface{} = nil
	subRequestMap := make(map[int]interface{}, 0)

	for j := 0; j < method.Type().NumIn(); j++ {
		reqType := method.Type().In(j)

		if reqType == reflect.TypeOf(&gin.Context{}) {
			continue
		}

		if reqType.Kind() == reflect.Ptr && reqType.Elem().Kind() == reflect.Struct {
			reqValue := reflect.New(reqType.Elem())

			for i := 0; i < reqType.Elem().NumField(); i++ {
				fieldValue := reqType.Elem().Field(i)

				switch fieldValue.Type.Kind() {
				case reflect.Slice:
					// silce with elem
					slice := reflect.MakeSlice(fieldValue.Type, 0, 4)
					subResponseValue := reflect.New(fieldValue.Type.Elem())
					slice = reflect.Append(slice, reflect.ValueOf(subResponseValue.Elem().Interface()))
					subRequestMap[i] = slice.Interface()

				case reflect.Struct:
					subResponseValue := reflect.New(fieldValue.Type)
					subRequestMap[i] = subResponseValue.Interface()
				}
			}
			request = reqValue.Interface()
		}

		if reqType.Kind() == reflect.Struct {
			request = reflect.New(reqType).Interface()

			for i := 0; i < reqType.NumField(); i++ {
				fieldValue := reqType.Field(i)

				switch fieldValue.Type.Kind() {
				case reflect.Slice:
					// slice elem
					subResponseValue := reflect.New(fieldValue.Type.Elem())
					subRequestMap[i] = subResponseValue.Interface()

					// slice
					subRequestMap[i] = reflect.New(fieldValue.Type).Interface()
				case reflect.Struct:
					subResponseValue := reflect.New(fieldValue.Type)
					subRequestMap[i] = subResponseValue.Interface()
				}
			}
		}
	}

	if len(subRequestMap) == 0 {
		return request, nil
	}

	return request, subRequestMap
}

func getResponse(method reflect.Value) (interface{}, interface{}) {
	var response interface{} = nil
	subResponseMap := make(map[int]interface{}, 0)

	for j := 0; j < method.Type().NumOut(); j++ {
		rspType := method.Type().Out(j)

		if rspType.Kind() == reflect.Ptr && rspType.Elem().Kind() == reflect.Struct {
			rspValue := reflect.New(rspType.Elem())

			for i := 0; i < rspType.Elem().NumField(); i++ {
				fieldValue := rspType.Elem().Field(i)

				switch fieldValue.Type.Kind() {
				case reflect.Slice:
					// slice elem
					// subResponseValue := reflect.New(fieldValue.Type.Elem())
					// subResponseMap[i] = subResponseValue.Interface()

					// silce with elem
					slice := reflect.MakeSlice(fieldValue.Type, 0, 4)
					subResponseValue := reflect.New(fieldValue.Type.Elem())
					slice = reflect.Append(slice, reflect.ValueOf(subResponseValue.Elem().Interface()))
					subResponseMap[i] = slice.Interface()

					// slice
					// subResponseMap[i] = reflect.New(fieldValue.Type).Elem().Interface()
				case reflect.Struct:
					subResponseValue := reflect.New(fieldValue.Type)
					subResponseMap[i] = subResponseValue.Interface()
				}
			}
			response = rspValue.Interface()
		}

		if rspType.Kind() == reflect.Struct {
			response = reflect.New(rspType).Interface()

			for i := 0; i < rspType.NumField(); i++ {
				fieldValue := rspType.Field(i)

				switch fieldValue.Type.Kind() {
				case reflect.Slice:
					// slice elem
					subResponseValue := reflect.New(fieldValue.Type.Elem())
					subResponseMap[i] = subResponseValue.Interface()

					// slice
					subResponseMap[i] = reflect.New(fieldValue.Type).Interface()
				case reflect.Struct:
					subResponseValue := reflect.New(fieldValue.Type)
					subResponseMap[i] = subResponseValue.Interface()
				}
			}
		}
	}

	if len(subResponseMap) == 0 {
		return response, nil
	}

	return response, subResponseMap
}
