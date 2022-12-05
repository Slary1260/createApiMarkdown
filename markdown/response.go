/*
 * @Author: tj
 * @Date: 2022-10-27 17:29:49
 * @LastEditors: tj
 * @LastEditTime: 2022-12-05 10:17:10
 * @FilePath: \createApiMarkdown\markdown\response.go
 */
package markdown

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Slary1260/createapimarkdown/document"
)

const (
	rspPrefix string = "rsp-"
)

func (m *Markdown) renderRspFields(content string, v *document.DocItem) string {
	if len(v.RspFields) <= 0 {
		content = strings.Replace(content, "{respTable}", "", 1)
		return content
	}

	// parse first layer
	respTable := m.renderRespTable("", 0, v.RspFields)

	// parse second layer
	subTable := ""
	for _, item := range m.subRespList {
		tpl := m.renderRespTable(item.Title, 1, item.Fields)
		subTable = fmt.Sprintf("%s%s", subTable, tpl)
	}

	// parse third layer
	for _, item := range m.subSubRespList {
		tpl := m.renderRespTable(item.Title, 1, item.Fields)
		subTable = fmt.Sprintf("%s%s", subTable, tpl)
	}

	respTable = fmt.Sprintf("%s%s", respTable, subTable)

	content = strings.Replace(content, "{respTable}", respTable, 1)

	return content
}

func (m *Markdown) renderRespTable(parentName string, index int, fields []*document.Field) string {
	ts := ""
	if parentName == "" {
		ts = TplRespTable
	} else {
		if index == 0 {
			ts = fmt.Sprintf("\n<a id=\"%s%d.%s\"></a> \n##### %s \n %s ", rspPrefix, m.index, parentName, parentName, TplRespTable)
		} else {
			ts = fmt.Sprintf("\n<a id=\"%s%d.%d.%s\"></a> \n##### %s \n %s ", rspPrefix, m.index, index, parentName, rspPrefix+parentName, TplRespTable)
		}
	}

	params := ""
	for _, v := range fields {
		tpl := m.renderRespParam(v, index)
		params = fmt.Sprintf("%s%s", params, tpl)
	}

	ts = strings.Replace(ts, "{params}", params, 1)

	return ts
}

func (m *Markdown) renderRespParam(v *document.Field, index int) string {
	ts := TplRespParam
	ts = strings.Replace(ts, "{name}", v.Name, 1)
	ts = strings.Replace(ts, "{kind}", v.Kind, 1)

	if len(v.List) > 0 {
		subTable := SubTable{
			Title:  v.Name,
			Fields: v.List,
		}

		if index == 0 {
			m.subRespList = append(m.subRespList, subTable)
			v.Description = fmt.Sprintf("%s [%s](#%s%d.%d.%s)", v.Description, rspPrefix+v.Name, rspPrefix, m.index, index+1, v.Name)
		} else {
			m.subSubRespList = append(m.subSubRespList, subTable)
			v.Description = fmt.Sprintf("%s [%s](#%s%d.%d.%s)", v.Description, rspPrefix+v.Name, rspPrefix, m.index, index, v.Name)
		}
	}

	ts = strings.Replace(ts, "{description}", v.Description, 1)

	return ts
}

func (m *Markdown) rspToJson(content string, v *document.DocItem) (string, error) {
	if v.Response != nil {
		respParam, err := json.MarshalIndent(v.Response, "", "\t")
		if err != nil {
			return "", err
		}

		content = strings.Replace(content, "{respParam}", fmt.Sprintf("```json\n %s \n```", string(respParam)), 1)
	} else {
		content = strings.Replace(content, "{respParam}", "", 1)
	}

	return content, nil
}
