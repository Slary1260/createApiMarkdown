/*
 * @Author: tj
 * @Date: 2022-11-22 20:05:33
 * @LastEditors: tj
 * @LastEditTime: 2022-12-05 10:17:35
 * @FilePath: \createApiMarkdown\manual.go
 */
package main

import (
	"github.com/Slary1260/createapimarkdown/document"
	"github.com/Slary1260/createapimarkdown/logger"
	"github.com/Slary1260/createapimarkdown/markdown"
	"github.com/sirupsen/logrus"
)

// 手动添加API对象，生成API接口文档
func normal() {
	// start log
	logger.DefaultLogrusLogger()

	headerMap := make(map[string]string, 8)
	headerMap["timestamp"] = ""
	headerMap["randStr"] = "length between 10 and 16"

	doc := document.NewDocument("api/", headerMap, document.WithParseReq(false), document.WithParseRsq(false))

	reqFields, err := doc.ParseReqOrRsp(req)
	if err != nil {
		log.Panic(err)
	}

	rspFields, err := doc.ParseReqOrRsp(rsp)
	if err != nil {
		log.Panic(err)
	}

	item.ReqFields = reqFields
	item.RspFields = rspFields
	err = doc.AddDocItem(item)
	if err != nil {
		log.Panic(err)
	}

	md := markdown.New(doc, markdown.WithMd2Html(true))
	err = md.Generate("doc.md")
	if err != nil {
		log.Panic(err)
	}

	log.Infoln("success")
}

var (
	req = &UpdateActivityRequest{
		Aid:        1,
		IsTop:      1,
		IsValid:    1,
		IsExtend:   1,
		ManualTime: 24,
		AiNum:      0,
		Picture:    "http://picture/url",
		StartTime:  "2022-11-1 17:27:41",
		EndTime:    "2022-11-1 17:27:44",
	}

	rsp = &Result{Data: &ActivityListResponse{
		TotalCount: 100,
		Counts:     []int64{123, 456, 789},
		Names:      []string{"123", "456", "789"},
		Details: []*ActivityListDetail{
			{
				Aid:       1,
				TitleName: []string{"手机"},
				StartTime: "2022-11-1 17:31:10",
				EndTime:   "2022-11-1 17:31:17",
				AiNum:     0,
				Status:    1,
			},
		},
		Goods: []*GoodName{
			{
				Gid:        1,
				GoodsName:  "GoodsName1",
				GoodsPrice: 125.5,
			},
		},
	}}

	item = &document.DocItem{
		Title:    "获取列表",
		Url:      "activity/list",
		Method:   "POST",
		Request:  req,
		Response: rsp,
		Author:   "1260",
	}

	log = logrus.WithFields(logrus.Fields{
		"main": "",
	})
)
