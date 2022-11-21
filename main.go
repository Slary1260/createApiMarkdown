/*
 * @Author: tj
 * @Date: 2022-11-16 10:12:44
 * @LastEditors: tj
 * @LastEditTime: 2022-11-21 17:09:08
 * @FilePath: \createApiMarkdown\main.go
 */
/*
 * @Author: tj
 * @Date: 2022-11-01 15:03:53
 * @LastEditors: tj
 * @LastEditTime: 2022-11-21 16:53:08
 * @FilePath: \createApiMarkdown\main.go
 */
package main

import (
	"github.com/Slary1260/createapimarkdown/gindemo"
)

func main() {
	// 使用gin自动注册路由,导出API接口文档
	gindemo.GinDemo()

	// 手动添加路由,导出API接口文档
	// normal()
}
