/*
 * @Author: tj
 * @Date: 2022-11-02 16:47:20
 * @LastEditors: tj
 * @LastEditTime: 2022-11-16 09:49:09
 * @FilePath: \createApiMarkdown\gindemo\route_api.go
 */
package gindemo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	register(NewApi())
}

type Api struct {
}

func NewApi() *Api {
	return &Api{}
}

func (a *Api) GetHtml(c *gin.Context) {
	c.HTML(http.StatusOK, apiDocHtml, nil)
}
