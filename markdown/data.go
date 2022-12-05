/*
 * @Author: tj
 * @Date: 2022-11-02 10:35:02
 * @LastEditors: tj
 * @LastEditTime: 2022-12-05 10:13:54
 * @FilePath: \createApiMarkdown\markdown\data.go
 */
package markdown

import "github.com/Slary1260/createapimarkdown/document"

type SubTable struct {
	Title  string
	Fields []*document.Field
}

type Markdown struct {
	doc            *document.Document
	index          int
	subReqList     []SubTable
	subSubReqList  []SubTable
	subRespList    []SubTable
	subSubRespList []SubTable
	is2html        bool
}
