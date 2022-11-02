/*
 * @Author: tj
 * @Date: 2022-10-26 18:07:11
 * @LastEditors: tj
 * @LastEditTime: 2022-11-02 12:17:38
 * @FilePath: \createApiMarkdown\markdown\markdown.go
 */
package markdown

import (
	"fmt"
	"os"
	"strings"

	"createmd/common"
	"createmd/document"

	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"markdown": "",
	})
)

func New(doc *document.Document, options ...Option) *Markdown {
	m := &Markdown{
		doc:         doc,
		index:       0,
		subReqList:  make([]SubTable, 0),
		subRespList: make([]SubTable, 0),
	}

	for _, option := range options {
		option(m)
	}

	return m
}

func (m *Markdown) Generate(file string) error {
	page, err := m.renderPage(m.doc)
	if err != nil {
		log.Errorln("Generate renderPage error:", err)
		return err
	}

	os.Remove(file)

	err = common.WriteFile(file, []byte(page))
	if err != nil {
		log.Errorln("Generate WriteFile error:", err)
		return err
	}

	if m.is2html {
		err := m.md2html(file)
		if err != nil {
			log.Errorln("Generate md2html error:", err)
			return err
		}
	}

	return nil
}

func (m *Markdown) renderPage(v *document.Document) (string, error) {
	ts := TplPage
	ts = strings.Replace(ts, "{title}", v.Title, 1)
	ts = strings.Replace(ts, "{version}", v.Version, 1)
	ts = strings.Replace(ts, "{Url}", v.Url, 1)

	body := ""
	for index, item := range v.GetItems() {
		tpl, err := m.renderBody(index+1, item)
		if err != nil {
			return "", err
		}

		body = fmt.Sprintf("%s%s", body, tpl)
	}

	ts = strings.Replace(ts, "{body}", body, 1)

	return ts, nil
}

func (m *Markdown) renderBody(index int, v *document.DocItem) (string, error) {
	m.subReqList = make([]SubTable, 0)
	m.subRespList = make([]SubTable, 0)
	m.index = index

	ts := TplBody
	ts = strings.Replace(ts, "{id}", fmt.Sprintf("%v", index), 1)
	ts = strings.Replace(ts, "{name}", v.Title, 1)
	ts = strings.Replace(ts, "{author}", v.Author, 1)
	ts = strings.Replace(ts, "{method}", string(v.Method), 1)
	ts = strings.Replace(ts, "{url}", string(v.Url), 1)

	// parse request struct fields table
	ts = m.renderReqFields(ts, v)

	// response struct fields table
	ts = m.renderRspFields(ts, v)

	// response data struct json
	ts, err := m.rspToJson(ts, v)
	if err != nil {
		return "", err
	}

	return ts, nil
}