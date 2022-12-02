/*
 * @Author: tj
 * @Date: 2022-11-07 14:10:26
 * @LastEditors: tj
 * @LastEditTime: 2022-12-02 10:41:01
 * @FilePath: \createApiMarkdown\gindemo\gindemo.go
 */
package gindemo

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"gindemo": "",
	})
)

func GinDemo() {
	r := gin.Default()
	isMd2Html := true
	initRouter(r, isMd2Html)
	r.Run()
}
