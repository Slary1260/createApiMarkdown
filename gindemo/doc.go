/*
 * @Author: tj
 * @Date: 2022-11-03 18:59:04
 * @LastEditors: tj
 * @LastEditTime: 2022-11-07 14:30:43
 * @FilePath: \createApiMarkdown\gindemo\doc.go
 */
package gindemo

import (
	"reflect"

	"github.com/Slary1260/createapimarkdown/document"
	"github.com/Slary1260/createapimarkdown/markdown"
)

func getDoc() (*document.Document, error) {
	doc := document.NewDocument("")

	for _, v := range Routes {
		item := &document.DocItem{
			Title:    "",
			Url:      v.Path,
			Method:   v.HttpMethod,
			Author:   "",
			Request:  v.Request,
			Response: &Result{Data: v.Response},
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

func createApiMd() error {
	doc, err := getDoc()
	if err != nil {
		log.Errorln("createApiMd getDoc error:", err)
		return err
	}

	md := markdown.New(doc, markdown.WithMd2Html(true))
	err = md.Generate("doc.md")
	if err != nil {
		log.Errorln("createApiMd Generate error:", err)
		return err
	}

	return nil
}
