/*
 * @Author: tj
 * @Date: 2022-11-02 10:05:45
 * @LastEditors: tj
 * @LastEditTime: 2022-11-04 15:34:20
 * @FilePath: \createApiMarkdown\document\data.go
 */
package document

type Field struct {
	Name        string   `json:"name"`        // 字段名称
	Kind        string   `json:"kind"`        // 字段类型
	Description string   `json:"description"` // 字段说明
	IsRequired  bool     `json:"required"`    // 是否必填
	List        []*Field `json:"list"`        // 字段列表
}

type DocItem struct {
	Title     string      `json:"title"`     // 标题
	Url       string      `json:"url"`       // 接口地址
	Method    string      `json:"method"`    // 请求类型
	Author    string      `json:"author"`    // 作者
	Request   interface{} `json:"request"`   // 请求参数
	Response  interface{} `json:"response"`  // 返回参数
	ReqFields []*Field    `json:"reqFields"` // 字段列表
	RspFields []*Field    `json:"rspFields"` // 字段列表
}

type Document struct {
	Title   string     `json:"title"`   // 文档标题
	Version string     `json:"version"` // 版本号
	Url     string     `json:"Url"`     // Url
	Items   []*DocItem `json:"items"`   // 接口列表

	mdKey          string
	isNeedParseReq bool
	isNeedParseRsq bool
}
