/*
 * @Author: tj
 * @Date: 2022-11-03 18:59:04
 * @LastEditors: tj
 * @LastEditTime: 2022-12-02 10:41:46
 * @FilePath: \createApiMarkdown\gindemo\doc.go
 */
package gindemo

import (
	"net/http"
	"reflect"

	"github.com/Slary1260/createapimarkdown/document"
	"github.com/Slary1260/createapimarkdown/markdown"
)

const (
	apiDocFile string = "doc.md"
	apiDocHtml string = "doc.html"
)

func createApiMd(isMd2Html bool) error {
	doc, err := getDoc()
	if err != nil {
		log.Errorln("createApiMd getDoc error:", err)
		return err
	}

	md := markdown.New(doc, markdown.WithMd2Html(isMd2Html))
	err = md.Generate(apiDocFile)
	if err != nil {
		log.Errorln("createApiMd Generate error:", err)
		return err
	}

	return nil
}

func getHeader() map[string]string {
	headerMap := make(map[string]string, 8)
	headerMap["timestamp"] = ""
	headerMap["randStr"] = "length between 10 and 16"

	return headerMap
}

func getDoc() (*document.Document, error) {
	doc := document.NewDocument("", getHeader())

	for _, v := range Routes {
		item := &document.DocItem{
			Title:   v.Title,
			Url:     v.Path,
			Method:  v.HttpMethod,
			Author:  v.Author,
			Request: v.Request,
		}

		if item.Method == http.MethodPost {
			item.Response = &Result{Data: v.Response}
		}

		if v.SubRequest != nil {
			reqType := reflect.TypeOf(v.Request)
			reqValue := reflect.ValueOf(v.Request).Elem()

			if reqType.Kind() == reflect.Ptr {
				reqType = reqType.Elem()
			}

			for key, detail := range v.SubRequest.(map[int]interface{}) {
				for i := 0; i < reqType.NumField(); i++ {
					fieldType := reqType.Field(i)

					if key == i && fieldType.Type == reflect.TypeOf(detail) {
						// rspValue.FieldByName(fieldType.Name).Set(reflect.ValueOf(detail))
						reqValue.FieldByName(fieldType.Name).Set(reflect.ValueOf(detail).Convert(fieldType.Type))
					}
				}
			}
		}

		if v.SubResponse != nil {
			rspType := reflect.TypeOf(v.Response)
			rspValue := reflect.ValueOf(v.Response).Elem()

			if rspType.Kind() == reflect.Ptr {
				rspType = rspType.Elem()
			}

			for key, detail := range v.SubResponse.(map[int]interface{}) {
				for i := 0; i < rspType.NumField(); i++ {
					fieldType := rspType.Field(i)

					if key == i && fieldType.Type == reflect.TypeOf(detail) {
						// rspValue.FieldByName(fieldType.Name).Set(reflect.ValueOf(detail))
						rspValue.FieldByName(fieldType.Name).Set(reflect.ValueOf(detail).Convert(fieldType.Type))
					}
				}
			}
		}

		err := doc.AddDocItem(item)
		if err != nil {
			return nil, err
		}
	}

	return doc, nil
}
