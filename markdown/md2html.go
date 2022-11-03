/*
 * @Author: tj
 * @Date: 2022-11-02 10:51:44
 * @LastEditors: tj
 * @LastEditTime: 2022-11-02 12:17:41
 * @FilePath: \createApiMarkdown\markdown\md2html.go
 */
package markdown

import (
	"createApiMarkdown/common"

	"github.com/Depado/bfchroma"
	"github.com/russross/blackfriday/v2"
)

func (m *Markdown) md2html(fileName string) error {
	bytes, err := common.ReadFile(fileName)
	if err != nil {
		return err
	}

	// 将markdown file内容转为html文件
	output := blackfriday.Run(bytes, blackfriday.WithRenderer(bfchroma.NewRenderer()))
	// 过滤不信任的内容
	// html := bluemonday.UGCPolicy().SanitizeBytes(output)

	htmlName := "doc.html"
	common.WriteFile(htmlName, output)
	if err != nil {
		return err
	}

	return nil
}
