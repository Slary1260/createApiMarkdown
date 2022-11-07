/*
 * @Author: tj
 * @Date: 2022-11-02 17:43:17
 * @LastEditors: tj
 * @LastEditTime: 2022-11-07 14:13:43
 * @FilePath: \createApiMarkdown\gindemo\route_activity.go
 */
package gindemo

import (
	"github.com/gin-gonic/gin"
)

func init() {
	register(NewActivity())
}

type Activity struct {
}

func NewActivity() *Activity {
	return &Activity{}
}

func (a *Activity) List(req *ActivityListRequest, c *gin.Context) (rsp *ActivityListResponse) {
	// do something
	return
}
