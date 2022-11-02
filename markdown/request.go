/*
 * @Author: tj
 * @Date: 2022-10-27 17:28:40
 * @LastEditors: tj
 * @LastEditTime: 2022-11-02 12:17:48
 * @FilePath: \createApiMarkdown\markdown\request.go
 */
package markdown

import (
	"fmt"
	"strings"

	"createmd/document"
)

const (
	reqPrefix string = "req-"
)

func (m *Markdown) renderReqFields(content string, v *document.DocItem) string {
	if len(v.ReqFields) <= 0 {
		content = strings.Replace(content, "{reqTable}", "", 1)
		return content
	}

	// parse first layer
	reqTable := m.renderReqTable("", 0, v.ReqFields)

	// parse second layer
	subTable := ""
	for _, item := range m.subReqList {
		tpl := m.renderReqTable(item.Title, 1, item.Fields)
		subTable = fmt.Sprintf("%s%s", subTable, tpl)
	}

	// parse third layer
	for _, item := range m.subSubReqList {
		tpl := m.renderReqTable(item.Title, 1, item.Fields)
		subTable = fmt.Sprintf("%s%s", subTable, tpl)
	}

	reqTable = fmt.Sprintf("%s%s", reqTable, subTable)

	content = strings.Replace(content, "{reqTable}", reqTable, 1)

	return content
}

func (m *Markdown) renderReqTable(parentName string, index int, fields []*document.Field) string {
	ts := ""
	if parentName == "" {
		ts = TplReqTable
	} else {
		if index == 0 {
			ts = fmt.Sprintf("\n<a id=\"%s%d.%s\"></a> \n##### %s \n %s ", reqPrefix, m.index, parentName, parentName, TplReqTable)
		} else {
			ts = fmt.Sprintf("\n<a id=\"%s%d.%d.%s\"></a> \n##### %s \n %s ", reqPrefix, m.index, index, parentName, reqPrefix+parentName, TplReqTable)
		}
	}

	params := ""
	for _, v := range fields {
		tpl := m.renderReqParam(v, index)
		params = fmt.Sprintf("%s%s", params, tpl)
	}
	ts = strings.Replace(ts, "{params}", params, 1)

	return ts
}

func (m *Markdown) renderReqParam(v *document.Field, index int) string {
	ts := TplReqParam

	isRequired := "是"
	if !v.IsRequired {
		isRequired = "否"
	}

	ts = strings.Replace(ts, "{name}", v.Name, 1)
	ts = strings.Replace(ts, "{kind}", v.Kind, 1)
	ts = strings.Replace(ts, "{required}", isRequired, 1)

	if len(v.List) > 0 {
		subTable := SubTable{
			Title:  v.Name,
			Fields: v.List,
		}

		if index == 0 {
			m.subReqList = append(m.subReqList, subTable)
			v.Description = fmt.Sprintf("%s [%s](#%s%d.%d.%s)", v.Description, reqPrefix+v.Name, reqPrefix, m.index, index+1, v.Name)
		} else {
			m.subSubReqList = append(m.subSubReqList, subTable)
			v.Description = fmt.Sprintf("%s [%s](#%s%d.%d.%s)", v.Description, reqPrefix+v.Name, reqPrefix, m.index, index, v.Name)
		}
	}

	ts = strings.Replace(ts, "{description}", v.Description, 1)

	return ts
}
