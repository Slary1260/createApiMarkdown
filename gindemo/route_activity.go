/*
 * @Author: tj
 * @Date: 2022-11-02 17:43:17
 * @LastEditors: tj
 * @LastEditTime: 2022-11-16 11:03:45
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

func (a *Activity) Add(req *AddActivityRequest, c *gin.Context) {
	// do something
}

func (a *Activity) Info(req *GetActivityRequest, c *gin.Context) (rsp *GetActivityResponse) {
	// do something
	return
}

func (a *Activity) Update(req *UpdateActivityRequest, c *gin.Context) {
	// do something
}

func (a *Activity) Close(req *CloseActivityRequest, c *gin.Context) {
	// do something
}
